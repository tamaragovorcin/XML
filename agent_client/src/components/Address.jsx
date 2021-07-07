import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import getAuthHeader from "../GetHeader";
import Axios from "axios";
import { BASE_URL_AGENT } from "../constants.js";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import { Carousel } from 'react-responsive-carousel';
import { YMaps, Map } from "react-yandex-maps";

import ModalDialog from "../components/ModalDialog";
const mapState = {
	center: [44, 21],
	zoom: 8,
	controls: [],
};
class Address extends Component {
	state = {
		quantity: "",
        errorHeader: "",
		errorMessage: "",
		hiddenErrorAlert: true,
		address: "",
		addressError: "none",
		addressNotFoundError: "none",
		openModal: false,
		coords: [],
	};
    
	constructor(props) {
		super(props);
		this.addressInput = React.createRef();
	}
	handleModalClose = ()=>{
		this.setState({openModal: false})
		window.location.reload();
	}

    onYmapsLoad = (ymaps) => {
		this.ymaps = ymaps;
		new this.ymaps.SuggestView(this.addressInput.current, {
			provider: {
				suggest: (request, options) => this.ymaps.suggest(request),
			},
		});
	};

    handleAddressChange = (event) => {
		this.setState({ address: event.target.value });
	};

    validateForm = (userDTO) => {
		this.setState({
			
			addressError: "none",
			addressNotFoundError: "none",
			
		});

	 if (this.addressInput.current.value === "") {
			this.setState({ addressError: "initial" });
			return false;
		} 
		
		return true;
	};
	componentDidMount() {
	}

	handleQuantityChange = (event) => {

        this.setState({ quantity: event.target.value });
	
    };
	

    handleReserveChange = () => {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1);
		let street;
		let city;
		let country;
		let latitude;
		let longitude;
		let found = true;
		this.ymaps
			.geocode(this.addressInput.current.value, {
				results: 1,
			})
			.then(function (res) {
				if (typeof res.geoObjects.get(0) === "undefined") found = false;
				else {
					var firstGeoObject = res.geoObjects.get(0),
						coords = firstGeoObject.geometry.getCoordinates();
					latitude = coords[0];
					longitude = coords[1];
					country = firstGeoObject.getCountry();
					street = firstGeoObject.getThoroughfare();
					city = firstGeoObject.getLocalities().join(", ");
				}
			})
			.then((res) => {
				let userDTO = {
					Location: { street, country, city, latitude, longitude },
					Products: this.props.products,
                    Buyer: id,
                    
				};
				console.log(userDTO);

				if (this.validateForm(userDTO)) {
					if (found === false) {
						this.setState({ addressNotFoundError: "initial" });
					} else {
						console.log(userDTO);
						Axios.post(BASE_URL_AGENT + "/api/purchase", userDTO, {  headers: { Authorization: getAuthHeader() } })
							.then((res) => {
								if (res.status === 409) {
									this.setState({
										errorHeader: "Resource conflict!",
										errorMessage: "Email already exist.",
										hiddenErrorAlert: false,
									});
								} else if (res.status === 500) {
									this.setState({ errorHeader: "Internal server error!", errorMessage: "Server error.", hiddenErrorAlert: false });
								} else {
									
									this.setState({openModal : true})

									Axios.get(BASE_URL_AGENT + "/api/removeCart/"+ id, {  headers: { Authorization: getAuthHeader() } })
									.then((res) => {
										if (res.status === 409) {
											this.setState({
												errorHeader: "Resource conflict!",
												errorMessage: "Email already exist.",
												hiddenErrorAlert: false,
											});
										} else if (res.status === 500) {
											this.setState({ errorHeader: "Internal server error!", errorMessage: "Server error.", hiddenErrorAlert: false });
										} else {
											console.log("Success");
											
											
										}
									})
									.catch((err) => {
										console.log(err);
									});



								}
							})
							.catch((err) => {
								console.log(err);
							});
					}
				}
			});
	};
	render() {

		return (
			<Modal
				show={this.props.show}
				dialogClassName="modal-80w-150h"
				aria-labelledby="contained-modal-title-vcenter"
				centered
				onHide={this.props.onCloseModal}
			>
				<Modal.Header closeButton>
					<Modal.Title id="contained-modal-title-vcenter">{this.props.header}</Modal.Title>
				</Modal.Header>
				<Modal.Body>
				
					
                <div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Address:</label>
										<input className="form-control" id="suggest" ref={this.addressInput} placeholder="Address" />
									</div>
									<YMaps
										query={{
											load: "package.full",
											apikey: "b0ea2fa3-aba0-4e44-a38e-4e890158ece2",
											lang: "en_RU",
										}}
									>
										<Map
											style={{ display: "none" }}
											state={mapState}
											onLoad={this.onYmapsLoad}
											instanceRef={(map) => (this.map = map)}
											modules={["coordSystem.geo", "geocode", "util.bounds"]}
										></Map>
									</YMaps>
									<div className="text-danger" style={{ display: this.state.addressError }}>
										Address must be entered.
									</div>
									<div className="text-danger" style={{ display: this.state.addressNotFoundError }}>
										Sorry. Address not found. Try different one.
									</div>
								</div>
					

					<div style={{ marginTop: "2rem", marginLeft: "12rem" }}>
						<Button className="mt-3" onClick={this.handleReserveChange}>
							{this.props.buttonName}
						</Button>
					</div>
				</Modal.Body>
				<ModalDialog
					show={this.state.openModal}
					onCloseModal={this.handleModalClose}
					header="Success"
					text="You have successfully placed an order."
				/>
			</Modal>
			
		);
	}
}

export default Address;

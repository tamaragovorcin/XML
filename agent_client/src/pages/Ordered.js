import React, { Component } from "react";
import TopBar from "../components/TopBar";
import Header from "../components/Header";
import Axios from "axios";
import { BASE_URL, BASE_URL_AGENT } from "../constants.js";
import "../App.js";
import { Redirect } from "react-router-dom";
import Order from "../components/Order";
import Address from "../components/Address";
import ModalDialog from "../components/ModalDialog";
import { NavItem } from "react-bootstrap";
import getAuthHeader from "../GetHeader";

class Ordered extends Component {
	state = {
		products: [],
		formShowed: false,
		name: "",
		city: "",
		gradeFrom: "",
		gradeTo: "",
		distanceFrom: "",
		distanceTo: "",
		showingSearched: false,
		showingSorted: false,
		currentLatitude: null,
		currentLongitude: null,
		sortIndicator: 0,
		redirect: false,
		redirectUrl: "",
		showOrderModal: false,
		handleOrderModalClose: false,
		openModal: false,


	};

	hasRole = (reqRole) => {
		let roles = JSON.parse(localStorage.getItem("keyRole"));

		if (roles === null) return false;

		if (reqRole === "*") return true;

		for (let role of roles) {
			if (role === reqRole) return true;
		}
		return false;
	};

	handleNameChange = (event) => {
		this.setState({ name: event.target.value });
	};

	getCurrentCoords = () => {
		if (navigator.geolocation) {
			navigator.geolocation.getCurrentPosition((position) => {
				this.setState({
					currentLatitude: position.coords.latitude,
					currentLongitude: position.coords.longitude,
				});
			});
		}
	};
    handleOrderModalClose = () => {
        this.setState({ showOrderModal: false });
    };
	componentDidMount() {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1);
		Axios.get(BASE_URL_AGENT + "/api/getOrder/" + id, {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ products: res.data });
				console.log(res.data);
			})
			.catch((err) => {
				console.log(err);
			});

	}

	hangleFormToogle = () => {
		this.setState({ formShowed: !this.state.formShowed });
	};

	handleDelete = (e, id) => {
		Axios.get(BASE_URL_AGENT + "/api/cart/remove/" + id, {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {

				this.setState({ openModal: true });
				window.location.reload();
			})
			.catch((err) => {
				console.log(err);
			});

	};



	handleGradeFromChange = (event) => {
		if (event.target.value < 0) this.setState({ gradeFrom: 0 });
		else this.setState({ gradeFrom: event.target.value });
	};

	handleGradeToChange = (event) => {
		if (event.target.value > 5) this.setState({ gradeTo: 5 });
		else this.setState({ gradeTo: event.target.value });
	};

	handleDistanceFromChange = (event) => {
		this.setState({ distanceFrom: event.target.value });
	};

	handleDistanceToChange = (event) => {
		this.setState({ distanceTo: event.target.value });
	};

	handleCityChange = (event) => {
		this.setState({ city: event.target.value });
	};



	handleClickOnPharmacy = (id) => {
		this.setState({ shirt: id });
		this.setState({ showOrderModal: true });
		this.setState({ colors: id.colors });

	};

	render() {
		if (this.state.redirect) return <Redirect push to={this.state.redirectUrl} />;

		return (
			<React.Fragment>
				<TopBar />
				<Header />

				<div className="container" style={{ marginTop: "10%" }}>
					<h5 className=" text-center mb-0 mt-2 text-uppercase">My orders</h5>


					<table className="table table-hover" style={{ width: "100%", marginTop: "3rem" }}>
						<tbody>
							{this.state.products.map((p) => (
								<div><label>Deliver to address: {p.Location.Country}, {p.Location.Street} </label>
								
									{p.Product.map((item) => (
								<tr
									id={p.id}
									key={p.id}
									style={{ cursor: "pointer" }}

								>
								
									<td width="130em">
										<img className="img-fluid" src={`data:image/jpg;base64,${item.Media[0]}`} width="70em" />
									</td>
									<td>
										<div>
											<b>Name: </b> {item.Name}
										</div>
										<div>
											<b>Price: </b> {item.Price}
										</div>
										<div>
											<b>Quantity: </b> {item.Quantity}
										</div>





									</td>
								
								</tr>
									))}
									</div>
							))}
						</tbody>
					</table>


			
				</div>
				<ModalDialog
					show={this.state.openModal}
					onCloseModal={this.handleModalClose}
					header="Success"
					text="You have successfully removed the item."
				/>

				<Address
					buttonName="Add"
					header="Add product to cart"
					show={this.state.showOrderModal}
					onCloseModal={this.handleOrderModalClose}
					handleAddress={this.handleAddressOrderChange}
					products={this.state.products}
				/>
			</React.Fragment>
		);
	}
}

export default Ordered;


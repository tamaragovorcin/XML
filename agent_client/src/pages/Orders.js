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
class TShirtsMen extends Component {
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
		show: true,


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
	handleModalClose = ()=>{
		this.setState({openModal: false})
		window.location.reload();
	}
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
		Axios.get(BASE_URL_AGENT + "/api/getAllCart/" + id)
			.then((res) => {
				this.setState({ products: res.data });
				console.log(res.data);
				
				if(res.data.length === 0){
					this.setState({ show: false });
				}
			})
			.catch((err) => {
				console.log(err);
			});

	}

	hangleFormToogle = () => {
		this.setState({ formShowed: !this.state.formShowed });
	};

	handleDelete = (e, id) => {
		Axios.get(BASE_URL_AGENT + "/api/cart/remove/" + id)
			.then((res) => {

				this.setState({ openModal: true });
				
			})
			.catch((err) => {
				console.log(err);
			});

	};


	handleOrder = () => {
		this.setState({ showOrderModal: true });



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
					<h5 className=" text-center mb-0 mt-2 text-uppercase">My cart</h5>


					<table className="table table-hover" style={{ width: "100%", marginTop: "3rem" }}>
						<tbody>
							{this.state.products.map((p) => (
								<tr
									id={p.id}
									key={p.id}
									style={{ cursor: "pointer" }}

								>
									<td width="130em">
										<img className="img-fluid" src={`data:image/jpg;base64,${p.Media[0]}`} width="70em" />
									</td>
									<td>
										<div>
											<b>Name: </b> {p.Product.Name}
										</div>
										<div>
											<b>Price: </b> {p.Product.Price}
										</div>
										<div>
											<b>Quantity: </b> {p.Quantity}
										</div>




										<div>  <button
											style={{
												background: "#1977cc",
												marginTop: "15px",
												marginLeft: "40%",
												width: "20%",
											}}
											onClick={(e) => this.handleDelete(e, p.Id)}
											className="btn btn-primary btn-xl"
											id="sendMessageButton"
											type="button"
										>
											Remove from cart
										</button></div>

									</td>
								</tr>
							))}
						</tbody>
					</table>


					<div><button
					hidden = {!this.state.show}
						style={{
							background: "#1977cc",
							marginTop: "15px",
							marginLeft: "40%",
							width: "20%",
						}}
						onClick={this.handleOrder}
						className="btn btn-primary btn-xl"
						id="sendMessageButton"
						type="button"
					>
						Add address and place an order
					</button></div>
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

export default TShirtsMen;


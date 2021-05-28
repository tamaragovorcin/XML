import React, { Component } from "react";
import Header from "./Header";
import TopBar from "./TopBar";
import { BASE_URL } from "../constants.js";
import Axios from "axios";
import { Redirect } from "react-router-dom";
import HeadingAlert from "./HeadingAlert";


class RegisterPage extends Component {
	state = {
		errorHeader: "",
		errorMessage: "",
		hiddenErrorAlert: true,
		email: "",
		password: "",
		repeatPassword: "",
		name: "",
		surname: "",
		phoneNumber: "",
		emailError: "none",
		passwordError: "none",
		repeatPasswordError: "none",
		repeatPasswordSameError: "none",
		nameError: "none",
		surnameError: "none",
		phoneError: "none",
		emailNotValid: "none",
		openModal: false,
		coords: [],
	};

	handleEmailChange = (event) => {
		this.setState({ email: event.target.value });
	};

	handlePasswordChange = (event) => {
		this.setState({ password: event.target.value });
	};
	
	handleRepeatPasswordChange = (event) => {
		this.setState({ repeatPassword: event.target.value });
	};

	handleNameChange = (event) => {
		this.setState({ name: event.target.value });
	};

	handleSurnameChange = (event) => {
		this.setState({ surname: event.target.value });
	};

	handlePhoneNumberChange = (event) => {
		this.setState({ phoneNumber: event.target.value });
	};

	validateForm = (completeDTO) => {
		this.setState({
			emailError: "none",
			emailNotValid: "none",
			nameError: "none",
			surnameError: "none",
			phoneError: "none",
			passwordError: "none",
			repeatPasswordError: "none",
			repeatPasswordSameError: "none",
		});

		if (completeDTO.user.email === "") {
			this.setState({ emailError: "initial" });
			return false;
		} else if (!completeDTO.user.email.includes("@")) {
			this.setState({ emailNotValid: "initial" });
			return false;
		} else if (completeDTO.user.name === "") {
			this.setState({ nameError: "initial" });
			return false;
		} else if (completeDTO.user.surname === "") {
			this.setState({ surnameError: "initial" });
			return false;
		}  else if (completeDTO.user.phoneNumber === "") {
			this.setState({ phoneError: "initial" });
			return false;
		} else if (completeDTO.user.password === "") {
			this.setState({ passwordError: "initial" });
			return false;
		} else if (this.state.repeatPassword === "") {
			this.setState({ repeatPasswordError: "initial" });
			return false;
		}else if (completeDTO.user.password !== this.state.repeatPassword) {
			this.setState({ repeatPasswordSameError: "initial" });
			return false;
		} 
		return true;
	};


	handleSignUp = () => {		
		let userDTO = {
			email: this.state.email,
			name: this.state.name,
			surname: this.state.surname,
			phoneNumber: this.state.phoneNumber,
			password: this.state.password,
		};
	
		if (this.validateForm(userDTO)) {
			
			Axios.post(BASE_URL + "/api/manager/signup", userDTO, { validateStatus: () => true })
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
							this.setState({ openModal: true });
							this.setState({ redirect: true })
						}
					})
					.catch((err) => {
						console.log(err);
					});
		}
		
	};

	handleCloseAlert = () => {
		this.setState({ hiddenErrorAlert: true });
	};

	render() {
		if (this.state.redirect) return <Redirect push to="/login" />;

		return (
			<React.Fragment>
				<TopBar />
				<Header />

				<div className="container" style={{ marginTop: "10%" }}>
					<HeadingAlert
						hidden={this.state.hiddenErrorAlert}
						header={this.state.errorHeader}
						message={this.state.errorMessage}
						handleCloseAlert={this.handleCloseAlert}
					/>
					<h5 className=" text-center  mb-0 text-uppercase" style={{ marginTop: "2rem" }}>
						Registration
					</h5>

					<div className="row section-design">
						<div className="col-lg-8 mx-auto">
							<br />
							<form id="contactForm" name="sentMessage" noValidate="novalidate">
							
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Email address:</label>
										<input
											placeholder="Email address"
											className="form-control"
											id="email"
											type="text"
											onChange={this.handleEmailChange}
											value={this.state.email}
										/>
									</div>
									<div className="text-danger" style={{ display: this.state.emailError }}>
										Email address must be entered.
									</div>
									<div className="text-danger" style={{ display: this.state.emailNotValid }}>
										Email address is not valid.
									</div>
								</div>
                                
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Manager name:</label>
										<input
											placeholder="Name"
											class="form-control"
											type="text"
											id="name"
											onChange={this.handleNameChange}
											value={this.state.name}
										/>
									</div>
									<div className="text-danger" style={{ display: this.state.nameError }}>
										Name must be entered.
									</div>
								</div>
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label> Manager surname:</label>
										<input
											placeholder="Surname"
											class="form-control"
											type="text"
											id="surname"
											onChange={this.handleSurnameChange}
											value={this.state.surname}
										/>
									</div>
									<div className="text-danger" style={{ display: this.state.surnameError }}>
										Surname must be entered.
									</div>
								</div>
								
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Phone number:</label>
										<input
											placeholder="Phone number"
											class="form-control"
											id="phone"
											type="text"
											onChange={this.handlePhoneNumberChange}
											value={this.state.phoneNumber}
										/>
									</div>
									<div className="text-danger" style={{ display: this.state.phoneError }}>
										Phone number must be entered.
									</div>
								</div>
								<div className="control-group">
									<label>Password:</label>
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<input
											placeholder="Password"
											class="form-control"
											type="password"
											onChange={this.handlePasswordChange}
											value={this.state.password}
										/>
									</div>
									<div className="text-danger" style={{ display: this.state.passwordError }}>
										Password must be entered.
									</div>
								</div>
								<div className="control-group">
									<label>Repeat password:</label>
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<input
											placeholder="Repeat password"
											class="form-control"
											type="password"
											onChange={this.handleRepeatPasswordChange}
											value={this.state.repeatPassword}
										/>
									</div>
									<div className="text-danger" style={{ display: this.state.repeatPasswordError }}>
										Repeat password must be entered.
									</div>
									<div className="text-danger" style={{ display: this.state.repeatPasswordSameError }}>
										Passwords are not the same.
									</div>
								</div>

								<div className="form-group">
									<button
										style={{
											background: "#1977cc",
											marginTop: "15px",
											marginLeft: "40%",
											width: "20%",
										}}
										onClick={this.handleSignUp}
										className="btn btn-primary btn-xl"
										id="sendMessageButton"
										type="button"
									>
										Sign Up
									</button>
								</div>
							</form>
						</div>
					</div>
				</div>
				
			</React.Fragment>
		);
	}
}

export default RegisterPage;

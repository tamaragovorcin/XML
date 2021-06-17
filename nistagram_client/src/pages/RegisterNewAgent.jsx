import React, { Component } from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import Axios from "axios";
import { Redirect } from "react-router-dom";
import HeadingAlert from "../components/HeadingAlert";
import ModalDialog from "../components/ModalDialog";

import { BASE_URL } from "../constants.js";
class RegisterNewAgent extends Component {
	state = {
		errorHeader: "",
		errorMessage: "",
		hiddenErrorAlert: true,
		email: "",
		date: "",
		password: "",
		repeatPassword: "",
		name: "",
		surname: "",
		phoneNumber: "",
		emailError: "none",
		passwordError: "none",
		dateError: "none",
		repeatPasswordError: "none",
		repeatPasswordSameError: "none",
		nameError: "none",
		surnameError: "none",
		phoneError: "none",
		emailNotValid: "none",
		openModal: false,
		coords: [],
		gender: "",
		biography: "",
		username: "",
		usernameError: "none",
		usernameNotValid: "none",
		selectedDate: "",
		private : false,
        textSuccessfulModal: ""

	};

	handleDateChange = (event) => {
		this.setState({ selectedDate: event.target.value });
	};
	handleEmailChange = (event) => {
		this.setState({ email: event.target.value });
	};
	handleUsernameChange = (event) => {
		this.setState({ username: event.target.value });
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
	handleModalClose = () => {
		this.setState({ openModal: false, redirect: true });
	};
	handlePhoneNumberChange = (event) => {
		this.setState({ phoneNumber: event.target.value });
	};
	handleBiographyChange = (event) => {
		this.setState({ biography: event.target.value });
	};
	validateForm = (userDTO) => {
		this.setState({
			emailError: "none",
			emailNotValid: "none",
			nameError: "none",
			dateError: "none",
			surnameError: "none",
			phoneError: "none",
			passwordError: "none",
			repeatPasswordError: "none",
			repeatPasswordSameError: "none",
			usernameError: "none",
			usernameNotValid: "none",
			website : "",
		});

		if (this.state.username === "") {
			this.setState({ usernameError: "initial" });
			return false;
		}

		else if (!this.state.email === "") {
			this.setState({ emailError: "initial" });
			return false;
		}
		else if (!this.state.email.includes("@")) {
			this.setState({ emailNotValid: "initial" });
			return false;
		}
		else if (!this.state.email.includes(".com")) {
			this.setState({ emailNotValid: "initial" });
			return false;
		}

		else if (this.state.name === "") {
			this.setState({ nameError: "initial" });
			return false;
		} else if (this.state.surname === "") {
			this.setState({ surnameError: "initial" });
			return false;


		} else if (this.state.phoneNumber === "") {
			this.setState({ phoneError: "initial" });
			return false;
		} else if (this.state.selectedDate === "") {
			this.setState({ dateError: "initial" });
			return false;
		} else if (this.state.password === "") {
			this.setState({ passwordError: "initial" });
			return false;
		} else if (this.state.repeatPassword === "") {
			this.setState({ repeatPasswordError: "initial" });
			return false;
		} else if (this.state.password !== this.state.repeatPassword) {
			this.setState({ repeatPasswordSameError: "initial" });
			return false;
		}
		return true;
	};


	
	handleRegisterAgent = () => {

		let userDTO = {
			Email: this.state.email,
			Username: this.state.username,
			Name: this.state.name,
			LastName: this.state.surname,
			PhoneNumber: this.state.phoneNumber,
			Gender: this.state.gender,
			DateOfBirth: this.state.selectedDate,
			Password: this.state.password,
			Biography: this.state.biography,
			Private: this.state.private,
			Website : this.state.website
		};

		if (this.validateForm(userDTO)) {
			Axios.post(BASE_URL + "/api/users/agent/byAdmin/", userDTO)
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
                        this.setState({ textSuccessfulModal: "You have successfully registred agent" });

					}
					const user1Id = {id: res.data}
					Axios.post(BASE_URL + "/api/userInteraction/api/createUser", user1Id)
					.then((res) => {
							console.log(res.data)
					})
					.catch ((err) => {
				console.log(err);
			});
				})
				.catch((err) => {
					if (err.response.status === 409) {
						this.setState({
							errorHeader: "Email taken",
							errorMessage: "Email already exist.",
							hiddenErrorAlert: false,
						});
					}
					else if (err.response.status === 500) {
						this.setState({ errorHeader: "Username taken", errorMessage: "User with this username already exists", hiddenErrorAlert: false });}

				});
		}


	};
	handleGenderChange(event) {

		this.setState({ gender: event.target.value });
	}
	handleBiographyChange(event) {

		this.setState({ biography: event.target.value });
	}
	handlePrivateChange(event) {
		if(this.state.private=== true) {
			this.setState({ private: false });
		}
		if(this.state.private=== false) {
			this.setState({ private: true });
		}
	}
	handleCloseAlert = () => {
		this.setState({ hiddenErrorAlert: true });
	};
	handleWebsiteChange = (event) => {
		this.setState({website: event.target.value})
	}

	render() {

		return (
			<React.Fragment>
				<TopBar />
				<Header />

				<div className="container" style={{ marginTop: "10%", border: "1px solid black" }}>
					<HeadingAlert
						hidden={this.state.hiddenErrorAlert}
						header={this.state.errorHeader}
						message={this.state.errorMessage}
						handleCloseAlert={this.handleCloseAlert}
					/>
					<h5 className=" text-center  mb-0 text-uppercase" style={{ marginTop: "2rem" }}>
						Agent registration
					</h5>

					<div className="row section-design" style={{ border: "1 solid black" }}>
						<div className="col-lg-8 mx-auto">
							<br />
							<form id="contactForm" name="sentMessage" noValidate="novalidate">

								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Username</label>
										<input
											placeholder="Username"
											className="form-control"
											id="username"
											type="text"
											onChange={this.handleUsernameChange}
											value={this.state.username}
										/>
									</div>
									<div className="text-danger" style={{ display: this.state.usernameError }}>
										Username must be entered.
									</div>
									<div className="text-danger" style={{ display: this.state.usernameNotValid }}>
										Username is not valid.
									</div>
								</div>
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
										<label>Biography:</label>
										<input
											readOnly={this.state.hiddenEditInfo}
											className={!this.state.hiddenEditInfo === false ? "form-control-plaintext" : "form-control"}
											placeholder="Biography"
											type="text"
											onChange={this.handleBiographyChange}
											value={this.state.biography}
										/>
									</div>

								</div>
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>First name:</label>
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
										<label>Surname:</label>
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
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Website:</label>
										<input
											placeholder="Website"
											class="form-control"
											type="text"
											id="website"
											onChange={this.handleWebsiteChange}
											value={this.state.website}
										/>
									</div>
									
								</div>
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Date of birth:</label>
										<input
											placeholder="Date of birth"
											class="form-control"
											id="date"
											type="date"
											onChange={this.handleDateChange}
											value={this.state.selectedDate}
										/>
									</div>
									<div className="text-danger" style={{ display: this.state.dateError }}>
										Date of birth must be entered.
									</div>
								</div>

								<div style={{ color: "#6c757d", opacity: 1 }}>
									<p><input type="radio" checked value="Male" name="gender" onChange={(e) => this.handleGenderChange(e)} /> Male</p>
									<p><input type="radio" value="Female" name="gender" onChange={(e) => this.handleGenderChange(e)} /> Female</p>
									<p><input type="radio" value="Other" name="gender" onChange={(e) => this.handleGenderChange(e)} /> Other </p>
								</div>
								<div>
									<p style={{ color: "#6c757d", opacity: 1 }} >Private </p>
									<label class="switch">
										<input type="checkbox" onChange={(e) => this.handlePrivateChange(e)} />
										<span class="slider round"></span>

									</label>
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
										onClick={this.handleRegisterAgent}
										className="btn btn-primary btn-xl"
										id="sendMessageButton"
										type="button"
									>
										Register agent
									</button>
								</div>
                              

							</form>
						</div>
					</div>
				</div>
				<ModalDialog
					show={this.state.openModal}
					onCloseModal={this.handleModalClose}
					header="Successful registration"
					text={this.state.textSuccessfulModal}
				/>
			</React.Fragment>
		);
	}
}

export default RegisterNewAgent;

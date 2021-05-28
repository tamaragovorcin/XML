import React, { Component } from "react";
import Header from "../Header";
import TopBar from "../TopBar";
import { BASE_URL } from "../../constants.js";
import Axios from "axios";
import PasswordChange from "../Users/PasswordChange";
import getAuthHeader from "../../GetHeader";
import { Redirect } from "react-router-dom";
import HeadingSuccessAlert from "../HeadingSuccessAlert";
import HeadingAlert from "../HeadingAlert";


class ProfileInfo extends Component {
	state = {
		id: "",
		email: "",
		password: "",
		name: "",
		surname: "",
		phoneNumber: "",
		nameError: "none",
		surnameError: "none",
		addressError: "none",
		phoneError: "none",
		oldPasswordEmptyError: "none",
		newPasswordEmptyError: "none",
		newPasswordRetypeEmptyError: "none",
		newPasswordRetypeNotSameError: "none",
		openModal: false,
		openPasswordModal: false,
		hiddenEditInfo: true,
		redirect: false,
		hiddenPasswordErrorAlert: true,
		errorPasswordHeader: "",
		errorPasswordMessage: "",
		hiddenSuccessAlert: true,
		successHeader: "",
		successMessage: "",
		hiddenFailAlert: true,
		failHeader: "",
		failMessage: "",
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

	
	componentDidMount() {
		if (!this.hasRole("ROLE_MANAGER") && !this.hasRole("ROLE_ADMIN") &&  !this.hasRole("ROLE_PLAYER") && !this.hasRole("ROLE_JUDGE") && !this.hasRole("ROLE_COACH")) {
			this.setState({ redirect: true });
		} else {

			Axios.get(BASE_URL + "/api/user", { validateStatus: () => true, headers: { Authorization: getAuthHeader() } })
				.then((res) => {
					if (res.status !== 401) {
						this.setState({
							id: res.data.Id,
							email: res.data.email,
							name: res.data.name,
							surname: res.data.surname,
							phoneNumber: res.data.phoneNumber,
						});

					} else {
						this.setState({ redirect: true });
					}
				})
				.catch((err) => {
					console.log(err);
				});
		}
	}

	handleEmailChange = (event) => {
		this.setState({ email: event.target.value });
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

	validateForm = (user) => {
		this.setState({
			nameError: "none",
			surnameError: "none",
			phoneError: "none",
		});

		if (user.name === "") {
			this.setState({ nameError: "initial" });
			return false;
		} else if (user.surname === "") {
			this.setState({ surnameError: "initial" });
			return false;
		} 
         else if (user.phoneNumber === "") {
			this.setState({ phoneError: "initial" });
			return false;
		}
		return true;
	};


	handleSuccessModalClose = () => {
		this.setState({ openSuccessModal: false });
	};

	handlePasswordModalClose = () => {
		this.setState({ openPasswordModal: false });
	};

	handleChangeInfo = () => {
		this.setState({
			hiddenSuccessAlert: true,
			successHeader: "",
			successMessage: "",
			hiddenFailAlert: true,
			failHeader: "",
			failMessage: "",
		});

        let userDTO = {
            name: this.state.name,
            surname: this.state.surname,
            phoneNumber: this.state.phoneNumber,
        };

        if (this.validateForm(userDTO)) {
           
                Axios.put(BASE_URL + "/api/user/update", userDTO, {
                    validateStatus: () => true,
                    headers: { Authorization: getAuthHeader() },
                })
                    .then((res) => {
                        if (res.status === 400) {
                            this.setState({ hiddenFailAlert: false, failHeader: "Bad request", failMessage: "Invalid argument." });
                        } else if (res.status === 500) {
                            this.setState({ hiddenFailAlert: false, failHeader: "Internal server error", failMessage: "Server error." });
                        } else if (res.status === 204) {
                            console.log("Success");
                            this.setState({
                                hiddenSuccessAlert: false,
                                successHeader: "Success",
                                successMessage: "You successfully updated your information.",
                                hiddenEditInfo: true,
                            });
                        }
                    })
                    .catch((err) => {
                        console.log(err);
                    });
            
        }
	
	};


	handlePasswordModal = () => {
		this.setState({ hiddenEditInfo: true, openPasswordModal: true });
	};


	changePassword = (oldPassword, newPassword, newPasswordRetype) => {
		console.log(oldPassword, newPassword, newPasswordRetype);

		this.setState({
			hiddenPasswordErrorAlert: true,
			errorPasswordHeader: "",
			errorPasswordMessage: "",
			hiddenEditInfo: true,
			oldPasswordEmptyError: "none",
			newPasswordEmptyError: "none",
			newPasswordRetypeEmptyError: "none",
			newPasswordRetypeNotSameError: "none",
			hiddenSuccessAlert: true,
			successHeader: "",
			successMessage: "",
		});

		if (oldPassword === "") {
			this.setState({ oldPasswordEmptyError: "initial" });
		} else if (newPassword === "") {
			this.setState({ newPasswordEmptyError: "initial" });
		} else if (newPasswordRetype === "") {
			this.setState({ newPasswordRetypeEmptyError: "initial" });
		} else if (newPasswordRetype !== newPassword) {
			this.setState({ newPasswordRetypeNotSameError: "initial" });
		} else {
			let passwordChangeDTO = { oldPassword, newPassword };
			Axios.post(BASE_URL + "/api/user/changePassword", passwordChangeDTO, {
				validateStatus: () => true,
				headers: { Authorization: getAuthHeader() },
			})
				.then((res) => {
					if (res.status === 403) {
						this.setState({
							hiddenPasswordErrorAlert: false,
							errorPasswordHeader: "Bad credentials",
							errorPasswordMessage: "You entered wrong password.",
						});
					} else if (res.status === 400) {
						this.setState({
							hiddenPasswordErrorAlert: false,
							errorPasswordHeader: "Invalid new password",
							errorPasswordMessage: "Invalid new password.",
						});
					} else if (res.status === 500) {
						this.setState({
							hiddenPasswordErrorAlert: false,
							errorPasswordHeader: "Internal server error",
							errorPasswordMessage: "Server error.",
						});
					} else if (res.status === 204) {
						this.setState({
							hiddenSuccessAlert: false,
							successHeader: "Success",
							successMessage: "You successfully changed your password.",
							openPasswordModal: false,
						});
					}
					console.log(res);
				})
				.catch((err) => {
					console.log(err);
				});
		}
	};

	handleEditInfoClick = () => {
		this.setState({ hiddenEditInfo: false });
	};

	handleCloseAlertPassword = () => {
		this.setState({ hiddenPasswordErrorAlert: true });
	};

	handleCloseAlertSuccess = () => {
		this.setState({ hiddenSuccessAlert: true });
	};

	handleCloseAlertFail = () => {
		this.setState({ hiddenFailAlert: true });
	};

	

	render() {
		if (this.state.redirect) return <Redirect push to="/unauthorized" />;

		return (
			<React.Fragment>
				<TopBar />
				<Header />

				<div className="container mt-15" style={{ marginTop: "8%" }}>
					<HeadingSuccessAlert
						hidden={this.state.hiddenSuccessAlert}
						header={this.state.successHeader}
						message={this.state.successMessage}
						handleCloseAlert={this.handleCloseAlertSuccess}
					/>
					<HeadingAlert
						hidden={this.state.hiddenFailAlert}
						header={this.state.failHeader}
						message={this.state.failMessage}
						handleCloseAlert={this.handleCloseAlertFail}
					/>
					<div className="row mt-10">
						<div className="col shadow p-3 bg-white rounded">
							<h5 className=" text-center text-uppercase">Personal Information</h5>
							<form id="contactForm" name="sentMessage">
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Email address:</label>
										<input
											readOnly
											placeholder="Email address"
											className="form-control-plaintext"
											id="name"
											type="text"
											onChange={this.handleEmailChange}
											value={this.state.email}
										/>
									</div>
								</div>
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Name:</label>
										<input
											readOnly={this.state.hiddenEditInfo}
											className={!this.state.hiddenEditInfo === false ? "form-control-plaintext" : "form-control"}
											placeholder="Name"
											type="text"
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
											readOnly={this.state.hiddenEditInfo}
											className={!this.state.hiddenEditInfo === false ? "form-control-plaintext" : "form-control"}
											placeholder="Surname"
											type="text"
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
											readOnly={this.state.hiddenEditInfo}
											className={!this.state.hiddenEditInfo === false ? "form-control-plaintext" : "form-control"}
											type="text"
											onChange={this.handlePhoneNumberChange}
											value={this.state.phoneNumber}
										/>
									</div>
									<div className="text-danger" style={{ display: this.state.phoneError }}>
										Phone number must be entered.
									</div>
								</div>
								<div className="form-group text-center" hidden={this.state.hiddenEditInfo}>
									<button
										style={{ background: "#1977cc", marginTop: "15px" }}
										onClick={this.handleChangeInfo}
										className="btn btn-primary btn-xl"
										id="sendMessageButton"
										type="button"
									>
										Change information
									</button>
								</div>
								<br />

								<div className="form-group">
									<div className="form-group controls mb-0 pb-2">
										<div className="form-row justify-content-center">
											<div className="form-col" hidden={!this.state.hiddenEditInfo}>
												<button
													onClick={this.handleEditInfoClick}
													className="btn btn-outline-primary btn-xl"
													id="sendMessageButton"
													type="button"
												>
													Edit Info
												</button>
											</div>
											<div className="form-col ml-3">
												<button
													onClick={this.handlePasswordModal}
													className="btn btn-outline-primary btn-xl"
													id="sendMessageButton"
													type="button"
												>
													Change Password
												</button>
											</div>
										</div>
									</div>
								</div>
							</form>
						</div>
					</div>
				</div>
				<PasswordChange
					handleCloseAlertPassword={this.handleCloseAlertPassword}
					hiddenPasswordErrorAlert={this.state.hiddenPasswordErrorAlert}
					errorPasswordHeader={this.state.errorPasswordHeader}
					errorPasswordMessage={this.state.errorPasswordMessage}
					oldPasswordEmptyError={this.state.oldPasswordEmptyError}
					newPasswordEmptyError={this.state.newPasswordEmptyError}
					newPasswordRetypeEmptyError={this.state.newPasswordRetypeEmptyError}
					newPasswordRetypeNotSameError={this.state.newPasswordRetypeNotSameError}
					show={this.state.openPasswordModal}
					changePassword={this.changePassword}
					onCloseModal={this.handlePasswordModalClose}
					header="Change password"
				/>
			</React.Fragment>
		);
	}
}

export default ProfileInfo;

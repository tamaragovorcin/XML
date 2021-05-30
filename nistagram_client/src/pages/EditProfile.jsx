import React, { Component } from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { BASE_URL } from "../constants.js";
import Axios from "axios";
import { Redirect } from "react-router-dom";
import HeadingAlert from "../components/HeadingAlert";
import HeadingSuccessAlert from "../components/HeadingSuccessAlert"
import SidebarSettings from "../components/SidebarSettings"



class EditProfile extends Component {
	state = {
		id: "",
		username: "mladenkak",
		name: "",
		surname : "",
		email: "",
		phoneNumber: "",
		gender : "",
		dateOfBirth : "",
		webSite : "",
		biography : "",
		nameError: "none",
		surnameError: "none",
		addressError: "none",
		phoneError: "none",
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
                    //  headers: { Authorization: getAuthHeader() },
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
		if (this.state.redirect) return <Redirect push to="/login" />;

		return (
			<React.Fragment>
				
				<TopBar />
				<Header />

				<div className="container=fluid" style={{ marginTop: "8%",marginLeft:"5%",marginRight:"5%",background: "#fcfafa"}}>
					<HeadingAlert
						hidden={this.state.hiddenErrorAlert}
						header={this.state.errorHeader}
						message={this.state.errorMessage}
						handleCloseAlert={this.handleCloseAlert}
					/>
                    <br/>
						<h2 className=" text-center  mb-0 text-uppercase" style={{ marginTop: "0", color:"#2c4964" }}>
						Profile settings
					</h2>
                    <br />
					<div className="row section-design" style={{ marginLeft: "5%",marginRight:"5%"}}>
							
                            <div className="col-md-2 padding-0">
                            <SidebarSettings/></div>
                            <div className="col-md-8 padding-0">
							<TopBar />
							<Header />

				<div className="container" style={{ marginTop: "0%" }}>
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
							<form id="contactForm" name="sentMessage">
								
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Username:</label>
										<input
											className="form-control-plaintext"
											id="name"
											type="text"
											onChange={this.handleEmailChange}
											value={this.state.username}
										/>
									</div>
								</div>
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Name:</label>
										<input
											readOnly={this.state.hiddenEditInfo}
											className={!this.state.hiddenEditInfo === false ? "form-control-plaintext" : "form-control"}
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
										<label>Email address:</label>
										<input
											readOnly
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
										<label>Phone number:</label>
										<input
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
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Date of birth:</label>
										<input
											readOnly={this.state.hiddenEditInfo}
											className={!this.state.hiddenEditInfo === false ? "form-control-plaintext" : "form-control"}
											type="date"
											onChange={this.handlePhoneNumberChange}
											value={this.state.dateOfBirth}
										/>
									</div>
									<div className="text-danger" style={{ display: this.state.phoneError }}>
										Date of birth must be entered.
									</div>
								</div>
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Link for web-site:</label>
										<input
											readOnly={this.state.hiddenEditInfo}
											className={!this.state.hiddenEditInfo === false ? "form-control-plaintext" : "form-control"}
											type="text"
											onChange={this.handlePhoneNumberChange}
											value={this.state.webSite}
										/>
									</div>
									
								</div>
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Biography:</label>
										<textarea
											readOnly={this.state.hiddenEditInfo}
											className={!this.state.hiddenEditInfo === false ? "form-control-plaintext" : "form-control"}
											onChange={this.handlePhoneNumberChange}
											value={this.state.biography}
										/>
									</div>
									
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
													Save
												</button>
											</div>
										
										</div>
									</div>
								</div>
							</form>
						</div>
					</div>
				</div>



                            </div>

                            </div>
							
                            
				</div>
				
			</React.Fragment>
		);
	}
}

export default EditProfile;

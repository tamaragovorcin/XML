import React, { Component } from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { BASE_URL, BASE_URL_USER } from "../constants.js";
import Axios from "axios";
import { Redirect } from "react-router-dom";
import HeadingAlert from "../components/HeadingAlert";
import HeadingSuccessAlert from "../components/HeadingSuccessAlert"
import SidebarSettings from "../components/SidebarSettings"



class EditProfile extends Component {
	state = {
		id: "",
		username: "",
		name: "",
		lastName : "",
		email: "",
		phoneNumber: "",
		gender : "Female",
		dateOfBirth : "",
		webSite : "",
		biography : "",
		private : true,
		usernameError : "none",
		nameError: "none",
		surnameError: "none",
		phoneError: "none",
		dateOfBirthError : "none",
		emailError : "none",
		emailNotValid: "none",
		genderError : "none",
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
	
		let id =localStorage.getItem("userId")
	Axios.get(BASE_URL_USER + "/api/" + id)
				.then((res) => {
					if (res.status === 401) {
						this.setState({ errorHeader: "Bad credentials!", errorMessage: "Wrong username or password.", hiddenErrorAlert: false });
					} else if (res.status === 500) {
						this.setState({ errorHeader: "Internal server error!", errorMessage: "Server error.", hiddenErrorAlert: false });
					} else {
						this.setState({
							id: res.data.Id,
							username : res.data.ProfileInformation.Username,
							name: res.data.ProfileInformation.Name,
							lastName : res.data.ProfileInformation.LastName,
							email : res.data.ProfileInformation.Email,
							phoneNumber : res.data.ProfileInformation.PhoneNumber,
							gender : res.data.ProfileInformation.Gender,
							dateOfBirth  : res.data.ProfileInformation.DateOfBirth,
							webSite : res.data.WebSite,
							biography : res.data.Biography,
							private : res.data.Private
						});
					}
				})
				.catch ((err) => {
			console.log(err);
		});

	}
	
	handleEmailChange = (event) => {
		this.setState({ email: event.target.value });
	};
	handleUsernameChange = (event) => {
		this.setState({ username: event.target.value });
	};

	handleNameChange = (event) => {
		this.setState({ name: event.target.value });
	};

	handleSurnameChange = (event) => {
		this.setState({ lastName: event.target.value });
	};
	handlePhoneNumberChange = (event) => {
		this.setState({ phoneNumber: event.target.value });
	};
	handleBiographyChange = (event) => {
		this.setState({ biography: event.target.value });
	};
	handlePrivateChange(event) {

		this.setState({ private: event.target.value });
	}
	handleWebSiteChange = (event) => {
		this.setState({ webSite: event.target.value });
	};
	handleDateOfBirthChange = (event) => {
		this.setState({ dateOfBirth: event.target.value });
	};
	handleGenderChange(event) {

		this.setState({ gender: event.target.value });
	}
	validateForm = (userDTO) => {
		this.setState({
			emailError: "none",
			emailNotValid: "none",
			nameError: "none",
			dateError: "none",
			surnameError: "none",
			phoneError: "none",
			usernameError: "none"
		});

		if (this.state.email === "") {
			this.setState({ emailError: "initial" });
			return false;
		} else if (!this.state.email.includes("@")) {
			this.setState({ emailNotValid: "initial" });
			return false;
		}
		else if (this.state.name === "") {
			this.setState({ nameError: "initial" });
			return false;
		} else if (this.state.date === "") {
			this.setState({ dateError: "initial" });
			return false;

		} else if (this.state.surname === "") {
			this.setState({ surnameError: "initial" });
			return false;
		} else if (this.state.phoneNumber === "") {
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
							Id : localStorage.getItem("userId"),
            				Username :this.state.username,
							Name: this.state.name,
							LastName : this.state.lastName,
							Email : this.state.email,
							PhoneNumber : this.state.phoneNumber,
							Gender : this.state.gender,
							DateOfBirth  : this.state.dateOfBirth,
							WebSite : this.state.webSite,
							Biography : this.state.biography,
							Private : this.state.private
        };

        if (this.validateForm(userDTO)) {
			Axios.post(`${BASE_URL_USER}/api/user/update/`, userDTO)
			.then((res) => {
                        if (res.status === 400) {
                            this.setState({ hiddenFailAlert: false, failHeader: "Bad request", failMessage: "Invalid argument." });
                        } else if (res.status === 500) {
                            this.setState({ hiddenFailAlert: false, failHeader: "Internal server error", failMessage: "Server error." });
                        } else if (res.status === 200) {
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
	handlePrivateChange(event) {

		this.setState({ private: true });
	}


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
					<div className="row section-design" style={{ marginLeft: "2%",marginRight:"2%"}}>
							
                            <div className="col-md-2 padding-0">
                            <SidebarSettings/></div>
                            <div className="col-md-8 padding-0">
							

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
											id="username"
											type="text"
											onChange={this.handleUsernameChange}
											value={this.state.username}
										/>
									</div>
								</div>
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Name:</label>
										<input
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
										<label>Lastname:</label>
										<input
											type="text"
											onChange={this.handleSurnameChange}
											value={this.state.lastName}
										/>
									</div>
									<div className="text-danger" style={{ display: this.state.surnameError }}>
										Lastname must be entered.
									</div>
								</div>
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Email address:</label>
										<input
											readOnly
											className="form-control-plaintext"
											id="email"
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
											type="date"
											onChange={this.handleDateOfBirthChange}
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
											type="text"
											onChange={this.handleWebSiteChange}
											value={this.state.webSite}
										/>
									</div>
									
								</div>
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Biography:</label>
										<textarea
											onChange={this.handleBiographyChange}
											value={this.state.biography}
										/>
									</div>
									
								</div>
								<div style={{ color: "#6c757d", opacity: 1 }}>
									<p><input type="radio" value="Male" name="gender" onChange={(e) => this.handleGenderChange(e)} /> Male</p>
									<p><input type="radio" value="Female" name="gender" onChange={(e) => this.handleGenderChange(e)} /> Female</p>
									<p><input type="radio" value="Other" name="gender" onChange={(e) => this.handleGenderChange(e)} /> Other </p>
								</div>
								
								<br />
								<div>
								<label>Private:</label>
									<br/>
									<label class="switch">
									<input type="checkbox" value={this.state.private} onChange={(e) => this.handlePrivateChange(e)}/>
										<span class="slider round"></span>
									
								</label>
								</div>

								<div className="form-group">
									<div className="form-group controls mb-0 pb-2">
										<div className="form-row justify-content-center">
											<div className="form-col" hidden={!this.state.hiddenEditInfo}>
												<button
													style={{ background: "#1977cc", marginTop: "15px" }}
													onClick={this.handleChangeInfo}
													className="btn btn-primary btn-xl"
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
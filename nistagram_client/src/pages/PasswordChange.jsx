import React, { Component } from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { BASE_URL } from "../constants.js";
import Axios from "axios";
import { Button, Modal } from "react-bootstrap";
import { Redirect } from "react-router-dom";
import HeadingAlert from "../components/HeadingAlert";
import SidebarSettings from "../components/SidebarSettings"


class PasswordChange extends Component {
	state = {
		errorHeader: "",
		errorMessage: "",
		hiddenErrorAlert: true,
		nameError: "none",
		surnameError: "none",
		phoneError: "none",
		emailNotValid: "none",
		openModal: false,
		coords: [],
	};

	state = {
		oldPassword: "",
		newPassword: "",
		newPasswordRetype: "",
	};

	handleOldPasswordChange = (event) => {
		this.setState({ oldPassword: event.target.value });
	};

	handleNewPasswordChange = (event) => {
		this.setState({ newPassword: event.target.value });
	};

	handleNewPasswordRetypeChange = (event) => {
		this.setState({ newPasswordRetype: event.target.value });
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
				// headers: { Authorization: getAuthHeader() },
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

	render() {
		//if (this.state.redirect) return <Redirect push to="/login" />;

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

							<div className="col shadow p-3 bg-white rounded">
							<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
											<label>Old password:</label>
											<input
												placeholder="Old password"
												className="form-control"
												type="password"
												onChange={this.handleOldPasswordChange}
												value={this.state.oldPassword}
											/>
							</div>
							<div className="text-danger" style={{ display: this.props.oldPasswordEmptyError }}>
								Old password must be entered.
							</div>
							<div className="text-danger" style={{ display: "none" }}>
								Old password is not valid.
							</div>
					
					<div className="control-group">
						<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
							<label>New password:</label>
							<input
								placeholder="New password"
								className="form-control"
								type="password"
								onChange={this.handleNewPasswordChange}
								value={this.state.newPassword}
							/>
						</div>
						<div className="text-danger" style={{ display: this.props.newPasswordEmptyError }}>
							New password must be entered.
						</div>
					</div>
					<div className="control-group">
						<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
							<label>Type again new password:</label>
							<input
								placeholder="Type again new password"
								className="form-control"
								type="password"
								onChange={this.handleNewPasswordRetypeChange}
								value={this.state.newPasswordRetype}
							/>
						</div>
						<div className="text-danger" style={{ display: this.props.newPasswordRetypeEmptyError }}>
							You need to enter again new password.
						</div>
						<div className="text-danger" style={{ display: this.props.newPasswordRetypeNotSameError }}>
							Passwords are not the same.
						</div>
					</div>
					<div className="form-group text-center">
						<button
							style={{ background: "#1977cc", marginTop: "15px" }}
							onClick={() => this.props.changePassword(this.state.oldPassword, this.state.newPassword, this.state.newPasswordRetype)}
							className="btn btn-primary btn-xl"
							id="sendMessageButton"
							type="button"
						>
							Change password
						</button>
					</div>
					</div>



</div>







                            </div>

                            </div>
							
                            
				
			</React.Fragment>
		);
	}
}

export default PasswordChange;

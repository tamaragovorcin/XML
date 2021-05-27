import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import HeadingAlert from "./HeadingAlert";

class PasswordChangeModal extends Component {
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

	render() {
		return (
			<Modal
				show={this.props.show}
				size="lg"
				dialogClassName="modal-80w-80h"
				aria-labelledby="contained-modal-title-vcenter"
				centered
				onHide={this.props.onCloseModal}
			>
				<Modal.Header closeButton>
					<Modal.Title id="contained-modal-title-vcenter">{this.props.header}</Modal.Title>
				</Modal.Header>
				<Modal.Body>
					<HeadingAlert
						hidden={this.props.hiddenPasswordErrorAlert}
						header={this.props.errorPasswordHeader}
						message={this.props.errorPasswordMessage}
						handleCloseAlert={this.props.handleCloseAlertPassword}
					/>
					<div className="control-group">
						<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
							<label>Old password:</label>
							<input
								placeholder="Old password"
								class="form-control"
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
					</div>
					<div className="control-group">
						<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
							<label>New password:</label>
							<input
								placeholder="New password"
								class="form-control"
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
								class="form-control"
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
				</Modal.Body>
				<Modal.Footer>
					<Button onClick={this.props.onCloseModal}>Close</Button>
				</Modal.Footer>
			</Modal>
		);
	}
}

export default PasswordChangeModal;

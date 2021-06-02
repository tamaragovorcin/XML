

import React, { Component } from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import Axios from "axios";
import { BASE_URL, BASE_URL_USER } from "../constants.js";
import { Button } from "react-bootstrap";
import { Redirect } from "react-router-dom";
import HeadingAlert from "../components/HeadingAlert";


class FollowerProfilePage extends Component {
	state = {
        userId: "",

	};

    fetchData = (id) => {
        this.setState({
            userId: id,
        });
    };



	


render() {
	//if (this.state.redirect) return <Redirect push to="/" />;
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
					Login
					</h5>

				<div className="row section-design">
					<div className="col-lg-8 mx-auto">
						<br />
						<form id="contactForm" name="sentMessage" noValidate="novalidate">
							<div className="control-group">
								<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
									<input
										placeholder="Email address"
										className="form-control"
										id="name"
										type="text"
										onChange={this.handleEmailChange}
										value={this.state.email}
									/>
								</div>
								<div className="text-danger" style={{ display: this.state.emailError }}>
									Email must be entered.
									</div>
							</div>

			

				
						</form>
					</div>
				</div>
			</div>
			
		</React.Fragment>
	);
	}
}

export default FollowerProfilePage;

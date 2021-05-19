import React, { Component } from "react";
import { Link } from "react-router-dom";
import Header from "../components/Header";
import TopBar from "../components/TopBar";

class Unauthorized extends Component {
	render() {
		return (
			<React.Fragment>
				<TopBar />
				<Header />
				<div className="container justify-content-center" style={{ marginTop: "10%" }}>
					<div className="text-center mt-5" style={{ fontSize: "6em", color: "#1977cc" }}>
						<b>401</b>
					</div>
					<div className="text-center mt-5" style={{ fontSize: "3em" }}>
						Unauthorized
					</div>
					<hr style={{ width: "20%", color: "#1977cc", backgroundColor: "#1977cc" }} className="mt-5" />
					<div className="form-row justify-content-center mt-5">
						<div className="form-col">
							<Link to="/login" className="btn btn-outline-primary btn-lg">
								Log in again
							</Link>
						</div>
					</div>
				</div>
			</React.Fragment>
		);
	}
}

export default Unauthorized;

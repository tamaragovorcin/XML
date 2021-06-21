import React from "react";
import { Link } from "react-router-dom";

class TopBar extends React.Component {
	hasRole = (reqRole) => {
		let roles = JSON.parse(localStorage.getItem("keyRole"));
		if (roles === null) return false;

		if (reqRole === "*") return true;

		for (let role of roles) {
			if (role === reqRole) return true;
		}
		return false;
	};
	handleLogout = () => {
		localStorage.removeItem("keyToken");
		localStorage.removeItem("keyRole");
		localStorage.removeItem("expireTime");

		window.location.href = "/login"
	};
	render() {
		return (
			<div id="topbar" className="d-none d-lg-flex align-items-center fixed-top">
			<div className="container d-flex">
				<div className="contact-info mr-auto">
					
				</div>

				<div className="register-login">
				<Link to="/registration" hidden={this.hasRole("*")}>
						Register
					</Link>
					<Link to="/login" hidden={this.hasRole("*")}>
						Login
					</Link>
					<Link onClick={this.handleLogout} to="/" hidden={!this.hasRole("*")}>
						Logout
					</Link>
				</div>
			</div>
		</div>
		);
	}
}

export default TopBar;
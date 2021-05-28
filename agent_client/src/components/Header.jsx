import React from "react";
import { Link } from "react-router-dom";

class Header extends React.Component {
	hasRole = (reqRole) => {
		let roles = JSON.parse(localStorage.getItem("keyRole"));

		if (roles === null) return false;

		if (reqRole === "*") return true;

		for (let role of roles) {
			if (role === reqRole) return true;
		}
		return false;
	};
	render() {
	
		return (
			<header id="header" className="fixed-top">
				<div className="container d-flex align-items-center">
					<h1 className="logo mr-auto">
						<Link to="/">NISTAGRAM</Link>
					</h1>

					<nav className="nav-menu d-none d-lg-block">
						<ul>
							<li className="active" >
								<Link to="/">Home</Link>
							</li>
						</ul>
					</nav>
				</div>
			</header>
			);

	}
}

export default Header;
import React from "react";
import { Link } from "react-router-dom";
import { CgProfile } from 'react-icons/cg';
import { BiBookmark } from 'react-icons/bi';

import { FiSettings, FiSend } from 'react-icons/fi';
import { VscHome } from 'react-icons/vsc';
import { AiOutlineHeart } from 'react-icons/ai';

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
							<li  >
								<Link to="/"><VscHome/></Link>
							</li>
							<li  >
								<Link to="/messages"><FiSend/></Link>
							</li>
							<li  >
								<Link to="/follows"><AiOutlineHeart/></Link>
							</li>

							<li className="drop-down">
								<a href="#"><CgProfile /></a>
								<ul>

									<li className="drop-down">
										<Link to="/profilePage" ><CgProfile /> Profile</Link>

									</li>
									<li>


										<Link to="/saved"><BiBookmark /> Saved </Link>



									</li>
									<li>


										<Link to="/settings"><FiSettings /> Settings </Link>


									</li>
									<li>


										<Link to="/login"> Log out </Link>


									</li>

								</ul>
							</li>
						</ul>
					</nav>
				</div>
			</header>
		);

	}
}

export default Header;
import React from "react";
import { Link } from "react-router-dom";
import { CgProfile } from 'react-icons/cg';
import { BiBookmark, BiSearch } from 'react-icons/bi';
import Select from 'react-select';

import { FiSettings, FiSend } from 'react-icons/fi';
import { VscHome } from 'react-icons/vsc';
import { AiOutlineHeart } from 'react-icons/ai';

class Header extends React.Component {
	state = {
		options:["mladenka","vojna"]
	}
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
					<label className="logo mr-auto" style={{fontFamily:"Trattatello, fantasy"}}>
						<Link to="/">Ništagram</Link>
					</label>
					<div class="input-group rounded" style={{marginLeft:"10%",marginRight:"10%"}}>
					<input type="search" class="form-control rounded" placeholder="Search" aria-label="Search"
						aria-describedby="search-addon" />
					<span class="input-group-text border-0" id="search-addon">
					<BiSearch/>
					</span>
					</div>

					<nav className="nav-menu d-none d-lg-block">
						<ul>
							<li>
								<Link to="" ><VscHome/></Link>
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

									<li >
										<Link to="/profilePage" ><CgProfile /> Profile</Link>

									</li>
									<li>


										<Link to="/favorites"><BiBookmark /> Saved </Link>



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
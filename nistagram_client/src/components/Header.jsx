import React from "react";
import { Link } from "react-router-dom";
import { CgProfile } from 'react-icons/cg';
import { BiBookmark } from 'react-icons/bi';

import { FiSettings, FiSend } from 'react-icons/fi';
import { VscHome } from 'react-icons/vsc';
import { AiOutlineHeart } from 'react-icons/ai';
import { Search } from "@material-ui/icons";



class Header extends React.Component {
	state = {
		username : "", 
		tags : "",
		location : "",
		usernameError: "none",
		tagsError: "none",
		locationError: "none",
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
	handlePeoples = (event) => {
		this.setState({ username: event.target.value });
	};
	handleTags = (event) => {
		this.setState({ tags: event.target.value });
	};
	handleLocation = (event) => {
		this.setState({ location: event.target.value });
	};
	validateForm = () => {
		if (this.state.username === "") {
			this.setState({ usernameError: "inline" });
			return false;
		} else if (this.state.tags === "") {
			this.setState({ tagsError: "inline" });
			return false;
		} else if (this.state.location === "") {
			this.setState({ locationError: "inline" });
			return false;
		}
	
		return true;
	};

	render() {

		return (
			<header id="header" className="fixed-top">
				<div className="container d-flex align-items-center">
					<h1 className="logo mr-auto">
						<Link to="/">NISTAGRAM</Link>
						<div class = "row justify-content-md-center">
						<div class="col-30">
						<input type="text" value={this.state.username} style={{ width : "80%" }} class="form-control" placeholder="Search people"></input>
                        <div className="text-danger" style={{ display: this.state.usernameError }}>
									Username must be entered.
									</div>
						<Link onClick={this.handlePeoples}><Search/></Link>
						</div>
						
						 
                                        
						<div class="col-30">
						<input type="text" value={this.state.tags} style={{ width : "80%" }} class="form-control" placeholder="Search tags"></input>
                        <Link onClick={this.handleTags}><Search/></Link>
						</div>
						
						 
						<div class="col-30">
						<input type="text" value={this.state.location} style={{ width : "80%" }} class="form-control" placeholder="Search location"></input>
						<Link onClick={this.handleLocation}><Search/></Link>
					    </div>
						
						</div>
						
					</h1>

					<nav className="nav-menu d-none d-lg-block">
						<ul>
							<li>
								<Link ><VscHome/></Link>
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
import React from "react";
import { Link } from "react-router-dom";
import { CgProfile } from 'react-icons/cg';
import { BiBookmark, BiSearch } from 'react-icons/bi';
import Axios from "axios";
import { FiSettings, FiSend } from 'react-icons/fi';
import { VscHome } from 'react-icons/vsc';
import { AiOutlineHeart } from 'react-icons/ai';
import { BASE_URL_USER } from "../constants.js";
import SelectSearch from 'react-select-search';
import Select from 'react-select';

import { Redirect } from "react-router-dom";
class Header extends React.Component {
	state = {
		options: ["mladenka", "vojna"],
		search: "",
		users: [],
		redirect: false,
		options: [],
		optionDTO: { value: "", label: "" }
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

	handleSearchChange = (event) => {
		this.setState({ search: event.target.value });
	};

	handleChange = (event) => {
		alert(event.value)
		this.setState({ redirect: true });
	};


	componentDidMount() {
		let help = []
		Axios.get(BASE_URL_USER + "/api/")
			.then((res) => {

				console.log(res.data)
				this.setState({ users: res.data });

				res.data.forEach((user) => {
					let optionDTO = { label: user.ProfileInformation.Username, value: user.Id }
					help.push(optionDTO)
				});

				this.setState({ options: help });
				console.log(help)
			})
			.catch((err) => {

				console.log(err)
			});



	};

	render() {
		if (this.state.redirect) return <Redirect push to="/login" />;
		return (
			<header id="header" className="fixed-top">
				<div className="container d-flex align-items-center">
					<label className="logo mr-auto" style={{ fontFamily: "Trattatello, fantasy" }}>
						<Link to="/">Ništagram</Link>
					</label>
					<div class="input-group rounded" style={{ marginLeft: "20%", marginRight: "10%" }}>

						<div style={{ width: '300px' }}>
							<Select
								style={{ width: `$500px` }}
								className="select-custom-class"
								label="Single select"
								options={this.state.options}
								onChange ={e => this.handleChange(e)}
							/>


						</div>

					</div>

						
					




					<nav className="nav-menu d-none d-lg-block">
						<ul>
							<li>
								<Link to="" ><VscHome /></Link>
							</li>
							<li  >
								<Link to="/messages"><FiSend /></Link>
							</li>
							<li  >
								<Link to="/follows"><AiOutlineHeart /></Link>
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
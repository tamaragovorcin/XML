import React from "react";
import { Link } from "react-router-dom";
import { CgProfile } from 'react-icons/cg';
import { BiBookmark, BiSearch } from 'react-icons/bi';
import Select from 'react-select';
import Axios from "axios";
import { FiSettings, FiSend } from 'react-icons/fi';
import { VscHome } from 'react-icons/vsc';
import { AiOutlineHeart } from 'react-icons/ai';
import { BASE_URL_USER } from "../constants.js";
import SelectSearch from 'react-select-search';
class Header extends React.Component {
	state = {
		options: ["mladenka", "vojna"],
		search: "",
		users: [],

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

	handleChange = () => {
		alert("skfjhksfu")
	};


	componentDidMount() {
		Axios.get(BASE_URL_USER + "/api/")
			.then((res) => {

				console.log(res.data)
				this.setState({ users: res.data });

			})
			.catch((err) => {

				console.log(err)
			});



	};

	render() {

		return (
			<header id="header" className="fixed-top">
				<div className="container d-flex align-items-center">
					<label className="logo mr-auto" style={{ fontFamily: "Trattatello, fantasy" }}>
						<Link to="/">Ništagram</Link>
					</label>
					<div class="input-group rounded" style={{ marginLeft: "10%", marginRight: "10%" }}>




						<input type="search" class="form-control rounded" placeholder="Search" aria-label="Search" onChange={this.handleSearchChange}
							aria-describedby="search-addon" />


					</div>
					<div>

						<script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.1/jquery.min.js"></script>
						<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js"></script>
						<link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" rel="stylesheet" />
						<script src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-select/1.10.0/js/bootstrap-select.min.js"></script>
						<link href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-select/1.10.0/css/bootstrap-select.min.css" rel="stylesheet" />

					

						<select class="selectpicker" data-live-search="true"  >
							{this.state.users.map((user) => (
								<option data-tokens="luna" onChange={this.handleChange} key={user.Id} value={user.ProfileInformation.Username} >{user.ProfileInformation.Username}</option>

							))}
						</select>

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
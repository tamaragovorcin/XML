import React from "react";
import { Link } from "react-router-dom";
import { CgProfile } from 'react-icons/cg';
import Axios from "axios";
import { FiSettings, FiSend } from 'react-icons/fi';
import { VscHome } from 'react-icons/vsc';
import { FaRegQuestionCircle, FaSearch } from 'react-icons/fa';
import {MdReportProblem} from 'react-icons/md';
import { GiThreeFriends } from 'react-icons/gi';
import { AiOutlineHeart,AiFillLike,AiFillDislike } from 'react-icons/ai';
import { BASE_URL } from "../constants.js";
import Select from 'react-select';

class Header extends React.Component {
	state = {
		search: "",
		users: [],
		options: [],
		optionDTO: { value: "", label: "" },
		userId: "",
	}

	hasRole = (reqRole) => {
		let roles = JSON.parse(localStorage.getItem("keyRole"));
		if (roles === null) return false;

		if (reqRole === "*") return true;

	
		if (roles.trim() === reqRole.trim()) 
		{
			return true;
		}
		return false;
	};

	handleSearchChange = (event) => {
		this.setState({ search: event.target.value });
	};

	handleChange = (event) => {
		this.setState({ userId: event.value });
		window.location = "#/followerProfilePage/" + event.value;

	};


	componentDidMount() {
		if(this.hasRole("*")) {
			let help = []
			let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

			Axios.get(BASE_URL + "/api/users/api/all/"+id)
				.then((res) => {
	
					console.log(res.data)
					this.setState({ users: res.data });
	
					res.data.forEach((user) => {
						let optionDTO = { id: user.ID, label: user.ProfileInformation.Username, value: user.Id }
						help.push(optionDTO)
					});
	
					this.setState({ options: help });
					console.log(help)
				})
				.catch((err) => {
	
					console.log(err)
				});
		}
		else {
			let help = []
			Axios.get(BASE_URL + "/api/users/api/")
				.then((res) => {
	
					console.log(res.data)
					this.setState({ users: res.data });
	
					res.data.forEach((user) => {
						let optionDTO = { id: user.ID, label: user.ProfileInformation.Username, value: user.Id }
						help.push(optionDTO)
					});
	
					this.setState({ options: help });
					console.log(help)
				})
				.catch((err) => {
	
					console.log(err)
				});
		}
		

	};

	render() {


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
								onChange={e => this.handleChange(e)}
							/>


						</div>

					</div>
					<nav className="nav-menu d-none d-lg-block">
						<ul>
							<li>
								<Link to="" ><VscHome /></Link>
							</li>
							<li  >
								<Link to="/seacrh"><FaSearch /></Link>
							</li>
							<li  hidden={!this.hasRole("USER")}>
								<Link to="/messages"><FiSend /></Link>
							</li>
							<li  hidden={!this.hasRole("USER")}>
								<Link to="/followRequest"><AiOutlineHeart /></Link>
							</li>
							<li  hidden={!this.hasRole("USER")}>
								<Link to="/verifyRequest"><FaRegQuestionCircle /></Link>
							</li>
							<li  hidden={!this.hasRole("ADMIN")}>
								<Link to="/reportedPosts"><MdReportProblem /></Link>
							</li>
							<li className="drop-down" hidden={!this.hasRole("USER")}>
								<a href="#"><CgProfile /></a>
								<ul>

									<li >
										<Link to="/profilePage" ><CgProfile /> Profile</Link>
									</li>
									
									<li>
										<Link to="/editProfile"><FiSettings /> Settings </Link>
									</li>

									<li>
										<Link to="/closeFriends"><GiThreeFriends /> Close friends </Link>
									</li>
									<li>
										<Link to="/likedPosts"><AiFillLike /> Liked posts </Link>
									</li>
									<li>
										<Link to="/dislikedPosts"><AiFillDislike /> Disliked posts </Link>
									</li>
									<li>
										<Link to="/login"> Log out </Link>
									</li>

								</ul>
							</li>
							
							<li className="drop-down" hidden={!this.hasRole("ADMIN")}>
								<a href="#"><CgProfile /></a>
								<ul>

									<li >
										<Link to="/profilePage" ><CgProfile /> Profile</Link>

									</li>
									
									<li>
										<Link to="/editProfile"><FiSettings /> Settings </Link>
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
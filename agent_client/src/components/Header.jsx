import React from "react";
import { Link } from "react-router-dom";
import { CgProfile } from 'react-icons/cg';
import { BiBookmark } from 'react-icons/bi';
import Axios from "axios";
import { FiTruck } from 'react-icons/fi';
import { VscHome } from 'react-icons/vsc';
import { FaSearch } from 'react-icons/fa';
import AddPostModal from "../components/Posts/AddPostModal";
import { RiAddCircleLine } from 'react-icons/ri';
import { AiOutlineOrderedList, AiOutlineShoppingCart } from 'react-icons/ai';
import { BASE_URL_AGENT, BASE_URL_USER } from "../constants.js";
import { BASE_URL } from "../constants.js";
import Select from 'react-select';
import ModalDialog from "../components/ModalDialog";
class Header extends React.Component {

	state = {
		search: "",
		users: [],
		options: [],
		optionDTO: { value: "", label: "" },
		userId: "",
		showImageModal : false,
		pictures: [],
		openModal: false,
		hiddenOne: true,
		hiddenMultiple: true,

	}


	hasRole = (reqRole) => {
		
		
		let roles =  ""
		if (localStorage === null) return false;

		roles = localStorage.getItem("keyRole").substring(3, localStorage.getItem('keyRole').length-3)
		
		if (roles === null) return false;

		if (reqRole === "*") return true;

	
		if (roles.trim() === reqRole) 
		{
			
			return true;
		}
		return false;
	};
	handlePostModalClose = () => {
		this.setState({ showImageModal: false });
	};
	handlePostModalOpen = () => {
		this.setState({ showImageModal: true });
		this.setState({pictures: []})
	};
	handleSearchChange = (event) => {
		this.setState({ search: event.target.value });
	};

	handleChange = (event) => {
		this.setState({ userId: event.value });
		window.location = "#/followerProfilePage/" + event.value;

	};


	

	render() {


		return (
			<React.Fragment>
			<header id="header" className="fixed-top">
				<div className="container d-flex align-items-center">
					<label className="logo mr-auto" style={{ fontFamily: "Trattatello, fantasy" }}>
						<Link to="/">Agent app</Link>
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
						<li  hidden={!this.hasRole("*")}>
								<Link to="/"><VscHome /></Link>
							</li>
							<li  hidden={!this.hasRole("Agent")}>
							<button  onClick={this.handlePostModalOpen} className="btn btn-outline-secondary btn-sm" style={{  border: "none", marginBottom: "1rem" }}><RiAddCircleLine /></button>
							</li>
							<li hidden={!this.hasRole("Agent")}>
							<Link to="/allAgents"><AiOutlineOrderedList /></Link>
							</li>
							<li>
							<Link to="/orders"><AiOutlineShoppingCart /></Link>
							</li>
							<li>
							<Link to="/ordered"><FiTruck /></Link>
							</li>
						
							
			

							
						</ul>
					</nav>
				</div>
			</header>










			<ModalDialog
						show={this.state.openModal}
						onCloseModal={this.handleModalClose}
						header="Successful"
						text={this.state.textSuccessfulModal}
						openModal = {this.state.openModal}
					/>


			<AddPostModal
						show={this.state.showImageModal}
						onCloseModal={this.handlePostModalClose}
						header="Add new product"
						hiddenMultiple = {this.state.hiddenMultiple}
						hiddenOne = {this.state.hiddenOne}
						noPicture = {this.state.noPicture}
						onDrop = {this.onDrop}

						handleAddFeedPost = {this.handleAddFeedPost}
						handleAddStoryPost = {this.handleAddStoryPost}
						handleAddStoryPostCloseFriends = {this.handleAddStoryPostCloseFriends}
						handleAddFeedPostAlbum = {this.handleAddFeedPostAlbum}
						handleAddStoryPostAlbum= {this.handleAddStoryPostAlbum}
						handleAddStoryPostAlbumCloseFriends = {this.handleAddStoryPostAlbumCloseFriends}
						addressNotFoundError = {this.state.addressNotFoundError}
						handleDescriptionChange = {this.handleDescriptionChange}
						handleHashtagsChange = {this.handleHashtagsChange}
						handleAddTagsModal = {this.handleAddTagsModal}


					/>


			</React.Fragment>

);


	}
}

export default Header;
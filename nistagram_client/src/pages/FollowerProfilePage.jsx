import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { Link } from "react-router-dom";
import playerLogo from "../static/coach.png";

import { BASE_URL_FEED } from "../constants.js";
import { BASE_URL_USER_INTERACTION } from "../constants.js";
import LikesModal from "../components/Posts/LikesModal"
import DislikesModal from "../components/Posts/DislikesModal"
import CommentsModal from "../components/Posts/CommentsModal"
import Axios from "axios";
import ModalDialog from "../components/ModalDialog";
import AddPostModal from "../components/Posts/AddPostModal";
import WriteCommentModal from "../components/Posts/WriteCommentModal"

import { BASE_URL_USER } from "../constants.js";
import { FiHeart } from "react-icons/fi";

import { FaHeartBroken, FaRegCommentDots } from "react-icons/fa"
import { BsBookmark } from "react-icons/bs"
import { Lock } from "@material-ui/icons";
import { Icon } from "@material-ui/core";
class FollowerProfilePage extends React.Component {
	constructor(props) {
		super(props);

		this.onDrop = this.onDrop.bind(this);
		this.addressInput = React.createRef();

	}
	state = {
		following: true,
		userId: "",
		id: "",
		username: "",
		numberPosts: 0,
		numberFollowing: 0,
		numberFollowers: 0,
		biography: "",
		highlihts: [],
		photos: [],
		pictures: [],
		picture: "",
		hiddenOne: true,
		hiddenMultiple: true,
		noPicture: true,
		peopleLikes: [],
		peopleDislikes: [],
		comments: [],
		coords: [],
		addressNotFoundError: "none",
		textSuccessfulModal: "",
		showLikesModal: false,
		showDislikesModal: false,
		showCommentsModal: false,
		showImageModal: false,
		openModal: false,
		addressLocation: null,
		foundLocation: true,
		description: "",
		hashtags: "",
		showWriteCommentModal: false
	}
	onYmapsLoad = (ymaps) => {
		this.ymaps = ymaps;
		new this.ymaps.SuggestView(this.addressInput.current, {
			provider: {
				suggest: (request, options) => this.ymaps.suggest(request),
			},
		});
	};

	fetchData = (id) => {
		this.setState({
			userId: id,
		});
	};

	onDrop(picture) {
		this.setState({
			pictures: this.state.pictures.concat(picture),
		});

		let pomoc = this.state.pictures.length;
		pomoc = pomoc + 1;
		if (pomoc === 0) {
			this.setState({
				noPicture: true,
			});
			this.setState({
				showImageModal: false,
			});
		}
		else {
			this.setState({
				noPicture: false,
			});
			this.setState({
				showImageModal: true,
			});
		}
		if (pomoc === 1) {
			this.setState({
				hiddenOne: false,
			});
			this.setState({
				hiddenMultiple: true,
			});
		}
		else if (pomoc >= 2) {
			this.setState({
				hiddenOne: true,
			});
			this.setState({
				hiddenMultiple: false,
			});
		}


	}



	test(pic, userId, feedId) {

		this.setState({
			fileUploadOngoing: true
		});

		const fileInput = document.querySelector("#fileInput");
		const formData = new FormData();

		formData.append("file", pic);
		formData.append("test", "StringValueTest");

		const options = {
			method: "POST",
			body: formData

		};
		fetch(BASE_URL_FEED + "/api/image/" + userId + "/" + feedId, options);
	}


	componentDidMount() {
		var sentence = window.location.toString()

		var s = []
		s = sentence.split("/");
		console.log(window.location.toString())


		this.fetchData(s[5]);
		let id = localStorage.getItem("userId")
		alert(s[5])
		Axios.get(BASE_URL_USER + "/api/" + s[5])
			.then((res) => {
				if (res.status === 401) {
					this.setState({ errorHeader: "Bad credentials!", errorMessage: "Wrong username or password.", hiddenErrorAlert: false });
				} else if (res.status === 500) {
					this.setState({ errorHeader: "Internal server error!", errorMessage: "Server error.", hiddenErrorAlert: false });
				} else {
					this.setState({ numberPosts: 10 });
					this.setState({ numberFollowing: 600 });
					this.setState({ numberFollowers: 750 });
					this.setState({ biography: res.data.Biography });
					this.setState({ username: res.data.ProfileInformation.Username });
					console.log(res.data)
				}
			})
			.catch((err) => {
				console.log(err);
			});

		//this.handleGetBasicInfo()
		this.handleGetHighlights()
		this.handleGetPhotos()

	}
	handleGetBasicInfo = () => {

		this.setState({ numberPosts: 10 });
		this.setState({ numberFollowing: 600 });
		this.setState({ numberFollowers: 750 });
		this.setState({ biography: "bla bla bla" });
		this.setState({ username: "USERNAME" });
	}

	handleGetHighlights = () => {
		let highliht1 = { id: 1, name: "ITALY" };
		let highliht2 = { id: 2, name: "AMERICA" };
		let highliht3 = { id: 3, name: "SERBIA" };

		let list = [];
		list.push(highliht1)
		list.push(highliht2)
		list.push(highliht3)

		this.setState({ highlihts: list });
	}

	handleGetPhotos = () => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1);
		Axios.get(BASE_URL_FEED + "/api/feed/usersImages/" + id)
			.then((res) => {
				this.setState({ photos: res.data });
				console.log(res.data);
			})
			.catch((err) => {
				console.log(err);
			});


	}
	handleDescriptionChange = (event) => {
		this.setState({ description: event.target.value });
	};
	handleHashtagsChange = (event) => {
		this.setState({ hashtags: event.target.value });
	}
	handleModalClose = () => {
		this.setState({ openModal: false });
	};
	handlePostModalClose = () => {
		this.setState({ showImageModal: false });
	};
	handlePostModalOpen = () => {
		this.setState({ showImageModal: true });
	};
	handleLikesModalOpen = () => {
		this.setState({ showLikesModal: true });
	}
	handleDislikesModalOpen = () => {
		this.setState({ showDislikesModal: true });
	}
	handleCommentsModalOpen = () => {
		this.setState({ showCommentsModal: true });
	}
	handleLikesModalClose = () => {
		this.setState({ showLikesModal: false });
	}
	handleDislikesModalClose = () => {
		this.setState({ showDislikesModal: false });
	}
	handleCommentsModalClose = () => {
		this.setState({ showCommentsModal: false });
	}

	handleWriteCommentModalClose = () => {
		this.setState({ showWriteCommentModal: false });
	}
	handleWriteCommentModal = () => {
		this.setState({ showWriteCommentModal: true });
	}
	handleLike = () => {

	}
	handleDislike = () => {

	}
	handleSave = () => {

	}
	handleAddFeedPost = () => {

		if (this.state.addressInput === "") {
			const feedPostDTO = {
				tagged: [],
				description: this.state.description,
				hashtags: this.state.hashtags,
				location: this.state.addressLocation
			};
			this.sendRequestForFeed(feedPostDTO);

		}
		else {
			let street;
			let city;
			let country;
			let latitude;
			let longitude;
			this.ymaps
				.geocode(this.addressInput.current.value, {
					results: 1,
				})
				.then(function (res) {

					if (typeof res.geoObjects.get(0) === "undefined") { this.setState({ foundLocation: false }); }
					else {
						var firstGeoObject = res.geoObjects.get(0),
							coords = firstGeoObject.geometry.getCoordinates();
						latitude = coords[0];
						longitude = coords[1];
						country = firstGeoObject.getCountry();
						street = firstGeoObject.getThoroughfare();
						city = firstGeoObject.getLocalities().join(", ");
					}
				}).then((res) => {
					var locationDTO = {
						street: street,
						country: country,
						town: city,
						latitude: latitude,
						longitude: longitude
					}
					let feedPostDTO = {
						tagged: [],
						description: this.state.description,
						hashtags: this.state.hashtags,
						location: locationDTO
					};

					if (this.state.foundLocation === false) {
						this.setState({ addressNotFoundError: "initial" });
					} else {
						this.sendRequestForFeed(feedPostDTO);
					}

				});


		}


	}
	handleAddStoryPost = () => {

	}
	handleAddFeedPostAlbum = () => {

	}
	handleAddStoryPostAlbum = () => {

	}
	handleFollow = () => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
		const user1Id = {id: id}
		const user2Id = {id : this.state.userId}
		alert(this.state.userId);
		
		Axios.post(BASE_URL_USER_INTERACTION + "/api/createUser", user1Id)
		.then((res) => {
			
				console.log(res.data)
				
			
		})
		.catch ((err) => {
	console.log(err);
});
	Axios.post(BASE_URL_USER_INTERACTION + "/api/createUser", user2Id)
				.then((res) => {
					
						console.log(res.data)
						
					
				})
				.catch ((err) => {
			console.log(err);
		});
		const followReguestDTO = { follower: id, following : this.state.userId};
		Axios.post(BASE_URL_USER_INTERACTION + "/api/followRequest", followReguestDTO)
				.then((res) => {
					
						console.log(res.data)
						this.setState({ redirect: true });
					
				})
				.catch ((err) => {
			console.log(err);
		});

	}



	sendRequestForFeed(feedPostDTO) {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1);;
		console.log(id)

		Axios.post(BASE_URL_FEED + "/api/feed/" + id, feedPostDTO)
			.then((res) => {
				if (res.status === 409) {
					this.setState({
						errorHeader: "Resource conflict!",
						errorMessage: "Email already exist.",
						hiddenErrorAlert: false,
					});
				} else if (res.status === 500) {
					this.setState({ errorHeader: "Internal server error!", errorMessage: "Server error.", hiddenErrorAlert: false });
				} else {
					this.setState({ openModal: true });
					this.setState({ redirect: true });
				}
				let feedId = res.data;
				console.log(res.data);
				console.log(res.status);
				let userid = localStorage.getItem("userId");
				let pics = [];

				this.state.pictures.forEach((p) => {
					pics.push(p.name);
				});
				this.state.pictures.forEach((pic) => {
					this.test(pic, userid, feedId);
				});

				this.setState({ pictures: [] });
				this.setState({ showImageModal: false, });
				this.setState({ openModal: true });
				this.setState({ textSuccessfulModal: "You have successfully added feed post." });

			})
			.catch((err) => {
				console.log(err);
			});
	}

	render() {
		return (
			<React.Fragment>
				<TopBar />
				<Header />

				<section id="hero" className="d-flex align-items-top">
					<div className="container">
						<div className="d-flex align-items-top">
							<div className="container" style={{ marginTop: "10rem", marginRight: "10rem" }}>
								<table className="table" style={{ width: "100%" }}>
									<tbody>

										<tr>
											<td width="130em">
												<img
													className="img-fluid"
													src={playerLogo}
													width="70em"
													alt="description"
												/>
											</td>

											<td>
												<div>
													<td>
														<label >{this.state.username}</label>
													</td>
													<td>
														<Link to="/userChangeProfile" className="btn btn-outline-secondary btn-sm">Edit profile</Link>
													<Link onClick={this.handleFollow} className="btn btn-outline-success btn-sm">Follow</Link>

													</td>

												</div>
												<div>
													<td>
														<label ><b>{this.state.numberPosts}</b> posts</label>
													</td>
													<td>
														<label ><b>{this.state.numberFollowers}</b> followers</label>
													</td>
													<td>
														<label ><b>{this.state.numberFollowing}</b> following</label>
													</td>

												</div>
												<div>
													<td>
														<label >{this.state.biography}</label>
													</td>
												</div>


											</td>


										</tr>
									</tbody>
								</table>
							</div>
						</div>





















						<div hidden={this.state.following}>

							<div className="container-fluid testimonial-group d-flex align-items-top">
								<div className="container-fluid scrollable" style={{ marginRight: "10rem", marginBottom: "5rem", marginTop: "5rem" }}>
									<table className="table-responsive" style={{ width: "100%" }}>
										<tbody>

											<tr >
												{this.state.highlihts.map((high) => (
													<td id={high.id} key={high.id} style={{ width: "60em", marginLeft: "10em" }}>
														<tr width="100em">
															<img
																className="img-fluid"
																src={playerLogo}
																style={{ borderRadius: "50%", margin: "2%" }}
																width="60em"
																alt="description"
															/>
														</tr>
														<tr>
															<label style={{ marginRight: "15px" }}>{high.username}</label>
														</tr>
													</td>

												))}
											</tr>


										</tbody>
									</table>
								</div>
							</div>
							<div className="d-flex align-items-top">
								<div className="container-fluid">

									<table className="table">
										<tbody>
											{this.state.photos.map((post) => (

												<tr id={post.id} key={post.id}>

													<tr style={{ width: "100%" }}>
														<td colSpan="3">
															<img
																className="img-fluid"
																src={`data:image/jpg;base64,${post.Media}`}
																width="100%"
																alt="description"
															/>
														</td>
													</tr>
													<tr style={{ width: "100%" }}>
														<td>
															<button onClick={this.handleLike} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height: "40px", marginLeft: "6rem" }}><FiHeart /></button>
														</td>
														<td>
															<button onClick={this.handleDislike} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height: "40px", marginLeft: "6rem" }}><FaHeartBroken /></button>

														</td>
														<td>
															<button onClick={this.handleWriteCommentModal} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height: "40px", marginLeft: "6rem" }}><FaRegCommentDots /></button>
														</td>
														<td>
															<button onClick={this.handleSave} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height: "40px" }}><BsBookmark /></button>
														</td>
													</tr>
													<tr style={{ width: "100%" }}>
														<td>
															<button onClick={this.handleLikesModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", marginLeft: "4rem" }}><label>likes</label></button>
														</td>
														<td>
															<button onClick={this.handleDislikesModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", marginLeft: "4rem" }}><label > dislikes</label></button>
														</td>
														<td>
															<button onClick={this.handleCommentsModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", marginLeft: "4rem" }}><label >Comments</label></button>
														</td>
													</tr>
													<br />
													<br />
													<br />
												</tr>

											))}

										</tbody>
									</table>
								</div>
							</div>










						</div>

					</div>

					<div hidden={!this.state.following}>

													NE PRATITE SE
						<div className="d-flex align-items-top p-3 mb-2 d-flex justify-content-center">
							
							<label><b>This Account is Private</b></label>
							
						</div>

						<div className="d-flex justify-content-center h-100">
							<Icon className="d-flex justify-content-center h-100 w-100"><Lock /></Icon>
						</div>

					</div>

				</section>
				<div>

				

				</div>

			</React.Fragment >
		);
	}
}

export default FollowerProfilePage;
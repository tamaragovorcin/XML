
import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { Link,Button } from "react-router-dom";
import playerLogo from "../static/coach.png";

import { BASE_URL, BASE_URL_FEED } from "../constants.js";

import LikesModal from "../components/Posts/LikesModal"
import DislikesModal from "../components/Posts/DislikesModal"
import CommentsModal from "../components/Posts/CommentsModal"
import Axios from "axios";
import ModalDialog from "../components/ModalDialog";
import AddPostModal from "../components/Posts/AddPostModal";

import { BASE_URL_USER } from "../constants.js";

class ProfilePage extends React.Component {
	constructor(props) {
		super(props);

		this.onDrop = this.onDrop.bind(this);
		this.addressInput = React.createRef();

	}
	state = {
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
		noPicture : true,
		peopleLikes : [],
		peopleDislikes : [],
		comments : [],
		coords: [],
		addressNotFoundError: "none",
		textSuccessfulModal : "",
		showLikesModal : false,
		showDislikesModal : false,
		showCommentsModal : false,
		showImageModal : false,
		openModal : false,
		addressLocation :null,
		foundLocation : true,
		description : "",
		hashtags :"",
	}
	onYmapsLoad = (ymaps) => {
		this.ymaps = ymaps;
		new this.ymaps.SuggestView(this.addressInput.current, {
			provider: {
				suggest: (request, options) => this.ymaps.suggest(request),
			},
		});
	};
	onDrop(picture) {
		this.setState({
			pictures: this.state.pictures.concat(picture),
		});

		let pomoc = this.state.pictures.length;
		pomoc = pomoc + 1;
		if(pomoc===0) {
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
		if(pomoc === 1){
			this.setState({
				hiddenOne: false,
			});
			this.setState({
				hiddenMultiple: true,
			});
		}
		else if(pomoc >= 2){
			this.setState({
				hiddenOne: true,
			});
			this.setState({
				hiddenMultiple: false,
			});
		}


	}

	

	test(pic,userId, feedId) {

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
		fetch(BASE_URL_FEED + "/api/image/"+userId+"/"+feedId , options);
	}


	componentDidMount() {

		let id =localStorage.getItem("userId")

		Axios.get(BASE_URL_USER + "/api/" + id)
				.then((res) => {
					if (res.status === 401) {
						this.setState({ errorHeader: "Bad credentials!", errorMessage: "Wrong username or password.", hiddenErrorAlert: false });
					} else if (res.status === 500) {
						this.setState({ errorHeader: "Internal server error!", errorMessage: "Server error.", hiddenErrorAlert: false });
					} else {
						
						console.log(res.data)
					}
				})
				.catch ((err) => {
			console.log(err);
		});

		this.handleGetBasicInfo()
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
		let list = []
		let comments1 = []
		let comments2 = []
		let comment1 = { id: 1, user: "USER 1 ", text: "very nice" }
		let comment11 = { id: 2, user: "USER 2 ", text: "cool" }
		let comment111 = { id: 3, user: "USER 3 ", text: "vau" }
		comments1.push(comment1)
		comments1.push(comment11)
		comments1.push(comment111)

		let comment2 = { id: 4, user: "USER 55443 ", text: "i like it" }
		let comment22 = { id: 5, user: "USER 11111 ", text: "ugly" }
		let comment222 = { id: 6, user: "USER 33333 ", text: "awesome" }
		comments2.push(comment2)
		comments2.push(comment22)
		comments2.push(comment222)

		let photo1 = { id: 1, photo: playerLogo, numLikes: 52, numDislikes: 2, comments: comments1 }
		let photo2 = { id: 2, photo: playerLogo, numLikes: 45, numDislikes: 0, comments: comments2 }
		list.push(photo1)
		list.push(photo2)

		this.setState({ photos: list });

	}
	handleDescriptionChange = (event) => {
		this.setState({ description: event.target.value });
	};
	handleHashtagsChange = (event)=> {
		this.setState({hashtags : event.target.value });
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
	handleLikesModalOpen = ()=> {
		this.setState({ showLikesModal: true });    
	}
	handleDislikesModalOpen = ()=> {
		this.setState({ showDislikesModal: true });    
	}
	handleCommentsModalOpen = ()=> {
		this.setState({ showCommentsModal: true });    
	}
	handleLikesModalClose = ()=> {
		this.setState({ showLikesModal: false });    
	}
	handleDislikesModalClose = ()=> {
		this.setState({ showDislikesModal: false });    
	}
	handleCommentsModalClose = ()=> {
		this.setState({ showCommentsModal: false });    
	}
	
	handleAddFeedPost = ()=> {
		
		if (this.state.addressInput === "") {
			const feedPostDTO = {
				tagged: [],
				description: this.state.description,
				hashtags: this.state.hashtags,
				location : this.state.addressLocation
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
	
					if (typeof res.geoObjects.get(0) === "undefined")  {this.setState({ foundLocation:false});}
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
						street : street,
						country : country,
						town : city,
						latitude : latitude,
						longitude : longitude
					}
					let feedPostDTO = {
						tagged: [],
						description: this.state.description,
						hashtags: this.state.hashtags,
						location : locationDTO
					};
					
					if (this.state.foundLocation === false) {
							this.setState({ addressNotFoundError: "initial" });
					} else {
							this.sendRequestForFeed(feedPostDTO);
					}

				});
				

		}
		
		
	}
	handleAddStoryPost = ()=> {

	}
	handleAddFeedPostAlbum = ()=> {

	}
	handleAddStoryPostAlbum = ()=> {
		
	}


	

	sendRequestForFeed(feedPostDTO) {
		let id = localStorage.getItem("userId");

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

											</td>
											<td>
												<button onClick={this.handlePostModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" }}>Add new post/video</button>

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
				

				<div className="d-flex align-items-top">
					<div className="container" style={{ marginRight: "10rem" }}>
						<table className="table" style={{ width: "100%" }}>
							<tbody>

								<tr >
									{this.state.highlihts.map((high) => (
										<td id={high.id} key={high.id} width="30em">
											<tr width="100em">
												<img
													className="img-fluid"
													src={playerLogo}
													width="40em"
													alt="description"
												/>
											</tr>
											<tr>
												<label>{high.name}</label>
											</tr>
										</td>
									))}
								</tr>


							</tbody>
						</table>
					</div>
				</div>
				<div className="d-flex align-items-top">
					<div className="container" style={{ marginLeft: "30rem" }}>
						<table className="table" style={{ width: "100%" }}>
							<tbody>
								{this.state.photos.map((photo) => (
									<tr id={photo.id} key={photo.id}>

										<td width="200em">
											<img
												className="img-fluid"
												src={photo.photo}
												width="100em"
												alt="description"
											/>
										</td>

										<td>
											<tr >
												<button onClick={this.handleLikesModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" }}><label><b>{photo.numLikes}</b>likes</label></button>
											</tr>
											<tr>
												<button onClick={this.handleDislikesModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" }}><label ><b>{photo.numDislikes}</b> dislikes</label></button>
											</tr>
											<tr>
												<button onClick={this.handleCommentsModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" }}><label >Comments</label></button>
											</tr>

										</td>
									</tr>
								))}

							</tbody>
						</table>
					</div>
				</div>
				<div>
                        
                    <LikesModal
					        show={this.state.showLikesModal}
					        onCloseModal={this.handleLikesModalClose}
					        header="People who liked the photo"
							peopleLikes = {this.state.peopleLikes}
				    />
                    <DislikesModal
                         show={this.state.showDislikesModal}
						 onCloseModal={this.handleDislikesModalClose}
						 header="People who disliked the photo"
						 peopleDislikes = {this.state.peopleDislikes}
				    />
                    <CommentsModal
                        show={this.state.showCommentsModal}
						onCloseModal={this.handleCommentsModalClose}
						header="Comments on the photo"
						comments = {this.state.comments}
                    />
                    <ModalDialog
						show={this.state.openModal}
						onCloseModal={this.handleModalClose}
						header="Successful publishing"
						text={this.state.textSuccessfulModal}
					/>
					<AddPostModal
						show={this.state.showImageModal}
						onCloseModal={this.handlePostModalClose}
						header="New post/story"
						hiddenMultiple = {this.state.hiddenMultiple}
						hiddenOne = {this.state.hiddenOne}
						noPicture = {this.state.noPicture}
						onDrop = {this.onDrop}
						addressInput = {this.addressInput}
						onYmapsLoad = {this.onYmapsLoad}
						handleAddFeedPost = {this.handleAddFeedPost}
						handleAddStoryPost = {this.handleAddStoryPost}
						handleAddFeedPostAlbum = {this.handleAddFeedPostAlbum}
						handleAddStoryPostAlbum= {this.handleAddStoryPostAlbum}
						addressNotFoundError = {this.state.addressNotFoundError}
						handleDescriptionChange = {this.handleDescriptionChange}
						handleHashtagsChange = {this.handleHashtagsChange}
					/>
                    </div>

			</React.Fragment>
		);
	}
}

export default ProfilePage;

import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { Link,Button } from "react-router-dom";
import playerLogo from "../static/coach.png";

import { BASE_URL } from "../constants.js";
import ImageUploader from 'react-images-upload';
import LikesModal from "../components/Posts/LikesModal"
import DislikesModal from "../components/Posts/DislikesModal"
import CommentsModal from "../components/Posts/CommentsModal"
import Axios from "axios";
import { BASE_URL_USER } from "../constants.js";

class ProfilePage extends React.Component {
	constructor(props) {
		super(props);

		this.onDrop = this.onDrop.bind(this);

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
		peopleLikes : [],
		peopleDislikes : [],
		comments : [],
		showLikesModal : false,
		showDislikesModal : false,
		showCommentsModal : false
	}
	onDrop(picture) {
		this.setState({
			pictures: this.state.pictures.concat(picture),
		});

		let pomoc = this.state.pictures.length;
		pomoc = pomoc + 1;
	
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

	

	test(pic) {

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
		fetch(BASE_URL + "/api/items/upload", options);
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

										<div style={{ marginLeft: "0rem" }}><ImageUploader
											withIcon={false}
											buttonText='Add new photo/video'
											onChange={this.onDrop}
											imgExtension={['.jpg', '.gif', '.png', '.gif']}
											withPreview={true}
										/>
										<div style={{ marginLeft: "19rem" }} hidden={this.state.hiddenOne}>
												<Link style={{ width: "10rem" }} to="/userChangeProfile" className="btn btn-outline-secondary btn-sm">Add as feed post </Link>
												<a style={{ padding: "25px" }}></a>
												<Link style={{ width: "10rem" }} to="/userChangeProfile" className="btn btn-outline-secondary btn-sm">Add as story </Link>
												</div>
									
										<div style={{ marginLeft: "19rem" }} hidden={this.state.hiddenMultiple}>
												<Link style={{ width: "10rem" }} to="/userChangeProfile" className="btn btn-outline-secondary btn-sm">Add as feed album </Link>
												<a style={{ padding: "25px" }}></a>
												<Link style={{ width: "10rem" }} to="/userChangeProfile" className="btn btn-outline-secondary btn-sm">Add as album story </Link>
												</div>
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
                        
                    </div>
			</React.Fragment>
		);
	}
}

export default ProfilePage;
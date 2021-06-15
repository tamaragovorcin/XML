import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { BASE_URL_FEED, BASE_URL_STORY } from "../constants.js";
import playerLogo from "../static/me.jpg";
import LikesModal from "../components/Posts/LikesModal"
import DislikesModal from "../components/Posts/DislikesModal"
import CommentsModal from "../components/Posts/CommentsModal"
import WriteCommentModal from "../components/Posts/WriteCommentModal"
import WriteCommentAlbumModal from "../components/Posts/WriteCommentAlbumModal"
import Axios from "axios";
import IconTabsHomePage from "../components/Posts/IconTabsHomePage"
import AddPostToCollection from "../components/Posts/AddPostToCollection";
import ModalDialog from "../components/ModalDialog";
import StoriesModal from "../components/Posts/StoriesModal.jsx";
//import $ from 'jquery';
import { BASE_URL } from "../constants.js";
class HomePage extends React.Component {


	state = {
		ss: [],
		photos: [],
		peopleLikes: [],
		peopleDislikes: [],
		peopleComments: [],
		albums: [],
		showLikesModal: false,
		showDislikesModal: false,
		showCommentsModal: false,

		showStories: false,
		showWriteCommentModal: false,
		showAddPostToCollection: false,
		selectedPostId: -1,
		collections: [],
		showWriteCommentModalAlbum: false,
		users: [],
		pics: [],
		image: [],
		converted: undefined,
		help: [],
		ubiucse: "",
		pictures: [],
		bla: [1, 2],
		imageUrl: "",
		helpImage: "",
		hid: true,
		ready: false,
		stories: [],
		convertedImage: "",
		count: 0,
		userIsLogged: false,
		ssAlbums: [],
		usern: "",
		brojac : 0,
		br: 0,
		myCollectionAlbums : [],
		showAddAlbumToCollectionAlbum : false,
		userIsLoggedIn : true,
		stoori: [],
		stt: "",

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

	handleConvertedImage = (converted, username) => {
		
		var hh = this.state.stories;
		this.setState({ br: this.state.br +1 });
		if (this.state.usern === "") {

			this.setState({
				usern: username.username,
			});

			let st = { id: this.state.br, stories: [] }
			let storiji = {
				url: converted, header: {
					heading: username.username,
					subheading: 'CLOSE FRIENDS',

				},
			}
			st.stories.push(storiji)
			hh.push(st)
			this.setState({
				stories: hh,
			});




			if (this.state.brojac === hh.length) {
				this.setState({
					ready: true,
				});
			}
		}

		else if (this.state.usern == username.username) {

			this.state.stories.forEach(l => {
				l.stories.forEach(ll => {
					console.log(ll)
					if (ll.header.heading === username.username) {
						
						console.log(ll)
						let storiji = {
							url: converted, header: {
								heading: username.username,
								subheading: 'CLOSE FRIENDS',

							},
						}
						
						l.stories.push(storiji)
						var pom =l
						hh.pop(l)
						hh.push(pom)

					}

				
					this.setState({
						stories: hh,
					});

				})
			})





			if (this.state.brojac === hh.length) {
				this.setState({
					ready: true,
				});
			}
		}
		else{
			this.setState({
				usern: username.username,
			});

			let st = { id:  this.state.br, stories: [] }
			let storiji = {
				url: converted, header: {
					heading: username.username,
					subheading: 'CLOSE FRIENDS',
				},
			}
			st.stories.push(storiji)
			hh.push(st)
			this.setState({
				stories: hh,
			});




			if (this.state.brojac === hh.length) {
				this.setState({
					ready: true,
				});
			}
			console.log(hh)
		}
	}

	handleLikesModalOpen = (postId) => {
		Axios.get(BASE_URL + "/api/feedPosts/api/feed/likes/" + postId)
			.then((res) => {
				this.setState({ peopleLikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showLikesModal: true });
	}
	handleDislikesModalOpen = (postId) => {
		Axios.get(BASE_URL + "/api/feedPosts/api/feed/dislikes/" + postId)
			.then((res) => {
				this.setState({ peopleDislikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showDislikesModal: true });
	}
	handleCommentsModalOpen = (postId) => {
		Axios.get(BASE_URL + "/api/feedPosts/api/feed/comments/" + postId)
			.then((res) => {
				this.setState({ peopleComments: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showCommentsModal: true });
	}
	handleLikesModalOpenAlbum = (postId) => {
		Axios.get(BASE_URL + "/api/feedPosts/api/albumFeed/likes/" + postId)
			.then((res) => {
				this.setState({ peopleLikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showLikesModal: true });
	}
	handleDislikesModalOpenAlbum = (postId) => {
		Axios.get(BASE_URL + "/api/feedPosts/api/albumFeed/dislikes/" + postId)
			.then((res) => {
				this.setState({ peopleDislikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showDislikesModal: true });
	}
	handleCommentsModalOpenAlbum = (postId) => {
		Axios.get(BASE_URL + "/api/feedPosts/api/albumFeed/comments/" + postId)
			.then((res) => {
				this.setState({ peopleComments: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showCommentsModal: true });
	}
	handleWriteCommentModal = (postId) => {
		this.setState({ selectedPostId: postId });
		this.setState({ showWriteCommentModal: true });
	}
	handleWriteCommentModalAlbum = (postId) => {
		this.setState({ selectedPostId: postId });
		this.setState({ showWriteCommentModalAlbum: true });
	}
	handleWriteCommentAlbumModalClose = () => {
		this.setState({ showWriteCommentModalAlbum: false });
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
	handleStoriesClose = () => {
		this.setState({ showStories: false });
	}
	handleWriteCommentModalClose = () => {
		this.setState({ showWriteCommentModal: false });
	}
	onClickImage = (e, stor) => {
		console.log(stor)
		this.setState({ stoori: [] });
		this.setState({ stt: stor });
		this.setState({ stoori: stor });
		this.setState({ showStories: true });
	}

	handleLike = (postId) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		let postReactionDTO = {
			PostId: postId,
			UserId: id
		}
		Axios.post(BASE_URL + "/api/feedPosts/api/feed/like/", postReactionDTO, {
		}).then((res) => {

			this.setState({ textSuccessfulModal: "You have successfully liked the photo." });
			this.setState({ openModal: true });

		})
			.catch((err) => {
				console.log(err);
			});
	}
	handleLikeAlbum = (postId) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		let postReactionDTO = {
			PostId: postId,
			UserId: id
		}
		Axios.post(BASE_URL + "/api/feedPosts/api/albumFeed/like/", postReactionDTO, {
		}).then((res) => {

			this.setState({ textSuccessfulModal: "You have successfully liked the album." });
			this.setState({ openModal: true });

		})
			.catch((err) => {
				console.log(err);
			});
	}
	handleDislike = (postId) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		let postReactionDTO = {
			PostId: postId,
			UserId: id
		}
		Axios.post(BASE_URL + "/api/feedPosts/api/feed/dislike/", postReactionDTO, {
		}).then((res) => {

			this.setState({ textSuccessfulModal: "You have successfully disliked the photo." });
			this.setState({ openModal: true });

		})
			.catch((err) => {
				console.log(err);
			});
	}
	handleDislikeAlbum = (postId) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		let postReactionDTO = {
			PostId: postId,
			UserId: id
		}
		Axios.post(BASE_URL + "/api/feedPosts/api/albumFeed/dislike/", postReactionDTO, {
		}).then((res) => {

			this.setState({ textSuccessfulModal: "You have successfully disliked the album." });
			this.setState({ openModal: true });

		})
			.catch((err) => {
				console.log(err);
			});
	}


	componentDidMount() {
		if (this.hasRole("*")) {

			this.setState({ userIsLogged: true });

			let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)
			/*Axios.get(BASE_URL_STORY + "/api/storyAlbum/homePage/" + id)

			Axios.get(BASE_URL + "/api/storyPosts/api/storyAlbum/homePage/" + id)

				.then((res) => {
					let list = [];
					let br = this.state.brojac;
					res.data.forEach(story => {
						br = br+1
						let luna = [];

						story.Albums.forEach(s => {

							s.Media.forEach(media => {
								let aa = `data:image/jpg;base64,${media}`
								luna.push(aa)
							});

						});

						let highliht1 = { id: this.state.br, username: story.UserUsername+"Album", storiess: luna };
						list.push(highliht1)
						this.setState({ br: br +1 });

					});
					this.setState({ ss: this.state.ss.concat(list) });
					this.setState({ brojac: br });
					console.log(list)

				})
				.catch((err) => {
					console.log(err);
				});*/

				Axios.get(BASE_URL + "/api/storyPosts/api/story/homePage/" + id)
	
					.then((res) => {
						let br = this.state.brojac;
						let list = [];
						let st = [];
						let luna = [];
						
						res.data.forEach(story => {
							br = br + 1
							let luna = [];
							story.Stories.forEach(s => {
								let aa = `data:image/jpg;base64,${s.Media}`
								luna.push(aa)
	
							});
	
							let highliht1 = { id: res.data.id, username: story.UserUsername, storiess: {s: luna[0], username: story.UserUsername} };
							list.push(highliht1)
						});
						this.setState({ ss: this.state.ss.concat(list) });
						this.setState({ brojac: br });
					})
					.catch((err) => {
						console.log(err);
					});
	

			this.handleGetPhotos(id)
			this.handleGetAlbums(id)
		} else {
			this.setState({ userIsLogged: false });
		}

	}

	handleAddAllDataCollection = (id) => {
		Axios.post(BASE_URL + "/api/feedPosts/api/collection/allData/" + id)
			.then((res) => {
				if (res.status === 409) {
					this.setState({
						errorHeader: "Resource conflict!",
						errorMessage: "Email already exist.",
						hiddenErrorAlert: false,
					});
				} else if (res.status === 500) {
					this.setState({ errorHeader: "Internal server error!", errorMessage: "Server error.", hiddenErrorAlert: false });
				}


			})
			.catch((err) => {
				console.log(err);
			});
	}

	handleGetPhotos = (id) => {

		Axios.get(BASE_URL + "/api/feedPosts/api/feed/homePage/" + id)
			.then((res) => {
				this.setState({ photos: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}
	handleGetAlbums = (id) => {

		Axios.get(BASE_URL + "/api/feedPosts/api/albumFeed/homePage/" + id)
			.then((res) => {
				this.setState({ albums: res.data });
			})
			.catch((err) => {
				console.log(err);
			});

	}
	handleOpenAddPostToCollectionModal = (postId) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		this.handleGetCollections(id)
		this.setState({ showAddPostToCollection: true });
		this.setState({ selectedPostId: postId });
	}
	handleGetCollections = (id) => {
		Axios.get(BASE_URL + "/api/feedPosts/api/collection/user/" + id)
			.then((res) => {
				this.setState({ collections: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}
	handleAddPostToCollectionModalClose = () => {
		this.setState({ showAddPostToCollection: false });
		this.setState({ showAddAlbumToCollectionAlbum: false });
		}
	addPostToCollection = (collectionId) => {
		let postCollectionDTO = {
			PostId: this.state.selectedPostId,
			CollectionId: collectionId
		}
		Axios.post(BASE_URL + "/api/feedPosts/api/collection/addPost/", postCollectionDTO, {
		}).then((res) => {

			this.setState({ showAddCollectionModal: false });
			this.setState({ textSuccessfulModal: "You have successfully added post to collection." });
			this.setState({ openModal: true });
			this.setState({ showAddPostToCollection: false });

		})
			.catch((err) => {
				console.log(err);
			});
	}
	handleModalClose = () => {
		this.setState({ openModal: false });
	};
	handleAddComment = (comment) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		let commentDTO = {
			PostId: this.state.selectedPostId,
			UserId: id,
			Content: comment

		}
		Axios.post(BASE_URL + "/api/feedPosts/api/feed/comment/", commentDTO, {
		}).then((res) => {

			this.setState({ textSuccessfulModal: "You have successfully commented the photo." });
			this.setState({ openModal: true });
			this.setState({ showWriteCommentModal: false });


		})
			.catch((err) => {
				console.log(err);
			});
	}
	handleAddCommentAlbum = (comment) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		let commentDTO = {
			PostId: this.state.selectedPostId,
			UserId: id,
			Content: comment

		}
		Axios.post(BASE_URL + "/api/feedPosts/api/albumFeed/comment/", commentDTO, {
		}).then((res) => {

			this.setState({ textSuccessfulModal: "You have successfully commented the album." });
			this.setState({ showWriteCommentModalAlbum: false });

			this.setState({ openModal: true });


		})
			.catch((err) => {
				console.log(err);
			});
	}
	handleOpenAddAlbumToCollectionAlbumModal = (postId)=> {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

			Axios.get(BASE_URL + "/api/feedPosts/api/collection/user/album/"+id)
				.then((res) => {
					this.setState({ myCollectionAlbums: res.data });
				})
				.catch((err) => {
					console.log(err);
				});
		
		this.setState({ showAddAlbumToCollectionAlbum: true });
		this.setState({ selectedPostId: postId });
	}
	addAlbumToCollectionAlbum = (collectionId) => {
		let postCollectionDTO = {
			PostId : this.state.selectedPostId,
			CollectionId : collectionId
		}
		Axios.post(BASE_URL + "/api/feedPosts/api/collection/album/addPost/", postCollectionDTO, {
		}).then((res) => {
			
			this.setState({ showAddCollectionAlbumModal: false });
			let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
			this.handleGetCollectionAlbums(id);
			this.setState({ textSuccessfulModal: "You have successfully added album to collection." });
			this.setState({ openModal: true });
			this.setState({ showAddAlbumToCollectionAlbum: false });

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

				<section id="hero" className="d-flex align-items-top" >
					<div className="container" hidden={!this.state.userIsLogged}>
						<div className="container-fluid testimonial-group d-flex align-items-top">
							<div className="container-fluid scrollable" style={{ marginRight: "10rem", marginBottom: "5rem", marginTop: "5rem" }}>
								<table className="table-responsive" style={{ width: "100%" }}>
									<thead></thead>
									<tbody>

										
									
										{this.state.ss.map((post) => (
											<td id="td" style={{ width: "15em", height: "15em", marginLeft: "8em" }}>
												<tr >
													<img
														class="td"
														src={post.storiess.s}
														style={{ borderRadius: "50%", margin: "2%" }}
														width="100em"
														height="100em"
														max-width="100%"
														max-height="100%"
														alt="description"
														onClick={e => this.onClickImage(e, post.storiess)}
													/>

												</tr>



											</td>
										))}


									</tbody>
								</table>
							</div>
						</div>


						<div className="d-flex align-items-top">
							<IconTabsHomePage
								photos={this.state.photos}
								handleLike={this.handleLike}
								handleDislike={this.handleDislike}
								handleWriteCommentModal={this.handleWriteCommentModal}
								handleLikesModalOpen={this.handleLikesModalOpen}
								handleDislikesModalOpen={this.handleDislikesModalOpen}
								handleCommentsModalOpen={this.handleCommentsModalOpen}
								albums={this.state.albums}
								handleLikeAlbum={this.handleLikeAlbum}
								handleDislikeAlbum={this.handleDislikeAlbum}
								handleWriteCommentModalAlbum={this.handleWriteCommentModalAlbum}
								handleLikesModalOpenAlbum={this.handleLikesModalOpenAlbum}
								handleDislikesModalOpenAlbum={this.handleDislikesModalOpenAlbum}
								handleCommentsModalOpenAlbum={this.handleCommentsModalOpenAlbum}

								handleOpenAddPostToCollectionModal={this.handleOpenAddPostToCollectionModal}
								handleOpenAddAlbumToCollectionAlbumModal = {this.handleOpenAddAlbumToCollectionAlbumModal}
								userIsLoggedIn = {this.state.userIsLoggedIn}

							/>
						</div>


					</div>
				</section>

				<div>

				</div>

				<div>
					<StoriesModal
						show={this.state.showStories}
						onCloseModal={this.handleStoriesClose}
						stories={this.state.stoori}
						stt= {this.state.stt}
						ready={this.state.ready}
						brojac = {this.state.brojac}
					/>
					<LikesModal
						show={this.state.showLikesModal}
						onCloseModal={this.handleLikesModalClose}
						header="People who liked"
						peopleLikes={this.state.peopleLikes}
					/>
					<DislikesModal
						show={this.state.showDislikesModal}
						onCloseModal={this.handleDislikesModalClose}
						header="People who disliked"
						peopleDislikes={this.state.peopleDislikes}
					/>
					<CommentsModal
						show={this.state.showCommentsModal}
						onCloseModal={this.handleCommentsModalClose}
						header="Comments"
						peopleComments={this.state.peopleComments}
					/>

					<WriteCommentModal
						show={this.state.showWriteCommentModal}
						onCloseModal={this.handleWriteCommentModalClose}
						header="Leave your comment"
						handleAddComment={this.handleAddComment}
					/>
					<WriteCommentAlbumModal
						show={this.state.showWriteCommentModalAlbum}
						onCloseModal={this.handleWriteCommentAlbumModalClose}
						header="Leave your comment"
						handleAddCommentAlbum={this.handleAddCommentAlbum}
					/>
					<AddPostToCollection
						show={this.state.showAddPostToCollection}
						onCloseModal={this.handleAddPostToCollectionModalClose}
						header="Add post to collection"
						addPostToCollection={this.addPostToCollection}
						collections={this.state.collections}

					/>
					 <AddPostToCollection
                          
						  show={this.state.showAddAlbumToCollectionAlbum}
						  onCloseModal={this.handleAddPostToCollectionModalClose}
						  header="Add album to collection album"
						  addPostToCollection={this.addAlbumToCollectionAlbum}
						  collections = {this.state.myCollectionAlbums}
					  />
					<ModalDialog
						show={this.state.openModal}
						onCloseModal={this.handleModalClose}
						header="Successful"
						text={this.state.textSuccessfulModal}
					/>

				</div>
			</React.Fragment>
		);
	}
}

export default HomePage;
import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { Link } from "react-router-dom";
import playerLogo from "../static/coach.png";
import IconTabsFollowerProfile from "../components/Posts/IconTabsFollowerProfile"
import { BASE_URL_FEED, BASE_URL_STORY,BASE_URL_USER_INTERACTION } from "../constants.js";
import Axios from "axios";
import ModalDialog from "../components/ModalDialog";
import WriteCommentModal from "../components/Posts/WriteCommentModal"
import LikesModal from "../components/Posts/LikesModal"
import DislikesModal from "../components/Posts/DislikesModal"
import CommentsModal from "../components/Posts/CommentsModal"
import { BASE_URL_USER } from "../constants.js";
import AddPostToCollection from "../components/Posts/AddPostToCollection";
import WriteCommentAlbumModal from "../components/Posts/WriteCommentAlbumModal"
import { Lock } from "@material-ui/icons";
import { Icon } from "@material-ui/core";
import { isCompositeComponentWithType } from "react-dom/test-utils";
class FollowerProfilePage extends React.Component {
	
	state = {
		following: true,
		userId: "",
		id: "",
		username : "",
		name: "",
		lastName : "",
		webSite : "",
		biography : "",
		private : false,
		numberOfPosts : "",
		numberOfFollowers : "",
		numberOfFollowings : "",
		photos : [],
		albums : [],
		highlights : [],
		peopleLikes: [],
		peopleDislikes: [],
		comments: [],
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
		peopleComments : [],
		coords: [],
		addressNotFoundError: "none",
		showWriteCommentModal : false,
		showAddHighLightModal : false,
		highlightNameError : "none",
		collectionNameError : "none",
		showAddStoryToHighLightModal : false,
		showAddPostToCollection : false,
		selectedStoryId : -1,
		selectedPostId : -1,
		hiddenStoriesForHighlight : true,
		storiesForHightliht : [],
		collections  :[],
		postsForCollection : [],
		hiddenStoriesForCollection : true,
		showAddCollectionModal : false,
		showWriteCommentModalAlbum : false,
		followingThisUser : false,
		mutedThisUser : false,
		allowPagePreview : false,
		ableToFollowThisUser : false,
		sentFollowRequest : false,
		privateUser : false,
		showAddAlbumToCollectionAlbum : false,
		collectionAlbums : [],
		highlightsAlbums : [],
		storiesForHightlihtAlbum : [],
		hiddenStoriesForHighlightalbum : false,
		myCollectionAlbums : [],
		myCollections : [],
		userIsLoggedIn : false
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
	fetchData = (id) => {
		this.setState({
			userId: id,
		});
	};

	componentDidMount() {
		var sentence = window.location.toString()

		var s = []
		s = sentence.split("/");
		console.log(window.location.toString())


		this.fetchData(s[5]);
		Axios.get(BASE_URL_USER + "/api/" + s[5])
			.then((res) => {
				this.setState({
					id: res.data.Id,
					username : res.data.ProfileInformation.Username,
					name: res.data.ProfileInformation.Name,
					lastName : res.data.ProfileInformation.LastName,
					webSite : res.data.WebSite,
					biography : res.data.Biography,
					private : res.data.Private,
					numberOfPosts : res.data.numberOfPosts,
					numberOfFollowers : res.data.numberOfFollowers,
					numberOfFollowings : res.data.numberOfFollowings
				});

			})
			.catch((err) => {
				console.log(err);
			});
		this.handleGetHighlights(s[5])
		this.handleGetFeedPosts(s[5])
		this.handleGetAlbums(s[5])
		this.handleSetAllowPagePreview(s[5])
		this.handleGetCollectionAlbums(s[5])
		this.handleGetHighlightAlbums(s[5])


	}
	handleSetAllowPagePreview = (id)=> {
		if(!this.hasRole("*")) {
			this.setState({ userIsLoggedIn: false });

			this.setState({ followingThisUser: false});
			this.setState({ sentFollowRequest: false});
			this.setState({ ableToFollowThisUser: false});
			Axios.get(BASE_URL_USER + "/api/user/privacy/"+id)
			.then((res2) => {
				this.setState({ privateUser: res2.data });
				if( res2.data==="private") {
					this.setState({ allowPagePreview: false });
				}
				else {
					this.setState({ allowPagePreview: true });
				}	
			})
			.catch((err) => {
				console.log(err);
			});
		
		}
		else {

			let loggedId = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
			const followReguestDTO = { follower: loggedId, following : id};
			this.setState({ userIsLoggedIn: true });
			Axios.get(BASE_URL_USER + "/api/checkIfMuted/"+loggedId+"/"+id)
							.then((response) => {
								this.setState({ mutedThisUser: response.data });
							}).catch((err) => {
								console.log(err);
							});

			Axios.post(BASE_URL_USER_INTERACTION + "/api/checkInteraction",followReguestDTO)
				.then((res) => {
				this.setState({ followingThisUser: res.data });
				Axios.get(BASE_URL_USER + "/api/user/privacy/"+id)
					.then((res2) => {
						this.setState({ privateUser: res2.data });
						if(!res.data && res2.data==="private") {
							this.setState({ allowPagePreview: false });
						}
						else {
							this.setState({ allowPagePreview: true });
						}
						if(!res.data) {
							Axios.post(BASE_URL_USER_INTERACTION + "/api/checkIfSentRequest",followReguestDTO)
							.then((res3) => {
								this.setState({ sentFollowRequest: res3.data });
								if(res3.data) {
									this.setState({ ableToFollowThisUser: false });
								}
								else {
									this.setState({ ableToFollowThisUser: true });
								}
							}).catch((err) => {
								console.log(err);
							});
						}
						else {
							this.setState({ sentFollowRequest: false });
							this.setState({ ableToFollowThisUser: false });
						}
						
			
					})
					.catch((err) => {
						console.log(err);
					});
				
				
			})
			.catch((err) => {
				console.log(err);
			});
		}
	}
	handleAddCollectionClick = () => {
		this.setState({ showAddCollectionModal: true });
	};
	handleOpenAddPostToCollectionModal = (postId)=> {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		Axios.get(BASE_URL_FEED + "/api/collection/user/"+id)
			.then((res) => {
				this.setState({ myCollections: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showAddPostToCollection: true });
		this.setState({ selectedPostId: postId });
	}
	handleAddPostToCollectionModalClose = ()=> {
		this.setState({ showAddPostToCollection: false });
		this.setState({ showAddAlbumToCollectionAlbum: false });
		}
	addPostToCollection = (collectionId) => {
		let postCollectionDTO = {
			PostId : this.state.selectedPostId,
			CollectionId : collectionId
		}
		Axios.post(BASE_URL_FEED + "/api/collection/addPost/", postCollectionDTO, {
		}).then((res) => {
			
			this.setState({ showAddCollectionModal: false });
			let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
			this.handleGetCollections(id);
			this.setState({ textSuccessfulModal: "You have successfully added post to collection." });
			this.setState({ openModal: true });
			this.setState({ showAddPostToCollection: false });
			
		})
		.catch((err) => {
			console.log(err);
		});
	}
	seeStoriesInHighlight = (stories)=> {
		this.setState({ hiddenStoriesForHighlight: false });
		this.setState({storiesForHightliht : stories})
	}
	handleLikesModalOpen = (postId)=> {
		Axios.get(BASE_URL_FEED + "/api/feed/likes/"+postId)
			.then((res) => {
				this.setState({ peopleLikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showLikesModal: true });    
	}
	handleDislikesModalOpen = (postId)=> {
		Axios.get(BASE_URL_FEED + "/api/feed/dislikes/"+postId)
			.then((res) => {
				this.setState({ peopleDislikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showDislikesModal: true });    
	}
	handleCommentsModalOpen = (postId)=> {
		Axios.get(BASE_URL_FEED + "/api/feed/comments/"+postId)
			.then((res) => {
				this.setState({ peopleComments: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
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
	handleWriteCommentModalClose = ()=>{
		this.setState({showWriteCommentModal : false});
	}
	
	handleLike = (postId)=>{
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		let postReactionDTO = {
			PostId : postId,
			UserId : id
		}
		Axios.post(BASE_URL_FEED + "/api/feed/like/", postReactionDTO, {
		}).then((res) => {

			this.setState({ textSuccessfulModal: "You have successfully liked the photo." });
			this.setState({ openModal: true });

		})
		.catch((err) => {
			console.log(err);
		});
	}
	handleDislike = (postId)=>{
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		let postReactionDTO = {
			PostId : postId,
			UserId : id
		}
		Axios.post(BASE_URL_FEED + "/api/feed/dislike/", postReactionDTO, {
		}).then((res) => {

			this.setState({ textSuccessfulModal: "You have successfully disliked the photo." });
			this.setState({ openModal: true });

		})
		.catch((err) => {
			console.log(err);
		});
	}
	
	handleAddComment =(comment) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		let commentDTO = {
			PostId : this.state.selectedPostId,
			UserId : id,
			Content : comment

		}
		Axios.post(BASE_URL_FEED + "/api/feed/comment/", commentDTO, {
		}).then((res) => {
			
			this.setState({ textSuccessfulModal: "You have successfully commented the photo." });
			this.setState({ openModal: true });
			this.setState({ showWriteCommentModal: false });


		})
		.catch((err) => {
			console.log(err);
		});
	}
	handleLikesModalOpenAlbum = (postId)=> {
		Axios.get(BASE_URL_FEED + "/api/albumFeed/likes/"+postId)
			.then((res) => {
				this.setState({ peopleLikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showLikesModal: true });    
	}
	handleDislikesModalOpenAlbum = (postId)=> {
		Axios.get(BASE_URL_FEED + "/api/albumFeed/dislikes/"+postId)
			.then((res) => {
				this.setState({ peopleDislikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showDislikesModal: true });    
	}
	handleCommentsModalOpenAlbum = (postId)=> {
		Axios.get(BASE_URL_FEED + "/api/albumFeed/comments/"+postId)
			.then((res) => {
				this.setState({ peopleComments: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showCommentsModal: true });    
	}
	handleAddCommentAlbum =(comment) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		let commentDTO = {
			PostId : this.state.selectedPostId,
			UserId : id,
			Content : comment

		}
		Axios.post(BASE_URL_FEED + "/api/albumFeed/comment/", commentDTO, {
		}).then((res) => {
			
			this.setState({ textSuccessfulModal: "You have successfully commented the album." });
			this.setState({ openModal: true });
			this.setState({ showWriteCommentModalAlbum: false });


		})
		.catch((err) => {
			console.log(err);
		});
	}
	handleLikeAlbum = (postId)=>{
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		let postReactionDTO = {
			PostId : postId,
			UserId : id
		}
		Axios.post(BASE_URL_FEED + "/api/albumFeed/like/", postReactionDTO, {
		}).then((res) => {

			this.setState({ textSuccessfulModal: "You have successfully liked the album." });
			this.setState({ openModal: true });

		})
		.catch((err) => {
			console.log(err);
		});
	}
	handleDislikeAlbum= (postId)=>{
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		let postReactionDTO = {
			PostId : postId,
			UserId : id
		}
		Axios.post(BASE_URL_FEED + "/api/albumFeed/dislike/", postReactionDTO, {
		}).then((res) => {

			this.setState({ textSuccessfulModal: "You have successfully disliked the album." });
			this.setState({ openModal: true });

		})
		.catch((err) => {
			console.log(err);
		});
	}
	handleWriteCommentModalAlbum = (postId)=>{
		this.setState({ selectedPostId: postId });
		this.setState({showWriteCommentModalAlbum : true});
	}
	handleWriteCommentAlbumModalClose = ()=>{
		this.setState({showWriteCommentModalAlbum : false});
	}


	handleGetHighlights = (id) => {
		Axios.get(BASE_URL_STORY + "/api/highlight/user/"+id)
			.then((res) => {
				this.setState({ highlights: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}
	

	handleGetFeedPosts = (id) => {
		Axios.get(BASE_URL_FEED + "/api/feed/usersImages/"+id)
			.then((res) => {
				this.setState({ photos: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}
	handleGetAlbums = (id) => {
		Axios.get(BASE_URL_FEED + "/api/feedAlbum/usersAlbums/"+id)
			.then((res) => {
				this.setState({ albums: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}

	
	
	
	handleWriteCommentModal = (postId)=>{
		this.setState({ selectedPostId: postId });
		this.setState({showWriteCommentModal : true});
	}
	handleAddComment =(comment) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		let commentDTO = {
			PostId : this.state.selectedPostId,
			UserId : id,
			Content : comment

		}
		Axios.post(BASE_URL_FEED + "/api/feed/comment/", commentDTO, {
		}).then((res) => {
			
			this.setState({ textSuccessfulModal: "You have successfully commented the photo." });
			this.setState({ openModal: true });
			this.setState({ showWriteCommentModal: false });


		})
		.catch((err) => {
			console.log(err);
		});
	}


	handleModalClose = () => {
		this.setState({ openModal: false });
	};

	


	handleFollow = () => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
	
		const followReguestDTO = { follower: id, following : this.state.userId};
		if(this.state.privateUser==="private") {

			Axios.post(BASE_URL_USER_INTERACTION + "/api/followRequest", followReguestDTO)
			.then((res) => {
				
				this.handleSetAllowPagePreview(this.state.userId)
				
				this.setState({ textSuccessfulModal: "You have successfully sent follow request." });
				this.setState({ openModal: true });

			})
			.catch ((err) => {
				console.log(err);
			});
		}else {
			Axios.post(BASE_URL_USER_INTERACTION + "/api/followPublic", followReguestDTO)
			.then((res) => {
				
				this.handleSetAllowPagePreview(this.state.userId)
				
				this.setState({ textSuccessfulModal: "You are now following this user." });
				this.setState({ openModal: true });

			})
			.catch ((err) => {
				console.log(err);
			});
		}
		

	}
	handleMute = () => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
	
		const dto = { Subject: id, Object : this.state.userId};
			Axios.post(BASE_URL_USER + "/api/mute/", dto)
			.then((res) => {
								
				this.setState({ textSuccessfulModal: "You have successfully muted this user." });
				this.setState({ openModal: true });

			})
			.catch ((err) => {
				console.log(err);
			});
			
		
		

	}
	handleUnMute = () => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
	
		const dto = { Subject: id, Object : this.state.userId};
			Axios.post(BASE_URL_USER + "/api/unmute/", dto)
			.then((res) => {
								
				this.setState({ textSuccessfulModal: "You have successfully unmuted this user." });
				this.setState({ openModal: true });

			})
			.catch ((err) => {
				console.log(err);
			});
			
		
		

	}
	handleBlock = () => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
	
		const dto = { Subject: id, Object : this.state.userId};
			Axios.post(BASE_URL_USER + "/api/block/", dto)
			.then((res) => {
								
				this.setState({ textSuccessfulModal: "You have successfully blocked this user." });
				this.setState({ openModal: true });

			})
			.catch ((err) => {
				console.log(err);
			});
			

	}
	handleOpenAddAlbumToCollectionAlbumModal = (postId)=> {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

			Axios.get(BASE_URL_FEED + "/api/collection/user/album/"+id)
				.then((res) => {
					this.setState({ myCollectionAlbums: res.data });
				})
				.catch((err) => {
					console.log(err);
				});
		
		this.setState({ showAddAlbumToCollectionAlbum: true });
		this.setState({ selectedPostId: postId });
	}

	handleGetCollectionAlbums = (id) => {
		Axios.get(BASE_URL_FEED + "/api/collection/user/album/"+id)
			.then((res) => {
				this.setState({ collectionAlbums: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}
	addAlbumToCollectionAlbum = (collectionId) => {
		let postCollectionDTO = {
			PostId : this.state.selectedPostId,
			CollectionId : collectionId
		}
		Axios.post(BASE_URL_FEED + "/api/collection/album/addPost/", postCollectionDTO, {
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
	handleGetHighlightAlbums = (id) => {
		Axios.get(BASE_URL_STORY + "/api/highlight/user/album/"+id)
			.then((res) => {
				this.setState({ highlightsAlbums: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}
	seeStoriesInHighlightAlbum = (stories)=> {
		this.setState({ hiddenStoriesForHighlightAlbum: false });
		this.setState({storiesForHightlihtAlbum : stories})
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
														<div hidden={!this.state.followingThisUser}>
															<button  className="btn btn-outline-success mt-1"  type="button"><i className="icofont-subscribe mr-1"></i>Following</button>
														</div>
														<div hidden={!this.state.ableToFollowThisUser}>
															<button  className="btn btn-outline-primary mt-1" onClick={() => this.handleFollow()} type="button"><i className="icofont-subscribe mr-1"></i>Follow</button>
														</div>
														<div hidden={!this.state.sentFollowRequest}>
															<button  className="btn btn-outline-warning mt-1"  type="button"><i className="icofont-subscribe mr-1"></i>Sent request</button>
														</div>
														<div hidden={!this.state.followingThisUser || this.state.mutedThisUser}>
															<button  className="btn btn-outline-primary mt-1" onClick={() => this.handleMute()} type="button"><i className="icofont-subscribe mr-1"></i>Mute</button>
														</div>
														<div hidden={!this.state.mutedThisUser}>
															<button  className="btn btn-outline-primary mt-1" onClick={() => this.handleUnMute()} type="button"><i className="icofont-subscribe mr-1"></i>Unmute</button>
														</div>
														<div>
															<button  className="btn btn-outline-primary mt-1" onClick={() => this.handleBlock()} type="button"><i className="icofont-subscribe mr-1"></i>Block</button>
														</div>
													</td>

												</div>
											
												<div>
													<td>
														<label >{this.state.biography}</label>
													</td>
													<td>
														<label >{this.state.webSite}</label>
													</td>
												</div>


											</td>


										</tr>
									</tbody>
								</table>
							</div>
						</div>
						<div hidden={!this.state.allowPagePreview}>
							<IconTabsFollowerProfile
								photos = {this.state.photos}
								handleLike = {this.handleLike}
								handleDislike = {this.handleDislike}
								handleWriteCommentModal = {this.handleWriteCommentModal}						
								handleSave = {this.handleSave}
								handleLikesModalOpen = {this.handleLikesModalOpen}
								handleDislikesModalOpen = {this.handleDislikesModalOpen}
								handleCommentsModalOpen = {this.handleCommentsModalOpen}

								albums ={this.state.albums}
								handleLikeAlbum = {this.handleLikeAlbum}
								handleDislikeAlbum  = {this.handleDislikeAlbum }
								handleWriteCommentModalAlbum  = {this.handleWriteCommentModalAlbum }						
								handleLikesModalOpenAlbum  = {this.handleLikesModalOpenAlbum }
								handleDislikesModalOpenAlbum  = {this.handleDislikesModalOpenAlbum}
								handleCommentsModalOpenAlbum  = {this.handleCommentsModalOpenAlbum }


								highlights = {this.state.highlights}
								seeStoriesInHighlight = {this.seeStoriesInHighlight}
								storiesForHightliht= {this.state.storiesForHightliht}
								hiddenStoriesForHighlight = {this.state.hiddenStoriesForHighlight}

								handleAddCollectionClick = {this.handleAddCollectionClick}
								collections = {this.state.collections}
								seePostsInCollection = {this.seePostsInCollection}
								postsForCollection = {this.state.postsForCollection}
								hiddenStoriesForCollection = {this.state.hiddenStoriesForCollection}
								handleOpenAddPostToCollectionModal = {this.handleOpenAddPostToCollectionModal}
								handleOpenAddAlbumToCollectionAlbumModal = {this.handleOpenAddAlbumToCollectionAlbumModal}

								highlightsAlbums = {this.state.highlightsAlbums}
								seeStoriesInHighlightAlbum = {this.seeStoriesInHighlightAlbum}
								storiesForHightlihtAlbum= {this.state.storiesForHightlihtAlbum}
								hiddenStoriesForHighlightalbum = {this.state.hiddenStoriesForHighlightAlbum}
								userIsLoggedIn = {this.state.userIsLoggedIn}
						/>
						</div>

						<div hidden={this.state.allowPagePreview}>

							<div className="d-flex align-items-top p-3 mb-2 d-flex justify-content-center">

								<label><b>This Account is Private</b></label>

							</div>

							<div className="d-flex justify-content-center h-100">
								<Icon className="d-flex justify-content-center h-100 w-100"><Lock /></Icon>
							</div>

						</div>

					</div>

					

				</section>
				<div>

					<LikesModal
					        show={this.state.showLikesModal}
					        onCloseModal={this.handleLikesModalClose}
					        header="People who liked"
							peopleLikes = {this.state.peopleLikes}
				    />
                    <DislikesModal
                         show={this.state.showDislikesModal}
						 onCloseModal={this.handleDislikesModalClose}
						 header="People who disliked"
						 peopleDislikes = {this.state.peopleDislikes}
				    />
                    <CommentsModal
                        show={this.state.showCommentsModal}
						onCloseModal={this.handleCommentsModalClose}
						header="Comments"
						peopleComments = {this.state.peopleComments}
                    />
					<WriteCommentModal
                        show={this.state.showWriteCommentModal}
						onCloseModal={this.handleWriteCommentModalClose}
						header="Leave your comment"
						handleAddComment = {this.handleAddComment}
                    />
					<WriteCommentAlbumModal
                        show={this.state.showWriteCommentModalAlbum}
						onCloseModal={this.handleWriteCommentAlbumModalClose}
						header="Leave your comment"
						handleAddCommentAlbum = {this.handleAddCommentAlbum}
                    />
                    <ModalDialog
						show={this.state.openModal}
						onCloseModal={this.handleModalClose}
						header="Successful"
						text={this.state.textSuccessfulModal}
					/>
				
					  <AddPostToCollection
                          
						  show={this.state.showAddPostToCollection}
						  onCloseModal={this.handleAddPostToCollectionModalClose}
						  header="Add post to collection"
						  addPostToCollection={this.addPostToCollection}
						  collections = {this.state.myCollections}
					  />
					  <AddPostToCollection
                          
						  show={this.state.showAddAlbumToCollectionAlbum}
						  onCloseModal={this.handleAddPostToCollectionModalClose}
						  header="Add album to collection album"
						  addPostToCollection={this.addAlbumToCollectionAlbum}
						  collections = {this.state.myCollectionAlbums}
					  />

				</div>

			</React.Fragment >
		);
	}
}

export default FollowerProfilePage;
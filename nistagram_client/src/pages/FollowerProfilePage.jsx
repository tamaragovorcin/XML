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
class FollowerProfilePage extends React.Component {
	constructor(props) {
		super(props);

	}
	state = {
		following: true,
		userId: "",
		id: "",
		username : "",
		name: "",
		lastName : "",
		webSite : "",
		biography : "",
		private : "",
		numberOfPosts : "",
		numberOfFollowers : "",
		numberOfFollowings : "",
		photos : [],
		albums : [],
		stories : [],
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
		peopleLikes : [],
		peopleDislikes : [],
		peopleComments : [],
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
		showWriteCommentModal : false,
		albums : [],
		stories : [],
		highlights : [],
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
		showWriteCommentModalAlbum : false
	}

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
		let id = localStorage.getItem("userId")
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
		this.handleGetStories(s[5])
		this.handleGetAlbums(s[5])

	}
	handleAddCollectionClick = () => {
		this.setState({ showAddCollectionModal: true });
	};
	handleOpenAddPostToCollectionModal = (postId)=> {
		this.setState({ showAddPostToCollection: true });
		this.setState({ selectedPostId: postId });
	}
	handleAddPostToCollectionModalClose = ()=> {
		this.setState({ showAddPostToCollection: false });
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
		alert("POGODIO")
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

	handleGetStories = (id)=> {
		Axios.get(BASE_URL_STORY + "/api/story/user/"+id)
		.then((res) => {
			this.setState({ stories: res.data });
		})
		.catch((err) => {
			console.log(err);
		});
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
													<Link onClick={this.handleFollow} className="btn btn-outline-success btn-sm">Follow</Link>

													</td>

												</div>
												<div>
													<td>
														<label ><b>{this.state.numberOfPosts}</b> posts</label>
													</td>
													<td>
														<label ><b>{this.state.numberOfFollowers}</b> followers</label>
													</td>
													<td>
														<label ><b>{this.state.numberOfFollowings}</b> following</label>
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

						stories = {this.state.stories}

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

						
					/>



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
						  collections = {this.state.collections}
					  />

				</div>

			</React.Fragment >
		);
	}
}

export default FollowerProfilePage;
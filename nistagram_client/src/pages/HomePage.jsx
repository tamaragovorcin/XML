import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { BASE_URL_FEED} from "../constants.js";
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

class HomePage extends React.Component {
	
	state = {
		stories: [],
		photos: [],
		peopleLikes : [],
		peopleDislikes : [],
		peopleComments : [],
		albums : [],
		showLikesModal : false,
		showDislikesModal : false,
		showCommentsModal : false,
		showWriteCommentModal : false,
		showAddPostToCollection : false,
		selectedPostId : -1,
		collections : [],
		showWriteCommentModalAlbum : false

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
	handleWriteCommentModal = (postId)=>{
		this.setState({ selectedPostId: postId });
		this.setState({showWriteCommentModal : true});
	}
	handleWriteCommentModalAlbum = (postId)=>{
		this.setState({ selectedPostId: postId });
		this.setState({showWriteCommentModalAlbum : true});
	}
	handleWriteCommentAlbumModalClose = ()=>{
		this.setState({showWriteCommentModalAlbum : false});
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
	componentDidMount() {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		
		this.handleGetStories(id)
		this.handleGetPhotos(id)
		this.handleGetAlbums(id)

	}
	

	handleGetStories= (id) => {
		let highliht1 = { id: 1, username: "mladenkak" };
		let highliht2 = { id: 2, username: "tamarag" };
		let highliht3 = { id: 3, username: "lunaz" };
		let highliht4 = { id: 4, username: "lunaz" };
		let highliht5 = { id: 5, username: "lunaz" };
		let highliht6 = { id: 6, username: "lunaz" };
		let highliht7 = { id: 7, username: "lunaz" };
		let highliht8 = { id: 8, username: "lunaz" };
		let highliht9 = { id: 9, username: "lunaz" };
		let highliht10 = { id: 10, username: "lunaz" };
		let highliht11 = { id: 11, username: "lunaz" };
		let highliht12 = { id: 12, username: "lunaz" };
		let highliht13 = { id: 13, username: "lunaz" };
		
		let list = [];
		list.push(highliht1)
		list.push(highliht2)
		list.push(highliht3)
		list.push(highliht4)
		list.push(highliht5)
		list.push(highliht6)
		list.push(highliht7)
		list.push(highliht8)
		list.push(highliht9)
		list.push(highliht10)
		list.push(highliht11)
		list.push(highliht12)
		list.push(highliht13)
		list.push(highliht13)
		list.push(highliht13)
		list.push(highliht13)
		list.push(highliht13)
		list.push(highliht13)
		list.push(highliht13)
		list.push(highliht1)
		list.push(highliht2)
		list.push(highliht3)
		list.push(highliht4)
		list.push(highliht5)
		list.push(highliht6)
		list.push(highliht7)
		list.push(highliht8)
		list.push(highliht9)
		list.push(highliht10)
		list.push(highliht11)
		list.push(highliht12)
		list.push(highliht13)
		list.push(highliht13)
		list.push(highliht13)
		list.push(highliht13)
		list.push(highliht13)
		list.push(highliht13)
		list.push(highliht13)

		this.setState({ stories: list });
	}

	handleGetPhotos = (id) => {
		
		Axios.get(BASE_URL_FEED + "/api/feed/homePage/"+id)
			.then((res) => {
				this.setState({ photos: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}
	handleGetAlbums = (id) => {
		
		Axios.get(BASE_URL_FEED + "/api/albumFeed/homePage/"+id)
			.then((res) => {
				this.setState({ albums: res.data });
			})
			.catch((err) => {
				console.log(err);
			});

	}
	handleOpenAddPostToCollectionModal = (postId)=> {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		this.handleGetCollections(id)
		this.setState({ showAddPostToCollection: true });
		this.setState({ selectedPostId: postId });
	}
	handleGetCollections = (id) => {
		Axios.get(BASE_URL_FEED + "/api/collection/user/"+id)
			.then((res) => {
				this.setState({ collections: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
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
			this.setState({ textSuccessfulModal: "You have successfully added post to highlight." });
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
			this.setState({ showWriteCommentModal: false });


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
					<div className="container-fluid testimonial-group d-flex align-items-top">
							<div className="container-fluid scrollable" style={{ marginRight: "10rem" , marginBottom:"5rem",marginTop:"5rem"}}>
								<table className="table-responsive" style={{ width: "100%" }}>
									<tbody>

										<tr >
											{this.state.stories.map((high) => (
												<td id={high.id} key={high.id} style={{width:"60em", marginLeft:"10em"}}>
													<tr width="100em">
														<img
															className="img-fluid"
															src={playerLogo}
															style ={{borderRadius:"50%",margin:"2%"}}
															width="60em"
															alt="description"
														/>
													</tr>
													<tr>
														<label style={{marginRight:"15px"}}>{high.username}</label>
													</tr>
												</td>
												
											))}
										</tr>


									</tbody>
								</table>
							</div>
					</div>
				
					<div className="d-flex align-items-top">
						<IconTabsHomePage
							photos = {this.state.photos}
							handleLike = {this.handleLike}
							handleDislike = {this.handleDislike}
							handleWriteCommentModal = {this.handleWriteCommentModal}						
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

							handleOpenAddPostToCollectionModal = {this.handleOpenAddPostToCollectionModal}

						/>
					</div>

				</div>
					
				</section>
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
					 <AddPostToCollection
						  show={this.state.showAddPostToCollection}
						  onCloseModal={this.handleAddPostToCollectionModalClose}
						  header="Add post to collection"
						  addPostToCollection={this.addPostToCollection}
						  collections = {this.state.collections}

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
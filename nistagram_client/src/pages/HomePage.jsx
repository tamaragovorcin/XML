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
import Stories from 'react-insta-stories';
import Axios from "axios";
import IconTabsHomePage from "../components/Posts/IconTabsHomePage"
import AddPostToCollection from "../components/Posts/AddPostToCollection";
import ModalDialog from "../components/ModalDialog";
import ConvertImage from "react-convert-image";
import StoriesModal from "../components/Posts/StoriesModal.jsx";
//import $ from 'jquery';
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




	}


	Change = () => {
		this.state.pictures.forEach(pic => {
			console.log(pic)
		}

		)

	}


	handleConvertedImage = (converted) => {

		var hh = this.state.stories;
		var username = this.state.ss[hh.length].username
		
		let st = { id: 1, stories: [] }
		let storiji = {url: converted, header: {
			heading: username,
			subheading: 'CLOSE FRIENDS',
			
		},}
		st.stories.push(storiji)
		hh.push(st)
		this.setState({
			stories: hh,
		});
		this.setState({
			convertedImage: converted,
		});



		if (this.state.ss.length === hh.length) {
			this.setState({
				ready: true,
			});
		}
		console.log(hh)
	}

	handleLikesModalOpen = (postId) => {
		Axios.get(BASE_URL_FEED + "/api/feed/likes/" + postId)
			.then((res) => {
				this.setState({ peopleLikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showLikesModal: true });
	}
	handleDislikesModalOpen = (postId) => {
		Axios.get(BASE_URL_FEED + "/api/feed/dislikes/" + postId)
			.then((res) => {
				this.setState({ peopleDislikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showDislikesModal: true });
	}
	handleCommentsModalOpen = (postId) => {
		Axios.get(BASE_URL_FEED + "/api/feed/comments/" + postId)
			.then((res) => {
				this.setState({ peopleComments: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showCommentsModal: true });
	}
	handleLikesModalOpenAlbum = (postId) => {
		Axios.get(BASE_URL_FEED + "/api/albumFeed/likes/" + postId)
			.then((res) => {
				this.setState({ peopleLikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showLikesModal: true });
	}
	handleDislikesModalOpenAlbum = (postId) => {
		Axios.get(BASE_URL_FEED + "/api/albumFeed/dislikes/" + postId)
			.then((res) => {
				this.setState({ peopleDislikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showDislikesModal: true });
	}
	handleCommentsModalOpenAlbum = (postId) => {
		Axios.get(BASE_URL_FEED + "/api/albumFeed/comments/" + postId)
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
	onClickImage = () => {
		this.setState({ showStories: true });
	}

	handleLike = (postId) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		let postReactionDTO = {
			PostId: postId,
			UserId: id
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
	handleLikeAlbum = (postId) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		let postReactionDTO = {
			PostId: postId,
			UserId: id
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
	handleDislike = (postId) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		let postReactionDTO = {
			PostId: postId,
			UserId: id
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
	handleDislikeAlbum = (postId) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		let postReactionDTO = {
			PostId: postId,
			UserId: id
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


		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		Axios.get(BASE_URL_STORY + "/api/story/homePage/" + id)

			.then((res) => {
				let list = [];
				let st = [];
				let luna = [];
				res.data.forEach(story => {
					let luna = [];
					story.Stories.forEach(s => {


						let aa = `data:image/jpg;base64,${s.Media}`
						luna.push(aa)
						//st.push(luna)



					});

					let highliht1 = { id: res.data.id, username: story.UserUsername, storiess: luna };
					list.push(highliht1)


				});
				//this.setState({ image: st });
				this.setState({ ss: list });
				//this.setState({ bla: luna });

			})
			.catch((err) => {
				console.log(err);
			});

		//this.setState({ showCommentsModal: true });

		this.handleGetPhotos(id)
		this.handleGetAlbums(id)

	}

	handleAddAllDataCollection = (id) => {
		Axios.post(BASE_URL_FEED + "/api/collection/allData/" + id)
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

		Axios.get(BASE_URL_FEED + "/api/feed/homePage/" + id)
			.then((res) => {
				this.setState({ photos: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}
	handleGetAlbums = (id) => {

		Axios.get(BASE_URL_FEED + "/api/albumFeed/homePage/" + id)
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
		Axios.get(BASE_URL_FEED + "/api/collection/user/" + id)
			.then((res) => {
				this.setState({ collections: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}
	handleAddPostToCollectionModalClose = () => {
		this.setState({ showAddPostToCollection: false });
	}
	addPostToCollection = (collectionId) => {
		let postCollectionDTO = {
			PostId: this.state.selectedPostId,
			CollectionId: collectionId
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
	handleAddComment = (comment) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		let commentDTO = {
			PostId: this.state.selectedPostId,
			UserId: id,
			Content: comment

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
	handleAddCommentAlbum = (comment) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		let commentDTO = {
			PostId: this.state.selectedPostId,
			UserId: id,
			Content: comment

		}
		Axios.post(BASE_URL_FEED + "/api/albumFeed/comment/", commentDTO, {
		}).then((res) => {

			this.setState({ textSuccessfulModal: "You have successfully commented the album." });
			this.setState({ showWriteCommentModalAlbum: false });

			this.setState({ openModal: true });


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
				{this.state.ss.map((user) => (
					<div hidden={this.state.hid}>
						{user.storiess.map((post) => (
							<div hidden={this.state.hid}>
								<ConvertImage
									image={post}
									onConversion={this.handleConvertedImage}

								/>
							</div>))}</div>

				))}

				<section id="hero" className="d-flex align-items-top">
					<div className="container">
						<div className="container-fluid testimonial-group d-flex align-items-top">
							<div className="container-fluid scrollable" style={{ marginRight: "10rem", marginBottom: "5rem", marginTop: "5rem" }}>
								<table className="table-responsive" style={{ width: "100%" }}>
									<thead></thead>
									<tbody>


										{this.state.ss.map((post) => (
											<td  id="td" style={{ width: "15em", height: "15em" ,marginLeft: "8em" }}>
												<tr >
													<img
														class="td"
														src={post.storiess[0]}
														style={{ borderRadius: "50%", margin: "2%" }}
														width="100em"
														height="100em"
														max-width= "100%"
														max-height= "100%"
														alt="description"
														onClick={this.onClickImage}
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

							/>
						</div>

					
						</div>
					</section>

<div>
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

						<StoriesModal
							show={this.state.showStories}
							onCloseModal={this.handleStoriesClose}
							stories={this.state.stories}
							ready={this.state.ready}
							count = {this.state.count}
						/>
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

						/>
					</div>

				</div>

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
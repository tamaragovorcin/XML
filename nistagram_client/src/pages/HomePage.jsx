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
		showWriteCommentModal: false,
		showAddPostToCollection: false,
		selectedPostId: -1,
		collections: [],
		showWriteCommentModalAlbum: false,
		users: [],
		pics: [],
		image: "",
		converted: undefined,
		help: [],
		pictures: [],
		bla: [],
		hid: true,
		stories: ["blob:http://localhost:3000/ac876899-5147-482c-9086-998ee05c765f"],
		image: ""

	}


	Change = () => {
		this.state.pictures.forEach(pic => {
			console.log(pic)
		}

		)

	}

	handleBlob = (im) => {

		
		var ab = new ArrayBuffer(im.length);
		var ia = new Uint8Array(ab);
	
		for (var i = 0; i < im.length; i++) {
			ia[i] = im.charCodeAt(i);
		}
		var blob =  new Blob([ab], { type: 'image/jpg' });
		var url = window.URL.createObjectURL(blob);
		alert(url)



		//const byteCharacters = atob(converted);
		/*const byteNumbers = new Array(im.length);
		for (let i = 0; i < im.length; i++) {
			byteNumbers[i] = im.charCodeAt(i);
		}
		const byteArray = new Uint8Array(byteNumbers);

		let image = new Blob([byteArray], { type: contentType });




		let imageUrl = URL.createObjectURL(image);
		alert(imageUrl)
		//imageUrl = imageUrl.replace("blob:", "");
		this.setState({ image: imageUrl });
		let hh = this.state.pictures;
		hh.push(imageUrl)
		console.log(hh)
		this.setState({
			pictures: hh,
		});*/


	}
	handleConvertedImage = (converted) => {




		/*	converted = converted.replace("webp", "jpg");
			console.log(converted)
			let hh = this.state.pictures;
			hh.push(converted)
			console.log(hh)
			this.setState({
				pictures: hh,
			});*/

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
	handleWriteCommentModalClose = () => {
		this.setState({ showWriteCommentModal: false });
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
					story.Stories.forEach(s => {
						luna.push(s.Media)
						this.handleBlob(s.Media)

					});
					let highliht1 = { id: res.data.id, username: story.UserUsername, storiess: luna };

					list.push(highliht1)

				});

				this.setState({ ss: list });
				this.setState({ bla: luna });

			})
			.catch((err) => {
				console.log(err);
			});

		//this.setState({ showCommentsModal: true });

		this.handleGetPhotos(id)
		this.handleGetAlbums(id)

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
			this.setState({ openModal: true });
			this.setState({ showWriteCommentModal: false });


		})
			.catch((err) => {
				console.log(err);
			});
	}
	render() {
		const { state: { converted } } = this;
		return (
			<React.Fragment>
				<TopBar />
				<Header />

				<div className="container-fluid testimonial-group d-flex align-items-top">
					<div className="container-fluid scrollable" style={{ marginRight: "10rem", marginBottom: "5rem", marginTop: "10rem" }}>
						<table style={{ width: "100%" }}>
							<thead></thead>
							<tbody>


								{this.state.bla.map((high) => (
									<tr style={{ width: "60em", marginLeft: "10em" }}>
										<td width="100em">
											<img
												className="img-fluid"
												src={playerLogo}
												style={{ borderRadius: "50%", margin: "2%" }}
												width="60em"
												alt="description"
											/>

										</td>

										<td>
										
										</td>


									</tr>))}


							</tbody>
						</table>

					</div>

					<div style={{ marginTop: "150px" }}>

						<button onClick={this.Change}>Button</button>
					</div>
					<div hidden={this.state.hid}>
						{this.state.bla.map((high) => (
							<ConvertImage
								image={`data:image/jpg;base64,${high}`}
								onConversion={this.handleConvertedImage}

							/>))}
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

				<div>

					<LikesModal
						show={this.state.showLikesModal}
						onCloseModal={this.handleLikesModalClose}
						header="People who liked the photo"
						peopleLikes={this.state.peopleLikes}
					/>
					<DislikesModal
						show={this.state.showDislikesModal}
						onCloseModal={this.handleDislikesModalClose}
						header="People who disliked the photo"
						peopleDislikes={this.state.peopleDislikes}
					/>
					<CommentsModal
						show={this.state.showCommentsModal}
						onCloseModal={this.handleCommentsModalClose}
						header="Comments on the photo"
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
import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";

import LikesModal from "../components/Posts/LikesModal"
import DislikesModal from "../components/Posts/DislikesModal"
import CommentsModal from "../components/Posts/CommentsModal"
import WriteCommentModal from "../components/Posts/WriteCommentModal"
import WriteCommentAlbumModal from "../components/Posts/WriteCommentAlbumModal"
import { YMaps, Map } from "react-yandex-maps";
import Axios from "axios";
import { BASE_URL_FEED } from "../constants.js";
import IconTabsHomePage from "../components/Posts/IconTabsHomePage"
import ModalDialog from "../components/ModalDialog";
import AddPostToCollection from "../components/Posts/AddPostToCollection";

const mapState = {
	center: [44, 21],
	zoom: 8,
	controls: [],
};
class Search extends React.Component {
	constructor(props) {
		super(props);
		this.addressInput = React.createRef();
	}
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
		showWriteCommentModalAlbum : false,
        coords: [],
		addressNotFoundError: "none",
        addressError: "none",
        hashtags :"",
        hashtagsError : "none"

	}

    onYmapsLoad = (ymaps) => {
		this.ymaps = ymaps;
		new this.ymaps.SuggestView(this.addressInput.current, {
			provider: {
				suggest: (request, options) => this.ymaps.suggest(request),
			},
		});
	};


	componentDidMount() {}
	
	handleHashTagsChange = (event) => {
		this.setState({ hashtags:  event.target.value });
	}

	handleSearchByLocation = ()=> {


        if (this.state.addressInput === "") {
			this.setState({ addressError: "initial" });
			return false;
		}
        let street;
		let city;
		let country;
		let found = true;
		this.ymaps
			.geocode(this.addressInput.current.value, {
				results: 1,
			})
			.then(function (res) {

				if (typeof res.geoObjects.get(0) === "undefined") found = false;
				else {
					var firstGeoObject = res.geoObjects.get(0);
				
					country = firstGeoObject.getCountry();
					street = firstGeoObject.getThoroughfare();
					city = firstGeoObject.getLocalities().join(", ");
             
                    if(country===undefined || country==="") {
                        country = "n"
                    }
                    if(street===undefined || street==="" ) {
                        street="n"
                    }
                    if(city===undefined || city ==="") {
                        city = "n"
                    }

				}
			})
			.then((res) => {
                if (found === false) {
                    this.setState({ addressNotFoundError: "initial" });
                } else {
                    this.setState({ photos: [] });
                    this.setState({ albums: [] });

                    Axios.get(BASE_URL_FEED + "/api/feed/searchByLocation/"+country + "/"+city+"/"+street)
                    .then((res) => {
                        this.setState({ photos: res.data });
                        this.setState({ hashtags: "" });
                    })
                    .catch((err) => {
                        console.log(err);
                    });
					Axios.get(BASE_URL_FEED + "/api/albumFeed/searchByLocation/"+country + "/"+city+"/"+street)
                    .then((res) => {
                        this.setState({ albums: res.data });
                        this.setState({ hashtags: "" });
                    })
                    .catch((err) => {
                        console.log(err);
                    });
                }
			});
    }

    handleSearchByHashTags = () => {
       var help = this.state.hashtags
       if(help==="") {
            help = "n"
       }
        let helpDTO = {
            HashTags : help
        }
		this.setState({ photos: [] });
		this.setState({ albums: [] });

       Axios.post(BASE_URL_FEED + "/api/feed/searchByHashTags/",helpDTO)
        .then((res) => {
            this.setState({ photos: res.data });
            this.setState({ hashtags: "" });
        })
        .catch((err) => {
            console.log(err);
        });
		Axios.post(BASE_URL_FEED + "/api/albumFeed/searchByHashTags/",helpDTO)
        .then((res) => {
            this.setState({ albums: res.data });
            this.setState({ hashtags: "" });
        })
        .catch((err) => {
            console.log(err);
        });
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

             <section id="hero" className="d-flex align-items-top">
				<div className="container">
                <div className="container" style={{ marginTop: "10rem", marginRight: "10rem" }}>
						<table className="table" style={{ width: "100%" }}>
							<tbody>

								<tr>
                                   
									<td width="150em">
                                        <td>
                                            <div className="control-group">
                                                <div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
                                                    <input className="form-control" id="suggest" ref={this.addressInput} placeholder="Address" style={{ width: '400px' }}/>
                                                </div>
                                                <YMaps
                                                    query={{
                                                        load: "package.full",
                                                        apikey: "b0ea2fa3-aba0-4e44-a38e-4e890158ece2",
                                                        lang: "en_RU",
                                                    }}
                                                >
                                                    <Map
                                                        style={{ display: "none" }}
                                                        state={mapState}
                                                        onLoad={this.onYmapsLoad}
                                                        instanceRef={(map) => (this.map = map)}
                                                        modules={["coordSystem.geo", "geocode", "util.bounds"]}
                                                    ></Map>
                                                </YMaps>
                                            
                                                <div className="text-danger" style={{ display: this.state.addressError }}>
                                                    Address must be entered.
                                                </div>
                                                <div className="text-danger" style={{ display: this.state.addressNotFoundError }}>
                                                    Sorry. Address not found. Try different one.
                                                </div>
                                            </div>
                                        </td>
                                        <td>
                                                <button onClick={this.handleSearchByLocation} className="btn btn-outline-secondary btn-sm" >Search</button>

                                        </td>
									</td>
                                    <td>
                                        <td>
                                            <div className="control-group">
                                                <div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
                                                    <input 
                                                        className="form-control" 
                                                        id="suggest" 
                                                        placeholder="HashTags" 
                                                        style={{ width: '400px' }}
                                                        onChange={this.handleHashTagsChange}/>
                                                </div>
                                            
                                            
                                                <div className="text-danger" style={{ display: this.state.hashtagsError }}>
                                                    Hashtags must be entered.
                                                </div>
                                            </div>
                                        </td>
                                        <td>
                                            <td>
                                                <button onClick={this.handleSearchByHashTags} className="btn btn-outline-secondary btn-sm" >Search</button>
                                            </td>
                                        </td>
									</td>
										
								</tr>
							</tbody>
						</table>
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

export default Search;
import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { Link } from "react-router-dom";
import { BASE_URL_FEED, BASE_URL_STORY } from "../constants.js";
import LikesModal from "../components/Posts/LikesModal"
import DislikesModal from "../components/Posts/DislikesModal"
import CommentsModal from "../components/Posts/CommentsModal"
import ImageUploader from 'react-images-upload';
import Axios from "axios";
import ModalDialog from "../components/ModalDialog";
import AddPostModal from "../components/Posts/AddPostModal";
import WriteCommentModal from "../components/Posts/WriteCommentModal"
import { BASE_URL_USER } from "../constants.js";
import AddHighlightModal from "../components/Posts/AddHighlightModal";
import AddStoryToHighlightModal from "../components/Posts/AddStoryToHighlightModal";
import IconTabsProfile from "../components/Posts/IconTabsProfile"
import AddCollectionModal  from "../components/Posts/AddCollectionModal";
import AddPostToCollection from "../components/Posts/AddPostToCollection";
import { BASE_URL_USER_INTERACTION } from "../constants.js";
import WriteCommentAlbumModal from "../components/Posts/WriteCommentAlbumModal"
import AddTagsModal from "../components/Posts/AddTagsModal";
import ConvertVideo from "react-convert-image";
import AddStoryAlbumToHighlightModal from "../components/Posts/AddStoryAlbumToHighlightModal";
 
class ProfilePage extends React.Component {
	constructor(props) {
		super(props);

		this.onDrop = this.onDrop.bind(this);
		this.addressInput = React.createRef();
		this.handleAddProfileImage = this.handleAddProfileImage.bind(this);

	}
	state = {
		id: "",
		username: "",
		name: "",
		lastName : "",
		email: "",
		phoneNumber: "",
		gender : "Female",
		dateOfBirth : "",
		webSite : "",
		biography : "",
		private : true,
		profilePicture : [],
		numberPosts: 0,
		numberFollowing: 0,
		numberFollowers: 0,
		highlihts: [],
		photos: [],
		pictures: [],
		videos : "",
		video : "",
		picture: "",
		hiddenOne: true,
		hiddenMultiple: true,
		noPicture : true,
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
		description : "",
		hashtags :"",
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
		showWriteCommentModalAlbum : false,
		selectedFile : "",
		loaded : "",
		followingUsers : [],
		storyAlbums : [],
		showTagsModal : false,
		taggedOnPost : [],
		stories: ["blob:http://localhost:3000/ac876899-5147-482c-9086-998ee05c765f"],
		urlVideo : "",
		highlightsAlbums : [],
		showAddStoryAlbumToHighLightModal : false,
		storiesForHightlihtAlbum : [],
		hiddenStoriesForHighlightalbum : [],
		showAddHighLightAlbumModal : false
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
			pictures: [],
		});
		this.setState({
			pictures: this.state.pictures.concat(picture),
		});

		let pomoc = picture.length;
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
		


	}
	handleAddProfileImage(picture) {
		let userid = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		this.setState({
			profilePicture: this.state.profilePicture.concat(picture),
		});
		this.testProfileImage(picture, userid);

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
	testProfileImage(pic,userId) {
		
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
		fetch(BASE_URL_USER + "/api/user/profileImage/"+userId , options);
	}
	testStory(pic,userId, storyId) {

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
		fetch(BASE_URL_STORY + "/api/image/"+userId+"/"+storyId , options);
	}


	componentDidMount() {

		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		Axios.get(BASE_URL_USER + "/api/" + id)
				.then((res) => {
					if (res.status === 401) {
						this.setState({ errorHeader: "Bad credentials!", errorMessage: "Wrong username or password.", hiddenErrorAlert: false });
					} else if (res.status === 500) {
						this.setState({ errorHeader: "Internal server error!", errorMessage: "Server error.", hiddenErrorAlert: false });
					} else {
						this.setState({
							id: res.data.Id,
							username : res.data.ProfileInformation.Username,
							name: res.data.ProfileInformation.Name,
							lastName : res.data.ProfileInformation.LastName,
							email : res.data.ProfileInformation.Email,
							phoneNumber : res.data.ProfileInformation.PhoneNumber,
							gender : res.data.ProfileInformation.Gender,
							dateOfBirth  : res.data.ProfileInformation.DateOfBirth,
							webSite : res.data.WebSite,
							biography : res.data.Biography,
							private : res.data.Private
						});
					}
				})
				.catch ((err) => {
			console.log(err);
		});

		this.handleGetHighlights(id)
		this.handleGetPhotos(id)
		this.handleGetAlbums(id)
		this.handleGetStoryAlbums(id)
		this.handleGetStories(id)
		this.handleGetCollections(id)
		// this.handleGetVideos(id)
		this.handleGetHighlightAlbums(id)

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
	handleGetHighlightAlbums = (id) => {
		Axios.get(BASE_URL_STORY + "/api/highlight/user/album/"+id)
			.then((res) => {
				this.setState({ highlightsAlbums: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
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

	handleGetPhotos = (id) => {
		Axios.get(BASE_URL_FEED + "/api/feed/usersImages/"+id)
			.then((res) => {
				this.setState({ photos: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}

	handleGetVideos = (id)=>{
		alert("ubisu se")
		Axios.get(BASE_URL_FEED + "/api/feed/usersVideos/" + id)
			.then((res) => {
				this.setState({ videos: res.data });

			})
			.catch((err) => {
				console.log(err);
			});
	}
	
	handleGetAlbums = (id) => {
		Axios.get(BASE_URL_FEED + "/api/feedAlbum/usersAlbums/"+id)
			.then((res) => {
				this.setState({ albums: res.data });
				console.log("sfsfsf" + res.data)
			})
			.catch((err) => {
				console.log(err);
			});
	}
	handleGetStoryAlbums = (id) => {
		Axios.get(BASE_URL_STORY + "/api/storyAlbum/usersAlbums/"+id)
			.then((res) => {
				this.setState({ storyAlbums: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
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
		this.setState({pictures: []})
	};
	getFollowing = ()=> {
		let help = []

        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = {id: id}
        Axios.post(BASE_URL_USER_INTERACTION + "/api/user/following", dto)
			.then((res) => {

				res.data.forEach((user) => {
					let optionDTO = { id: user.Id, label: user.Username, value: user.Id }
					help.push(optionDTO)
				});
				

				this.setState({ followingUsers: help });
			})
			.catch((err) => {
				console.log(err)
			});
    }
	handleAddTagsModal = ()=> {
		this.getFollowing()
		this.setState({ showTagsModal: true });
	}
	
	handleAddStoryPostCloseFriends =()=>{
		if (this.state.addressInput === "") {
			const storyPostDTO = {
				tagged: [],
				description: this.state.description,
				hashtags: this.state.hashtags,
				location : this.state.addressLocation,
				onlyCloseFriends : true
			};
			this.sendRequestForStory(storyPostDTO);

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
					let storyPostDTO = {
						tagged: [],
						description: this.state.description,
						hashtags: this.state.hashtags,
						location : locationDTO,
						onlyCloseFriends : true
					};
					
					if (this.state.foundLocation === false) {
							this.setState({ addressNotFoundError: "initial" });
					} else {
							this.sendRequestForStory(storyPostDTO);
					}

				});
				

		}

	};
	
	handleAddFeedPost = ()=> {
		var taggedHelp = []
		this.state.taggedOnPost.forEach((user) => {
			taggedHelp.push(user.id)
		});
		if (this.state.addressInput === "") {
			const feedPostDTO = {
				tagged: taggedHelp,
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
						tagged: taggedHelp,
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
		var taggedHelp = []
		this.state.taggedOnPost.forEach((user) => {
			taggedHelp.push(user.id)
		});
		if (this.state.addressInput === "") {
			const storyPostDTO = {
				tagged: taggedHelp,
				description: this.state.description,
				hashtags: this.state.hashtags,
				location : this.state.addressLocation,
				onlyCloseFriends : false
			};
			this.sendRequestForStory(storyPostDTO);

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
					let storyPostDTO = {
						tagged: taggedHelp,
						description: this.state.description,
						hashtags: this.state.hashtags,
						location : locationDTO,
						onlyCloseFriends : false
					};
					if (this.state.foundLocation === false) {
							this.setState({ addressNotFoundError: "initial" });
					} else {
							this.sendRequestForStory(storyPostDTO);
					}

				});
				

		}

	}
	handleAddFeedPostAlbum = ()=> {
		var taggedHelp = []
		this.state.taggedOnPost.forEach((user) => {
			taggedHelp.push(user.id)
		});
		if (this.state.addressInput === "") {
			const feedPostDTO = {
				tagged: taggedHelp,
				description: this.state.description,
				hashtags: this.state.hashtags,
				location : this.state.addressLocation
			};
			this.sendRequestForFeedAlbum(feedPostDTO);

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
						tagged: taggedHelp,
						description: this.state.description,
						hashtags: this.state.hashtags,
						location : locationDTO
					};
					
					if (this.state.foundLocation === false) {
							this.setState({ addressNotFoundError: "initial" });
					} else {
							this.sendRequestForFeedAlbum(feedPostDTO);
					}

				});
				

		}

	}
	handleAddStoryPostAlbum = ()=> {
		var taggedHelp = []
		this.state.taggedOnPost.forEach((user) => {
			taggedHelp.push(user.id)
		});
		if (this.state.addressInput === "") {
			const storyPostDTO = {
				tagged: taggedHelp,
				description: this.state.description,
				hashtags: this.state.hashtags,
				location : this.state.addressLocation,
				onlyCloseFriends : false
			};
			this.sendRequestForAlbumStory(storyPostDTO);

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
					let storyPostDTO = {
						tagged: taggedHelp,
						description: this.state.description,
						hashtags: this.state.hashtags,
						location : locationDTO,
						onlyCloseFriends : false
					};
					
					if (this.state.foundLocation === false) {
							this.setState({ addressNotFoundError: "initial" });
					} else {
							this.sendRequestForAlbumStory(storyPostDTO);
					}

				});
				

		}
	}
	handleAddStoryPostAlbumCloseFriends = ()=> {
		var taggedHelp = []
		this.state.taggedOnPost.forEach((user) => {
			taggedHelp.push(user.id)
		});
		if (this.state.addressInput === "") {
			const storyPostDTO = {
				tagged: taggedHelp,
				description: this.state.description,
				hashtags: this.state.hashtags,
				location : this.state.addressLocation,
				onlyCloseFriends : true
			};
			this.sendRequestForAlbumStory(storyPostDTO);

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
					let storyPostDTO = {
						tagged: taggedHelp,
						description: this.state.description,
						hashtags: this.state.hashtags,
						location : locationDTO,
						onlyCloseFriends : true
					};
					
					if (this.state.foundLocation === false) {
							this.setState({ addressNotFoundError: "initial" });
					} else {
							this.sendRequestForAlbumStory(storyPostDTO);
					}

				});
				

		}
	}


	

	sendRequestForFeed(feedPostDTO) {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
				
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
				
				let userid = localStorage.getItem("userId");
			
				this.state.pictures.forEach((pic) => {
					this.test(pic, userid, feedId);
				});

				this.setState({ showImageModal: false, });
				this.setState({ openModal: true });
				this.setState({ textSuccessfulModal: "You have successfully added feed post." });
				this.handleGetPhotos(id)

			})
			.catch((err) => {
				console.log(err);
			});
	}
	sendRequestForFeedAlbum(feedPostDTO){
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		Axios.post(BASE_URL_FEED + "/api/feedAlbum/" + id, feedPostDTO)
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
				
				let userid = localStorage.getItem("userId");
				
				this.state.pictures.forEach((pic) => {
					this.test(pic, userid, feedId);
				});

				this.setState({ pictures: [] });
				this.setState({ showImageModal: false, });
				this.setState({ openModal: true });
				this.setState({ textSuccessfulModal: "You have successfully added album feed post." });
				this.handleGetAlbums(id)

			})
			.catch((err) => {
				console.log(err);
			});
	}
	sendRequestForStory(storyPostDTO) {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		Axios.post(BASE_URL_STORY + "/api/story/" + id, storyPostDTO)
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
				let storyId = res.data;
			
				let userid = localStorage.getItem("userId");
				let pics = [];

				this.state.pictures.forEach((p) => {
					pics.push(p.name);
				});
				this.state.pictures.forEach((pic) => {
					this.testStory(pic, userid, storyId);
				});

				this.setState({ pictures: [] });
				this.setState({ showImageModal: false, });
				this.setState({ openModal: true });
				this.setState({ textSuccessfulModal: "You have successfully added story post." });
				this.handleGetStories(id)
			})
			.catch((err) => {
				console.log(err);
			});
	}
	sendRequestForAlbumStory(storyPostDTO){
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		Axios.post(BASE_URL_STORY + "/api/storyAlbum/" + id, storyPostDTO)
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
				let storyId = res.data;
				
				let userid = localStorage.getItem("userId");
				let pics = [];

				this.state.pictures.forEach((p) => {
					pics.push(p.name);
				});
				this.state.pictures.forEach((pic) => {
					this.testStory(pic, userid, storyId);
				});

				this.setState({ pictures: [] });
				this.setState({ showImageModal: false, });
				this.setState({ openModal: true });
				this.setState({ textSuccessfulModal: "You have successfully added album story post." });
				this.handleGetStoryAlbums(id)
			})
			.catch((err) => {
				console.log(err);
			});
	}
	
	handleAddHighLightClick = () => {
		this.setState({ showAddHighLightModal: true });
	};
	handleAddHighLightAlbumClick = () => {
		this.setState({ showAddHighLightAlbumModal: true });
	};
    handleAddHighLightModalClose = () => {
		this.setState({ showAddHighLightModal: false });
		this.setState({ showAddHighLightAlbumModal: false });
	};
	handleAddCollectionModalClose = () => {
		this.setState({ showAddCollectionModal: false });
	};
	handleAddHighlight = (name)=> {
		this.setState({highlightNameError: "none"});

        if (name === "") {
			this.setState({ highlightNameError: "initial" });
		} 
        else {
			let highlightDTO = {
                name: name,
            };
			let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

			Axios.post(BASE_URL_STORY + "/api/highlight/"+id, highlightDTO, {
				}).then((res) => {
					
                    this.setState({ showAddHighLightModal: false });
                    this.handleGetHighlights(id);
                    
                })
                .catch((err) => {
                    console.log(err);
                });
		}
	}
	handleAddHighlightAlbum = (name)=> {
		this.setState({highlightNameError: "none"});

        if (name === "") {
			this.setState({ highlightNameError: "initial" });
		} 
        else {
			let highlightDTO = {
                name: name,
            };
			let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

			Axios.post(BASE_URL_STORY + "/api/highlight/album/"+id, highlightDTO, {
				}).then((res) => {
					
                    this.setState({ showAddHighLightAlbumModal: false });
                    this.handleGetHighlightAlbums(id);
                    
                })
                .catch((err) => {
                    console.log(err);
                });
		}
	}
	handleAddCollection = (name)=> {
		this.setState({collectionNameError: "none"});

        if (name === "") {
			this.setState({ collectionNameError: "initial" });
		} 
        else {
			let collectionDTO = {
                name: name,
            };
			let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

			Axios.post(BASE_URL_FEED + "/api/collection/"+id, collectionDTO, {
				}).then((res) => {
					
                    this.setState({ showAddCollectionModal: false });
                    this.handleGetCollections(id);
                })
                .catch((err) => {
                    console.log(err);
                });
		}
	}
	handleOpenAddStoryToHighlightModal = (storyId)=> {
		this.setState({ showAddStoryAlbumToHighLightModal: true });
		this.setState({ selectedStoryId: storyId });
	}
	handleOpenAddStoryAlbumToHighlightModal = (highlightId)=> {
		this.setState({ showAddStoryAlbumToHighLightModal: true });
		this.setState({ selectedStoryId: highlightId });
	}
	handleOpenAddPostToCollectionModal = (postId)=> {
		this.setState({ showAddPostToCollection: true });
		this.setState({ selectedPostId: postId });
	}
	handleAddStoryToHighlightModalClose = ()=> {
		this.setState({ showAddStoryToHighLightModal: false });
	}
	handleAddStoryAlbumToHighlightModalClose = ()=> {
		this.setState({ showAddStoryAlbumToHighLightModal: false });
	}
	handleAddPostToCollectionModalClose = ()=> {
		this.setState({ showAddPostToCollection: false });
	}
	addStoryToHighlight = (highlightId) => {
		let storyHighlightDTO = {
			StoryId : this.state.selectedStoryId,
			HighlightId : highlightId
		}
		Axios.post(BASE_URL_STORY + "/api/highlight/addStory/", storyHighlightDTO, {
		}).then((res) => {
			
			this.setState({ showAddHighLightModal: false });
			let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
			this.handleGetHighlights(id);
			this.setState({ textSuccessfulModal: "You have successfully added story to highlight." });
			this.setState({ openModal: true });
			this.setState({ showAddStoryToHighLightModal: false });

		})
		.catch((err) => {
			console.log(err);
		});
	}
	addStoryAlbumToHighlight = (highlightId) => {
		let storyHighlightDTO = {
			StoryId : this.state.selectedStoryId,
			HighlightId : highlightId
		}
		Axios.post(BASE_URL_STORY + "/api/highlight/addStoryAlbum/", storyHighlightDTO, {
		}).then((res) => {
			
			this.setState({ showAddHighLightModal: false });
			this.setState({ showAddHighLightAlbumModal: false });

			let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
			this.handleGetHighlightAlbums(id);
			this.setState({ textSuccessfulModal: "You have successfully added story album to highlight." });
			this.setState({ openModal: true });
			this.setState({ showAddStoryAlbumToHighLightModal: false });

		})
		.catch((err) => {
			console.log(err);
		});
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
	seeStoriesInHighlightAlbum = (stories)=> {
		this.setState({ hiddenStoriesForHighlightAlbum: false });
		this.setState({storiesForHightlihtAlbum : stories})
		console.log(stories)
	}
	seePostsInCollection = (posts)=> {
		this.setState({ hiddenStoriesForCollection: false });
		this.setState({postsForCollection : posts})
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
	handleWriteCommentModal = (postId)=>{
		this.setState({ selectedPostId: postId });
		this.setState({showWriteCommentModal : true});
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
	onChangeHandler = (event) => {
        this.setState({
            selectedFile: event.target.files[0],
            loaded: 0,
        });
        console.log(event.target.files[0]);
    };

    handleSubmit = (event) => {
		
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		event.preventDefault();
        const formData = new FormData();
        const { selectedFile } = this.state;	
		formData.append('inputFile', selectedFile);
		Axios.post(BASE_URL_FEED + "/api/video/" + id, formData)
			.then((res) => {

				this.setState({ showImageModal: false, });
				this.setState({ openModal: true });
				this.setState({ textSuccessfulModal: "You have successfully added video feed post." });
				// this.handleGetVideos(id)

			})
			.catch((err) => {
				console.log(err);
			});
        event.preventDefault();
        
    };
	handleTagsModalClose = () =>{
		this.setState({ showTagsModal: false });

	}
	handleChangeTags = (event) => {
	
		let optionDTO = { id: event.value, label: event.label, value: event.value }
		let helpDto = this.state.taggedOnPost.concat(optionDTO)
		
		this.setState({ taggedOnPost: helpDto });

		const newList2 = this.state.followingUsers.filter((item) => item.Id !== event.value);
		this.setState({ followingUsers: newList2 });		
	};
	handleConvertedImage = (converted) => {
		console.log(converted)
		alert("helo")
		alert(converted)

		/*	converted = converted.replace("webp", "jpg");
			console.log(converted)
			let hh = this.state.pictures;
			hh.push(converted)
			console.log(hh)
			this.setState({
				pictures: hh,
			});*/

	}
	render() {
		return (
			<React.Fragment>
				<TopBar />
				<Header />
			

				<section id="hero" className="d-flex align-items-top">
				< div >
                <form onSubmit={this.handleSubmit}>
                    <label>
                        Upload a file: <br /><br />
                        <input type="file" name="file" onChange={this.onChangeHandler} />
                    </label>
                    <br /><br />
                    <button type="submit">
                        Upload
                    </button>
                </form >
            </div >
				<div className="container">
				<div className="d-flex align-items-top">
					<div className="container" style={{ marginTop: "10rem", marginRight: "10rem" }}>
						<table className="table" style={{ width: "100%" }}>
							<tbody>

								<tr>
									<td width="130em">
										<img
											className="img-fluid"
											src={this.state.profilePhoto}
											width="70em"
											alt="description"
										/>
										<ImageUploader
											withIcon={false}
											buttonText='Add profile picture'
											onChange={this.handleAddProfileImage}
											imgExtension={['.jpg', '.gif', '.png', '.gif']}
											withPreview={true}
										/>
									</td>

									<td>
										<div>
											<td>
												<label >{this.state.username}</label>
											</td>
											<td>
												<Link to="/settings" className="btn btn-outline-secondary btn-sm">Edit profile</Link>

											</td>
											
											<td>
												<button onClick={this.handlePostModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" }}>Add new post/video</button>

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
				
				<div>
					<IconTabsProfile
						photos = {this.state.photos}
						videos = {this.state.videos}
						urlVideo = {this.state.urlVideo}
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
						handleOpenAddStoryToHighlightModal = {this.handleOpenAddStoryToHighlightModal}
						handleOpenAddStoryAlbumToHighlightModal = {this.handleOpenAddStoryAlbumToHighlightModal}
						

						storyAlbums = {this.state.storyAlbums}

						handleAddHighLightClick = {this.handleAddHighLightClick}
						highlights = {this.state.highlights}
						seeStoriesInHighlight = {this.seeStoriesInHighlight}
						storiesForHightliht= {this.state.storiesForHightliht}
						hiddenStoriesForHighlight = {this.state.hiddenStoriesForHighlight}

						handleAddHighLightAlbumClick = {this.handleAddHighLightAlbumClick}
						highlightsAlbums = {this.state.highlightsAlbums}
						seeStoriesInHighlightAlbum = {this.seeStoriesInHighlightAlbum}
						storiesForHightlihtAlbum= {this.state.storiesForHightlihtAlbum}
						hiddenStoriesForHighlightalbum = {this.state.hiddenStoriesForHighlightAlbum}
						
						handleAddCollectionClick = {this.handleAddCollectionClick}
						collections = {this.state.collections}
						seePostsInCollection = {this.seePostsInCollection}
						postsForCollection = {this.state.postsForCollection}
						hiddenStoriesForCollection = {this.state.hiddenStoriesForCollection}
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
                    <ModalDialog
						show={this.state.openModal}
						onCloseModal={this.handleModalClose}
						header="Successful"
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
						handleAddStoryPostCloseFriends = {this.handleAddStoryPostCloseFriends}
						handleAddFeedPostAlbum = {this.handleAddFeedPostAlbum}
						handleAddStoryPostAlbum= {this.handleAddStoryPostAlbum}
						handleAddStoryPostAlbumCloseFriends = {this.handleAddStoryPostAlbumCloseFriends}
						addressNotFoundError = {this.state.addressNotFoundError}
						handleDescriptionChange = {this.handleDescriptionChange}
						handleHashtagsChange = {this.handleHashtagsChange}
						handleAddTagsModal = {this.handleAddTagsModal}
					/>
					 <AddHighlightModal
                          
                            highlightNameError={this.state.highlightNameError}
                        
					        show={this.state.showAddHighLightModal}
					        onCloseModal={this.handleAddHighLightModalClose}
					        header="Add new highlight"
                            handleAddHighlight={this.handleAddHighlight}
				        />
					<AddHighlightModal
                          
						  highlightNameError={this.state.highlightNameError}
					  
						  show={this.state.showAddHighLightAlbumModal}
						  onCloseModal={this.handleAddHighLightModalClose}
						  header="Add new highlight album"
						  handleAddHighlight={this.handleAddHighlightAlbum}
					  />
						<AddCollectionModal
                          collectionNameError={this.state.collectionNameError}
                        
						  show={this.state.showAddCollectionModal}
						  onCloseModal={this.handleAddCollectionModalClose}
						  header="Add new collection"
						  handleAddCollection={this.handleAddCollection}
					  />
					<AddStoryToHighlightModal
                          
					  
						  show={this.state.showAddStoryToHighLightModal}
						  onCloseModal={this.handleAddStoryToHighlightModalClose}
						  header="Add story to highlight"
						  addStoryToHighlight={this.addStoryToHighlight}
						  highlights = {this.state.highlights}
					  />
					  <AddStoryAlbumToHighlightModal
                          
					  
						  show={this.state.showAddStoryAlbumToHighLightModal}
						  onCloseModal={this.handleAddStoryAlbumToHighlightModalClose}
						  header="Add story album to highlight album"
						  addStoryAlbumToHighlight={this.addStoryAlbumToHighlight}
						  highlightsAlbums = {this.state.highlightsAlbums}
					  />
					  <AddPostToCollection
                          
					  
						  show={this.state.showAddPostToCollection}
						  onCloseModal={this.handleAddPostToCollectionModalClose}
						  header="Add post to collection"
						  addPostToCollection={this.addPostToCollection}
						  collections = {this.state.collections}
					  />
					   <AddTagsModal
                          
					  
						  show={this.state.showTagsModal}
						  onCloseModal={this.handleTagsModalClose}
						  header="Add tags"
						  followingUsers = {this.state.followingUsers}
						  taggedOnPost = {this.state.taggedOnPost}
						  tagUserOnPost={this.tagUserOnPost}
						  handleChangeTags = {this.handleChangeTags}
					  />
                    </div>

			</React.Fragment>
		);
	}
}

export default ProfilePage;
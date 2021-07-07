import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { Link } from "react-router-dom";
import { BASE_URL_FEED, BASE_URL_STORY } from "../constants.js";
import LikesModal from "../components/Posts/LikesModal"
import DislikesModal from "../components/Posts/DislikesModal"
import CommentsModal from "../components/Posts/CommentsModal"
import ImageUploader from 'react-images-upload';
import playerLogo from "../static/coach.png";
import Axios from "axios";
import ModalDialog from "../components/ModalDialog";
import AddPostModal from "../components/Posts/AddPostModal";
import AddVideoPostModal from "../components/Posts/AddVideoPostModal";
import WriteCommentModal from "../components/Posts/WriteCommentModal"
import AddHighlightModal from "../components/Posts/AddHighlightModal";
import AddStoryToHighlightModal from "../components/Posts/AddStoryToHighlightModal";
import IconTabsProfile from "../components/Posts/IconTabsProfile"
import AddCollectionModal  from "../components/Posts/AddCollectionModal";
import AddPostToCollection from "../components/Posts/AddPostToCollection";
import EditCampaignModal from "../components/EditCampaignModal"
import WriteCommentAlbumModal from "../components/Posts/WriteCommentAlbumModal"
import AddTagsModal from "../components/Posts/AddTagsModal";
import AddStoryAlbumToHighlightModal from "../components/Posts/AddStoryAlbumToHighlightModal";
import { BASE_URL } from "../constants.js";
import VerifyModal from "../pages/VerifyModal";
import AddCampaignModal from "../components/AddCampaignModal";
import AddVideoCampaignModal from "../components/AddVideoCampaignModal";
import OneTimeCampaignModal from "../components/OneTimeCampaignModal";
import MultipleTimeCampaignModal from "../components/MultipleTimeCampaignModal";
import AddInfluencerModal from "../components/AddInfluencerModal";
import TargetGroupModal from "../components/TargetGroupModal";
import {GoVerified} from 'react-icons/go'
import getAuthHeader from "../GetHeader";

class ProfilePage extends React.Component {
	constructor(props) {
		super(props);

		this.onDrop = this.onDrop.bind(this);
		this.onDropCampaign = this.onDropCampaign.bind(this);
		this.addressInput = React.createRef();
		this.addressInputCampaign = React.createRef();

		this.handleAddProfileImage = this.handleAddProfileImage.bind(this);

	}
	state = {
		id: "",
		username: "",
		name: "",
		lastName : "",
		email: "",
		phoneNumber: "",
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
		campaigns : [],
		pictures: [],
		videos : [],
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
		showVerifyModal : false,
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
		showAddCollectionAlbumModal : false,
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
		showAddHighLightAlbumModal : false,
		collectionAlbums : [],
		hiddenPostsForCollectionAlbums: false,
		postsForCollectionAlbum : [],
		showAddAlbumToCollectionAlbum : false,
		followingsThatAllowTags : [],
		category : "",
		showCampaignModal : false,
		showVideoCampaignModal : false,
		link : "",
		showOneTimeCampaignModal : false,
		showMultipleTimeCampaignModal : false,
		showInfluencersModal : false,
		influencers : [],
		choosenInfluencers : [],
		showTargetGroupModal : false,
		selectedGender : "MALE",
        selectedDateOne : "",
        selectedDateTwo : "",
		campaignTime : "",
		campaignDate : "",
		campaignLink : "",
		campaignDescription : "",
		campaignId : "",
		targetGroup : {},
		showEditCampaignModal : false,
		isAgent : false,
		isInfluencer : false,
		oneTimeCampaignsInfluencer : [],
        multipleCampaignsInfluencer : [],
		campaignStartTime : "",
		campaignEndTime : "",
		campaignDesiredNumber : "",
		campaignType : "",
		isVerified : false,
		categoryString : "",
		
	}
	hasRole = (reqRole) => {
		let roles = JSON.parse(localStorage.getItem("keyRole"));
		if (roles === null) return false;

		if (reqRole === "*") return true;

	
		if (roles.trim() === reqRole.trim()) 
		{
			return true;
		}
		return false;
	};
	onYmapsLoad = (ymaps) => {
		this.ymaps = ymaps;
		new this.ymaps.SuggestView(this.addressInput.current, {
			provider: {
				suggest: (request, options) => this.ymaps.suggest(request),
			},
		});
		
		
	};
	
	onDrop1(picture) {

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
			
		}


	}
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
	onDropCampaign(picture) {

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
			});
			
			
		}
		

	}
	handleAddProfileImage(picture) {
		let userid = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		this.setState({
			profilePicture: this.state.profilePicture.concat(picture),
		});
		this.testProfileImage(picture, userid);

	}

	

	testVideo(pic,userId, feedId) {
		const formData = new FormData();

		formData.append("file", pic);
		formData.append("test", "StringValueTest");

		const options = {
			method: "POST",
			body: formData

		};
		fetch(BASE_URL + "/api/feedPosts/api/image/"+userId+"/"+feedId , options,  {  headers: { Authorization: getAuthHeader() } });
	}
	testVideoCampaign(pic,userId, campaignId) {
		const formData = new FormData();

		formData.append("file", pic);
		formData.append("test", "StringValueTest");

		const options = {
			method: "POST",
			body: formData

		};
		fetch(BASE_URL + "/api/campaign/api/image/"+userId+"/"+campaignId , options,  {  headers: { Authorization: getAuthHeader() } });
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

	
		fetch( BASE_URL + "/api/feedPosts/api/image/"+userId+"/"+feedId, options,  {  headers: { Authorization: getAuthHeader() } });
	}
	testCampaign(pic,userId, campaignId) {
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

	
		fetch( BASE_URL + "/api/campaign/api/image/"+userId+"/"+campaignId, options,  {  headers: { Authorization: getAuthHeader() } });
	}
	
	testVerification(pic,userId, requestId) {
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

	
		fetch( BASE_URL + "/api/users/api/image/"+userId+"/"+requestId, options,  {  headers: { Authorization: getAuthHeader() } });
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
		fetch(BASE_URL + "/api/users/api/user/profileImage/"+userId , options,  {  headers: { Authorization: getAuthHeader() } });
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
		fetch(BASE_URL + "/api/storyPosts/api/image/"+userId+"/"+storyId , options,  {  headers: { Authorization: getAuthHeader() } });
	}


	componentDidMount() {

		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		Axios.get(BASE_URL + "/api/users/api/" + id)
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
							dateOfBirth  : res.data.ProfileInformation.DateOfBirth,
							webSite : res.data.WebSite,
							biography : res.data.Biography,
							private : res.data.Private,
							isVerified : res.data.Verified,
							categoryString : res.data.Category, 
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
		this.handleGetHighlightAlbums(id)
		this.handleGetCollectionAlbums(id)
		this.getFollowingThatCanBeTagged()
		this.handeleGetCampaigns(id)
		this.handeleGetMultipleCampaignsInfluencer(id)
		this.handeleGetOneTimeCampaignsInfluencer(id)
		this.handleGetCategoryUser(id)

	}
	handeleGetMultipleCampaignsInfluencer = (id)=> {
		Axios.get(BASE_URL + "/api/campaign/promoteMultiple/"+id,  {  headers: { Authorization: getAuthHeader() } })
		.then((res) => {
			this.setState({ multipleCampaignsInfluencer: res.data });
		})
		.catch((err) => {
			console.log(err);
		});	
	}
	handeleGetOneTimeCampaignsInfluencer = (id)=> {
		Axios.get(BASE_URL + "/api/campaign/promoteOneTime/"+id,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ oneTimeCampaignsInfluencer: res.data });
			})
			.catch((err) => {
				console.log(err);
			});	
	}
	handleGetCategoryUser = (id)=> {

		var role = this.hasRole("AGENT")
		this.setState({isAgent : role});

		Axios.get(BASE_URL + "/api/users/api/user/username/category/"+id,  {  headers: { Authorization: getAuthHeader() } })
				.then((res) => {
					
					if (res.data.trim()!=="not") {
						this.setState({ isInfluencer: true });

						return true;
					}
					return false;
				})
				.catch((err) => {
					console.log(err);
				});

	}
	handleGetStories = (id)=> {
		Axios.get(BASE_URL + "/api/storyPosts/api/story/user/"+id,  {  headers: { Authorization: getAuthHeader() } })
		.then((res) => {
			this.setState({ stories: res.data });
		})
		.catch((err) => {
			console.log(err);
		});
	}

	handleGetHighlights = (id) => {
		Axios.get(BASE_URL + "/api/storyPosts/api/highlight/user/"+id,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ highlights: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}
	handleGetHighlightAlbums = (id) => {
		Axios.get(BASE_URL + "/api/storyPosts/api/highlight/user/album/"+id,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ highlightsAlbums: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}
	
	handleGetCollectionAlbums = (id) => {
		Axios.get(BASE_URL + "/api/feedPosts/api/collection/user/album/"+id,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ collectionAlbums: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}
	handleGetCollections = (id) => {
		Axios.get(BASE_URL + "/api/feedPosts/api/collection/user/"+id,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ collections: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
	}

	handleGetPhotos = (id) => {
		Axios.get(BASE_URL + "/api/feedPosts/api/feed/usersImages/"+id,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ photos: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		
			
	}
	handeleGetCampaigns = (id) => {
		Axios.get(BASE_URL + "/api/campaign/api/getUsersCampaigns/"+id,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ campaigns: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		
			
	}
	handleGetVideos = (id)=>{
		Axios.get(BASE_URL + "/api/feedPosts/api/feed/usersVideos/" + id,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ videos: res.data });

			})
			.catch((err) => {
				console.log(err);
			});
	}
	
	handleGetAlbums = (id) => {
		Axios.get(BASE_URL + "/api/feedPosts/api/feedAlbum/usersAlbums/"+id,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ albums: res.data });
				console.log("sfsfsf" + res.data)
			})
			.catch((err) => {
				console.log(err);
			});
	}
	handleGetStoryAlbums = (id) => {
		Axios.get(BASE_URL + "/api/storyPosts/api/storyAlbum/usersAlbums/"+id,  {  headers: { Authorization: getAuthHeader() } })
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
	handleLinkChange = (event) => {
		this.setState({ link: event.target.value });
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
	handleVideoPostModalClose = () => {
		this.setState({ showVideoModal: false });
		this.setState({selectedFile: ""})
	};
	handleVideoPostModalOpen = () => {
		this.setState({ showVideoModal: true });
		this.setState({selectedFile: ""})
	};
	handleVerifyModalOpen = () => {
		this.setState({ showVerifyModal: true });
		this.setState({pictures: []})
	};
	handleAddCampaignModal = ()=>{
		this.setState({ showCampaignModal: true });
	}
	handleAddCampaignModalClose = ()=>{
		this.setState({ showCampaignModal: false });
	}
	handleAddVideoCampaignModal = ()=>{
		this.setState({ showVideoCampaignModal: true });
		this.setState({selectedFile: ""})

	}
	handleAddVideoCampaignModalClose = ()=>{
		this.setState({ showVideoCampaignModal: false });
		this.setState({selectedFile: ""})

	}
	handleOneTimeCampaignModalOpen = ()=>{
		this.setState({ showOneTimeCampaignModal: true });
	}
	handleOneTimeCampaignModalClose = ()=>{
		this.setState({ showOneTimeCampaignModal: false });
	}
	handleMultipleTimeCampaignModalOpen = ()=>{
		this.setState({ showMultipleTimeCampaignModal: true });
	}
	handleMultipleTimeCampaignModalClose = ()=>{
		this.setState({ showMultipleTimeCampaignModal: false });
	}
	handleVerifyModalClose = () => {
		this.setState({ showVerifyModal: false });
	}
	getFollowing = ()=> {
		let help = []

        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = {id: id}
        Axios.post(BASE_URL + "/api/userInteraction/api/user/following", dto,  {  headers: { Authorization: getAuthHeader() } })
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
	getFollowingThatCanBeTagged = ()=> {
		let help = []

        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = {id: id}
        Axios.post(BASE_URL + "/api/userInteraction/api/user/following/tagged", dto,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				res.data.forEach((user) => {
					let optionDTO = { id: user.Id, label: user.Username, value: user.Id }
					help.push(optionDTO)
				});
				
				this.setState({ followingsThatAllowTags: help });
			})
			.catch((err) => {
				console.log(err)
			});
    }
	
	getInfluencers = ()=> {
		let help = []

        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = {id: id}
        Axios.post(BASE_URL + "/api/userInteraction/api/user/following/category", dto,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {

				res.data.forEach((user) => {
					let optionDTO = { id: user.Id, label: user.Username, value: user.Id }
					help.push(optionDTO)
				});
				

				this.setState({ influencers: help });
			})
			.catch((err) => {
				console.log(err)
			});
    }
	handleAddTagsModal = ()=> {
		this.getFollowing()
		this.setState({ showTagsModal: true });
	}
	handleAddInfluencersModal = ()=> {
		this.getInfluencers()
		this.setState({ showInfluencersModal: true });
	}
	
	handleDefineTargetGroupModal = ()=> {
		this.setState({ showTargetGroupModal: true });
	}
	handleAddOneTimeCampaignModal =()=>{
		this.setState({ showOneTimeCampaignModal: true });

	}
	handleAddMultipleTimeCampaignModal =() =>{
		this.setState({ showMultipleTimeCampaignModal: true });

	}
	
	getpartnershipsRequests = ()=> {
		var choosenInfluencersHelp = []
		this.state.choosenInfluencers.forEach((user) => {
			choosenInfluencersHelp.push(user.id)
		});
		return choosenInfluencersHelp
	}
	handleAddOneTimeCampaign =(date,time, type)=>{

        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
		let partnershipsRequestsList = this.getpartnershipsRequests();
		const campaignDTO = {
			User : id,
			TargetGroup : this.state.targetGroup,
			Link : this.state.link,
			Date : date,
			Time : time,
			Description : this.state.description,
			PartnershipsRequests : partnershipsRequestsList,
			Type : type
		}
		Axios.post(BASE_URL + "/api/campaign/oneTimeCampaign/", campaignDTO,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ showOneTimeCampaignModal: false, showCampaignModal : false });
				this.setState({ openModal: true });
				this.setState({ textSuccessfulModal: "Campaign is successfully created." });
				
				let campaignId = res.data;
			
				let userid = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
				let pics = [];

				this.state.pictures.forEach((p) => {
					pics.push(p.name);
				}); 
				this.state.pictures.forEach((pic) => {
					this.testCampaign(pic, userid, campaignId);
				});
				
				if(this.state.selectedFile != ""){
				this.testVideoCampaign(this.state.selectedFile, userid, campaignId)
				}
				this.setState({selectedFile : ""});
				this.setState({ pictures: [] });
			})
			.catch((err) => {
				console.log(err);
			});

	}
	handleAddMultipleTimeCampaign =(startDate,endDate,numberOfRepetitions, type) =>{

		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
		let partnershipsRequestsList = this.getpartnershipsRequests();
		const campaignDTO = {
			User : id,
			TargetGroup : this.state.targetGroup,
			Link : this.state.link,
			StartTime : startDate,
			EndTime : endDate,
			Description : this.state.description,
			PartnershipsRequests : partnershipsRequestsList,
			DesiredNumber : numberOfRepetitions,
			Type : type
		}
		Axios.post(BASE_URL + "/api/campaign/multipleTimeCampaign/", campaignDTO,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ showOneTimeCampaignModal: false, showCampaignModal : false });
				this.setState({ openModal: true });
				this.setState({ textSuccessfulModal: "Campaign is successfully created." });
				
				let campaignId = res.data;
			
				let userid = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
				let pics = [];

				this.state.pictures.forEach((p) => {
					pics.push(p.name);
				}); 
				this.state.pictures.forEach((pic) => {
					this.testCampaign(pic, userid, campaignId);
				});
				if(this.state.selectedFile != ""){
				this.testVideoCampaign(this.state.selectedFile, userid, campaignId)
				}
				this.setState({selectedFile : ""});
				this.setState({ pictures: [] });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showMultipleTimeCampaignModal: false, showCampaignModal : false });

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

	handleSendRequestVerification = (name,surname,category)=> {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
			const verificationDTO = {
				Id : id,
				Name: name,
				LastName: surname,
				Category : category
				
			};
			this.sendRequestForVerification(verificationDTO);

		
			
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


	sendRequestForVerification(verificationDTO) {
				
		Axios.post(BASE_URL + "/api/users/api/verificationRequest", verificationDTO,  {  headers: { Authorization: getAuthHeader() } })
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
				let requestId = res.data;
				
				let userid = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
			
				this.state.pictures.forEach((pic) => {
					this.testVerification(pic, userid, requestId);
				});
				
				this.setState({ selectedFile : ""});
				this.setState({ pictures: [] });
				this.setState({ showVerifyModal: false, });
				this.setState({ showVideoModal: false, });
				this.setState({ openModal: true });
				this.setState({ textSuccessfulModal: "You successfully sent request." });

			})
			.catch((err) => {
				console.log(err);
			});
	}


	sendRequestForFeed(feedPostDTO) {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
				
		Axios.post(BASE_URL + "/api/feedPosts/api/feed/" + id, feedPostDTO,  {  headers: { Authorization: getAuthHeader() } })
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
				
				let userid = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
			
				this.state.pictures.forEach((pic) => {
					this.test(pic, userid, feedId);
				});
				if(this.state.selectedFile != ""){
				this.testVideo(this.state.selectedFile, userid, feedId)
				}
				this.setState({ selectedFile : ""});
				this.setState({ pictures: [] });
				this.setState({ showImageModal: false, });
				this.setState({ showVideoModal: false, });
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
		Axios.post(BASE_URL + "/api/feedPosts/api/feedAlbum/" + id, feedPostDTO,  {  headers: { Authorization: getAuthHeader() } })
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

		Axios.post(BASE_URL + "/api/storyPosts/api/story/" + id, storyPostDTO,  {  headers: { Authorization: getAuthHeader() } })
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
			
				let userid = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
				let pics = [];

				this.state.pictures.forEach((p) => {
					pics.push(p.name);
				}); 
				this.state.pictures.forEach((pic) => {
					this.testStory(pic, userid, storyId);
				});
				if(this.state.selectedFile != ""){
				this.testStory(this.state.selectedFile, userid, storyId)
				}
				this.setState({selectedFile : ""});
				this.setState({ pictures: [] });
				this.setState({ showImageModal: false, });
				this.setState({ showVideoModal: false, });
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
		Axios.post(BASE_URL + "/api/storyPosts/api/storyAlbum/" + id, storyPostDTO)
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
	handleAddCollectionClick = () => {
		this.setState({ showAddCollectionModal: true });
	};
	handleAddCollectionAlbumClick = () => {
		this.setState({ showAddCollectionAlbumModal: true });
	};
    handleAddHighLightModalClose = () => {
		this.setState({ showAddHighLightModal: false });
		this.setState({ showAddHighLightAlbumModal: false });
	};
	handleAddCollectionModalClose = () => {
		this.setState({ showAddCollectionModal: false });
		this.setState({ showAddCollectionAlbumModal: false });

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

			Axios.post(BASE_URL + "/api/storyPosts/api/highlight/"+id, highlightDTO, {
				},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
					
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

			Axios.post(BASE_URL + "/api/storyPosts/api/highlight/album/"+id, highlightDTO, {
				},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
					
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

			Axios.post(BASE_URL + "/api/feedPosts/api/collection/"+id, collectionDTO, {
				},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
					
                    this.setState({ showAddCollectionModal: false });
                    this.handleGetCollections(id);
                })
                .catch((err) => {
                    console.log(err);
                });
		}
	}
	
	handleAddCollectionAlbum = (name)=> {
		this.setState({collectionNameError: "none"});

        if (name === "") {
			this.setState({ collectionNameError: "initial" });
		} 
        else {
			let collectionDTO = {
                name: name,
            };
			let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

			Axios.post(BASE_URL + "/api/feedPosts/api/collection/album/"+id, collectionDTO, {
				},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
					
                    this.setState({ showAddCollectionModal: false });
					this.setState({ showAddCollectionAlbumModal: false });

                    this.handleGetCollectionAlbums(id);
                })
                .catch((err) => {
                    console.log(err);
                });
		}
	}
	handleOpenAddStoryToHighlightModal = (storyId)=> {
		this.setState({ showAddStoryToHighLightModal: true });
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
	
	handleOpenAddAlbumToCollectionAlbumModal = (postId)=> {
		this.setState({ showAddAlbumToCollectionAlbumToCollection: true });
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
		this.setState({ showAddAlbumToCollectionAlbumToCollection: false });
	}
	addStoryToHighlight = (highlightId) => {
		let storyHighlightDTO = {
			StoryId : this.state.selectedStoryId,
			HighlightId : highlightId
		}
		Axios.post(BASE_URL + "/api/storyPosts/api/highlight/addStory/", storyHighlightDTO,  {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
			
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
		Axios.post(BASE_URL + "/api/storyPosts/api/highlight/addStoryAlbum/", storyHighlightDTO, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
			
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
		Axios.post(BASE_URL + "/api/feedPosts/api/collection/addPost/", postCollectionDTO, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
			
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
	
	addAlbumToCollectionAlbum = (collectionId) => {
		let postCollectionDTO = {
			PostId : this.state.selectedPostId,
			CollectionId : collectionId
		}
		Axios.post(BASE_URL + "/api/feedPosts/api/collection/album/addPost/", postCollectionDTO, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
			
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
	seeStoriesInHighlight = (stories)=> {
		this.setState({ hiddenStoriesForHighlight: false });
		this.setState({storiesForHightliht : stories})
	}
	seeStoriesInHighlightAlbum = (stories)=> {
		this.setState({ hiddenStoriesForHighlightAlbum: false });
		this.setState({storiesForHightlihtAlbum : stories})
	}
	seePostsInCollection = (posts)=> {
		this.setState({ hiddenStoriesForCollection: false });
		this.setState({postsForCollection : posts})
	}
	
	seePostsInCollectionAlbum = (albums)=> {
		this.setState({ hiddenPostsForCollectionAlbums: false });
		this.setState({postsForCollectionAlbum : albums})
	}
	handleLikesModalOpen = (postId)=> {
		Axios.get(BASE_URL + "/api/feedPosts/api/feed/likes/"+postId,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ peopleLikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showLikesModal: true });    
	}
	handleDislikesModalOpen = (postId)=> {
		Axios.get(BASE_URL + "/api/feedPosts/api/feed/dislikes/"+postId,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ peopleDislikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showDislikesModal: true });    
	}
	handleCommentsModalOpen = (postId)=> {
		Axios.get(BASE_URL + "/api/feedPosts/api/feed/comments/"+postId,  {  headers: { Authorization: getAuthHeader() } })
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
		Axios.post(BASE_URL + "/api/feedPosts/api/feed/like/", postReactionDTO, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {

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
		Axios.post(BASE_URL + "/api/feedPosts/api/feed/dislike/", postReactionDTO, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {

			this.setState({ textSuccessfulModal: "You have successfully disliked the photo." });
			this.setState({ openModal: true });

		})
		.catch((err) => {
			console.log(err);
		});
	}
	handleDeleteCampaign =(id, type)=>{
		if (type === "oneTime"){
		Axios.get(BASE_URL + "/api/campaign/api/campaign/delete/"+ id, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
			
			this.setState({ textSuccessfulModal: "You have successfully deleted campaign." });
			this.setState({ openModal: true });
			this.setState({ showWriteCommentModal: false });
		

		})
		.catch((err) => {
			console.log(err);
		});
		let user = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		this.handeleGetCampaigns(user)
	}else if(type === "multiple"){
		Axios.get(BASE_URL + "/api/campaign/api/campaign/delete/multiple/"+ id, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
			
			this.setState({ textSuccessfulModal: "You have successfully deleted campaign." });
			this.setState({ openModal: true });
			this.setState({ showWriteCommentModal: false });
		

		})
		.catch((err) => {
			console.log(err);
		});
		let user = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		this.handeleGetCampaigns(user)
	}
		
	}
	handleCampaignDateChange = (event) =>{
        this.setState({ campaignDate: event.target.value });
        
    }
        handleCampaignTimeChange = (event) =>{
        this.setState({ campaignTime: event.target.value });
        
        };
        handleCampaignDescriptionChange = (event) => {
        this.setState({ campaignDescription: event.target.value });
      };
      handleCampaignLinkChange = (event) => {
        this.setState({ campaignLink: event.target.value });
      };
	  handleCampaignStartTimeChange = (event) =>{
        this.setState({ campaignStartTime: event.target.value });
        
        };
        handleCampaignEndTimeChange = (event) => {
        this.setState({ campaignEndTime: event.target.value });
      };
      handleCampaignDesiredNumberChange = (event) => {
        this.setState({ campaignDesiredNumber: event.target.value });
      };
   
   
	handleEditCampaignModal = (id, type) =>{
		if( type === "oneTime"){
			this.setState({ campaignType: type });
		Axios.get(BASE_URL + "/api/campaign/api/campaign/id/"+id,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
                this.setState({
                campaignDate : res.data.Date,
                campaignLink : res.data.Campaign.Link,
                campaignDescription : res.data.Campaign.Description,
                campaignTime : res.data.Time,
				campaignForEdit : res.data.Id
            });
			})
			.catch((err) => {
				console.log(err)
			});
		this.setState({ showEditCampaignModal: true });
		}else if(type ==="multiple"){
			this.setState({campaignType : "multiple"})
			Axios.get(BASE_URL + "/api/campaign/api/campaign/multiple/id/"+id,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
                this.setState({
                campaignEndTime : res.data.EndTime,
                campaignLink : res.data.Campaign.Link,
                campaignDescription : res.data.Campaign.Description,
                campaignStartTime : res.data.StartTime,
				campaignForEdit : res.data.Id,
				campaignDesiredNumber : res.data.DesiredNumber
            });
			})
			.catch((err) => {
				console.log(err)
			});
		this.setState({ showEditCampaignModal: true });
		}
		
	}
	handleEditCampaignModalClose = () =>{
		this.setState({ showEditCampaignModal: false });
	}
	handleChangeCampaign = (id,date,time,link,des)=>{
		let user = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		let dto = {
			Id : id,
			Link : link,
			Description : des,
			Time : time,
			Date : date,
			User :  user
		}
		Axios.post(BASE_URL + "/api/campaign/api/campaign/update", dto, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
			
			this.setState({ textSuccessfulModal: "You have successfully updated campaign." });
			this.setState({ openModal: true });
			this.setState({ showEditCampaignModal: false });



		})
		.catch((err) => {
			
			console.log(err);
		});
		this.handeleGetCampaigns(user)
	}
	handleChangeMultipleCampaign = (id,start,end,number,link,des)=>{
		let user = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		let dto = {
			Id : id,
			Link : link,
			Description : des,
			StartTime : start,
			EndTime : end,
			DesiredNumber : number,
			User : user,
		}
		Axios.post(BASE_URL + "/api/campaign/api/campaign/multiple/update", dto, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
			
			this.setState({ textSuccessfulModal: "You have successfully updated campaign." });
			this.setState({ openModal: true });
			this.setState({ showWriteCommentModal: false });
			this.setState({ showEditCampaignModal: false });



		})
		.catch((err) => {
			this.setState({ textSuccessfulModal: "Campaign can not be updated. 24 hours have not yet passed since the last modification" });
			this.setState({ openModal: true });
			this.setState({ showEditCampaignModal: false });
			console.log(err);
		});
		this.handeleGetCampaigns(user)
	}
	handleAddComment =(comment) => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

		let commentDTO = {
			PostId : this.state.selectedPostId,
			UserId : id,
			Content : comment

		}
		Axios.post(BASE_URL + "/api/feedPosts/api/feed/comment/", commentDTO, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
			
			this.setState({ textSuccessfulModal: "You have successfully commented the photo." });
			this.setState({ openModal: true });
			this.setState({ showWriteCommentModal: false });


		})
		.catch((err) => {
			console.log(err);
		});
	}
	handleLikesModalOpenAlbum = (postId)=> {
		Axios.get(BASE_URL + "/api/feedPosts/api/albumFeed/likes/"+postId,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ peopleLikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showLikesModal: true });    
	}
	handleDislikesModalOpenAlbum = (postId)=> {
		Axios.get(BASE_URL + "/api/feedPosts/api/albumFeed/dislikes/"+postId,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ peopleDislikes: res.data });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showDislikesModal: true });    
	}
	handleCommentsModalOpenAlbum = (postId)=> {
		Axios.get(BASE_URL + "/api/feedPosts/api/albumFeed/comments/"+postId,  {  headers: { Authorization: getAuthHeader() } })
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
		Axios.post(BASE_URL + "/api/feedPosts/api/albumFeed/comment/", commentDTO, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {
			
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
		Axios.post(BASE_URL + "/api/feedPosts/api/albumFeed/like/", postReactionDTO, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {

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
		Axios.post(BASE_URL + "/api/feedPosts/api/albumFeed/dislike/", postReactionDTO, {
		},  {  headers: { Authorization: getAuthHeader() } }).then((res) => {

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
		
    };

    handleSubmit = (event) => {
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
	
	handleInfluencersModalClose = () =>{
		this.setState({ showInfluencersModal: false });

	}
	handleTargetGroupModalClose = () =>{
		console.log(this.addressInput)
		if (this.state.addressInput === "") {
			const dto = {
				Gender : this.state.selectedGender,
				DateOne : this.state.selectedDateOne,
				DateTwo : this.state.selectedDateTwo,
				Location : this.state.addressLocation,
			}
			this.setState({ targetGroup: dto });

			return dto;

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
					const dto = {
						Gender : this.state.selectedGender,
						DateOne : this.state.selectedDateOne,
						DateTwo : this.state.selectedDateTwo,
						Location : locationDTO
					}
					this.setState({ targetGroup: dto });				
				});
				

				}


		this.setState({ showTargetGroupModal: false });

	}
	handleChangeInfluencers = (event) => {
	
		let optionDTO = { id: event.value, label: event.label, value: event.value }
		let helpDto = this.state.choosenInfluencers.concat(optionDTO)
		
		this.setState({ choosenInfluencers: helpDto });

		const newList2 = this.state.influencers.filter((item) => item.Id !== event.value);
		this.setState({ influencers: newList2 });		
	};
	handleGenderChange=(event) =>{
        this.setState({  selectedGender: event.target.value });
    }
    handleDateOneChange = (event) => {
		this.setState({ selectedDateOne: event.target.value });
	};
    handleDateTwoChange = (event) => {
		this.setState({ selectedDateTwo: event.target.value });
	};
	handleLikesModalOpenCampaign = (likes) => {
		
		this.setState({ peopleLikes: likes });
			
		this.setState({ showLikesModal: true });
	}
	handleDislikesModalOpenCampaign = (dislikes) => {
		
		this.setState({ peopleDislikes: dislikes });
		
		this.setState({ showDislikesModal: true });
	}
	handleCommentsModalOpenCampaign = (comments) => {
		
		this.setState({ peopleComments: comments });
		
		this.setState({ showCommentsModal: true });
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
                                width="100%"
                                alt="description"
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
												<button onClick={this.handlePostModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" }}>Add image post</button>

											</td>
											<td>
												<button onClick={this.handleVideoPostModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" }}>Add video post</button>

											</td>
											<td>
												<button onClick={this.handleVerifyModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" }}>Send verification request</button>

											</td>
											<td hidden={!this.hasRole("AGENT")}>
												<button onClick={this.handleAddCampaignModal} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" }}>Publish image campaign</button>

											</td>
											<td hidden={!this.hasRole("AGENT")}>
												<button onClick={this.handleAddVideoCampaignModal} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" }}>Publish video campaign</button>

											</td>
										</div>
										
										<div>
											<td>
												<label >{this.state.biography}</label>
											</td>
										</div>

										<div hidden={!this.state.isVerified}>
											<td>
												<GoVerified />
											</td>
											<td>
												<label >{this.state.categoryString}</label>
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
						campaigns = {this.state.campaigns}
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

						handleAddCollectionAlbumClick = {this.handleAddCollectionAlbumClick}
						collectionAlbums = {this.state.collectionAlbums}
						seePostsInCollectionAlbum = {this.seePostsInCollectionAlbum}
						hiddenPostsForCollectionAlbums= {this.state.hiddenPostsForCollectionAlbums}
						postsForCollectionAlbum = {this.state.postsForCollectionAlbum}
						handleOpenAddAlbumToCollectionAlbumModal = {this.handleOpenAddAlbumToCollectionAlbumModal}

						handleDeleteCampaign = {this.handleDeleteCampaign}
						handleEditCampaignModal = {this.handleEditCampaignModal}

						isAgent = {this.state.isAgent}
						isInfluencer = {this.state.isInfluencer}

						oneTimeCampaignsInfluencer = {this.state.oneTimeCampaignsInfluencer}
						multipleCampaignsInfluencer = {this.state.multipleCampaignsInfluencer}

						handleLikesModalOpenCampaign={this.handleLikesModalOpenCampaign}
						handleDislikesModalOpenCampaign={this.handleDislikesModalOpenCampaign}
						handleCommentsModalOpenCampaign={this.handleCommentsModalOpenCampaign}
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
					<AddCampaignModal
						show={this.state.showCampaignModal}
						onCloseModal={this.handleAddCampaignModalClose}
						header="New campaign"
						hiddenMultiple = {this.state.hiddenMultiple}
						hiddenOne = {this.state.hiddenOne}
						noPicture = {this.state.noPicture}
						onDrop = {this.onDropCampaign}

					
						handleAddOneTimeCampaign = {this.handleAddOneTimeCampaignModal}
						handleAddMultipleTimeCampaign = {this.handleAddMultipleTimeCampaignModal}
						handleLinkChange = {this.handleLinkChange}
						handleDescriptionChange = {this.handleDescriptionChange}
						handleAddInfluencersModal = {this.handleAddInfluencersModal}
						handleDefineTargetGroupModal = {this.handleDefineTargetGroupModal}

					/>
					<AddVideoCampaignModal
						show={this.state.showVideoCampaignModal}
						onCloseModal={this.handleAddVideoCampaignModalClose}
						header="New video campaign"
						hiddenMultiple = {this.state.hiddenMultiple}
						hiddenOne = {this.state.hiddenOne}
						noPicture = {this.state.noPicture}
						onDrop = {this.onDropCampaign}

						handleSubmit = {this.handleSubmit}
						onChangeHandler = {this.onChangeHandler}

					
						handleAddOneTimeCampaign = {this.handleAddOneTimeCampaignModal}
						handleAddMultipleTimeCampaign = {this.handleAddMultipleTimeCampaignModal}
						handleLinkChange = {this.handleLinkChange}
						handleDescriptionChange = {this.handleDescriptionChange}
						handleAddInfluencersModal = {this.handleAddInfluencersModal}
						handleDefineTargetGroupModal = {this.handleDefineTargetGroupModal}

					/>
					<EditCampaignModal
						show={this.state.showEditCampaignModal}
						onCloseModal={this.handleEditCampaignModalClose}
						header="Edit campaign data"

						campaignForEdit = {this.state.campaignForEdit}
						campaignDate = {this.state.campaignDate}
						handleCampaignDateChange = {this.handleCampaignDateChange}
						campaignTime= {this.state.campaignTime}
						handleCampaignTimeChange = {this.handleCampaignTimeChange}
						campaignLink = {this.state.campaignLink}
						handleCampaignLinkChange = {this.handleCampaignLinkChange}
						campaignDescription = {this.state.campaignDescription}
						handleCampaignDescriptionChange = {this.handleCampaignDescriptionChange}
						handleChangeCampaign = {this.handleChangeCampaign}
						handleChangeMultipleCampaign = {this.handleChangeMultipleCampaign}
						campaignType = {this.state.campaignType}
						campaignDesiredNumber = {this.state.campaignDesiredNumber}
						campaignStartTime = {this.state.campaignStartTime}
						campaignEndTime = {this.state.campaignEndTime}
						handleCampaignDesiredNumberChange = {this.handleCampaignDesiredNumberChange}
						handleCampaignStartTimeChange = {this.handleCampaignStartTimeChange}
						handleCampaignEndTimeChange = {this.handleCampaignEndTimeChange}



					/>
					<OneTimeCampaignModal
						show={this.state.showOneTimeCampaignModal}
						onCloseModal={this.handleOneTimeCampaignModalClose}
						header="New one time campaign"

						handleAddOneTimeCampaign = {this.handleAddOneTimeCampaign}
					/>
					<MultipleTimeCampaignModal
						show={this.state.showMultipleTimeCampaignModal}
						onCloseModal={this.handleMultipleTimeCampaignModalClose}
						header="New multiple time campaign"

					
						handleAddMultipleTimeCampaign = {this.handleAddMultipleTimeCampaign}
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
					<VerifyModal
						show={this.state.showVerifyModal}
						onCloseModal={this.handleVerifyModalClose}
						header="Verify your profile"
						onDrop = {this.onDrop}
						handleSendRequestVerification = {this.handleSendRequestVerification}
					/>
					<AddVideoPostModal
						show={this.state.showVideoModal}
						onCloseModal={this.handleVideoPostModalClose}
						header="New video post/story"
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

						handleSubmit = {this.handleSubmit}
						onChangeHandler = {this.onChangeHandler}

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
					  <AddCollectionModal
                          collectionNameError={this.state.collectionNameError}
                        
						  show={this.state.showAddCollectionAlbumModal}
						  onCloseModal={this.handleAddCollectionModalClose}
						  header="Add new collection"
						  handleAddCollection={this.handleAddCollectionAlbum}
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
					   <AddPostToCollection
                          
					  
						  show={this.state.showAddAlbumToCollectionAlbumToCollection}
						  onCloseModal={this.handleAddPostToCollectionModalClose}
						  header="Add album to collection album"
						  addPostToCollection={this.addAlbumToCollectionAlbum}
						  collections = {this.state.collectionAlbums}
					  />
					   <AddTagsModal
                          
					  
						  show={this.state.showTagsModal}
						  onCloseModal={this.handleTagsModalClose}
						  header="Add tags"
						  followingUsers = {this.state.followingsThatAllowTags}
						  taggedOnPost = {this.state.taggedOnPost}
						  handleChangeTags = {this.handleChangeTags}
					  />
					   <AddInfluencerModal
                          
					  
						  show={this.state.showInfluencersModal}
						  onCloseModal={this.handleInfluencersModalClose}
						  header="Hire influencers"
						  influencers = {this.state.influencers}
						  choosenInfluencers = {this.state.choosenInfluencers}
						  handleChangeInfluencers = {this.handleChangeInfluencers}
					  />
					   <TargetGroupModal
                          
					  
						  show={this.state.showTargetGroupModal}
						  onCloseModal={this.handleTargetGroupModalClose}
						  header="Define target group"
						  addressInput = {this.addressInput}
						  onYmapsLoad = {this.onYmapsLoad}
						  handleGenderChange = {this.handleGenderChange}
						  selectedGender = {this.state.selectedGender}
						  handleDateOneChange = {this.handleDateOneChange}
						  handleDateTwoChange = {this.handleDateTwoChange}
						  selectedDateOne = {this.state.selectedDateOne}
						  selectedDateTwo = {this.state.selectedDateTwo}

					  />
                    </div>

			</React.Fragment>
		);
	}
}

export default ProfilePage;
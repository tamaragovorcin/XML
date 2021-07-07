
import React from "react";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import TopBar from "../components/TopBar";
import Header from "../components/Header";
import Axios from "axios";
import { BASE_URL } from "../constants.js";
import { Carousel } from 'react-responsive-carousel';
import ModalDialog from "../components/ModalDialog";
import AddCampaignModal from "../components/AddCampaignModal";
import AddVideoCampaignModal from "../components/AddVideoCampaignModal";
import OneTimeCampaignModal from "../components/OneTimeCampaignModal";
import MultipleTimeCampaignModal from "../components/MultipleTimeCampaignModal";
import AddInfluencerModal from "../components/AddInfluencerModal";
import TargetGroupModal from "../components/TargetGroupModal";
import getAuthHeader from "../GetHeader";
class NewCampaigns extends React.Component {

    constructor(props) {
		super(props);
		this.onDropCampaign = this.onDropCampaign.bind(this);
		this.addressInput = React.createRef();
	}
    state = {
        showCampaignModal : false,
		showVideoCampaignModal : false,
		link : "",
		showOneTimeCampaignModal : false,
		showMultipleTimeCampaignModal : false,
        openModal : false,
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
        selectedFile : "",
        pictures: [],
		videos : [],
		video : "",
		picture: "",
		hiddenOne: true,
		hiddenMultiple: true,
		noPicture : true,
        addressLocation :null,
		foundLocation : true,
        showImageModal : false,
        description : "",
		token: "",

    }
	onYmapsLoad = (ymaps) => {
		this.ymaps = ymaps;
		new this.ymaps.SuggestView(this.addressInput.current, {
			provider: {
				suggest: (request, options) => this.ymaps.suggest(request),
			},
		});
		
		
	};
		
	handleInfluencersModalClose = () =>{
		this.setState({ showInfluencersModal: false });

	}
    handlePostModalClose = () => {
		this.setState({ showImageModal: false });
	};
	handlePostModalOpen = () => {
		this.setState({ showImageModal: true });
		this.setState({pictures: []})
	};
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
    componentDidMount() {
       
    }
	getInfluencers = ()=> {
		let help = []

        const dto = {id: this.state.token}
        Axios.post(BASE_URL + "/api/userInteraction/api/user/following/category/token", dto, {  headers: { Authorization: getAuthHeader() } })
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
    handleModalClose = () => {
		this.setState({ openModal: false });
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
    handleDefineTargetGroupModal = ()=> {
		this.setState({ showTargetGroupModal: true });
	}
	handleAddOneTimeCampaignModal =()=>{
		this.setState({ showOneTimeCampaignModal: true });

	}
	handleAddMultipleTimeCampaignModal =() =>{
		this.setState({ showMultipleTimeCampaignModal: true });

	}
    handleAddInfluencersModal = ()=> {
		this.getInfluencers()
		this.setState({ showInfluencersModal: true });
	}
    handleDescriptionChange = (event) => {
		this.setState({ description: event.target.value });
	};
	handleLinkChange = (event) => {
		this.setState({ link: event.target.value });
	};
	handleTokenChange = (event) => {
		this.setState({ token: event.target.value });
	};
    onChangeHandler = (event) => {
		this.setState({
            selectedFile: event.target.files[0],
            loaded: 0,
        });
		
    };

    handleSubmit = (event) => {
        event.preventDefault();
    };
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
		Axios.post(BASE_URL + "/api/campaign/oneTimeCampaign/"+this.state.token, campaignDTO, {  headers: { Authorization: getAuthHeader() } })
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
					this.testCampaign(pic, campaignId);
				});
				
				if(this.state.selectedFile != ""){
				this.testVideoCampaign(this.state.selectedFile, campaignId)
				}
				this.setState({selectedFile : ""});
				this.setState({ pictures: [] });
			})
			.catch((err) => {
				console.log(err);
			});

	}
	getpartnershipsRequests = ()=> {
		var choosenInfluencersHelp = []
		this.state.choosenInfluencers.forEach((user) => {
			choosenInfluencersHelp.push(user.id)
		});
		return choosenInfluencersHelp
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
		Axios.post(BASE_URL + "/api/campaign/multipleTimeCampaign/"+this.state.token, campaignDTO, {  headers: { Authorization: getAuthHeader() } })
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
					this.testCampaign(pic, campaignId);
				});
				if(this.state.selectedFile != ""){
				this.testVideoCampaign(this.state.selectedFile, campaignId)
				}
				this.setState({selectedFile : ""});
				this.setState({ pictures: [] });
			})
			.catch((err) => {
				console.log(err);
			});
		this.setState({ showMultipleTimeCampaignModal: false, showCampaignModal : false });

	}
    testCampaign(pic, campaignId) {
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

	
		fetch( BASE_URL + "/api/campaign/api/image/"+this.state.token+"/"+campaignId, options, {  headers: { Authorization: getAuthHeader() } });
	}
    testVideoCampaign(pic, campaignId) {
		const formData = new FormData();

		formData.append("file", pic);
		formData.append("test", "StringValueTest");

		const options = {
			method: "POST",
			body: formData

		};
		fetch(BASE_URL + "/api/campaign/api/image/"+this.state.token+"/"+campaignId , options, {  headers: { Authorization: getAuthHeader() } });
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
render(){
    return (
        <React.Fragment>
				<TopBar />
				<Header />
         <div className="container" style={{ marginTop: "10%" }}>
					<h5 className=" text-center mb-0 mt-2 text-uppercase">Create new campaign</h5>
              <div className="d-flex align-items-top">
                <div className="container-fluid">
                  
                  <div className="form-group">
                     <button 
                        onClick={this.handleAddCampaignModal} 
                        className="btn btn-primary btn-xl"
                        type="button"

                        style={{
                            background: "#1977cc",
                            marginTop: "15px",
                            marginLeft: "20%",
                            width: "20%",
                        }}>
                             Publish image campaign
                    </button>
                    <button 
                        onClick={this.handleAddVideoCampaignModal} 
                        className="btn btn-primary btn-xl"
                        type="button"

                        style={{
                            background: "#1977cc",
                            marginTop: "15px",
                            marginLeft: "20%",
                            width: "20%",
                        }}>
                             Publish video campaign
                      </button>
									
									
		        	</div>
                </div>
            </div>
         </div>

            <ModalDialog
                    show={this.state.openModal}
					onCloseModal={this.handleModalClose}
					header="Successful"
					text={this.state.textSuccessfulModal}
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

						description = {this.state.description}
						link = {this.state.link}
						token = {this.state.token}
					
						handleAddOneTimeCampaign = {this.handleAddOneTimeCampaignModal}
						handleAddMultipleTimeCampaign = {this.handleAddMultipleTimeCampaignModal}
						handleLinkChange = {this.handleLinkChange}
						handleTokenChange = {this.handleTokenChange}

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
        </React.Fragment>

    );

	}
}
export default NewCampaigns;
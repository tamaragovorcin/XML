import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { Link } from "react-router-dom";
import playerLogo from "../static/me.jpg";
import profileImage from "../static/profileImage.jpg"
import LikesModal from "../components/Posts/LikesModal"
import DislikesModal from "../components/Posts/DislikesModal"
import CommentsModal from "../components/Posts/CommentsModal"
import WriteCommentModal from "../components/Posts/WriteCommentModal"
import { FiHeart } from "react-icons/fi";
import {FaHeartBroken,FaRegCommentDots} from "react-icons/fa"
import {BsBookmark} from "react-icons/bs"
import { YMaps, Map } from "react-yandex-maps";
import Axios from "axios";
import { BASE_URL_FEED } from "../constants.js";

const mapState = {
	center: [44, 21],
	zoom: 8,
	controls: [],
};
class HomePage extends React.Component {
	constructor(props) {
		super(props);
		this.addressInput = React.createRef();
	}
	state = {
		stories: [],
		photos: [],
		pictures: [],
		picture: "",
		hiddenOne: true,
		hiddenMultiple: true,
		peopleLikes : [],
		peopleDislikes : [],
		comments : [],
		showLikesModal : false,
		showDislikesModal : false,
		showCommentsModal : false,
		showWriteCommentModal : false,
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
	handleLikesModalOpen = ()=> {
		this.setState({ showLikesModal: true });    
	}
	handleDislikesModalOpen = ()=> {
		this.setState({ showDislikesModal: true });    
	}
	handleCommentsModalOpen = ()=> {
		this.setState({ showCommentsModal: true });    
	}
	handleWriteCommentModal = ()=>{
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

	
	handleLike = ()=>{
		
	}
	handleDislike = ()=>{
		
	}
	handleSave = ()=>{

	}
	componentDidMount() {
       
	}
	
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

                    Axios.get(BASE_URL_FEED + "/api/feed/searchByLocation/"+country + "/"+city+"/"+street)
                    .then((res) => {
                        this.setState({ photos: res.data });
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

       Axios.post(BASE_URL_FEED + "/api/feed/searchByHashTags/",helpDTO)
        .then((res) => {
            this.setState({ photos: res.data });
            this.setState({ hashtags: "" });
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
					<div className="container-fluid">
						
                    <table className="table">
							<tbody>
								{this.state.photos.map((post) => (
									
									<tr id={post.id} key={post.id}>
										
										<tr  style={{ width: "100%"}}>
											<td colSpan="3">
											<img
												className="img-fluid"
												src={`data:image/jpg;base64,${post.Media}`}
												width="100%"
												alt="description"
											/>
											</td>
										</tr>
										<tr  style={{ width: "100%" }}>
												<td>
												<button onClick={this.handleLike}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FiHeart/></button>
												</td>
												<td>
												<button onClick={this.handleDislike}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaHeartBroken/></button>

												</td>
												<td>
												<button onClick={this.handleWriteCommentModal}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaRegCommentDots/></button>
												</td>
												<td>
												<button onClick={this.handleSave}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px" }}><BsBookmark/></button>
												</td>
										</tr>
										<tr  style={{ width: "100%" }}>
												<td>
												<button onClick={this.handleLikesModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" , marginLeft:"4rem"}}><label>likes</label></button>
												</td>
												<td>
												<button onClick={this.handleDislikesModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label > dislikes</label></button>
												</td>
												<td>
												<button onClick={this.handleCommentsModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label >Comments</label></button>
												</td>
										</tr>
										<br/>
										<br/>
										<br/>
									</tr>
									
								))}

							</tbody>
						</table>
					</div>
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
						comments = {this.state.comments}
                    />
					<WriteCommentModal
                        show={this.state.showWriteCommentModal}
						onCloseModal={this.handleWriteCommentModalClose}
						header="Leave your comment"
                    />
                        
                    </div>
			</React.Fragment>
		);
	}
}

export default HomePage;
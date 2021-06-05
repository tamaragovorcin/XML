import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { Link } from "react-router-dom";
import { BASE_URL_FEED} from "../constants.js";
import playerLogo from "../static/me.jpg";
import profileImage from "../static/profileImage.jpg"
import LikesModal from "../components/Posts/LikesModal"
import DislikesModal from "../components/Posts/DislikesModal"
import CommentsModal from "../components/Posts/CommentsModal"
import WriteCommentModal from "../components/Posts/WriteCommentModal"
import { FiHeart } from "react-icons/fi";
import {FaHeartBroken,FaRegCommentDots} from "react-icons/fa"
import {BsBookmark} from "react-icons/bs"
import Axios from "axios";
import IconTabsHomePage from "../components/Posts/IconTabsHomePage"

class HomePage extends React.Component {
	
	state = {
		stories: [],
		photos: [],
		peopleLikes : [],
		peopleDislikes : [],
		comments : [],
		albums : [],
		showLikesModal : false,
		showDislikesModal : false,
		showCommentsModal : false,
		showWriteCommentModal : false
	}
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
							handleSave = {this.handleSave}
							handleLikesModalOpen = {this.handleLikesModalOpen}
							handleDislikesModalOpen = {this.handleDislikesModalOpen}
							handleCommentsModalOpen = {this.handleCommentsModalOpen}

							albums ={this.state.albums}
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
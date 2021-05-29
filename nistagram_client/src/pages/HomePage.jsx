import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { Link } from "react-router-dom";
import playerLogo from "../static/me.jpg";
import LikesModal from "../components/Posts/LikesModal"
import DislikesModal from "../components/Posts/DislikesModal"
import CommentsModal from "../components/Posts/CommentsModal"
import WriteCommentModal from "../components/Posts/WriteCommentModal"
import { FiHeart } from "react-icons/fi";
import {FaHeartBroken,FaRegCommentDots} from "react-icons/fa"


class HomePage extends React.Component {
	
	state = {
		highlihts: [],
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
	componentDidMount() {
		this.handleGetBasicInfo()
		this.handleGetHighlights()
		this.handleGetPhotos()

	}
	handleGetBasicInfo = () => {
		this.setState({ username: "USERNAME" });
	}

	handleGetHighlights = () => {
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

		this.setState({ highlihts: list });
	}

	handleGetPhotos = () => {
		let list = []
		let comments1 = []
		let comments2 = []
		let comment1 = { id: 1, user: "USER 1 ", text: "very nice" }
		let comment11 = { id: 2, user: "USER 2 ", text: "cool" }
		let comment111 = { id: 3, user: "USER 3 ", text: "vau" }
		comments1.push(comment1)
		comments1.push(comment11)
		comments1.push(comment111)

		let comment2 = { id: 4, user: "USER 55443 ", text: "i like it" }
		let comment22 = { id: 5, user: "USER 11111 ", text: "ugly" }
		let comment222 = { id: 6, user: "USER 33333 ", text: "awesome" }
		comments2.push(comment2)
		comments2.push(comment22)
		comments2.push(comment222)

		let photo1 = { id: 1, username:"mladenkak", photo: playerLogo, numLikes: 52, numDislikes: 2, comments: comments1 }
		let photo2 = { id: 2, username:"mladenkak", photo: playerLogo, numLikes: 45, numDislikes: 0, comments: comments2 }
		let photo3 = { id: 3, username:"mladenkak", photo: playerLogo, numLikes: 52, numDislikes: 2, comments: comments1 }
		let photo4 = { id: 4, username:"mladenkak", photo: playerLogo, numLikes: 45, numDislikes: 0, comments: comments2 }
		list.push(photo1)
		list.push(photo2)
		list.push(photo3)
		list.push(photo4)

		this.setState({ photos: list });

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

										<tr>
											{this.state.highlihts.map((high) => (
												<td id={high.id} key={high.id} width="30em">
													<tr width="100em">
														<img
															className="img-fluid"
															src={playerLogo}
															width="30em"
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
					<div className="container-fluid">
						
						<table className="table">
							<tbody>
								{this.state.photos.map((photo) => (
									
									<tr id={photo.id} key={photo.id}>
										<img
												className="img-fluid"
												src={photo.profilePhoto}
												style={{marginRight:"2%"}}
												alt="profile"
											/>
										<label style={{fontSize:"20px",fontWeight:"bold"}}>{photo.username}</label>
										<tr  style={{ width: "100%"}}>
											<td colSpan="3">
											<img
												className="img-fluid"
												src={photo.photo}
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
										</tr>
										<tr  style={{ width: "100%" }}>
												<td>
												<button onClick={this.handleLikesModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" , marginLeft:"4rem"}}><label><b>{photo.numLikes}</b>likes</label></button>
												</td>
												<td>
												<button onClick={this.handleDislikesModalOpen} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label ><b>{photo.numDislikes}</b> dislikes</label></button>
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

import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { Link } from "react-router-dom";
import playerLogo from "../static/coach.png";
import { CgProfile } from 'react-icons/cg';
class HomePage extends React.Component {
	state = {
		username : "",
		numberPosts : 0,
		numberFollowing : 0,
		numberFollowers : 0,
		biography : "",
		highlihts : [],
		photos : []
	}
	componentDidMount() {
		this.handleGetBasicInfo()
		this.handleGetHighlights()
		this.handleGetPhotos()
		
    }
	handleGetBasicInfo = ()=> {
		this.setState({ numberPosts: 10});  
		this.setState({ numberFollowing: 600 }); 
		this.setState({ numberFollowers: 750 });   
		this.setState({ biography: "bla bla bla" });     
		this.setState({ username: "USERNAME"});  
	}

	handleGetHighlights =() => {
		let highliht1 = {id: 1,name : "ITALY"};
		let highliht2 = {id:2, name : "AMERICA"};
		let highliht3 = {id:3,name : "SERBIA"};

		let list = [];
		list.push(highliht1)
		list.push(highliht2)
		list.push(highliht3)
	
		this.setState({ highlihts: list}); 
	}

	handleGetPhotos = ()=> {
		let list = []
		let comments1 = []
		let comments2 = []
		let comment1 = {id:1, user: "USER 1 ", text : "very nice"}
		let comment11 = {id:2, user: "USER 2 ", text : "cool"}
		let comment111 = {id:3, user: "USER 3 ", text : "vau"}
		comments1.push(comment1)
		comments1.push(comment11)
		comments1.push(comment111)

		let comment2 = {id:4, user: "USER 55443 ", text : "i like it"}
		let comment22 = {id:5, user: "USER 11111 ", text : "ugly"}
		let comment222 = {id:6, user: "USER 33333 ", text : "awesome"}
		comments2.push(comment2)
		comments2.push(comment22)
		comments2.push(comment222)

		let photo1 = {id:1, photo: playerLogo, numLikes: 52, numDislikes : 2, comments : comments1}
		let photo2 = {id:2, photo: playerLogo, numLikes: 45, numDislikes : 0, comments : comments2}
		list.push(photo1)
		list.push(photo2)

		this.setState({ photos: list}); 

	}
	
	render() {
		return (
			<React.Fragment>
				<TopBar />
				<Header />

				
				<div className="d-flex align-items-top">
					<div className="container"  style={{ marginTop: "10rem", marginRight: "10rem"}}>
						<table className="table" style={{ width: "100%" }}>
								<tbody>
									
										<tr>
											<td width="130em">
												<img
														className="img-fluid"
														src={playerLogo}
														width="70em"
														alt="description"
													/>
											</td>

											<td>
												<div>
													<td>
														<label >{this.state.username}</label>
													</td>
													<td>
														<Link to="/userChangeProfile" className="btn btn-outline-secondary btn-sm">Edit profile</Link>
											
													</td>
												</div>
												<div>
													<td>
														<label ><b>{this.state.numberPosts}</b> posts</label>
													</td>
													<td>
														<label ><b>{this.state.numberFollowers}</b> followers</label>
													</td>
													<td>
														<label ><b>{this.state.numberFollowing}</b> following</label>
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

				<div className="d-flex align-items-top">
					<div className="container"  style={{  marginRight: "10rem"}}>
					<table className="table" style={{ width: "100%" }}>
								<tbody>
								
                                    <tr >
									{this.state.highlihts.map((high) => (
										<td id={high.id} key={high.id} width="30em">
											<tr width="100em">
												<img
													className="img-fluid"
													src={playerLogo}
													width="40em"
													alt="description"
												/>
											</tr>
											<tr>
												<label>{high.name}</label>
											</tr>
										</td>
										))}
                                    </tr>

                                
								</tbody>
							</table>
					</div>
				</div>
				<div className="d-flex align-items-top">
					<div className="container"  style={{  marginLeft: "30rem"}}>
					<table className="table" style={{ width: "100%" }}>
								<tbody>
									{this.state.photos.map((photo) => (
										<tr id={photo.id} key={photo.id}>

											<td width="200em">
												<img
													className="img-fluid"
													src={photo.photo}
													width="100em"
													alt="description"
												/>
											</td>

											<td>
													<tr>
														<label ><b>{photo.numLikes}</b> likes</label>
													</tr>
													<tr>
														<label ><b>{photo.numDislikes}</b> dislikes</label>
													</tr>
													<tr>
														<label >comments</label>
													</tr>
													
											</td>
										</tr>
									))}
									
								</tbody>
							</table>
					</div>
				</div>
				
			</React.Fragment>
		);
	}
}

export default HomePage;

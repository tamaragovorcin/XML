import React, { Component } from "react";
import Axios from "axios";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { BASE_URL_USER_INTERACTION, BASE_URL_USER } from "../constants.js";
import ModalDialog from "../components/ModalDialog";
import { BASE_URL } from "../constants.js";
import getAuthHeader from "../GetHeader";
class FollowRequest extends Component {
	state = {
        followerRequests: [],
        following : [],
        followers : [],
        followerRequestsByMe : [],
        blockedUsers : [],
        textSuccessfulModal : "",
        openModal : false,
        followRecommendations : []
    };

    componentDidMount() {
        this.getFollowRequests()
        this.getFollowers()
        this.getFollowing()
        this.getFollowRequestsByMe()
        this.getBlockedUsers()
        this.getFollowRecommendations()
	}
    hasRole = (reqRole) => {
		let roles = JSON.parse(localStorage.getItem("keyRole"));
		if (roles === null) return false;

		if (reqRole === "*") return true;

		for (let role of roles) {
			if (role === reqRole) return true;
		}
		return false;
	};
    handleAccept  = (followerId) => {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
            const acceptedDTO = { following: id, follower: followerId};
            Axios.post(BASE_URL  + "/api/userInteraction/api/acceptFollowRequest", acceptedDTO,  {  headers: { Authorization: getAuthHeader() } })
                    .then((res) => {
                        this.setState({ openModal: true });
                        this.setState({ textSuccessfulModal: "You have successfully accepted follow request." });			
                        this.getFollowRequests()
                        this.getFollowers()
                    })
                    .catch ((err) => {
                console.log(err);
            });
    }

    handleDelete  = (followerId) => {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = { following: id, follower: followerId};
        Axios.post(BASE_URL  + "/api/userInteraction/api/deleteFollowRequest", dto,  {  headers: { Authorization: getAuthHeader() } })
                .then((res) => {
                    this.setState({ openModal: true });
                    this.setState({ textSuccessfulModal: "You have successfully deleted follow request." });			
                    this.getFollowRequests()
                })
                .catch ((err) => {
            console.log(err);
        });
    }

	getFollowRequests = ()=> {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = {id: id}
        Axios.post(BASE_URL + "/api/userInteraction/api/user/followRequests", dto,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ followerRequests: res.data });
			})
			.catch((err) => {
				console.log(err)
			});
    }
    getFollowRequestsByMe = ()=> {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = {id: id}
        Axios.post(BASE_URL + "/api/userInteraction/api/user/followRequestsByMe", dto,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ followerRequestsByMe: res.data });
			})
			.catch((err) => {
				console.log(err)
			});
    }
    getFollowers = ()=> {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = {id: id}
        Axios.post(BASE_URL + "/api/userInteraction/api/user/followers", dto,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ followers: res.data });
			})
			.catch((err) => {
				console.log(err)
			});
    }

    getFollowing = ()=> {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = {id: id}
        Axios.post(BASE_URL + "/api/userInteraction/api/user/following", dto, {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ following: res.data });
			})
			.catch((err) => {
				console.log(err)
			});
    }
    getBlockedUsers = ()=> {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        Axios.get(BASE_URL + "/api/users/api/user/blockedUsers/"+id,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ blockedUsers: res.data });
                console.log(res.data)
			})
			.catch((err) => {
				console.log(err)
			});
    }
    handleUnblock = (followerId) =>{
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = { Subject: id, Object: followerId};
        Axios.post(BASE_URL_USER  + "/api/unblock/", dto,  {  headers: { Authorization: getAuthHeader() } })
                .then((res) => {
                    this.setState({ openModal: true });
                    this.setState({ textSuccessfulModal: "You have successfully unblocked user" });			
                    this.getBlockedUsers();
                })
                .catch ((err) => {
            console.log(err);
        });
    }
    getFollowRecommendations = ()=> {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = {id: id}
        Axios.post(BASE_URL + "/api/userInteraction/followRecommendations", dto,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				this.setState({ followRecommendations: res.data });
			})
			.catch((err) => {
				console.log(err)
			});
    }
    handleModalClose = () => {
		this.setState({ openModal: false });
	};

    handleFollowUser = (userId, privacy)=> {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
	
		const followReguestDTO = { follower: id, following : userId};
		if(privacy==="private") {

			Axios.post(BASE_URL + "/api/userInteraction/api/followRequest", followReguestDTO,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				                this.getFollowRequestsByMe()
                this.getFollowRecommendations()

				this.setState({ textSuccessfulModal: "You have successfully sent follow request." });
				this.setState({ openModal: true });

			})
			.catch ((err) => {
				console.log(err);
			});
		}else {
			Axios.post(BASE_URL + "/api/userInteraction/api/followPublic", followReguestDTO,  {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
				
                this.getFollowing()
                this.getFollowRecommendations()
				this.setState({ textSuccessfulModal: "You are now following this user." });
				this.setState({ openModal: true });

			})
			.catch ((err) => {
				console.log(err);
			});
		}
    }
    handleTransferToFollowerPage = (userId) => {
        window.location = "#/followerProfilePage/" + userId;
    }

	render() {
		return (
            <React.Fragment>
				<TopBar />
				<Header />

                    <div className="container">
                        <div className="container" style={{ marginTop: "10rem", marginRight: "10rem" }}>
                            <h3>
                                Follow requests
                            </h3>
                            <table className="table" style={{ width: "100%" }}>
                                        <tbody>
                                        {this.state.followerRequests.map((follower) => (
                                                <tr id={follower.Id} key={follower.Id}>
                                                    
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Username: </b> {follower.Username}
                                                        </div>
                                                    </td>
                                                    <td >
                                                        <div style={{marginLeft:'55%'}}>
                                                            <td>
                                                                <button  className="btn btn-success mt-1" onClick={() => this.handleAccept(follower.Id)} type="button"><i className="icofont-subscribe mr-1"></i>Accept</button>
                                                            </td>
                                                            <td>
                                                                <button className="btn btn-danger mt-1" onClick={() => this.handleDelete(follower.Id)} type="button"><i className="icofont-subscribe mr-1"></i>Delete</button>
                                                            </td>
                                                            
                                                        </div>
                                                    </td>
                                                </tr>

                                            ))}
                                            <tr>
                                                <td></td>
                                                <td></td>
                                                <td></td>
                                            </tr>
                                        </tbody>
                                    </table>
                            </div>
                        </div>
                        <div className="container">
                        <div className="container" style={{ marginTop: "10rem", marginRight: "10rem" }}>
                            <h3>
                                Follow requests that I have sent
                            </h3>
                            <table className="table" style={{ width: "100%" }}>
                                        <tbody>
                                        {this.state.followerRequestsByMe.map((follower) => (
                                                <tr id={follower.Id} key={follower.Id}>
                                                    
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Username: </b> {follower.Username}
                                                        </div>
                                                    </td>
                                                </tr>

                                            ))}
                                            <tr>
                                                <td></td>
                                                <td></td>
                                                <td></td>
                                            </tr>
                                        </tbody>
                                    </table>
                            </div>
                        </div>
                        <div className="container">
                        <div className="container" style={{ marginTop: "10rem", marginRight: "10rem" }}>
                            <h3>
                                People i am following
                            </h3>
                            <table className="table" style={{ width: "100%" }}>
                                        <tbody>
                                        {this.state.following.map((follower) => (
                                                <tr id={follower.Id} key={follower.Id}>
                                                    
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Username: </b> {follower.Username}
                                                        </div>
                                                    </td>
                                                </tr>

                                            ))}
                                            <tr>
                                                <td></td>
                                                <td></td>
                                                <td></td>
                                            </tr>
                                        </tbody>
                                    </table>
                            </div>
                        </div>
                        <div className="container">
                        <div className="container" style={{ marginTop: "10rem", marginRight: "10rem" }}>
                            <h3>
                                People who follow me
                            </h3>
                            <table className="table" style={{ width: "100%" }}>
                                        <tbody>
                                        {this.state.followers.map((follower) => (
                                                <tr id={follower.Id} key={follower.Id}>
                                                    
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Username: </b> {follower.Username}
                                                        </div>
                                                    </td>
                                                </tr>

                                            ))}
                                            <tr>
                                                <td></td>
                                                <td></td>
                                                <td></td>
                                            </tr>
                                        </tbody>
                                    </table>
                            </div>
                        </div>
                        <div className="container">
                        <div className="container" style={{ marginTop: "10rem", marginRight: "10rem" }}>
                            <h3>
                               Blocked users:
                            </h3>
                            <table className="table" style={{ width: "100%" }}>
                                        <tbody>
                                        {this.state.blockedUsers.map((follower) => (
                                                <tr id={follower.Id} key={follower.Id}>
                                                    
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Username: </b> {follower.Username}
                                                        </div>
                                                    </td>
                                                    <td>
                                                    <button  className="btn btn-success mt-1" onClick={() => this.handleUnblock(follower.Id)} type="button"><i className="icofont-subscribe mr-1"></i>Unblock</button>

                                                    </td>
                                                </tr>

                                            ))}
                                            <tr>
                                                <td></td>
                                                <td></td>
                                                <td></td>
                                            </tr>
                                        </tbody>
                                    </table>
                            </div>
                        </div>
                        <div className="container">
                        <div className="container" style={{ marginTop: "10rem", marginRight: "10rem" }}>
                            <h3>
                               Follow recommendations:
                            </h3>
                            <table className="table" style={{ width: "100%" }}>
                                        <tbody>
                                        {this.state.followRecommendations.map((user) => (
                                                <tr id={user.Id} key={user.Id}>
                                                    
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Username: </b> {user.Username}
                                                        </div>
                                                    </td>
                                                    <td>
                                                        <button  className="btn btn-success mt-1" onClick={() => this.handleTransferToFollowerPage(user.Id)} type="button"><i className="icofont-subscribe mr-1"></i>See profile</button>
                                                    </td>
                                                    <td>
                                                        <button  className="btn btn-success mt-1" onClick={() => this.handleFollowUser(user.Id, user.Privacy)} type="button"><i className="icofont-subscribe mr-1"></i>Follow</button>
                                                    </td>
                                                </tr>

                                            ))}
                                            <tr>
                                                <td></td>
                                                <td></td>
                                                <td></td>
                                            </tr>
                                        </tbody>
                                    </table>
                            </div>
                        </div>
                        <div>
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

export default FollowRequest;
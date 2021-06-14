import React, { Component } from "react";
import Axios from "axios";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { BASE_URL_USER_INTERACTION, BASE_URL_USER } from "../constants.js";
import ModalDialog from "../components/ModalDialog";

class FollowRequest extends Component {
	state = {
        followerRequests: [],
        following : [],
        followers : [],
        followerRequestsByMe : [],
        blockedUsers : [],
        textSuccessfulModal : "",
        openModal : false,

    };

    componentDidMount() {
        this.getFollowRequests()
        this.getFollowers()
        this.getFollowing()
        this.getFollowRequestsByMe()
        this.getBlockedUsers()
	}
    handleAccept  = (followerId) => {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
            const acceptedDTO = { following: id, follower: followerId};
            Axios.post(BASE_URL_USER_INTERACTION  + "/api/acceptFollowRequest", acceptedDTO)
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
        Axios.post(BASE_URL_USER_INTERACTION  + "/api/deleteFollowRequest", dto)
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
        Axios.post(BASE_URL_USER_INTERACTION + "/api/user/followRequests", dto)
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
        Axios.post(BASE_URL_USER_INTERACTION + "/api/user/followRequestsByMe", dto)
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
        Axios.post(BASE_URL_USER_INTERACTION + "/api/user/followers", dto)
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
        Axios.post(BASE_URL_USER_INTERACTION + "/api/user/following", dto)
			.then((res) => {
				this.setState({ following: res.data });
			})
			.catch((err) => {
				console.log(err)
			});
    }
    getBlockedUsers = ()=> {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        Axios.get(BASE_URL_USER + "/api/user/blockedUsers/"+id)
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
        Axios.post(BASE_URL_USER  + "/api/unblock/", dto)
                .then((res) => {
                    this.setState({ openModal: true });
                    this.setState({ textSuccessfulModal: "You have successfully unblocked user" });			
                    this.getBlockedUsers();
                })
                .catch ((err) => {
            console.log(err);
        });
    }
    handleModalClose = () => {
		this.setState({ openModal: false });
	};
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
                                                            <b>Username: </b> {follower.ProfileInformation.Username}
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
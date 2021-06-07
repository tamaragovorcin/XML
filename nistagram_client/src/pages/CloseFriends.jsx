import React, { Component } from "react";
import TopBar from "../components/TopBar";
import Header from "../components/Header";
import Axios from "axios";
import { BASE_URL_USER } from "../constants.js";
import PharmacyLogo from "../static/coach.png";
import "../App.js";
import { Redirect } from "react-router-dom";
import { BASE_URL_USER_INTERACTION } from "../constants.js";

import ModalDialog from "../components/ModalDialog";
class CloseFriends extends Component {
	state = {
		closeFriends : [],
		notCloseFriends : [],
		openModal : false,
		textSuccessfulModal : ""
	};

	hasRole = (reqRole) => {
        let roles = JSON.parse(localStorage.getItem("keyRole"));

        if (roles === null) return false;

        if (reqRole === "*") return true;

        for (let role of roles) {
            if (role === reqRole) return true;
        }
        return false;
    };
	componentDidMount() {
		this.getCloseFriendsDTO()
	};

	getCloseFriendsDTO = ()=> {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
		const dto = {id: id}
		Axios.post(BASE_URL_USER_INTERACTION + "/api/user/followingCloseFriends", dto)
			.then((res) => {
				this.setState({ closeFriends: res.data.CloseFriends });
				this.setState({ notCloseFriends: res.data.NotCloseFriends });

			})
			.catch((err) => {
				console.log(err)
			});
	}	

	hangleFormToogle = () => {
		this.setState({ formShowed: !this.state.formShowed });
	};

	handleAddToCloseFriends = (e,id) => {
        let userid = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
        
        let CloseFriendsDTO = {
            IdLogged: userid,
            IdClose: id 
        }
        Axios.post(BASE_URL_USER + "/api/user/addToCloseFriends/", CloseFriendsDTO)
        .then((res) => {
			this.setState({ openModal: true });
			this.setState({ textSuccessfulModal: "You have successfully added user to close friends." });

			this.getCloseFriendsDTO()
        })
        .catch((err) => {

            console.log(err)
        });

	};


    handleRemoveFromCloseFriends = (e,id) => {
        let userid = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
        
        let CloseFriendsDTO = {
            IdLogged: userid,
            IdClose: id 
        }
        Axios.post(BASE_URL_USER + "/api/user/removeFromCloseFriends/", CloseFriendsDTO)
        .then((res) => {

			this.setState({ openModal: true });
			this.setState({ textSuccessfulModal: "You have successfully removed user from close friends." });

			this.getCloseFriendsDTO()
		})
        .catch((err) => {

            console.log(err)
        });

	};
	handleModalClose = () => {
		this.setState({ openModal: false });
	};


	render() {
		if (this.state.redirect) return <Redirect push to={this.state.redirectUrl} />;

		return (
			<React.Fragment>
				<TopBar />
				<Header />

				<div className="container" style={{ marginTop: "10%" }}>
					<h5 className=" text-center mb-0 mt-2 text-uppercase">Add to close friends</h5>

					
					<table className="table table-hover" style={{ width: "100%", marginTop: "3rem" }}>
						<tbody>
							{this.state.notCloseFriends.map((user) => (
								<tr
									id={user.Id}
									key={user.Id}
									style={{ cursor: "pointer" }}
							
								>
									<td width="130em">
										<img className="img-fluid" src={user.files?.[0] ?? PharmacyLogo} width="70em" />
									</td>
									<td>
										<div>
											{user.Username}
										</div>
										

									<div>  <button
											style={{
												background: "#1977cc",
												marginTop: "15px",
												marginLeft: "40%",
												width: "20%",
											}}
											onClick={(e)=>this.handleAddToCloseFriends(e,user.Id)}
											className="btn btn-primary btn-xl"
											id="sendMessageButton"
											type="button"
										>
											Add to close friends
									</button></div>

									</td>
								</tr>
							))}
						</tbody>
					</table>
				</div>


                <div className="container" style={{ marginTop: "10%" }}>
					<h5 className=" text-center mb-0 mt-2 text-uppercase">My close friends</h5>

					
					<table className="table table-hover" style={{ width: "100%", marginTop: "3rem" }}>
						<tbody>
							{this.state.closeFriends.map((user) => (
								<tr
									id={user.Id}
									key={user.Id}
									style={{ cursor: "pointer" }}
							
								>
									<td width="130em">
										<img className="img-fluid" src={user.files?.[0] ?? PharmacyLogo} width="70em" />
									</td>
									<td>
										<div>
											{user.Username}
										</div>
										

									<div>  <button
											style={{
												background: "#1977cc",
												marginTop: "15px",
												marginLeft: "40%",
												width: "20%",
											}}
											onClick={(e)=>this.handleRemoveFromCloseFriends(e,user.Id)}
											className="btn btn-primary btn-xl"
											id="sendMessageButton"
											type="button"
										>
											Remove from close friends
									</button></div>

									</td>
								</tr>
							))}
						</tbody>
					</table>
				</div>
				<ModalDialog
                    show={this.state.openModal}
					onCloseModal={this.handleModalClose}
					header="Successful"
					text={this.state.textSuccessfulModal}
                />
			
			</React.Fragment>
		);
	}
}

export default CloseFriends;


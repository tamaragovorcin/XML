import React, { Component } from "react";
import TopBar from "../components/TopBar";
import Header from "../components/Header";
import Axios from "axios";
import { BASE_URL_USER } from "../constants.js";
import PharmacyLogo from "../static/coach.png";
import "../App.js";
import { Redirect } from "react-router-dom";

import ModalDialog from "../components/ModalDialog";
class CloseFriends extends Component {
	state = {
        users: [],
        options: [],
        myCloseFriends: [],


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
		
        let help = []
		Axios.get(BASE_URL_USER + "/api/")
			.then((res) => {

				console.log(res.data)
				this.setState({ users: res.data });

				res.data.forEach((user) => {
					let optionDTO = { id: user.Id, label: user.ProfileInformation.Username, value: user.Id }
					help.push(optionDTO)
				});

				this.setState({ options: help });
				console.log(help)
			})
			.catch((err) => {

				console.log(err)
			});







	};


	

	hangleFormToogle = () => {
		this.setState({ formShowed: !this.state.formShowed });
	};

	handleAddToCloseFriends = (e,id) => {
        let userid = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
        
        let CloseFriendsDTO = {
            IdLogged: userid,
            IdClose: id 
        }
        Axios.get(BASE_URL_USER + "/api/addToCloseFriends/" + CloseFriendsDTO)
        .then((res) => {

            console.log(res.data)
			window.location.reload();
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
        Axios.get(BASE_URL_USER + "/api/removeFromCloseFriends/" + CloseFriendsDTO)
        .then((res) => {

            console.log(res.data)
            window.location.reload();
        })
        .catch((err) => {

            console.log(err)
        });

	};



	handleOrderModalClose = () => {
		this.setState({ showOrderModal: false });
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
							{this.state.users.map((user) => (
								<tr
									id={user.id}
									key={user.id}
									style={{ cursor: "pointer" }}
							
								>
									<td width="130em">
										<img className="img-fluid" src={user.files?.[0] ?? PharmacyLogo} width="70em" />
									</td>
									<td>
										<div>
											{user.ProfileInformation.Username}
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
							{this.state.users.map((user) => (
								<tr
									id={user.id}
									key={user.id}
									style={{ cursor: "pointer" }}
							
								>
									<td width="130em">
										<img className="img-fluid" src={user.files?.[0] ?? PharmacyLogo} width="70em" />
									</td>
									<td>
										<div>
											{user.ProfileInformation.Username}
										</div>
										

									<div>  <button
											style={{
												background: "#1977cc",
												marginTop: "15px",
												marginLeft: "40%",
												width: "20%",
											}}
											onClick={(e)=>this.handleRemoveFromCloseFriends(e,user.id)}
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
                    header="Success"
                    text="You have successfully removed the item."
                />
			
			</React.Fragment>
		);
	}
}

export default CloseFriends;


import React, { Component } from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { BASE_URL, BASE_URL_USER } from "../constants.js";
import Axios from "axios";
import { Redirect } from "react-router-dom";
import HeadingAlert from "../components/HeadingAlert";
import HeadingSuccessAlert from "../components/HeadingSuccessAlert"
import SidebarSettings from "../components/SidebarSettings"



class NotificationSettings extends Component {
	state = {
		id: "",
        messages : false,
        comments :  false,
		openModal: false,
		hiddenEditInfo: true,
		redirect: false,
		hiddenSuccessAlert: true,
		successHeader: "",
		successMessage: "",
		hiddenFailAlert: true,
		failHeader: "",
		failMessage: "",
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
    handleMessagesChange(event) {
		if(this.state.messages=== true) {
			this.setState({ messages: false });
		}
		if(this.state.messages=== false) {
			this.setState({ messages: true });
		}
	}
    handleCommentsChange(event) {
		if(this.state.comments=== true) {
			this.setState({ comments: false });
		}
		if(this.state.comments=== false) {
			this.setState({ comments: true });
		}
	}
	
	componentDidMount() {
	
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
	Axios.get(BASE_URL + "/api/users/api/user/notificationSettings/" + id)
				.then((res) => {
						this.setState({
							id: res.data.Id,
                            messages : res.data.NotificationsMessages,
                            comments : res.data.NotificationsComments
						});
				})
				.catch ((err) => {
			console.log(err);
		});

	}

	

	handleSuccessModalClose = () => {
		this.setState({ openSuccessModal: false });
	};


	handleChangeInfo = () => {
		this.setState({
			hiddenSuccessAlert: true,
			successHeader: "",
			successMessage: "",
			hiddenFailAlert: true,
			failHeader: "",
			failMessage: "",
		});

        const dto = {
                        User : localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1),
                        Messages :  this.state.messages,
                        Comments : this.state.comments
        };

		Axios.post(BASE_URL + "/api/users/api/updateNotifications", dto)
		.then((res2) => {
			alert("success")	
		})
		.catch((err) => {
			console.log(err);
		});
            
        
	
	};


	handleEditInfoClick = () => {
		this.setState({ hiddenEditInfo: false });
	};


	handleCloseAlertSuccess = () => {
		this.setState({ hiddenSuccessAlert: true });
	};

	handleCloseAlertFail = () => {
		this.setState({ hiddenFailAlert: true });
	};



	render() {
		//if (this.state.redirect) return <Redirect push to="/login" />;

		return (
			<React.Fragment>
				
				<TopBar />
				<Header />

				<div className="container=fluid" style={{ marginTop: "8%",marginLeft:"5%",marginRight:"5%",background: "#fcfafa"}}>
				
                    <br/>
						<h2 className=" text-center  mb-0 text-uppercase" style={{ marginTop: "0", color:"#2c4964" }}>
						Profile settings
					</h2>
                    <br />
					<div className="row section-design" style={{ marginLeft: "2%",marginRight:"2%"}}>
							
                            <div className="col-md-2 padding-0">
                            <SidebarSettings/></div>
                            <div className="col-md-8 padding-0">
							

				<div className="container" style={{ marginTop: "0%" }}>
					<HeadingSuccessAlert
						hidden={this.state.hiddenSuccessAlert}
						header={this.state.successHeader}
						message={this.state.successMessage}
						handleCloseAlert={this.handleCloseAlertSuccess}
					/>
					<HeadingAlert
						hidden={this.state.hiddenFailAlert}
						header={this.state.failHeader}
						message={this.state.failMessage}
						handleCloseAlert={this.handleCloseAlertFail}
					/>
					<div className="row mt-10">
						<div className="col shadow p-3 bg-white rounded">
							<form id="contactForm" name="sentMessage">
						
                          
                                <div>
									<p style={{ color: "#6c757d", opacity: 1 }} >Enable notifications for messages: </p>
									<label class="switch" >
										<input checked ={this.state.messages} type="checkbox"  onChange={(e) => this.handleMessagesChange(e)}/>
										<span  class="slider round"></span>

									</label>
								</div>
                                <div>
									<p style={{ color: "#6c757d", opacity: 1 }} >Enable notifications for comments: </p>
									<label class="switch" >
										<input checked ={this.state.comments} type="checkbox"  onChange={(e) => this.handleCommentsChange(e)}/>
										<span  class="slider round"></span>

									</label>
								</div>

								<div className="form-group">
									<div className="form-group controls mb-0 pb-2">
										<div className="form-row justify-content-center">
											<div className="form-col" hidden={!this.state.hiddenEditInfo}>
												<button
													style={{ background: "#1977cc", marginTop: "15px" }}
													onClick={this.handleChangeInfo}
													className="btn btn-primary btn-xl"
													id="sendMessageButton"
													type="button"
												>
													Save
												</button>
											</div>
										
										</div>
									</div>
								</div>
							</form>
						</div>
					</div>
				</div>



                            </div>

                            </div>
							
                            
				</div>
				
			</React.Fragment>
		);
	}
}

export default NotificationSettings;
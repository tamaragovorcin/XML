import React, { Component } from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { BASE_URL, BASE_URL_USER } from "../constants.js";
import Axios from "axios";
import { Redirect } from "react-router-dom";
import HeadingAlert from "../components/HeadingAlert";
import HeadingSuccessAlert from "../components/HeadingSuccessAlert"
import SidebarSettings from "../components/SidebarSettings"



class PrivacySettings extends Component {
	state = {
		id: "",
        private : false,
        messages : false,
        tags :  false,
		openModal: false,
		openPasswordModal: false,
		hiddenEditInfo: true,
		redirect: false,
		hiddenPasswordErrorAlert: true,
		errorPasswordHeader: "",
		errorPasswordMessage: "",
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
    handlePrivateChange(event) {
		if(this.state.private=== true) {
			this.setState({ private: false });
		}
		if(this.state.private=== false) {
			this.setState({ private: true });
		}
	}
    handleMessagesChange(event) {
		if(this.state.messages=== true) {
			this.setState({ messages: false });
		}
		if(this.state.messages=== false) {
			this.setState({ messages: true });
		}
	}
    handleTagsChange(event) {
		if(this.state.tags=== true) {
			this.setState({ tags: false });
		}
		if(this.state.tags=== false) {
			this.setState({ tags: true });
		}
	}
	
	componentDidMount() {
	
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
	Axios.get(BASE_URL + "/api/users/api/user/privacySettings/" + id)
				.then((res) => {
					if (res.status === 401) {
						this.setState({ errorHeader: "Bad credentials!", errorMessage: "Wrong username or password.", hiddenErrorAlert: false });
					} else if (res.status === 500) {
						this.setState({ errorHeader: "Internal server error!", errorMessage: "Server error.", hiddenErrorAlert: false });
					} else {
						this.setState({
							id: res.data.Id,
                            messages : res.data.AcceptMessages,
                            tags : res.data.AllowTags
						});
					}
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

        let dto = {
                          UserId : localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1),
                          AcceptMessages :  this.state.messages,
                         AcceptTags : this.state.tags
        };

			Axios.post(`${BASE_URL}/api/users/api/user/privacySettings/`, dto)
			.then((res) => {
                        if (res.status === 400) {
                            this.setState({ hiddenFailAlert: false, failHeader: "Bad request", failMessage: "Invalid argument." });
                        } else if (res.status === 500) {
                            this.setState({ hiddenFailAlert: false, failHeader: "Internal server error", failMessage: "Server error." });
                        } else if (res.status === 200) {
                            console.log("Success");
                            this.setState({
                                hiddenSuccessAlert: false,
                                successHeader: "Success",
                                successMessage: "You successfully updated your settings.",
                                hiddenEditInfo: true,
                            });
                        }
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
									<p style={{ color: "#6c757d", opacity: 1 }} >Recieve messages from non-followers: </p>
									<label class="switch" >
										<input checked ={this.state.messages} type="checkbox"  onChange={(e) => this.handleMessagesChange(e)}/>
										<span  class="slider round"></span>

									</label>
								</div>
                                <div>
									<p style={{ color: "#6c757d", opacity: 1 }} >Allow tags on posts and comments: </p>
									<label class="switch" >
										<input checked ={this.state.tags} type="checkbox"  onChange={(e) => this.handleTagsChange(e)}/>
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

export default PrivacySettings;
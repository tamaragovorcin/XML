import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import ImageUploader from 'react-images-upload';
import { YMaps, Map } from "react-yandex-maps";
import {GiThreeFriends} from "react-icons/gi"
import {MdPublic} from "react-icons/md"
import {CgFeed} from "react-icons/cg"
import { BASE_URL_USER } from "../constants.js";
import Axios from "axios";

const mapState = {
	center: [44, 21],
	zoom: 8,
	controls: [],
};
class VerifyModal extends Component {
   
    state = {
        name: "",
		lastName : "",
		category : ""
    }
	handleGenderChange(event) {

		this.setState({ category: event.target.value });
	}

    componentDidMount() {

		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		Axios.get(BASE_URL_USER + "/api/" + id)
				.then((res) => {
					if (res.status === 401) {
						this.setState({ errorHeader: "Bad credentials!", errorMessage: "Wrong username or password.", hiddenErrorAlert: false });
					} else if (res.status === 500) {
						this.setState({ errorHeader: "Internal server error!", errorMessage: "Server error.", hiddenErrorAlert: false });
					} else {
						this.setState({
							id: res.data.Id,
							username : res.data.ProfileInformation.Username,
							name: res.data.ProfileInformation.Name,
							lastName : res.data.ProfileInformation.LastName,
							email : res.data.ProfileInformation.Email,
							phoneNumber : res.data.ProfileInformation.PhoneNumber,
							gender : res.data.ProfileInformation.Gender,
							dateOfBirth  : res.data.ProfileInformation.DateOfBirth,
							webSite : res.data.WebSite,
							biography : res.data.Biography,
							private : res.data.Private
						});
					}
				})
				.catch ((err) => {
			console.log(err);
		});

	}
	render() {
	
		return (
			<Modal
				show={this.props.show}
				size="lg"
				dialogClassName="modal-60w-60h"
				aria-labelledby="contained-modal-title-vcenter"
				centered
				onHide={this.props.onCloseModal}
			>
				<Modal.Header closeButton>
                   <Modal.Title id="contained-modal-title-vcenter">{this.props.header}</Modal.Title>
				</Modal.Header>
				<Modal.Body>
				<div style={{marginBottom: "2rem"}}>
					<div style={{ marginLeft: "0rem" }}>
						<ImageUploader
											withIcon={false}
											buttonText='Add identitly document'
											//	onChange={this.props.onDrop}
											imgExtension={['.jpg', '.gif', '.png', '.gif']}
											withPreview={true}
						/>
					<div className="row section-design"  style={{ border:"1 solid black", }}>
						<div className="col-lg-8 mx-auto">
								<div className="control-group">
									
                                            <td>
												<label ><b>First name:</b>{this.state.name}</label>
											</td>
					
								</div>
								
								<div className="control-group">
									
                                             <td>
												<label ><b>Surname:</b>{this.state.lastName}</label>
											</td>
								</div>
                                
								<div className="control-group">
								<div style={{ color: "#6c757d", opacity: 1 }}>
									<p><input type="radio" checked value="influencer" name="category" onChange={(e) => this.handleGenderChange(e)} />Influencer</p>
									<p><input type="radio" value="sports" name="category" onChange={(e) => this.handleGenderChange(e)} />Sports</p>
									<p><input type="radio" value="business" name="category" onChange={(e) => this.handleGenderChange(e)} /> Business </p>
									<p><input type="radio" value="brand" name="category" onChange={(e) => this.handleGenderChange(e)} /> Brand </p>
									<p><input type="radio" value="new/media" name="category" onChange={(e) => this.handleGenderChange(e)} /> New/media </p>
									

								
								</div>
								</div>
								
								</div>

						</div>
						<div className="form-group text-center">
                        
							<div>
						
								<button style={{ width: "10rem", margin : "1rem",background:"#42DA61" }}  onClick={this.props.handleSendRequestVerification} className="btn btn-outline-secondary btn-sm">Send verify request<br/> <CgFeed/> </button>
								
							</div>
					
							
						</div>
					</div>

				</div>
				</Modal.Body>
				<Modal.Footer>
					<Button onClick={this.props.onCloseModal}>Close</Button>
				</Modal.Footer>
			</Modal>
		);
	}
}

export default VerifyModal;

import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import ImageUploader from 'react-images-upload';
import { YMaps, Map } from "react-yandex-maps";
import {GiThreeFriends} from "react-icons/gi"
import {MdPublic} from "react-icons/md"
import {CgFeed} from "react-icons/cg"
import { BASE_URL_USER } from "../constants.js";
import Axios from "axios";


class VerifyModal extends Component {
   
    state = {
        name: "",
		surname : "",
		category : ""
    }
	handleCategoryChange(event) {
		this.setState({ category: event.target.value });
	}
	handleNameChange = (event) => {
		this.setState({ name: event.target.value });
	};

	handleSurnameChange = (event) => {
		this.setState({ surname: event.target.value });
	};
    componentDidMount() {

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
											onChange={this.props.onDrop}
											imgExtension={['.jpg', '.gif', '.png', '.gif']}
											withPreview={true}
						/>
					<div className="row section-design"  style={{ border:"1 solid black", }}>
						<div className="col-lg-8 mx-auto">
						<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>First name:</label>
										<input
											placeholder="Name"
											class="form-control"
											type="text"
											id="name"
											onChange={this.handleNameChange}
											value={this.state.name}
										/>
									</div>
									
								</div>
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Surname:</label>
										<input
											placeholder="Surname"
											class="form-control"
											type="text"
											id="surname"
											onChange={this.handleSurnameChange}
											value={this.state.surname}
										/>
									</div>
									
								</div>
                                
								<div className="control-group">
								<div style={{ color: "#6c757d", opacity: 1 }}>
									<p><input type="radio"  value="INFLUENCER" name="category" onChange={(e) => this.handleCategoryChange(e)} />Influencer</p>
									<p><input type="radio" value="SPORTS" name="category" onChange={(e) => this.handleCategoryChange(e)} /> Sports</p>
									<p><input type="radio" value="NEW_MEDIA" name="category" onChange={(e) => this.handleCategoryChange(e)} /> New media </p>
									<p><input type="radio" value="BUSINESS" name="category" onChange={(e) => this.handleCategoryChange(e)} />Business</p>
									<p><input type="radio" value="BRAND" name="category" onChange={(e) => this.handleCategoryChange(e)} /> Brand</p>
									<p><input type="radio" value="ORGANIZATION" name="category" onChange={(e) => this.handleCategoryChange(e)} /> Organization </p>
								</div>
								</div>
								
								</div>

						</div>
						<div className="form-group text-center">
                        
							<div>
						
								<button 
									style={{ width: "10rem", margin : "1rem",background:"#42DA61" }} 
									 onClick={() => this.props.handleSendRequestVerification(this.state.name,this.state.surname,this.state.category)}
									 className="btn btn-outline-secondary btn-sm">
									 Send verify request<br/> 
									 <CgFeed/> </button>
								
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

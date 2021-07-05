import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import ImageUploader from 'react-images-upload';
import {GiThreeFriends} from "react-icons/gi"
import {CgFeed} from "react-icons/cg"


class AddVideoCampaignModal extends Component {
   

	
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
                            <label>
                                Upload a file: <br /><br />
                                <input type="file" className="btn btn-outline-secondary btn-sm" name="file" onChange={this.props.onChangeHandler} />
                            </label>
                            <br /><br />
					<div className="row section-design"  style={{ border:"1 solid black", }}>
						<div className="col-lg-8 mx-auto">
								
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Description:</label>
										<input
											placeholder="Description"
											className="form-control"
											id="email"
											
											type="text"
											onChange={this.props.handleDescriptionChange}
											value={this.props.description}
										/>
									</div>
								</div>
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Link to your web shop or article:</label>
										<input
											placeholder="Description"
											className="form-control"
											id="email"
											
											type="text"
											onChange={this.props.handleLinkChange}
											value={this.props.link}
										/>
									</div>
								</div>
								
								
							</div>
						</div>
						<div className="form-group text-center" >

									<button style={{ width: "10rem", margin : "1rem" }}  onClick={this.props.handleAddInfluencersModal} className="btn btn-outline-secondary btn-sm">Add influencers<br/> <CgFeed/> </button>		
									<button style={{ width: "10rem", margin : "1rem" }}  onClick={this.props.handleDefineTargetGroupModal} className="btn btn-outline-secondary btn-sm">Define target group<br/> <CgFeed/> </button>		
						</div>
					</div>
						<div className="form-group text-center">
                        						
								<button style={{ width: "10rem", margin : "1rem", background:"#37FF33" }} onClick={this.props.handleAddOneTimeCampaign} className="btn btn-outline-primary btn-sm">Add as one time campaign<GiThreeFriends/> </button>
								<button style={{ width: "10rem", margin : "1rem", background:"#37FF33" }} onClick={this.props.handleAddMultipleTimeCampaign} className="btn btn-outline-primary btn-sm">Add as multiple time campaign<GiThreeFriends/> </button>

						
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

export default AddVideoCampaignModal;

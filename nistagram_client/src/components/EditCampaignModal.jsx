import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import { BASE_URL } from "../constants.js";
import Axios from "axios";
class EditCampaignModal extends Component {
    state = {
      
	};
   

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
				
					 <table className="table" style={{ width: "100%", marginTop: "3rem" }}>
                            <tbody>
                            <div className="control-group">
                                <label>Date of publising:</label>
								<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
									<input
										placeholder="Your comment"
										className="form-control"
										id="comment"
										type="date"
										onChange={this.props.handleCampaignDateChange}
										value={this.props.campaignDate}
									/>
								</div>
                                <label>Time of publising:</label>
								<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
									<input
										placeholder="Your comment"
										className="form-control"
										id="comment"
										type="time"
										onChange={this.props.handleCampaignTimeChange}
										value={this.props.campaignTime}
									/>
								</div>
                                <label>Link to web shop or article:</label>
								<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
									<input
										placeholder="Your comment"
										className="form-control"
										id="comment"
										type="text"
										onChange={this.props.handleCampaignLinkChange}
										value={this.props.campaignLink}
									/>
								</div>
                                <label>Description:</label>
								<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
									<input
										placeholder="Your comment"
										className="form-control"
										id="comment"
										type="text"
										onChange={this.props.handleCampaignDescriptionChange}
										value={this.props.campaignDescription}
									/>
								</div>
                                <Button
									style={{ background: "#1977cc", marginTop: "15px", marginLeft: "40%", width: "20%" }}
                                    onClick={() =>  this.props.handleChangeCampaign(this.props.campaignForEdit, this.props.campaignDate, this.props.campaignTime,this.props.campaignLink, this.props.campaignDescription)}
									className="btn btn-primary btn-md"
									id="sendMessageButton"
									type="button"
								>
									Save changes
									</Button>
								
							</div>
                            </tbody>
                        </table>
						
				</Modal.Body>
				
			</Modal>
		);
	}
}

export default EditCampaignModal;

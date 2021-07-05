import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";

class OneTimeCampaignModal extends Component {
    state = {
		campaignDate: "",
        campaignTime : "",
        campaignTargetGroup : ""
	};
   
    handleCampaignDateChange = (event) =>{
        this.setState({ campaignDate: event.target.value });

    }
    handleCampaignTimeChange = (event) =>{
        this.setState({ campaignTime: event.target.value });

    }
    handleCampaignTargetGroupChange = (event) =>{
        this.setState({ campaignTargetGroup: event.target.value });

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
				
					 <table className="table" style={{ width: "100%", marginTop: "3rem" }}>
                            <tbody>
                            <div className="control-group">
                            <label>Select the date on which campaign will be displayed to users:</label>
								<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
									<input
										placeholder="Your comment"
										className="form-control"
										id="comment"
										type="date"
										onChange={this.handleCampaignDateChange}
										value={this.state.campaignDate}
									/>
								</div>
                                <label>Select time:</label>
								<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
									<input
										placeholder="Your comment"
										className="form-control"
										id="comment"
										type="time"
										onChange={this.handleCampaignTimeChange}
										value={this.state.campaignTime}
									/>
								</div>
                           
                                <Button
									style={{ background: "#1977cc", marginTop: "15px", marginLeft: "40%", width: "20%" }}
									onClick={() => this.props.handleAddOneTimeCampaign(this.state.campaignDate,this.state.campaignTime, "feed")}
									className="btn btn-primary btn-md"
									id="sendMessageButton"
									type="button"
								>
									Publish as feed
									</Button>
								<Button
									style={{ background: "#1977cc", marginTop: "15px", marginLeft: "40%", width: "20%" }}
									onClick={() => this.props.handleAddOneTimeCampaign(this.state.campaignDate,this.state.campaignTime,"story")}
									className="btn btn-primary btn-md"
									id="sendMessageButton"
									type="button"
								>
									Publish as story
									</Button>
							</div>
                            </tbody>
                        </table>
						
				</Modal.Body>
				
			</Modal>
		);
	}
}

export default OneTimeCampaignModal;

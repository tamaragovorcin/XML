import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";

class MultipleTimeCampaignModal extends Component {
    state = {
		campaignStartDate: "",
        campaignEndDate : "",
        campaignNumberOfRepetition : 0,
	};
   
    handleCampaignStartDateChange = (event) =>{
        this.setState({ campaignStartDate: event.target.value });

    }
    handleCampaignEndDateChange = (event) =>{
        this.setState({ campaignEndDate: event.target.value });

    }
    handleCampaignNumberOfRepetitionChange = (event) =>{
        this.setState({ campaignNumberOfRepetition: event.target.value });

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
                            <label>Select the date on which campaign will start to be displayed to users:</label>
								<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
									<input
										placeholder="Your comment"
										className="form-control"
										id="comment"
										type="date"
										onChange={this.handleCampaignStartDateChange}
										value={this.state.campaignStartDate}
									/>
								</div>
                                <label>End date:</label>
								<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
									<input
										placeholder="Your comment"
										className="form-control"
										id="comment"
										type="date"
										onChange={this.handleCampaignEndDateChange}
										value={this.state.campaignEndDate}
									/>
								</div>
                                <label>Destired number of campaign repetitions in one day:</label>
								<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
									<input
										placeholder="Your comment"
										className="form-control"
										id="comment"
										type="number"
										onChange={this.handleCampaignNumberOfRepetitionChange}
										value={this.state.campaignNumberOfRepetition}
									/>
								</div>
                              
                                <Button
									style={{ background: "#1977cc", marginTop: "15px", marginLeft: "40%", width: "20%" }}
									onClick={() => this.props.handleAddMultipleTimeCampaign(this.state.campaignStartDate,this.state.campaignEndDate,this.state.campaignNumberOfRepetition,"feed")}
									className="btn btn-primary btn-md"
									id="sendMessageButton"
									type="button"
								>
									Publish ass feed 
								</Button>
									<Button
									style={{ background: "#1977cc", marginTop: "15px", marginLeft: "40%", width: "20%" }}
									onClick={() => this.props.handleAddMultipleTimeCampaign(this.state.campaignStartDate,this.state.campaignEndDate,this.state.campaignNumberOfRepetition,"story")}
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

export default MultipleTimeCampaignModal;

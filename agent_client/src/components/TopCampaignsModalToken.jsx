import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import {GiThreeFriends} from "react-icons/gi"


class TopCampaignsModalToken extends Component {
   

	
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
						
					<div className="row section-design"  style={{ border:"1 solid black", }} hidden={this.props.noPicture}>
						<div className="col-lg-8 mx-auto">
								
			
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Token:</label>
										<input
											placeholder="Token"
											className="form-control"
											id="email"
											
											type="text"
											onChange={this.props.handleTokenChange}
											value={this.props.token}
										/>
									</div>
								</div>
								
							</div>
						</div>
				
					</div>
						<div className="form-group text-center">
                        						
								<button style={{ width: "10rem", margin : "1rem", background:"#37FF33" }} onClick={this.props.hadleGetTopCampaigns} className="btn btn-outline-primary btn-sm">Get top campaigns<GiThreeFriends/> </button>
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

export default TopCampaignsModalToken;

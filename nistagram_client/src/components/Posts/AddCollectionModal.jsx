import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";

class AddCollectionModal extends Component {
	state = {
		name: "",
	};

	handleNameChange = (event) => {
		this.setState({ name: event.target.value });
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
					
					<div className="control-group">
									
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Collection name:</label>
										<input
											placeholder="Name"
											className="form-control"
											type="text"
											id="name"
											onChange={this.handleNameChange}
											value={this.state.name}
										/>
									</div>
									<div className="text-danger" style={{ display: this.props.collectionNameError }}>
										Name must be entered.
									</div>
								</div>
								
                                <div className="form-group text-center">
									
                                    <button
                                        style={{ background: "#1977cc", marginTop: "15px" }}
                                        onClick={() => this.props.handleAddCollection(this.state.name)}
                                        className="btn btn-primary btn-xl"
                                        id="sendMessageButton"
                                        type="button"
                                    >
                                        Add
                                    </button>
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

export default AddCollectionModal;
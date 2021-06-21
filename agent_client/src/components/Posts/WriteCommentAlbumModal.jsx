import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";

class WriteCommentAlbumModal extends Component {
    state = {
		comment: ""
	};
    handleCommentChange = (event) => {
		this.setState({ comment: event.target.value });
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
                                <label>Write your comment:</label>
								<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
									<input
										placeholder="Your comment"
										className="form-control"
										id="comment"
										type="text"
										onChange={this.handleCommentChange}
										value={this.state.comment}
									/>
								</div>
                                <Button
									style={{ background: "#1977cc", marginTop: "15px", marginLeft: "40%", width: "20%" }}
									onClick={() => this.props.handleAddCommentAlbum(this.state.comment)}
									className="btn btn-primary btn-md"
									id="sendMessageButton"
									type="button"
								>
									Comment
									</Button>
								
							</div>
                            </tbody>
                        </table>
						
				</Modal.Body>
				
			</Modal>
		);
	}
}

export default WriteCommentAlbumModal;

import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";

class DisposableImageModal extends Component {
   

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
                            <img
							  	style={{margin:"5%" }}
								
                                className="img-fluid"
                                src={"http://localhost:80/api/messages/api/disposableImage/file/"+this.props.media}
                                width="500px"
								height="500px"
                                alt="description"
                              />
                              
                            </tbody>
                        </table>
						
				</Modal.Body>
				<Modal.Footer>
					<Button onClick={this.props.onCloseModal}>Close</Button>
				</Modal.Footer>
			</Modal>
		);
	}
}

export default DisposableImageModal;

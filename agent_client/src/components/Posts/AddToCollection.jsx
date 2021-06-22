import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";

class AddToCollection extends Component {
   

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
                                {this.props.collections.map((collection) => (
                                    <tr id={collection.id} key={collection.id}>
                                        <tr>
                                            <label><b>{collection.name}</b></label>
                                        </tr>
                                    </tr>

                                ))}
                              
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

export default AddToCollection;

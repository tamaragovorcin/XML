import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import Select from 'react-select';

class AddTagsModal extends Component {
   

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
					<Modal.Title id="contained-modal-title-vcenter"></Modal.Title>
				</Modal.Header>
				<Modal.Body>
				
                    <div className="container d-flex align-items-center">
                        <div class="input-group rounded" style={{ marginLeft: "20%", marginRight: "10%" }}>

                            <div style={{ width: '300px' }}>
                                <Select
                                    style={{ width: `$500px` }}
                                    className="select-custom-class"
                                    label="Single select"
                                    options={this.props.followingUsers}
                                    onChange ={e => this.props.handleChangeTags(e)}
                                />

                            </div>

                        </div>
                    </div>
                    <div>
                        <label className="logo mr-auto" style={{ fontFamily: "Trattatello, fantasy" }}>
                            Tagged:
                        </label>
                         <table className="table" style={{ width: "100%" }}>
                             <tbody>
                                     {this.props.taggedOnPost.map((follower) => (
                                                    <tr id={follower.value} key={follower.value}>
                                                        
                                                        <td >
                                                            <div style={{ marginTop: "1rem"}}>
                                                                <b>Username: </b> {follower.label}
                                                            </div>
                                                        </td>
                                                    </tr>

                                                ))}
                                                <tr>
                                            <td></td>
                                            <td></td>
                                        <td></td>
                                    </tr>
                                </tbody>
                         </table>
                    </div>

				</Modal.Body>
				<Modal.Footer>
					<Button onClick={this.props.onCloseModal}>Close</Button>
				</Modal.Footer>
			</Modal>
		);
	}
}

export default AddTagsModal;
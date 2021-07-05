import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import Select from 'react-select';

class AddInfluencerModal extends Component {
   

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
                                    options={this.props.influencers}
                                    onChange ={e => this.props.handleChangeInfluencers(e)}
                                />

                            </div>

                        </div>
                    </div>
                    <div>
                        <label className="logo mr-auto" style={{ fontFamily: "Trattatello, fantasy" }}>
                            Chosen influencers:
                        </label>
                         <table className="table" style={{ width: "100%" }}>
                             <tbody>
                                     {this.props.choosenInfluencers.map((user) => (
                                                    <tr id={user.value} key={user.value}>
                                                        
                                                        <td >
                                                            <div style={{ marginTop: "1rem"}}>
                                                                <b>Username: </b> {user.label}
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

export default AddInfluencerModal;
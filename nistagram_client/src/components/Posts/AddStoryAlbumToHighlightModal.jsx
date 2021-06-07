import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";

class AddStoryAlbumToHighlightModal extends Component {
   

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
				
						<table className="table" style={{ width: "100%", marginTop: "3rem" }}>
                            <tbody>
                                {this.props.highlightsAlbums.map((highlight) => (
                                    <tr id={highlight.Id} key={highlight.Id}>
                                       
                                        <td>
                                            <div>
                                                <br></br>
                                                <b>Name: </b> {highlight.Name}
                                            </div>
                                           
                                
                                        </td>
                                        <td >
                                        <div style={{marginLeft:'55%'}}>
                                               <td>
                                                    <br></br>
                                                    <button style={{height:'30px'},{verticalAlign:'center'},{marginTop:'2%'}} className="btn btn-outline-secondary mt-1"
                                                        onClick={() => this.props.addStoryAlbumToHighlight(highlight.Id)}
                                                         type="button">
                                                            <i className="icofont-subscribe mr-1"> </i>
                                                             Add to highlight album
                                                             </button >
                                                </td>
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

				</Modal.Body>
				<Modal.Footer>
					<Button onClick={this.props.onCloseModal}>Close</Button>
				</Modal.Footer>
			</Modal>
		);
	}
}

export default AddStoryAlbumToHighlightModal;
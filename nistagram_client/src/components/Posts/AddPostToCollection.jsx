import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";

class AddPostToCollection extends Component {
   

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
					<Modal.Title id="contained-modal-title-vcenter">{this.props.selectedPlayerName}</Modal.Title>
				</Modal.Header>
				<Modal.Body>
				
						<table className="table" style={{ width: "100%", marginTop: "3rem" }}>
                            <tbody>
                                {this.props.collections.map((collection) => (
                                    <tr id={collection.Id} key={collection.Id}>
                                       
                                        <td>
                                            <div>
                                                <br></br>
                                                <b>Name: </b> {collection.Name}
                                            </div>
                                           
                                
                                        </td>
                                        <td >
                                        <div style={{marginLeft:'55%'}}>
                                               <td>
                                                    <br></br>
                                                    <button style={{height:'30px'},{verticalAlign:'center'},{marginTop:'2%'}} className="btn btn-outline-secondary mt-1"
                                                        onClick={() => this.props.addPostToCollection(collection.Id)}
                                                         type="button">
                                                            <i className="icofont-subscribe mr-1"> </i>
                                                             Add to collection</button
                                                             >
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

export default AddPostToCollection;
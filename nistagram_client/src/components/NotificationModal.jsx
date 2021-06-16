import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import Select from 'react-select';

class NotificationModal extends Component {
   

	render() {
		return (
			<Modal
				show={this.props.show}
				size="lg"
				dialogClassName="modal-40w-40h"
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
                               

                            </div>

                        </div>
                    </div>
                    <div>
                    <div>
									<p style={{ color: "#6c757d", opacity: 1 }} >Enable notifications for posts: </p>
									<label class="switch" >
										<input checked ={this.props.postsNotification} type="checkbox"  onChange={ this.props.handlePostsNotificationChange}/>
										<span  class="slider round"></span>

									</label>
								</div>
                                <div>
									<p style={{ color: "#6c757d", opacity: 1 }} >Enable notifications for stories: </p>
									<label class="switch" >
										<input checked ={this.props.storiesNotification} type="checkbox"  onChange={ this.props.handleStoriesNotificationChange}/>
										<span  class="slider round"></span>

									</label>
								</div>
                                <button  className="btn btn-primary mt-1" onClick={() => this.props.handleNotifications()} type="button"><i className="icofont-subscribe mr-1"></i>Save</button>

                    </div>

				</Modal.Body>
				<Modal.Footer>
					<Button onClick={this.props.onCloseModal}>Close</Button>
				</Modal.Footer>
			</Modal>
		);
	}
}

export default NotificationModal;
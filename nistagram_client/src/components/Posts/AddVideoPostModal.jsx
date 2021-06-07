import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import ImageUploader from 'react-images-upload';
import { YMaps, Map } from "react-yandex-maps";
import {GiThreeFriends} from "react-icons/gi"
import {MdPublic} from "react-icons/md"
import {CgFeed} from "react-icons/cg"

const mapState = {
	center: [44, 21],
	zoom: 8,
	controls: [],
};
class AddVideoPostModal extends Component {
   

	
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
                        <form onSubmit={this.props.handleSubmit}>
                            <label>
                                Upload a file: <br /><br />
                                <input type="file" className="btn btn-outline-secondary btn-sm" name="file" onChange={this.props.onChangeHandler} />
                            </label>
                            <br /><br />
                        </form >
					<div className="row section-design"  style={{ border:"1 solid black", }}>
						<div className="col-lg-8 mx-auto">
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Address:</label>
										<input className="form-control" id="suggest" ref={this.props.addressInput} placeholder="Address" />
									</div>
									<YMaps
										query={{
											load: "package.full",
											apikey: "b0ea2fa3-aba0-4e44-a38e-4e890158ece2",
											lang: "en_RU",
										}}
									>
										<Map
											style={{ display: "none" }}
											state={mapState}
											onLoad={this.props.onYmapsLoad}
											instanceRef={(map) => (this.map = map)}
											modules={["coordSystem.geo", "geocode", "util.bounds"]}
										></Map>
									</YMaps>
								
									<div className="text-danger" style={{ display: this.props.addressNotFoundError }}>
										Sorry. Address not found. Try different one.
									</div>
								</div>
								
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Description:</label>
										<input
											placeholder="Description"
											className="form-control"
											id="email"
											
											type="text"
											onChange={this.props.handleDescriptionChange}
											value={this.props.description}
										/>
									</div>
								</div>
                                
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Hashtags:</label>
										<input
											placeholder="Hashtags"
											class="form-control"
											type="text"
											id="name"
											onChange={this.props.handleHashtagsChange}
											value={this.props.hashtags}
										/>
									</div>
								</div>
								<div className="control-group">
									<button style={{ width: "10rem", margin : "1rem" }}  onClick={this.props.handleAddTagsModal} className="btn btn-outline-secondary btn-sm">Add tags<br/> <CgFeed/> </button>
								</div>
								</div>

						</div>
						<div className="form-group text-center">
                        
							<div>
						
								<button style={{ width: "10rem", margin : "1rem",background:"#37FF33" }}  onClick={this.props.handleAddFeedPost} className="btn btn-outline-secondary btn-sm">Add as feed post<br/> <CgFeed/> </button>
								<button style={{ width: "10rem", margin : "1rem" ,background:"#37FF33"}} onClick={this.props.handleAddStoryPost} className="btn btn-outline-secondary btn-sm">Add as story<br/> <MdPublic/> </button>
								<button style={{ width: "10rem", margin : "1rem", background:"#37FF33" }} onClick={this.props.handleAddStoryPostCloseFriends} className="btn btn-outline-secondary btn-sm">Story for close friends<GiThreeFriends/> </button>
							</div>
					
							<div  hidden={this.props.hiddenMultiple}>

								<button style={{ width: "10rem" , margin : "1rem", background:"#37FF33" }} onClick={this.props.handleAddFeedPostAlbum} className="btn btn-outline-secondary btn-sm">Feed album<br/> <CgFeed/> </button>
								<button style={{ width: "10rem", margin : "1rem", background:"#37FF33"  }} onClick={this.props.handleAddStoryPostAlbum} className="btn btn-outline-secondary btn-sm">Story album<br/> <MdPublic/> </button>
								<button style={{ width: "10rem", margin : "1rem", background:"#37FF33"  }} onClick={this.props.handleAddStoryPostAlbumCloseFriends} className="btn btn-outline-secondary btn-sm">Story album close friends<GiThreeFriends/> </button>
							</div>
						</div>
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

export default AddVideoPostModal;

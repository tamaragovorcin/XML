import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import { YMaps, Map } from "react-yandex-maps";

const mapState = {
	center: [44, 21],
	zoom: 8,
	controls: [],
};
class TargetGroupModal extends Component {
   

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
				
                <div className="row section-design"  style={{ border:"1 solid black", }} >
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
								
                                
								</div>
								
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Gender:</label>
										<div style={{ color: "#6c757d", opacity: 1 }}>
									<p><input type="radio"  value="MALE" name="gender" onChange={(e) => this.props.handleGenderChange(e)} />Male</p>
									<p><input type="radio" value="FEMALE" name="gender" onChange={(e) => this.props.handleGenderChange(e)} /> Female</p>
									<p><input type="radio" value="OTHER" name="gender" onChange={(e) => this.props.handleGenderChange(e)} /> Other </p>
									
								</div>
									</div>
								</div>
                                
								<div className="control-group">
									<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
										<label>Person born betwen</label>
										<input
											placeholder="First date"
											class="form-control"
											id="date"
											type="date"
											onChange={this.props.handleDateOneChange}
											value={this.props.selectedDateOne}
										/>
                                        <input
											placeholder="Second date"
											class="form-control"
											id="date"
											type="date"
											onChange={this.props.handleDateTwoChange}
											value={this.props.selectedDateTwo}
										/>
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

export default TargetGroupModal;
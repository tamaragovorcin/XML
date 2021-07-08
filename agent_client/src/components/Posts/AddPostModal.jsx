import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import ImageUploader from 'react-images-upload';
import { YMaps, Map } from "react-yandex-maps";
import { GiThreeFriends } from "react-icons/gi"
import { MdPublic } from "react-icons/md"
import { CgFeed } from "react-icons/cg"
import Axios from "axios";
import { BASE_URL_AGENT } from "../../constants.js";
import ModalDialog from "../../components/ModalDialog";
import getAuthHeader from "../GetHeader";
class AddPostModal extends Component {
	constructor(props) {
		super(props);
		this.state = {
			name: "",
			price: "",
			pictures: [],
			quantity: "",
			openModal: false,
			help: [],
			fileUploadOngoing: false,

	
		}
		this.onDrop = this.onDrop.bind(this);

	}
	
	onDrop(picture) {
		this.setState({
			pictures: [],
		});
		this.setState({
			pictures: this.state.pictures.concat(picture),
		});

	}

	test(pic, userIdd, feedId) {
		this.setState({
			fileUploadOngoing: true
		});

		const fileInput = document.querySelector("#fileInput");
		const formData = new FormData();

		formData.append("file", pic);
		formData.append("test", "StringValueTest");

		const options = {
			method: "POST",
			body: formData

		};
		fetch(BASE_URL_AGENT + "/api/image/" + userIdd , options, {  headers: { Authorization: getAuthHeader() } });
	}
	sendRequestForFeedAlbum(feedPostDTO) {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)

		Axios.post(BASE_URL_AGENT + "/api/product/" + id, feedPostDTO, {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
			
					//this.props.redirect=true
				
				let feedId = res.data;
				this.state.pictures.forEach((pic) => {
					this.test(pic, id, feedId);
				});

		
				this.setState({ textSuccessfulModal: "You have successfully added album feed post." });
				

			})
			.catch((err) => {
				console.log(err);
			});

			//this.props.redirect=true
		
		this.setState({ openModal: true });
	}
	handleAddFeedPostAlbum = () => {
		let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)
	
		let h = []
		this.state.pictures.forEach((pic) => {
			h.push(pic.name)
			this.setState({
				help: this.state.help.concat(pic.name),
			});
		});



		const product = {
			user: id,
			name: this.state.name,
			price: this.state.price,
			quantity: this.state.quantity,
			media: h
		};
		this.sendRequestForFeedAlbum(product);







	}



	handleQuantityChange = (e) => {
		this.setState({ quantity: e.target.value });
	}
	handleNameChange = (e) => {
		this.setState({ name: e.target.value });
	}

	handlePriceChange = (e) => {
		this.setState({ price: e.target.value });
	}

	handleModalClose = ()=>{
		this.setState({openModal: false})
	}
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
					<div style={{ marginBottom: "2rem" }}>
						<div style={{ marginLeft: "0rem" }}>
							<ImageUploader
								withIcon={false}
								buttonText='Add new photo/video'
								onChange={this.onDrop}
								imgExtension={['.jpg', '.gif', '.png', '.gif']}
								withPreview={true}
							/>
							<div className="row section-design" style={{ border: "1 solid black", }} hidden={this.state.noPicture}>
								<div className="col-lg-8 mx-auto">



									<div className="control-group">
										<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
											<label>Product name:</label>
											<input
												placeholder="Product name"
												className="form-control"
												id="name"

												type="text"
												onChange={this.handleNameChange}
												value={this.state.name}
											/>
										</div>
									</div>
									<div className="control-group">
										<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
											<label>Price:</label>
											<input
												placeholder="Price"
												className="form-control"
												id="price"

												type="text"
												onChange={this.handlePriceChange}
												value={this.state.price}
											/>
										</div>
									</div>

									<div className="control-group">
										<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
											<label>Quantity:</label>
											<input
												placeholder="Quantity"
												className="form-control"
												id="quantity"

												type="text"
												onChange={this.handleQuantityChange}
												value={this.state.quantity}
											/>
										</div>
									</div>
								</div>

							</div>
							<div className="form-group text-center">

								<div>

									<button style={{ width: "10rem", margin: "1rem", background: "#37FF33" }} onClick={this.handleAddFeedPostAlbum} className="btn btn-outline-secondary btn-sm">Add<br /> </button>

								</div>
							</div>
						</div>

					</div>
					<ModalDialog
                    show={this.state.openModal}
                    onCloseModal={this.handleModalClose}
                    header="Success"
                    text="You have successfully added new item."
                />
				</Modal.Body>
				<Modal.Footer>
					<Button onClick={this.props.onCloseModal}>Close</Button>
				</Modal.Footer>
			</Modal>
		);
	}
}

export default AddPostModal;

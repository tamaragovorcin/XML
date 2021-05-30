
import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { Link,Button } from "react-router-dom";
import playerLogo from "../static/me.jpg";
import profileImage from "../static/profileImage.jpg"

import { BASE_URL } from "../constants.js";
import ImageUploader from 'react-images-upload';

class Favorites extends React.Component {
	constructor(props) {
		super(props);

		this.onDrop = this.onDrop.bind(this);

	}
	state = {
        favorites : [],
		username: "",
		numberPosts: 0,
		numberFollowing: 0,
		numberFollowers: 0,
		biography: "",
		photos: [],
		pictures: [],
		picture: "",
		hiddenOne: true,
		hiddenMultiple: true,
	}
	onDrop(picture) {
		this.setState({
			pictures: this.state.pictures.concat(picture),
		});

		let pomoc = this.state.pictures.length;
		pomoc = pomoc + 1;
	
		if(pomoc === 1){
			this.setState({
				hiddenOne: false,
			});
			this.setState({
				hiddenMultiple: true,
			});
		}
		else if(pomoc >= 2){
			this.setState({
				hiddenOne: true,
			});
			this.setState({
				hiddenMultiple: false,
			});
		}


	}

	

	test(pic) {

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
		fetch(BASE_URL + "/api/items/upload", options);
	}


	componentDidMount() {
		this.handleGetBasicInfo()
		this.handleGetFavorites()

	}
	handleGetBasicInfo = () => {
		this.setState({ numberPosts: 10 });
		this.setState({ numberFollowing: 600 });
		this.setState({ numberFollowers: 750 });
		this.setState({ biography: "bla bla bla" });
		this.setState({ username: "USERNAME" });
	}
    handleGetFavorites=()=>{
        let list = []
        let headline1 = "Food"
        let headLine2 = "Workout"
        let headline3 = "Makeup"
        list.push({id: 1, headline :"Food", images:[playerLogo,profileImage]})
        list.push({id: 2, headline :"MakeUp", images:[playerLogo,profileImage]})

        this.setState({ favorites: list });

    }

	

	handleGetPhotos = () => {
		let list = []
		let comments1 = []
		let comments2 = []
		let comment1 = { id: 1, user: "USER 1 ", text: "very nice" }
		let comment11 = { id: 2, user: "USER 2 ", text: "cool" }
		let comment111 = { id: 3, user: "USER 3 ", text: "vau" }
		comments1.push(comment1)
		comments1.push(comment11)
		comments1.push(comment111)

		let comment2 = { id: 4, user: "USER 55443 ", text: "i like it" }
		let comment22 = { id: 5, user: "USER 11111 ", text: "ugly" }
		let comment222 = { id: 6, user: "USER 33333 ", text: "awesome" }
		comments2.push(comment2)
		comments2.push(comment22)
		comments2.push(comment222)

		let photo1 = { id: 1, photo: playerLogo, numLikes: 52, numDislikes: 2, comments: comments1 }
		let photo2 = { id: 2, photo: playerLogo, numLikes: 45, numDislikes: 0, comments: comments2 }
		list.push(photo1)
		list.push(photo2)

		this.setState({ photos: list });

	}
	

	render() {
		return (
			<React.Fragment>
				<TopBar />
				<Header />


				<div className="d-flex align-items-top">
					<div className="container-fluid" style={{ marginTop: "10rem", marginRight: "11rem" }}>
						<table className="table" style={{ width: "100%" }}>
							<tbody>

								<tr>
									<td width="130em">
										<img
											className="img-fluid"
											src={playerLogo}
											width="70em"
											alt="description"
										/>
									</td>

									<td>
										<div>
											<td>
												<label >{this.state.username}</label>
											</td>
											<td>
												<Link to="/userChangeProfile" className="btn btn-outline-secondary btn-sm">Edit profile</Link>

											</td>
										</div>
										<div>
											<td>
												<label ><b>{this.state.numberPosts}</b> posts</label>
											</td>
											<td>
												<label ><b>{this.state.numberFollowers}</b> followers</label>
											</td>
											<td>
												<label ><b>{this.state.numberFollowing}</b> following</label>
											</td>

										</div>
										<div>
											<td>
												<label >{this.state.biography}</label>
											</td>
										</div>

										<div style={{ marginLeft: "0rem" }}><ImageUploader
											withIcon={false}
											buttonText='Add new photo/video'
											onChange={this.onDrop}
											imgExtension={['.jpg', '.gif', '.png', '.gif']}
											withPreview={true}
										/>
										<div style={{ marginLeft: "19rem" }} hidden={this.state.hiddenOne}>
												<Link style={{ width: "10rem" }} to="/userChangeProfile" className="btn btn-outline-secondary btn-sm">Add as feed post </Link>
												<a style={{ padding: "25px" }}></a>
												<Link style={{ width: "10rem" }} to="/userChangeProfile" className="btn btn-outline-secondary btn-sm">Add as story </Link>
												</div>
									
										<div style={{ marginLeft: "19rem" }} hidden={this.state.hiddenMultiple}>
												<Link style={{ width: "10rem" }} to="/userChangeProfile" className="btn btn-outline-secondary btn-sm">Add as feed album </Link>
												<a style={{ padding: "25px" }}></a>
												<Link style={{ width: "10rem" }} to="/userChangeProfile" className="btn btn-outline-secondary btn-sm">Add as album story </Link>
												</div>
												</div>
									</td>
									
								</tr>
							</tbody>
						</table>
					</div>


					





				</div>
				<div
  id="carouselMultiItemExample"
  class="carousel slide carousel-dark text-center"
  data-mdb-ride="carousel"
>
  <div class="d-flex justify-content-center mb-4">
    <button
      class="carousel-control-prev position-relative"
      type="button"
      data-mdb-target="#carouselMultiItemExample"
      data-mdb-slide="prev"
    >
      <span class="carousel-control-prev-icon" aria-hidden="true"></span>
      <span class="visually-hidden">Previous</span>
    </button>
    <button
      class="carousel-control-next position-relative"
      type="button"
      data-mdb-target="#carouselMultiItemExample"
      data-mdb-slide="next"
    >
      <span class="carousel-control-next-icon" aria-hidden="true"></span>
      <span class="visually-hidden">Next</span>
    </button>
  </div>
  <div class="carousel-inner py-4">
    <div class="carousel-item active">
      <div class="container">
        <div class="row">
          <div class="col-lg-4">
            <div class="card">
              <img
                src="https://mdbootstrap.com/img/new/standard/nature/181.jpg"
                class="card-img-top"
                alt="..."
              />
              <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">
                  Some quick example text to build on the card title and make up the bulk
                  of the card's content.
                </p>
                <a href="#!" class="btn btn-primary">Button</a>
              </div>
            </div>
          </div>

          <div class="col-lg-4 d-none d-lg-block">
            <div class="card">
              <img
                src="https://mdbootstrap.com/img/new/standard/nature/182.jpg"
                class="card-img-top"
                alt="..."
              />
              <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">
                  Some quick example text to build on the card title and make up the bulk
                  of the card's content.
                </p>
                <a href="#!" class="btn btn-primary">Button</a>
              </div>
            </div>
          </div>

          <div class="col-lg-4 d-none d-lg-block">
            <div class="card">
              <img
                src="https://mdbootstrap.com/img/new/standard/nature/183.jpg"
                class="card-img-top"
                alt="..."
              />
              <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">
                  Some quick example text to build on the card title and make up the bulk
                  of the card's content.
                </p>
                <a href="#!" class="btn btn-primary">Button</a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

  

    <div class="carousel-item">
      <div class="container">
        <div class="row">
          <div class="col-lg-4 col-md-12 mb-4 mb-lg-0">
            <div class="card">
              <img
                src="https://mdbootstrap.com/img/new/standard/nature/187.jpg"
                class="card-img-top"
                alt="..."
              />
              <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">
                  Some quick example text to build on the card title and make up the bulk
                  of the card's content.
                </p>
                <a href="#!" class="btn btn-primary">Button</a>
              </div>
            </div>
          </div>

          <div class="col-lg-4 mb-4 mb-lg-0 d-none d-lg-block">
            <div class="card">
              <img
                src="https://mdbootstrap.com/img/new/standard/nature/188.jpg"
                class="card-img-top"
                alt="..."
              />
              <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">
                  Some quick example text to build on the card title and make up the bulk
                  of the card's content.
                </p>
                <a href="#!" class="btn btn-primary">Button</a>
              </div>
            </div>
          </div>

          <div class="col-lg-4 mb-4 mb-lg-0 d-none d-lg-block">
            <div class="card">
              <img
                src="https://mdbootstrap.com/img/new/standard/nature/189.jpg"
                class="card-img-top"
                alt="..."
              />
              <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">
                  Some quick example text to build on the card title and make up the bulk
                  of the card's content.
                </p>
                <a href="#!" class="btn btn-primary">Button</a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>


                    
                        
			</React.Fragment>
		);
	}
}

export default Favorites;

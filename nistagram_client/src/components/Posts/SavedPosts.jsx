
import React from "react";


import { BASE_URL } from "../../constants.js";
import ImageUploader from 'react-images-upload';
import { Collections } from "@material-ui/icons";

class SavedPosts extends React.Component {
	constructor(props) {
		super(props);

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


	render() {
		return (
			<React.Fragment>
				


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
	{this.props.collections.map((collection) => (
                               
      <div id={collection.Id} key = {collection.Id} class="container">
        <div class="row">
          <div class="col-lg-4">
            <div class="card">
              <img
                src={collection.Image}
                class="card-img-top"
                
              />
              <div class="card-body">
                <h5 class="card-title">{collection.Name}</h5>
              </div>
            </div>
          </div>
		  
        </div>
      </div>
	  ))}
    </div>

  

   
  </div>
</div>


                    
                        
			</React.Fragment>
		);
	}
}

export default SavedPosts;


import React from "react";
import collection from "../../static/collection.png";


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
      

                        
			</React.Fragment>
		);
	}
}

export default SavedPosts;

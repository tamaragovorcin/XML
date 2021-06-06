import React, { Component } from "react";
import Axios from "axios";

import { BASE_URL_USER_INTERACTION } from "../constants.js";
class FollowRequest extends Component {
	state = {
	followers: []	

    };


  handleAccept  = (followerId) => {
    let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
		const acceptedDTO = { following: id, follower: followerId};
		Axios.post(BASE_URL_USER_INTERACTION  + "/api/acceptFollowRequest", acceptedDTO)
				.then((res) => {
					
						console.log(res.data)
					
				})
				.catch ((err) => {
			console.log(err);
		});
  
    
}
handleDelete  = (followerId) => {
    let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
    const dto = { following: id, follower: followerId};
    Axios.post(BASE_URL_USER_INTERACTION  + "/api/deleteFollowRequest", dto)
            .then((res) => {
                
                    console.log(res.data)
                
            })
            .catch ((err) => {
        console.log(err);
    });
}
    componentDidMount() {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = {id: id}
        Axios.post(BASE_URL_USER_INTERACTION + "/api/user/followRequests", dto)
			.then((res) => {
				this.setState({ followers: res.data });
               
				
			})
			.catch((err) => {

				console.log(err)
			});

	}
	

	render() {
		return (
            <div class="container-fluid">
            <b class="tab"></b>     
            <b class="tab"></b>  
                 <h3>Your follow requests</h3>
                 <table className="table" style={{ width: "100%", marginTop: "3rem" }}>
                            <tbody>
                            {this.state.followers.map((follower) => (
                                    <tr id={follower.id} key={follower.id}>
                                        

                                        <td>
                                            
                                                <b>Username: </b> {follower.id}
                                            
                                
                                        </td>
                                        <td >
                                            <div style={{marginLeft:'55%'}}>
                                                <td>
                                                    <br></br>
                                                    <button  className="btn btn-outline-succes mt-1" onClick={() => this.handleAccept(follower.id)} type="button"><i className="icofont-subscribe mr-1"></i>Accept</button>
                                                </td>
                                                <td>
                                                    <br></br>
                                                    <button className="btn btn-outline-danger mt-1" onClick={() => this.handleDelete(follower.id)} type="button"><i className="icofont-subscribe mr-1"></i>Delete</button>
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

                       



    </div>

);
}
}

export default FollowRequest;
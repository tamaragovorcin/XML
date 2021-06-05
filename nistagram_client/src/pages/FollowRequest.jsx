import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import HeadingAlert from "../components/HeadingAlert";
import Axios from "axios";

import { BASE_URL_USER } from "../constants.js";
import { BASE_URL_USER_INTERACTION } from "../constants.js";
class FollowRequest extends Component {
	state = {
		followers: [],
    username : "",
	};

    handleFollows = () => {
    let highliht1 = { username: "follower1" };
		let highliht2 = { username: "follower111" };
		let highliht3 = { username: "follower11" };

		let list = [];
		list.push(highliht1)
		list.push(highliht2)
		list.push(highliht3)

		this.setState({ followers: list });
		
	}
  handleAccept  = () => {
    let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
		const followReguestDTO = { followingId: id, followerId: id};
		Axios.post(BASE_URL_USER_INTERACTION  + "/api/followApproved", followReguestDTO)
				.then((res) => {
					
						console.log(res.data)
						this.setState({ redirect: true });
					
				})
				.catch ((err) => {
			console.log(err);
		});
  
    
}
handleDelete  = () => {

}
    componentDidMount() {
		this.handleFollows()
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = {id: id}
        Axios.post(BASE_URL_USER_INTERACTION + "/api/allFollowRequest", dto)
			.then((res) => {

				console.log(res.data)
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
                                {this.state.followers.map((people) => (
                                    <tr id={people.id} key={people.id}>
                                        

                                        <td>
                                            <div>
                                                <br></br>
                                                <b>Username: </b> {people.username}
                                            </div>
                                
                                        </td>
                                        <td >
                                            <div style={{marginLeft:'55%'}}>
                                                <td>
                                                    <br></br>
                                                    <button style={{height:'30px'},{verticalAlign:'center'},{marginTop:'2%'}} className="btn btn-outline-succes mt-1" onClick={() => this.handleAccept(people.id)} type="button"><i className="icofont-subscribe mr-1"></i>Accept</button>
                                                </td>
                                                <td>
                                                    <br></br>
                                                    <button style={{height:'30px'},{verticalAlign:'center'},{marginTop:'2%'}} className="btn btn-outline-danger mt-1" onClick={() => this.handleDelete(people.id)} type="button"><i className="icofont-subscribe mr-1"></i>Delete</button>
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
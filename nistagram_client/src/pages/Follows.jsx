import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import HeadingAlert from "../components/HeadingAlert";

class Follows extends Component {
	state = {
		follower: "",
	};

	

	render() {
		return (
            <div class="container-fluid">
            <b class="tab"></b>     
            <b class="tab"></b>  
                 <h3>Your follow requests</h3>

                        <table class="table table-striped table-light">
                                <thead class="thead-dark">
                                  <tr>
                                    <th scope="col"></th>
                                    <th scope="col">New follower</th>
                                    <th scope="col">Accept</th>
                                    <th scope="col">Delete</th>
                                    
                                  </tr>
                                </thead>
                                <tbody>
                                <tr>
                                            <td></td>
                                            <td></td>
                                            <td><button class="btn btn-success btn-lg">Accept</button></td>
                                           <td><button class="btn btn-danger btn-lg">Delete</button></td>
                                        </tr>
                                 </tbody>
                        </table>



    </div>

);
}
}

export default Follows;
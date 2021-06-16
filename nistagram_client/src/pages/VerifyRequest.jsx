import React, { Component } from "react";
import Axios from "axios";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { BASE_URL_USER_INTERACTION, BASE_URL_USER } from "../constants.js";
import ModalDialog from "../components/ModalDialog";
import { BASE_URL } from "../constants.js";
class FollowRequest extends Component {
	state = {
        verificationRequests: [],
       

    };

    componentDidMount() {
        alert("ok")
        Axios.get(BASE_URL  + "/api/users/api/verificationRequest")
        .then((res) => {
            console.log(res.data)

        })
        .catch ((err) => {
    console.log(err);
});	}
   

 

	
  
   
 
    handleModalClose = () => {
		this.setState({ openModal: false });
	};
	render() {
		return (
            <React.Fragment>
				<TopBar />
				<Header />

                    <div className="container">
                        <div className="container" style={{ marginTop: "10rem", marginRight: "10rem" }}>
                            <h3>
                                Follow requests
                            </h3>
                            <table className="table" style={{ width: "100%" }}>
                                        <tbody>
                                        {this.state.verificationRequests.map((req) => (
                                                <tr id={req.Id} key={req.Id}>
                                                    
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Name: </b> {req.Name}
                                                        </div>
                                                    </td>
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Surname: </b> {req.Lastname}
                                                        </div>
                                                    </td>
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Category: </b> {req.Category}
                                                        </div>
                                                    </td>
                                                    <td >
                                                        <div style={{marginLeft:'55%'}}>
                                                            <td>
                                                                <button  className="btn btn-success mt-1" onClick={() => this.handleAccept(req.Id)} type="button"><i className="icofont-subscribe mr-1"></i>Approve</button>
                                                            </td>
                                                            <td>
                                                                <button className="btn btn-danger mt-1" onClick={() => this.handleDelete(req.Id)} type="button"><i className="icofont-subscribe mr-1"></i>Delete</button>
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
                        </div>
                        
                        
                        <div>
                            <ModalDialog
                            show={this.state.openModal}
                            onCloseModal={this.handleModalClose}
                            header="Successful"
                            text={this.state.textSuccessfulModal}
                            />
                        </div>
    </React.Fragment>
);
}
}

export default FollowRequest;
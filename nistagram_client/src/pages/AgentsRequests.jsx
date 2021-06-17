import React from "react";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import TopBar from "../components/TopBar";
import Header from "../components/Header";
import Axios from "axios";
import { BASE_URL } from "../constants.js";
import { Carousel } from 'react-responsive-carousel';
import ModalDialog from "../components/ModalDialog";

class AgentRequests extends React.Component {
    state = {
        posts : [],
        albums : [],
        openModal : false,

    }

    
    componentDidMount() {
        this.handleGetRequestAgents()
        
    }
    handleGetRequestAgents = () => {
        Axios.get(BASE_URL + "/api/users/agentRequests/")
            .then((res) => {
                this.setState({ posts: res.data });
            })
            .catch((err) => {
                console.log(err);
            });
        
            
    }
     handleAccept  = (followerId) => {
        const dto = {
           
            UserId : followerId
        }
        Axios.post(BASE_URL + "/api/users/agents/accept/",dto)
                .then((res) => {
                    this.setState({ openModal: true });
                    this.setState({ textSuccessfulModal: "You have successfully accepted agent request." });
        
                    this.handleGetRequestAgents()
                })
                .catch ((err) => {
            console.log(err);
            
        });
       
    }

    handleDelete  = (followerId) => {
        Axios.delete(BASE_URL + "/api/users/agent/"+followerId)
        .then((res) => {
            this.setState({ openModal: true });
			this.setState({ textSuccessfulModal: "You have successfully removed agent's request." });

           
        })
        .catch ((err) => {
             console.log(err);
    
        });
    }
    handleModalClose = () => {
		this.setState({ openModal: false });
	};

render(){
    return (
        <React.Fragment>
				<TopBar />
				<Header />
        
                <div className="container">
                        <div className="container" style={{ marginTop: "10rem", marginRight: "10rem" }}>
                            <h3>
                                Registration requests from agents
                            </h3>
                            <table className="table" style={{ width: "100%" }}>
                                        <tbody>
                                        {this.state.posts.map((follower) => (
                                                <tr id={follower.Id} key={follower.Id}>
                                                    
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Username: </b> {follower.ProfileInformation.Username}
                                                        </div>
                                                    </td>
                                                    <td >
                                                        <div style={{marginLeft:'55%'}}>
                                                            <td>
                                                                <button  className="btn btn-success mt-1" onClick={() => this.handleAccept(follower.Id)} type="button"><i className="icofont-subscribe mr-1"></i>Accept</button>
                                                            </td>
                                                            <td>
                                                                <button className="btn btn-danger mt-1" onClick={() => this.handleDelete(follower.Id)} type="button"><i className="icofont-subscribe mr-1"></i>Delete</button>
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
export default AgentRequests;
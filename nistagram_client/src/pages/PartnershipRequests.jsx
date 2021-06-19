import React, { Component } from "react";
import Axios from "axios";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import ModalDialog from "../components/ModalDialog";
import { BASE_URL } from "../constants.js";
class PartnershipRequests extends Component {
	state = {
        textSuccessfulModal : "",
        openModal : false,
        campaigns : [],
        multipleCampaigns : [],


    };

    componentDidMount() {
        this.handeleGetCampaigns()
        this.handeleGetMultipleCampaigns()

	}
    handeleGetCampaigns = () => {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);

		Axios.get(BASE_URL + "/api/campaign/partnershipRequests/"+id)
			.then((res) => {
				this.setState({ campaigns: res.data });
			})
			.catch((err) => {
				console.log(err);
			});	
	}
    handeleGetMultipleCampaigns = () => {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);

		Axios.get(BASE_URL + "/api/campaign/partnershipRequestsMultiple/"+id)
			.then((res) => {
				this.setState({ multipleCampaigns: res.data });
			})
			.catch((err) => {
				console.log(err);
			});	
	}
    handleAccept  = (postId) => {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
            const acceptedDTO = { CampaignId: postId, UserId: id};
            Axios.post(BASE_URL  + "/api/campaign/acceptPartnership", acceptedDTO)
                    .then((res) => {
                        this.setState({ openModal: true });
                        this.setState({ textSuccessfulModal: "You have successfully accepted partnership request." });			
                        this.handeleGetCampaigns()
                    })
                    .catch ((err) => {
                console.log(err);
            });
    }

    handleDelete = (postId) => {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = { CampaignId: postId, UserId: id};
        Axios.post(BASE_URL  + "/api/campaign/deletePartnership", dto)
                .then((res) => {
                    this.setState({ openModal: true });
                    this.setState({ textSuccessfulModal: "You have successfully deleted partnership request." });			
                    this.handeleGetCampaigns()
                })
                .catch ((err) => {
            console.log(err);
        });
    }
    handleAcceptMultiple  = (postId) => {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
            const acceptedDTO = { CampaignId: postId, UserId: id};
            Axios.post(BASE_URL  + "/api/campaign/acceptPartnershipMultiple", acceptedDTO)
                    .then((res) => {
                        this.setState({ openModal: true });
                        this.setState({ textSuccessfulModal: "You have successfully accepted partnership request." });			
                        this.handeleGetMultipleCampaigns()
                    })
                    .catch ((err) => {
                console.log(err);
            });
    }

    handleDeleteMultiple = (postId) => {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        const dto = { CampaignId: postId, UserId: id};
        Axios.post(BASE_URL  + "/api/campaign/deletePartnershipMultiple", dto)
                .then((res) => {
                    this.setState({ openModal: true });
                    this.setState({ textSuccessfulModal: "You have successfully deleted partnership request." });			
                    this.handeleGetMultipleCampaigns()
                })
                .catch ((err) => {
            console.log(err);
        });
    }
    
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
                               Partnership requests
                            </h3>
                            <table className="table" style={{ width: "100%" }}>
                                <tbody>
                                        {this.state.campaigns.map((post) => (
                                            
                                            <tr id={post.Id} key={post.Id}>
                                             <tr>
                                                <td colSpan="3">
                                                <label>Agent: &nbsp;</label>{post.AgentUsername}
                                                </td>
                                            </tr>
                                            <tr  style={{ width: "100%"}}>
                                                <td colSpan="3">
                                                {post.ContentType === "image/jpeg" ? (
                                                    <img
                                                    className="img-fluid"
                                                    src={"http://localhost:80/api/campaign/api/file/"+post.Id}
                                                    width="100%"
                                                    alt="description"
                                                /> ) : (
                                                <video width="100%"  controls autoPlay loop muted>
                                                <source src={"http://localhost:80/api/campaign/api/file/"+post.Id} type ="video/mp4"></source>
                                                </video>)}
                                                </td>
                                            </tr>
                                           
                                            <tr>
                                                <td colSpan="3">
                                                <label>Link to webasite/article: &nbsp;</label><a href={post.Link}>{post.Link}</a>
                                                </td>
                                            </tr>
                                            <tr>
                                                <td colSpan="3">
                                                <label>Description: &nbsp;</label>{post.Description}
                                                </td>
                                            </tr>
                                            <tr>
                                                <td colSpan="3">
                                                        <label>Date: &nbsp;</label>{post.Date}
                                                </td>
                                            </tr>
                                            <tr>
                                                <td colSpan="3">
                                                    <label>Time: &nbsp;</label>{post.Time}
                                                </td>
                                            </tr>
                                            <tr>
                                           
                                            
                                        </tr>
                                        <tr>
                                            <td>
                                            <button onClick={() =>  this.handleAccept(post.Id)} className="btn btn-primary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem",fontcolor:"white" }}><label >Accept</label></button>
                                            </td>
                                            <td>
                                            <button onClick={() =>  this.handleDelete(post.Id)} className="btn btn-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem", color:"white" }}><label > Remove</label></button>
                                            </td>
                                        </tr>
                                            <br/>
                                            <br/>
                                            <br/>
                                            </tr>
                                            
                                        ))}

                                        </tbody>
                                    </table>
                                     <table className="table" style={{ width: "100%" }}>
                                        <tbody>
                                                {this.state.multipleCampaigns.map((post) => (
                                                    
                                                    <tr id={post.Id} key={post.Id}>
                                                    <tr>
                                                        <td colSpan="3">
                                                        <label>Agent: &nbsp;</label>{post.AgentUsername}
                                                        </td>
                                                    </tr>
                                                    <tr  style={{ width: "100%"}}>
                                                        <td colSpan="3">
                                                        {post.ContentType === "image/jpeg" ? (
                                                            <img
                                                            className="img-fluid"
                                                            src={"http://localhost:80/api/campaign/api/file/"+post.Id}
                                                            width="100%"
                                                            alt="description"
                                                        /> ) : (
                                                        <video width="100%"  controls autoPlay loop muted>
                                                        <source src={"http://localhost:80/api/campaign/api/file/"+post.Id} type ="video/mp4"></source>
                                                        </video>)}
                                                        </td>
                                                    </tr>
                                                
                                                    <tr>
                                                        <td colSpan="3">
                                                        <label>Link to webasite/article: &nbsp;</label><a href={post.Link}>{post.Link}</a>
                                                        </td>
                                                    </tr>
                                                    <tr>
                                                        <td colSpan="3">
                                                        <label>Description: &nbsp;</label>{post.Description}
                                                        </td>
                                                    </tr>
                                                
                                                    <tr>
                                                  
                                                    <tr>
                                                        <td >
                                                        <label>Start date: &nbsp;</label>{post.StartTime}
                                                        </td>
                                                    </tr>
                                                    <tr>
                                                        <td >
                                                        <label>End date: &nbsp;</label>{post.EndTime}
                                                        </td>
                                                    </tr>
                                                    <tr>
                                                        <td >
                                                        <label>Desired Number: &nbsp;</label>{post.DesiredNumber}
                                                        </td>
                                                    </tr>
                                                </tr>
                                                <tr>
                                                    <td>
                                                    <button onClick={() =>  this.handleAcceptMultiple(post.Id)} className="btn btn-primary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem",fontcolor:"white" }}><label >Accept</label></button>
                                                    </td>
                                                    <td>
                                                    <button onClick={() =>  this.handleDeleteMultiple(post.Id)} className="btn btn-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem", color:"white" }}><label > Remove</label></button>
                                                    </td>
                                                </tr>
                                                    <br/>
                                                    <br/>
                                                    <br/>
                                                    </tr>
                                                    
                                                ))}

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

export default PartnershipRequests;
import React from "react";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import { BASE_URL } from "../constants.js";
import Axios from "axios";
import TopBar from "../components/TopBar";
import Header from "../components/Header";
import ModalDialog from "../components/ModalDialog";

class VerifyRequest extends React.Component {
    state = {
		verifications : [],
        openModal : false,

	}

    componentDidMount() {

		this.getVerificationRequests()

    }
    getVerificationRequests = ()=> {
        Axios.get(BASE_URL + "/api/users/verificationRequestAll")
                .then((res) => {
                   
                    this.setState({ verifications: res.data });
                })
                .catch ((err) => {
            console.log(err);
            
        });
    }
    handleAccept  = (userId, requestId) => {
        const dto = {
            RequestId : requestId,
            UserId : userId
        }
        Axios.post(BASE_URL + "/api/users/verificationRequest/accept/",dto)
                .then((res) => {
                    this.setState({ openModal: true });
                    this.setState({ textSuccessfulModal: "You have successfully accepted verification request." });
        
                    this.getVerificationRequests()
                })
                .catch ((err) => {
            console.log(err);
            
        });
    }

    handleDelete  = (id) => {
        Axios.delete(BASE_URL + "/api/users/verificationRequest/"+id)
        .then((res) => {
            this.setState({ openModal: true });
			this.setState({ textSuccessfulModal: "You have successfully removed verification request." });

            this.getVerificationRequests()
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

        <div className="container" style={{ marginTop: "10%" }}>
            <h5 className=" text-center mb-0 mt-2 text-uppercase">Verification requests</h5>

            
              <div className="d-flex align-items-top">
                <div className="container-fluid">
                  
                  <table className="table">
                  <tbody>
                        {this.state.verifications.map((request) => (
                            
                            <tr id={request.Id} key={request.Id}>
                            
                            <tr  style={{ width: "100%"}}>
                                <td colSpan="3">
                               
                                    <img
                                    className="img-fluid"
                                    src={`data:image/jpg;base64,${request.Media}`}
                                    width="100%"
                                    alt="description"
                                  /> 
                                
                                </td>
                            </tr>
                            <tr>
                                <td colSpan="3">
                                    {request.Name}
                                </td>
                            </tr>
                            <tr>
                                <td colSpan="3">
                                    {request.LastName}
                                </td>
                            </tr>
                            <tr>
                                <td colSpan="3">
                                    {request.Category}
                                </td>
                            </tr>
                            <tr>
                           
                            
                          </tr>
                          <tr  >
                            <td>
                              <button onClick={() =>  this.handleAccept(request.UserId,request.Id)} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" , marginLeft:"4rem"}}><label>Accept</label></button>
                            </td>
                            <td>
                            <button onClick={() =>  this.handleDelete(request.Id)} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label > Delete</label></button>
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
export default VerifyRequest;
import React, { Component } from "react";
import Axios from "axios";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { BASE_URL_USER_INTERACTION, BASE_URL_USER } from "../constants.js";
import ModalDialog from "../components/ModalDialog";
import { BASE_URL } from "../constants.js";
import getAuthHeader from "../GetHeader";
class Notifications extends Component {
	state = {
        postsNotifications : [],
        commentsNotifications : [],
        textSuccessfulModal : "",
        openModal : false,

    };

    componentDidMount() {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1);
        Axios.get(BASE_URL + "/api/users/api/getPostNotifications/"+ id, {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
                this.setState({ postsNotifications: res.data });
			})
			.catch((err) => {
				console.log(err)
			});	
            Axios.get(BASE_URL + "/api/users/api/getCommentNotification/"+ id, {  headers: { Authorization: getAuthHeader() } })
			.then((res) => {
                this.setState({ commentsNotifications: res.data });
			})
			.catch((err) => {
				console.log(err)
			});	
        
        
        
        
        }
   
        hasRole = (reqRole) => {
            let roles = JSON.parse(localStorage.getItem("keyRole"));
            if (roles === null) return false;
    
            if (reqRole === "*") return true;
    
            for (let role of roles) {
                if (role === reqRole) return true;
            }
            return false;
        };

   
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
                            Notifications
                        </h3>
                        <table className="table" style={{ width: "100%" }}>
                                    <tbody>
                                    {this.state.postsNotifications.map((notification) => (
                                            <tr id={notification.Id} key={notification.Id}>
                                                
                                                <td >
                                                    <div style={{ marginTop: "1rem"}}>
                                                    Your following <b> {notification.Username}</b> has added new {notification.Posted}.
                                                    </div>
                                                </td>
    
                                            </tr>

                                        ))}
                                      
                                    </tbody>
                                </table>
                                <table className="table" style={{ width: "100%" }}>
                                    <tbody>
                                    {this.state.commentsNotifications.map((notification) => (
                                            <tr id={notification.Id} key={notification.Id}>
                                                
                                                <td >
                                                    <div style={{ marginTop: "1rem"}}>
                                                    <b> {notification.Username} </b>commented your post: &nbsp;<b>{notification.Posted}</b>
                                                    </div>
                                                </td>
    
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

export default Notifications;
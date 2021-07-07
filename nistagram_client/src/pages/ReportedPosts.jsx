
import React from "react";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import TopBar from "../components/TopBar";
import Header from "../components/Header";
import Axios from "axios";
import { BASE_URL } from "../constants.js";
import { Carousel } from 'react-responsive-carousel';
import ModalDialog from "../components/ModalDialog";
import getAuthHeader from "../GetHeader";
class ReportedPosts extends React.Component {
    state = {
        posts : [],
        albums : [],
        openModal : false,

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
    
    componentDidMount() {
        this.handleGetReportedPosts()
        this.handleGetReporteddAlbums()
    }
    handleModalClose = () => {
		this.setState({ openModal: false });
	};
    handleGetReportedPosts = () => {
        Axios.get(BASE_URL + "/api/feedPosts/feed/reports", {  headers: { Authorization: getAuthHeader() } })
            .then((res) => {
                this.setState({ posts: res.data });
            })
            .catch((err) => {
                console.log(err);
            });
        
            
    }
    handleGetReporteddAlbums = () => {
        Axios.get(BASE_URL + "/api/feedPosts/albumFeed/reports", {  headers: { Authorization: getAuthHeader() } })
            .then((res) => {
                this.setState({ albums: res.data });
            })
            .catch((err) => {
                console.log(err);
            });
    }

    ignoreReport = (reportId) => {
       
        Axios.delete(BASE_URL + "/api/feedPosts/report/remove/"+reportId, {  headers: { Authorization: getAuthHeader() } })
        .then((res) => {
            this.setState({ openModal: true });
			this.setState({ textSuccessfulModal: "You have successfully removed this report" });
            this.handleGetReporteddAlbums()
            this.handleGetReportedPosts()
        })
        .catch((err) => {
            console.log(err);
        });
    }
    removeUser = (userId,username,reportId) => {
       
        Axios.delete(BASE_URL + "/api/users/remove/"+userId, {  headers: { Authorization: getAuthHeader() } })
        .then((res) => {
            this.setState({ openModal: true });
			this.setState({ textSuccessfulModal: "You have successfully removed user "+username });
            Axios.delete(BASE_URL + "/api/feedPosts/report/remove/"+reportId)
                .then((res) => {
                    
                    Axios.delete(BASE_URL + "/api/feedPosts/removeUserId/"+userId, {  headers: { Authorization: getAuthHeader() } })
                    .then((res) => {
                        Axios.delete(BASE_URL + "/api/storyPosts/removeUserId/"+userId, {  headers: { Authorization: getAuthHeader() } })
                            .then((res) => {
                                const userDTO = { Id: userId};

                                Axios.post(BASE_URL + "/api/userInteraction/removeUser",userDTO, {  headers: { Authorization: getAuthHeader() } })
                                    .then((res) => {
                                        this.handleGetReporteddAlbums()
                                        this.handleGetReportedPosts()
                                    })
                                    .catch((err) => {
                                        console.log(err);
                                    });
                              
                            })
                            .catch((err) => {
                                console.log(err);
                            });
                    })
                    .catch((err) => {
                        console.log(err);
                    });
                })
                .catch((err) => {
                    console.log(err);
                });
        })
        .catch((err) => {
            console.log(err);
        });
    }
    removePost = (postId,reportId) => {
      
        Axios.delete(BASE_URL + "/api/feedPosts/feed/remove/"+postId+"/"+reportId, {  headers: { Authorization: getAuthHeader() } })
        .then((res) => {
            this.setState({ openModal: true });
			this.setState({ textSuccessfulModal: "You have successfully removed this feed post" });
            this.handleGetReporteddAlbums()
            this.handleGetReportedPosts()
        })
        .catch((err) => {
            console.log(err);
        });
    }
    removeAlbum = (postId,reportId) => {
      
        Axios.delete(BASE_URL + "/api/feedPosts/albumFeed/remove/"+postId+"/"+reportId, {  headers: { Authorization: getAuthHeader() } })
        .then((res) => {
            this.setState({ openModal: true });
			this.setState({ textSuccessfulModal: "You have successfully removed this album" });
            this.handleGetReporteddAlbums()
            this.handleGetReportedPosts()
        })
        .catch((err) => {
            console.log(err);
        });
    }

render(){
    return (
        <React.Fragment>
				<TopBar />
				<Header />
         <div className="container" style={{ marginTop: "10%" }}>
					<h5 className=" text-center mb-0 mt-2 text-uppercase">Reported posts</h5>
              <div className="d-flex align-items-top">
                <div className="container-fluid">
                  
                  <table className="table">
                    <tbody>
                      {this.state.posts.map((post) => (
                        
                        <tr id={post.Id} key={post.Id}>
                        <tr>
                            <label style={{fontSize:"20px",fontWeight:"bold"}}>{post.Username}</label>
                        </tr>
                        <tr  style={{ width: "100%"}}>
                                <td colSpan="3">
                                {post.ContentType === "image/jpeg" ? (
                                  <img
                                  className="img-fluid"
                                  src={"http://localhost:80/api/feedPosts/api/feed/file/"+post.Id}
                                  width="100%"
                                  alt="description"
                                />
                                ) : (
                                  
                                  <video width="100%"  controls autoPlay loop muted><source src={"http://localhost:80/api/feedPosts/api/feed/file/"+post.Id} type ="video/mp4"></source></video>
                                  
                                )}
    
                                </td>
                              </tr>
                              <tr></tr>
                            <tr>
                            <td colSpan="3">
                                {post.Location}
                            </td>
                            </tr>
                            <tr>
                            <td colSpan="3">
                                {post.Description}
                            </td>
                            </tr>
                            <tr>
                            <td colSpan="3">
                                {post.Hashtags}
                            </td>
                            </tr>
                            <tr>
                                <td colSpan="3">
                                    {post.Tagged}
                                </td>
                                    
                            </tr>
                            <tr>
                                <td>
                                    <button onClick={() =>  this.ignoreReport(post.ReportId)} className="btn btn-outline-secondary btn-sm"><label >Ignore</label></button>
                                </td>

                                <td>
                                    <button onClick={() =>  this.removePost(post.Id, post.ReportId)} className="btn btn-outline-secondary btn-sm"><label >Remove post</label></button>
                                </td>

                                <td>
                                    <button onClick={() =>  this.removeUser(post.UserId,post.Username,post.ReportId)} className="btn btn-outline-secondary btn-sm"><label >Remove user</label></button>
                                </td>

                            </tr>
                        </tr>
                      ))}

                    </tbody>
                  </table>
                </div>
              </div>
          
            <div className="d-flex align-items-top">
                 <div className="container-fluid">
                
                    <table className="table">
                    <tbody>
                        {this.state.albums.map((post) => (
                        
                        <tr id={post.Id} key={post.Id}>
                            <tr>
                                <label style={{fontSize:"20px",fontWeight:"bold"}}>{post.Username}</label>
                            </tr>
                            <tr  style={{ width: "100%"}}>
                            <td colSpan="3">
                            <Carousel dynamicHeight={true}>
                                {post.Media.map(img => (<div>
                                    <img
                                    className="img-fluid"
                                    src={`data:image/jpg;base64,${img}`}
                                    width="100%"
                                    alt="description"
                                    />		
                                </div>))}
                                </Carousel>
                            </td>
                            </tr>
                            <tr>
                            <td colSpan="3">
                                {post.Location}
                            </td>
                            </tr>
                            <tr>
                            <td colSpan="3">
                                {post.Description}
                            </td>
                            </tr>
                            <tr>
                            <td colSpan="3">
                                {post.Hashtags}
                            </td>
                            </tr>
                            <tr>
                                <td colSpan="3">
                                        {post.Tagged}
                                </td>
                                    
                             </tr>
                            <tr>
                                <td>
                                    <button onClick={() =>  this.ignoreReport(post.ReportId)} className="btn btn-outline-secondary btn-sm"><label >Ignore</label></button>
                                 </td>
                                <td>
                                     <button onClick={() =>  this.removeAlbum(post.Id,post.ReportId)} className="btn btn-outline-secondary btn-sm"><label >Remove album</label></button>
                                 </td>
                                <td>
                                    <button onClick={() =>  this.removeUser(post.UserId,post.Username,post.ReportId)} className="btn btn-outline-secondary btn-sm"><label >Remove user</label></button>
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

            <ModalDialog
                    show={this.state.openModal}
					onCloseModal={this.handleModalClose}
					header="Successful"
					text={this.state.textSuccessfulModal}
                />
        </React.Fragment>

    );

	}
}
export default ReportedPosts;
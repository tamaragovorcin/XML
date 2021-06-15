
import React from "react";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import TopBar from "../components/TopBar";
import Header from "../components/Header";
import Axios from "axios";
import { BASE_URL } from "../constants.js";
import { Carousel } from 'react-responsive-carousel';
import ModalDialog from "../components/ModalDialog";

class ReportedPosts extends React.Component {
    state = {
        posts : [],
        albums : [],
        openModal : false,

    }

    
    componentDidMount() {
        this.handleGetReportedPosts()
        this.handleGetReporteddAlbums()
    }
    handleModalClose = () => {
		this.setState({ openModal: false });
	};
    handleGetReportedPosts = (id) => {
        Axios.get(BASE_URL + "/api/feedPosts/feed/reports")
            .then((res) => {
                this.setState({ posts: res.data });
            })
            .catch((err) => {
                console.log(err);
            });
        
            
    }
    handleGetReporteddAlbums = (id) => {
        Axios.get(BASE_URL + "/api/feedPosts/albumFeed/reports")
            .then((res) => {
                this.setState({ albums: res.data });
            })
            .catch((err) => {
                console.log(err);
            });
    }

    ignoreReport = (reportId) => {
       
        Axios.delete(BASE_URL + "/api/feedPosts/report/remove/"+reportId)
        .then((res) => {
            this.setState({ openModal: true });
			this.setState({ textSuccessfulModal: "You have successfully removed this report" });
        })
        .catch((err) => {
            console.log(err);
        });
    }
    removeUser = (userId,username) => {
       
        Axios.delete(BASE_URL + "/api/users/remove/"+userId)
        .then((res) => {
            this.setState({ openModal: true });
			this.setState({ textSuccessfulModal: "You have successfully removed user "+username });
        })
        .catch((err) => {
            console.log(err);
        });
    }
    removePost = (postId) => {
      
        Axios.delete(BASE_URL + "/api/feedPosts/feed/remove/"+postId)
        .then((res) => {
            this.setState({ openModal: true });
			this.setState({ textSuccessfulModal: "You have successfully removed this feed post" });
        })
        .catch((err) => {
            console.log(err);
        });
    }
    removeAlbum = (postId) => {
      
        Axios.delete(BASE_URL + "/api/feedPosts/albumFeed/remove/"+postId)
        .then((res) => {
            this.setState({ openModal: true });
			this.setState({ textSuccessfulModal: "You have successfully removed this album" });
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
                                <button onClick={() =>  this.ignoreReport(post.ReportId)} className="btn btn-outline-secondary btn-sm"><label >Ignore</label></button>
                            </tr>
                            <tr>
                                <button onClick={() =>  this.removePost(post.Id)} className="btn btn-outline-secondary btn-sm"><label >Remove post</label></button>
                            </tr>
                            <tr>
                                <button onClick={() =>  this.removeUser(post.UserId,post.Username)} className="btn btn-outline-secondary btn-sm"><label >Remove user</label></button>
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
                                <button onClick={() =>  this.ignoreReport(post.ReportId)} className="btn btn-outline-secondary btn-sm"><label >Ignore</label></button>
                            </tr>
                            <tr>
                                <button onClick={() =>  this.removeAlbum(post.Id)} className="btn btn-outline-secondary btn-sm"><label >Remove album</label></button>
                            </tr>
                            <tr>
                                <button onClick={() =>  this.removeUser(post.UserId,post.Username)} className="btn btn-outline-secondary btn-sm"><label >Remove user</label></button>
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
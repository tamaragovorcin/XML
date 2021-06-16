
import React from "react";
import {Tabs, Tab} from 'react-bootstrap';
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import TopBar from "../components/TopBar";
import Header from "../components/Header";
import Axios from "axios";
import { BASE_URL } from "../constants.js";
import { Carousel } from 'react-responsive-carousel';

class DislikedPosts extends React.Component {
    state = {
    
        dislikedPhotos : [],
        dislikedAlbums : [],
    }

    


componentDidMount() {

    let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

    this.handleGetDislikedPosts(id)
    this.handleGetDislikedAlbums(id)
}
handleGetDislikedPosts = (id) => {
    Axios.get(BASE_URL + "/api/feedPosts/feed/disliked/"+id)
        .then((res) => {
            this.setState({ dislikedPhotos: res.data });
        })
        .catch((err) => {
            console.log(err);
        });
    
        
}
handleGetDislikedAlbums = (id) => {
    Axios.get(BASE_URL + "/api/feedPosts/albumFeed/disliked/"+id)
        .then((res) => {
            this.setState({ dislikedAlbums: res.data });
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
					<h5 className=" text-center mb-0 mt-2 text-uppercase">Disliked posts</h5>

              <div className="d-flex align-items-top">
                <div className="container-fluid">
                  
                  <table className="table">
                    <tbody>
                      {this.state.dislikedPhotos.map((post) => (
                        
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
                        {this.state.dislikedAlbums.map((post) => (
                        
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
        </React.Fragment>

    );
 }
	
}
export default DislikedPosts;

import React from "react";
import {Tabs, Tab} from 'react-bootstrap';
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import TopBar from "../components/TopBar";
import Header from "../components/Header";
import Axios from "axios";
import { BASE_URL } from "../constants.js";
import { Carousel } from 'react-responsive-carousel';

class LikedAndDislikedPosts extends React.Component {
    state = {
        likedPhotos : [],
        dislikedPhotos : [],
        likedAlbums : [],
        dislikedAlbums : [],
    }

    constructor(props){
        super(props);
        this.state = {
            key: 1 | props.activeKey
        }
    this.handleSelect = this.handleSelect.bind(this);
}

handleSelect (key) {
    this.setState({key})
}
componentDidMount() {

    let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)

    this.handleGetLikedPosts(id)
    this.handleGetLikedAlbums(id)

    this.handleGetDislikedPosts(id)
    this.handleGetDislikedAlbums(id)

}
handleGetLikedPosts = (id) => {
    Axios.get(BASE_URL + "/api/feedPosts/feed/liked/"+id)
        .then((res) => {
            this.setState({ likedPhotos: res.data });
        })
        .catch((err) => {
            console.log(err);
        });
    
        
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
handleGetLikedAlbums = (id) => {
    Axios.get(BASE_URL + "/api/feedPosts/likedAlbums/"+id)
        .then((res) => {
            this.setState({ likedAlbums: res.data });
        })
        .catch((err) => {
            console.log(err);
        });
}
handleGetDislikedAlbums = (id) => {
    Axios.get(BASE_URL + "/api/feedPosts/dislikedAlbums/"+id)
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
         <Tabs
            activeKey={this.state.key}
            onSelect={this.handleSelect}
            id="controlled-tab-example"
            style={{ width: "100%" }}
            >
            <Tab eventKey={1} title="Liked posts">
              <div className="d-flex align-items-top">
                <div className="container-fluid">
                  
                  <table className="table">
                    <tbody>
                      {this.state.liked.map((post) => (
                        
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
            </Tab>
            <Tab eventKey={2} title="Disliked posts">
              <div className="d-flex align-items-top">
                <div className="container-fluid">
                  
                  <table className="table">
                    <tbody>
                      {this.state.disliked.map((post) => (
                        
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
            </Tab>
            <Tab eventKey={3} title="Liked albums">
                 <div className="d-flex align-items-top">
                 <div className="container-fluid">
                
                    <table className="table">
                    <tbody>
                        {this.state.likedAlbums.map((post) => (
                        
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
      </Tab>
      <Tab eventKey={4} title="Disliked albums">
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
          </Tab>
           
           
        </Tabs>
        </React.Fragment>

    );

	}
}
export default LikedAndDislikedPosts;
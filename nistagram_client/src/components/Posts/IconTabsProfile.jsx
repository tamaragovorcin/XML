
import React from "react";
import {Tabs, Tab} from 'react-bootstrap';
import SavedPosts from "./SavedPosts"
import { FiHeart } from "react-icons/fi";
import {FaHeartBroken,FaRegCommentDots} from "react-icons/fa"
import {BsBookmark} from "react-icons/bs"
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import { Carousel } from 'react-responsive-carousel';
import playerLogo from "../../static/coach.png";
import logo from "../../static/collection.png";

class IconTabsProfile extends React.Component {
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
render(){
    return (
         <Tabs
            activeKey={this.state.key}
            onSelect={this.handleSelect}
            id="controlled-tab-example"
            style={{ width: "100%" }}
            >
            <Tab eventKey={1} title="Posts">
              <div className="d-flex align-items-top">
                <div className="container-fluid">
                  
                  <table className="table">
                    <tbody>
                      {this.props.photos.map((post) => (
                        
                        <tr id={post.id} key={post.id}>
                          
                          <tr  style={{ width: "100%"}}>
                            <td colSpan="3">
                            <img
                              className="img-fluid"
                              src={`data:image/jpg;base64,${post.Media}`}
                              width="100%"
                              alt="description"
                            />
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
                          <tr  style={{ width: "100%" }}>
                           <td>
                              <button onClick={() =>  this.props.handleLike(post.Id)}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FiHeart/></button>
                            </td>
                            <td>
                              <button onClick={() =>  this.props.handleDislike(post.Id)}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaHeartBroken/></button>
                            </td>
                            <td>
                              <button onClick={() =>  this.props.handleWriteCommentModal(post.Id)}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaRegCommentDots/></button>
                            </td>
                            <td>
                                  <button onClick={() =>  this.props.handleOpenAddPostToCollectionModal(post.Id)} style={{ marginBottom: "1rem", height:"40px" }} className="btn btn-outline-secondary btn-sm"><label ><BsBookmark/></label></button>
                            </td>
                          </tr>
                          <tr  style={{ width: "100%" }}>
                            <td>
                              <button onClick={() =>  this.props.handleLikesModalOpen(post.Id)} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" , marginLeft:"4rem"}}><label>likes</label></button>
                            </td>
                            <td>
                            <button onClick={() =>  this.props.handleDislikesModalOpen(post.Id)} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label > dislikes</label></button>
                            </td>
                            <td>
                              <button onClick={() =>  this.props.handleCommentsModalOpen(post.Id)} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label >Comments</label></button>
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
            <Tab eventKey={2} title="Albums">
            <div className="d-flex align-items-top">
                <div className="container-fluid">
                  
                  <table className="table">
                    <tbody>
                      {this.props.albums.map((post) => (
                        
                        <tr id={post.id} key={post.id}>
                          
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
                          <tr  style={{ width: "100%" }}>
                                <td>
                                  <button onClick={() =>  this.props.handleLikeAlbum(post.Id)}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FiHeart/></button>
                                </td>
                                <td>
                                  <button onClick={() =>  this.props.handleDislikeAlbum(post.Id)}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaHeartBroken/></button>
                                </td>
                                <td>
                                  <button onClick={() =>  this.props.handleWriteCommentModalAlbum(post.Id)}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaRegCommentDots/></button>
                                </td>
                                <td>
                                      <button onClick={() =>  this.props.handleOpenAddPostToCollectionModal(post.Id)} style={{ marginBottom: "1rem", height:"40px" }} className="btn btn-outline-secondary btn-sm"><label ><BsBookmark/></label></button>
                                </td>
                              
                            </tr>
                            <tr  style={{ width: "100%" }}>
                                <td>
                                  <button onClick={() =>  this.props.handleLikesModalOpenAlbum(post.Id)} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" , marginLeft:"4rem"}}><label>likes</label></button>
                                </td>
                                <td>
                                <button onClick={() =>  this.props.handleDislikesModalOpenAlbum(post.Id)} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label > dislikes</label></button>
                                </td>
                                <td>
                                  <button onClick={() =>  this.props.handleCommentsModalOpenAlbum(post.Id)} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label >Comments</label></button>
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
            <Tab eventKey={3} title="All stories">
                <div className="d-flex align-items-top">
                    <div className="container-fluid">
                    
                    <table className="table">
                        <tbody>
                        {this.props.stories.map((post) => (
                            
                            <tr id={post.Id} key={post.Id}>
                            
                            <tr  style={{ width: "100%"}}>
                                <td colSpan="3">
                                <img
                                className="img-fluid"
                                src={`data:image/jpg;base64,${post.Media}`}
                             
                                width="100%"
                                alt="description"
                                />
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
                                    <button onClick={() =>  this.props.handleOpenAddStoryToHighlightModal(post.Id)} className="btn btn-outline-secondary btn-sm"><label >Add story to highlight</label></button>
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
            <Tab eventKey={4} title="Story Albums">
            <div className="d-flex align-items-top">
                <div className="container-fluid">
                  
                  <table className="table">
                    <tbody>
                      {this.props.storyAlbums.map((post) => (
                        
                        <tr id={post.id} key={post.id}>
                          
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
                                    <button onClick={() =>  this.props.handleOpenAddStoryAlbumToHighlightModal(post.Id)} className="btn btn-outline-secondary btn-sm"><label >Add story to highlight</label></button>
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
            <Tab eventKey={5} title="Highlights">
                 <button onClick={() =>  this.props.handleAddHighLightClick()} className="btn btn-outline-secondary btn-sm" style={{ marginTop: "3rem",marginLeft:"2rem",marginBottom: "2rem" }}><label >Add new highlight</label></button>
                 <div className="container">
                    <div className="container-fluid testimonial-group d-flex align-items-top">
                        <div className="container-fluid scrollable" style={{ marginRight: "10rem" , marginBottom:"5rem",marginTop:"5rem"}}>
                          <table className="table-responsive" style={{ width: "100%" }}>
                            <tbody>

                              <tr >
                                {this.props.highlights.map((high) => (
                                  <td id={high.Id} key={high.Id} style={{width:"60em", marginLeft:"10em"}}>
                                    <tr width="100em">
                                      <button onClick={() =>  this.props.seeStoriesInHighlight(high.Stories)} className="btn btn-outline-secondary btn-sm" style={{ marginTop: "3rem",marginLeft:"2rem",marginBottom: "3rem" }}>

                                        <img
                                          className="img-fluid"
                                          src={playerLogo}
                                          style ={{borderRadius:"50%",margin:"2%"}}
                                          width="60em"
                                          alt="description"
                                        />
                                        </button>
                                    </tr>
                                    <tr>
                                      <label style={{marginRight:"15px"}}>{high.Name}</label>
                                    </tr>
                                  </td>
                                  
                                ))}
                              </tr>


                            </tbody>
                          </table>
                        </div>
                  </div>
                </div>
                <div className="d-flex align-items-top" hidden={this.props.hiddenStoriesForHighlight}>
                    <div className="container-fluid">
                    
                    <table className="table">
                        <tbody>
                        {this.props.storiesForHightliht.map((post) => (
                            
                            <tr id={post.Id} key={post.Id}>
                            
                            <tr  style={{ width: "100%"}}>
                                <td colSpan="3">
                                <img
                                className="img-fluid"
                                src={`data:image/jpg;base64,${post.Media}`}
                                width="100%"
                                alt="description"
                                />
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
            <Tab eventKey={5} title="HighlightsAlbums">
                 <button onClick={() =>  this.props.handleAddHighLightClick()} className="btn btn-outline-secondary btn-sm" style={{ marginTop: "3rem",marginLeft:"2rem",marginBottom: "2rem" }}><label >Add new highlight</label></button>
                 <div className="container">
                    <div className="container-fluid testimonial-group d-flex align-items-top">
                        <div className="container-fluid scrollable" style={{ marginRight: "10rem" , marginBottom:"5rem",marginTop:"5rem"}}>
                          <table className="table-responsive" style={{ width: "100%" }}>
                          <tbody>
                               {this.props.highlightsAlbum.map((high) => (
                        
                        <tr id={high.id} key={high.id}>
                          
                          <tr  style={{ width: "100%"}}>
                            <td colSpan="3">
                             <Carousel dynamicHeight={true}>
							                	{highlightsAlbum.Media.map(img => (<div>
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
                                {high.Location}
                            </td>
                          </tr>
                          <tr>
                            <td colSpan="3">
                                {high.Description}
                            </td>
                          </tr>
                          <tr>
                            <td colSpan="3">
                                {high.Hashtags}
                            </td>
                          </tr>
                          <tr>
                            <td colSpan="3">
                                {high.Tagged}
                            </td>
                            
                          </tr>
                          <tr>
                                    <button onClick={() =>  this.props.handleOpenAddStoryAlbumToHighlightModal(high.Id)} className="btn btn-outline-secondary btn-sm"><label >Add story album to highlight</label></button>
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
                <div className="d-flex align-items-top" hidden={this.props.hiddenStoriesAlbumsForHighlight}>
                    <div className="container-fluid">
                    
                    <table className="table">
                        <tbody>
                        {this.props.storiesAlbumsForHightliht.map((high) => (
                            
                            <tr id={high.Id} key={high.Id}>
                            
                            <tr  style={{ width: "100%"}}>
                                <td colSpan="3">
                                <img
                                className="img-fluid"
                                src={`data:image/jpg;base64,${high.Media}`}
                                width="100%"
                                alt="description"
                                />
                                </td>
                            </tr>
                            <tr>
                                <td colSpan="3">
                                    {high.Location}
                                </td>
                            </tr>
                            <tr>
                                <td colSpan="3">
                                    {high.Description}
                                </td>
                            </tr>
                            <tr>
                                <td colSpan="3">
                                    {high.Hashtags}
                                </td>
                            </tr>
                            <tr>
                            <td colSpan="3">
                                {high.Tagged}
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
            <Tab eventKey={6} title="Saved posts">
            <button onClick={() =>  this.props.handleAddCollectionClick()} className="btn btn-outline-secondary btn-sm" style={{ marginTop: "3rem",marginLeft:"2rem",marginBottom: "2rem" }}><label >Add new collection</label></button>
                 <div className="container">
                    <div className="container-fluid testimonial-group d-flex align-items-top">
                        <div className="container-fluid scrollable" style={{ marginRight: "10rem" , marginBottom:"5rem",marginTop:"5rem"}}>
                          <table className="table-responsive" style={{ width: "100%" }}>
                            <tbody>

                              <tr >
                                {this.props.collections.map((collection) => (
                                  <td id={collection.Id} key={collection.Id} style={{width:"60em", marginLeft:"10em"}}>
                                    <tr width="100em">
                                      <button onClick={() =>  this.props.seePostsInCollection(collection.Posts)} className="btn btn-outline-secondary btn-sm" style={{ marginTop: "3rem",marginLeft:"2rem",marginBottom: "3rem" }}>

                                        <img
                                          className="img-fluid"
                                          src={logo}
                                          style ={{borderRadius:"50%",margin:"2%"}}
                                          width="60em"
                                          alt="description"
                                        />
                                        </button>
                                    </tr>
                                    <tr>
                                      <label style={{marginRight:"15px"}}>{collection.Name}</label>
                                    </tr>
                                  </td>
                                  
                                ))}
                              </tr>


                            </tbody>
                          </table>
                        </div>
                  </div>
                </div>
                <div className="d-flex align-items-top" hidden={this.props.hiddenPostsForCollection}>
                    <div className="container-fluid">
                    <table className="table">
                        <tbody>
                        {this.props.postsForCollection.map((post) => (
                            
                            <tr id={post.Id} key={post.Id}>
                            
                            <tr  style={{ width: "100%"}}>
                                <td colSpan="3">
                                <img
                                className="img-fluid"
                                src={`data:image/jpg;base64,${post.Media}`}
                                width="100%"
                                alt="description"
                                />
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
    );

	}
}
export default IconTabsProfile;
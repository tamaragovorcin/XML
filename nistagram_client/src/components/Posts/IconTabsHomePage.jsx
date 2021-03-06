
import React from "react";
import {Tabs, Tab} from 'react-bootstrap';
import {FaHeartBroken,FaRegCommentDots} from "react-icons/fa"
import {BsBookmark} from "react-icons/bs"
import { FiHeart, FiSend } from "react-icons/fi";
import {MdReportProblem} from "react-icons/md"
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import { Carousel } from 'react-responsive-carousel';

class IconTabsHomePage extends React.Component {
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
                        <td>
                                      <button onClick={() =>  this.props.handleOpenForwardModal(post.Id,"post")} style={{ marginBottom: "1rem", height:"40px" }} className="btn btn-outline-secondary btn-sm"><label ><FiSend/></label></button>
                                </td>
                        <td>
                              <button onClick={() =>  this.props.handleReportPost(post.Id,"post")} style={{ marginBottom: "1rem", height:"40px" }} className="btn btn-outline-secondary btn-sm"><label ><MdReportProblem/></label></button>
                        </td>
                    </tr>
                    <tr  style={{ width: "100%" }} hidden={!this.props.userIsLoggedIn}>
                        <td>
                          <button onClick={(e) =>  this.props.handleLikesModalOpen(e,post.Id)} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" , marginLeft:"4rem"}}><label>likes</label></button>
                        </td>
                        <td>
                         <button onClick={(e) =>  this.props.handleDislikesModalOpen(e,post.Id)} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label > dislikes</label></button>
                        </td>
                        <td>
                          <button onClick={(e) =>  this.props.handleCommentsModalOpen(e,post.Id)} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label >Comments</label></button>
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
                    <tr  style={{ width: "100%" }} hidden={!this.props.userIsLoggedIn}>
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
                            <button onClick={() =>  this.props.handleOpenAddAlbumToCollectionAlbumModal(post.Id)} style={{ marginBottom: "1rem", height:"40px" }} className="btn btn-outline-secondary btn-sm"><label ><BsBookmark/></label></button>
                        </td>
                        <td>
                                      <button onClick={() =>  this.props.handleOpenForwardModal(post.Id,"post")} style={{ marginBottom: "1rem", height:"40px" }} className="btn btn-outline-secondary btn-sm"><label ><FiSend/></label></button>
                                </td>
                        <td>
                              <button onClick={() =>  this.props.handleReportPost(post.Id, "album")} style={{ marginBottom: "1rem", height:"40px" }} className="btn btn-outline-secondary btn-sm"><label ><MdReportProblem/></label></button>
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
      <Tab eventKey={3} title="Campaigns" >
            <div className="d-flex align-items-top">
                <div className="container-fluid">
                  
                    <table className="table" style={{ width: "100%" }}>
                                <tbody>
                                        {this.props.oneTimeCampaigns.map((post) => (
                                            
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
                                                  <button class="astext" onClick={() =>  this.props.handleClickOnLink(post.Id,"oneTime",post.AgentId,post.Link)}>Link to webasite/article: {post.Link}</button>
                                                </td>
                                            </tr>
                                            <tr>
                                                <td colSpan="3">
                                                <label>Description: &nbsp;</label>{post.Description}
                                                </td>
                                            </tr>
                                            <tr  style={{ width: "100%" }}>
                                                <td>
                                                  <button onClick={() =>  this.props.handleLikeCampaign(post.Id,"oneTime")}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FiHeart/></button>
                                                </td>
                                                <td>
                                                  <button onClick={() =>  this.props.handleDislikeCampaign(post.Id, "oneTime")}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaHeartBroken/></button>
                                                </td>
                                                <td>
                                                  <button onClick={() =>  this.props.handleWriteCommentModalCampaign(post.Id,"oneTime")}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaRegCommentDots/></button>
                                                </td>
                                            </tr>
                                            <tr  style={{ width: "100%" }}>
                                                  <td>
                                                    <button onClick={() =>  this.props.handleLikesModalOpenCampaign(post.Id,"oneTime")} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" , marginLeft:"4rem"}}><label>likes</label></button>
                                                  </td>
                                                  <td>
                                                  <button onClick={() =>  this.props.handleDislikesModalOpenCampaign(post.Id,"oneTime")} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label > dislikes</label></button>
                                                  </td>
                                                  <td>
                                                    <button onClick={() =>  this.props.handleCommentsModalOpenCampaign(post.Id,"oneTime")} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label >Comments</label></button>
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
                                        {this.props.oneTimeCampaignsPromotion.map((post) => (
                                            
                                            <tr id={post.Id} key={post.Id}>
                                             <tr>
                                                <td colSpan="3">
                                                <label>Influencer: &nbsp;</label>{post.AgentUsername}
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
                                                   <button class="astext" onClick={() =>  this.props.handleClickOnLink(post.Id,"oneTime",post.AgentId,post.Link)}>Link to webasite/article: {post.Link}</button>
                                                </td>
                                            </tr>
                                            <tr>
                                                <td colSpan="3">
                                                <label>Description: &nbsp;</label>{post.Description}
                                                </td>
                                            </tr>
                                            <tr  style={{ width: "100%" }}>
                                                <td>
                                                  <button onClick={() =>  this.props.handleLikeCampaign(post.Id,"oneTime")}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FiHeart/></button>
                                                </td>
                                                <td>
                                                  <button onClick={() =>  this.props.handleDislikeCampaign(post.Id, "oneTime")}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaHeartBroken/></button>
                                                </td>
                                                <td>
                                                  <button onClick={() =>  this.props.handleWriteCommentModalCampaign(post.Id,"oneTime")}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaRegCommentDots/></button>
                                                </td>
                                            </tr>
                                            <tr  style={{ width: "100%" }}>
                                                  <td>
                                                    <button onClick={() =>  this.props.handleLikesModalOpenCampaign(post.Id,"oneTime")} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" , marginLeft:"4rem"}}><label>likes</label></button>
                                                  </td>
                                                  <td>
                                                  <button onClick={() =>  this.props.handleDislikesModalOpenCampaign(post.Id,"oneTime")} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label > dislikes</label></button>
                                                  </td>
                                                  <td>
                                                    <button onClick={() =>  this.props.handleCommentsModalOpenCampaign(post.Id,"oneTime")} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label >Comments</label></button>
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
                                                {this.props.multipleCampaigns.map((post) => (
                                                    
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
                                                           <button class="astext" onClick={() =>  this.props.handleClickOnLink(post.Id,"multiple",post.AgentId,post.Link)}>Link to webasite/article: {post.Link}</button>
                                                        </td>
                                                    </tr>
                                                    <tr>
                                                        <td colSpan="3">
                                                        <label>Description: &nbsp;</label>{post.Description}
                                                        </td>
                                                    </tr>
                                                    <tr  style={{ width: "100%" }}>
                                                        <td>
                                                          <button onClick={() =>  this.props.handleLikeCampaign(post.Id,"multiple")}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FiHeart/></button>
                                                        </td>
                                                        <td>
                                                          <button onClick={() =>  this.props.handleDislikeCampaign(post.Id, "multiple")}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaHeartBroken/></button>
                                                        </td>
                                                        <td>
                                                          <button onClick={() =>  this.props.handleWriteCommentModalCampaign(post.Id,"multiple")}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaRegCommentDots/></button>
                                                        </td>
                                                    </tr>
                                                    <tr  style={{ width: "100%" }}>
                                                          <td>
                                                            <button onClick={() =>  this.props.handleLikesModalOpenCampaign(post.Id,"multiple")} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" , marginLeft:"4rem"}}><label>likes</label></button>
                                                          </td>
                                                          <td>
                                                          <button onClick={() =>  this.props.handleDislikesModalOpenCampaign(post.Id,"multiple")} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label > dislikes</label></button>
                                                          </td>
                                                          <td>
                                                            <button onClick={() =>  this.props.handleCommentsModalOpenCampaign(post.Id,"multiple")} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label >Comments</label></button>
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
                                                {this.props.multipleCampaignsPromotion.map((post) => (
                                                    
                                                    <tr id={post.Id} key={post.Id}>
                                                    <tr>
                                                        <td colSpan="3">
                                                        <label>Influencer: &nbsp;</label>{post.AgentUsername}
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
                                                           <button class="astext" onClick={() =>  this.props.handleClickOnLink(post.Id,"multiple",post.AgentId,post.Link)}>Link to webasite/article: {post.Link}</button>

                                                        </td>
                                                    </tr>
                                                    <tr>
                                                        <td colSpan="3">
                                                        <label>Description: &nbsp;</label>{post.Description}
                                                        </td>
                                                    </tr>
                                                    <tr  style={{ width: "100%" }}>
                                                        <td>
                                                          <button onClick={() =>  this.props.handleLikeCampaign(post.Id,"multiple")}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FiHeart/></button>
                                                        </td>
                                                        <td>
                                                          <button onClick={() =>  this.props.handleDislikeCampaign(post.Id, "multiple")}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaHeartBroken/></button>
                                                        </td>
                                                        <td>
                                                          <button onClick={() =>  this.props.handleWriteCommentModalCampaign(post.Id,"multiple")}  className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem", height:"40px",marginLeft:"6rem" }}><FaRegCommentDots/></button>
                                                        </td>
                                                    </tr>
                                                    <tr  style={{ width: "100%" }}>
                                                          <td>
                                                            <button onClick={() =>  this.props.handleLikesModalOpenCampaign(post.Id,"multiple")} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem" , marginLeft:"4rem"}}><label>likes</label></button>
                                                          </td>
                                                          <td>
                                                          <button onClick={() =>  this.props.handleDislikesModalOpenCampaign(post.Id,"multiple")} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label > dislikes</label></button>
                                                          </td>
                                                          <td>
                                                            <button onClick={() =>  this.props.handleCommentsModalOpenCampaign(post.Id,"multiple")} className="btn btn-outline-secondary btn-sm" style={{ marginBottom: "1rem",marginLeft:"4rem" }}><label >Comments</label></button>
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
export default IconTabsHomePage;
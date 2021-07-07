
import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import playerLogo from "../static/me.jpg";
import profileImage from "../static/profileImage.jpg"
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import Axios from "axios";
import { BASE_URL } from "../constants.js";
import ImageUploader from 'react-images-upload';
import Select from 'react-select';
import { Button, Modal } from "react-bootstrap";
import { FiSend } from 'react-icons/fi';
import DisposableImageModal from "../components/DisposableImageModal";
import getAuthHeader from "../GetHeader";



class Chat extends React.Component {
	constructor(props) {
		super(props);


	}
	handleSubmit = (event) => {
        event.preventDefault();
    };
    hasRole = (reqRole) => {
		
		let roles = JSON.parse(localStorage.getItem("keyRole"));
		if (roles === null) return false;

		if (reqRole === "*") return true;

	
		if (roles.trim() === reqRole.trim()) 
		{
			return true;
		}
		return false;
	};
	state = {
        search: "",
		users: [],
		options: [],
		optionDTO: { value: "", label: "" },
		userId: "",
		message : "",
		messages : [],
		username : "",
		sentMessages : [],
		receivedMessages : [],
		chat : [],
		user1 : "",
		user2 : "",
		time : "",
		albumImages : [],
		selectedFile : "",
		showDisposable : false,
		disposableForOpenMedia : "",
		disposableForOpenId : "",
		following : false,
		deletedChat : false

	}
	
	handleOpenDisposable = (media,id)=>{
		this.setState({ disposableForOpenMedia: media });
		this.setState({ showDisposable: true });
		Axios.get(BASE_URL + "/api/messages/api/openDisposable/"+id,  {  headers: { Authorization: getAuthHeader() } })
				.then((res) => {
					this.handleGetChat();
					console.log(res)

				})
				.catch((err) => {
	
					console.log(err)
				});
  setTimeout(() => this.handleDisposableClose(), 5000)


	}
	handleDisposableClose = ()=>{
		this.setState({ showDisposable: false });
	}
	handleMessageChange = (event) => {
		this.setState({ message: event.target.value });
	};
	handleDeleteChat =()=>{
		let sender = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		var sentence = window.location.toString()
		var s = []
		let help = []
		s = sentence.split("/");
		const dto = { follower: sender, following : s[6]};

		Axios.get(BASE_URL + "/api/messages/api/deleteChat/"+sender+"/"+s[6],  {  headers: { Authorization: getAuthHeader() } })
		.then((res) => {
			this.isChatDeleted()
			this.handleGetChat()

	})
		.catch((err) => {
			console.log(err);
		});	
	}
	handleSendDisposable = (event)=>{
		var sentence = window.location.toString()

		var s = []
		s = sentence.split("/");
		console.log(window.location.toString())
		let sender = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		if (this.state.selectedFile == ""){
		const dto = {
			Sender: sender,
			Receiver: s[6],
			Text : this.state.message,
		    };

			Axios.post(BASE_URL + "/api/messages/api/send",dto,  {  headers: { Authorization: getAuthHeader() } })
				.then((res) => {

					this.handleGetChat();
					this.setState({ message: "" });

				})
				.catch((err) => {
	
					console.log(err)
				});

			}else{
				const formData = new FormData();

					formData.append("file", this.state.selectedFile);
					formData.append("test", "StringValueTest");
					const options = {
						method: "POST",
						body: formData

					};
				fetch(BASE_URL + "/api/messages/api/send/disposableImage/"+sender+"/"+s[6], options,  {  headers: { Authorization: getAuthHeader() } });
				this.setState({ selectedFile: "" });
				this.handleGetChat();


			}
				
	}
	handleSendMessage = (event)=>{
		var sentence = window.location.toString()

		var s = []
		s = sentence.split("/");
		console.log(window.location.toString())
		let sender = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		if (this.state.selectedFile == ""){
		const dto = {
			Sender: sender,
			Receiver: s[6],
			Text : this.state.message,
		    };

			Axios.post(BASE_URL + "/api/messages/api/send",dto)
				.then((res) => {

					this.handleGetChat();
					this.setState({ message: "" });

				})
				.catch((err) => {
	
					console.log(err)
				});

			}
				

	}
	testVideo(pic,messageId) {
		alert(messageId)
		
	}
	onChangeHandler = (event) => {
		this.setState({
            selectedFile: event.target.files[0],
            loaded: 0,
        });
		
    };
	amIFollowingThisUser() {
		let sender = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		var sentence = window.location.toString()
		var s = []
		let help = []
		s = sentence.split("/");
		const dto = { follower: sender, following : s[6]};

		Axios.post(BASE_URL + "/api/userInteraction/api/checkInteraction",dto,  {  headers: { Authorization: getAuthHeader() } })
		.then((res) => {
			this.setState({ following:res.data });

	})
		.catch((err) => {
			console.log(err);
		});	
	}
	isChatDeleted() {
		let sender = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		var sentence = window.location.toString()
		var s = []
		let help = []
		s = sentence.split("/");
		const dto = { follower: sender, following : s[6]};

		Axios.get(BASE_URL + "/api/messages/api/isChatDeleted/"+sender+"/"+s[6],  {  headers: { Authorization: getAuthHeader() } })
		.then((res) => {
			this.setState({ deletedChat:res.data.Deleted });
			this.setState({ deletedForUser:res.data.ForUser });
		


	})
		.catch((err) => {
			console.log(err);
		});	
	}
	handleGetChat(){
		let sender = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		var sentence = window.location.toString()
		var s = []
		let help = []
		s = sentence.split("/");
		console.log(window.location.toString())

		this.setState({ user2: s[6] });

		Axios.get(BASE_URL + "/api/messages/api/getMessages/"+sender+"/"+s[6],  {  headers: { Authorization: getAuthHeader() } })
		.then((res) => {
			res.data.forEach((mess) => {
				let time = mess.DateTime.split("T")[1].slice(0,5)
				this.handleGetAlbumImages(mess.AlbumPost)
				let feedUser = ""
				if( mess.FeedPost != "000000000000000000000000"){
				
				Axios.get(BASE_URL + "/api/feedPosts/api/feed/username/"+mess.FeedPost)
				.then((res2) => {
					feedUser = res2.data
					Axios.get(BASE_URL + "/api/users/api/user/username/"+mess.Sender)
				.then((res1) => {
					let optionDTO = { id: mess.Id, username: res1.data , text: mess.Text, time : time, feedPost : mess.FeedPost,feedUser : feedUser, storyPost : mess.StoryPost, disposableImage : mess.DisposableImage, albumPost: mess.AlbumPost, media : mess.DisposableImage.split("/")[1]}
					help.push(optionDTO)
				this.setState({ messages: help });
				})
				.catch((err) => {
					console.log(err);
				});	
					
				})
				.catch((err) => {
					console.log(err);
				});	
			}else if( mess.AlbumPost != "000000000000000000000000"){
				
				Axios.get(BASE_URL + "/api/feedPosts/api/album/username/"+mess.AlbumPost,  {  headers: { Authorization: getAuthHeader() } })
				.then((res2) => {
					feedUser = res2.data
					Axios.get(BASE_URL + "/api/users/api/user/username/"+mess.Sender,  {  headers: { Authorization: getAuthHeader() } })
				.then((res1) => {
					let optionDTO = { id: mess.Id, username: res1.data , text: mess.Text, time : time, feedPost : mess.FeedPost,feedUser : feedUser, storyPost : mess.StoryPost, disposableImage : mess.DisposableImage, albumPost: mess.AlbumPost, media : mess.DisposableImage.split("/")[1]}
					help.push(optionDTO)
				this.setState({ messages: help });
				})
				.catch((err) => {
					console.log(err);
				});	
					
				})
				.catch((err) => {
					console.log(err);
				});	
			}
			else if( mess.StoryPost != "000000000000000000000000"){
				Axios.get(BASE_URL + "/api/storyPosts/api/story/username/"+mess.StoryPost,  {  headers: { Authorization: getAuthHeader() } })
				.then((res2) => {
					feedUser = res2.data
					Axios.get(BASE_URL + "/api/users/api/user/username/"+mess.Sender,  {  headers: { Authorization: getAuthHeader() } })
				.then((res1) => {
					let optionDTO = { id: mess.Id, username: res1.data , text: mess.Text, time : time, feedPost : mess.FeedPost,feedUser : feedUser, storyPost : mess.StoryPost, disposableImage : mess.DisposableImage, albumPost: mess.AlbumPost, media : mess.DisposableImage.split("/")[1]}
					help.push(optionDTO)
				this.setState({ messages: help });
				})
				.catch((err) => {
					console.log(err);
				});	
					
				})
				.catch((err) => {
					console.log(err);
				});	
			}
			else if(mess.FeedPost === "000000000000000000000000" && mess.AlbumPost == "000000000000000000000000" && mess.StoryPost == "000000000000000000000000"){
				Axios.get(BASE_URL + "/api/users/api/user/username/"+mess.Sender,  {  headers: { Authorization: getAuthHeader() } })
				.then((res1) => {
					let optionDTO = { id: mess.Id,senderId:mess.Sender, username: res1.data , text: mess.Text, time : time, feedPost : mess.FeedPost, storyPost : mess.StoryPost, disposableImage : mess.DisposableImageId,openedDisposable : mess.OpenedDisposable, albumPost: mess.AlbumPost, media : mess.DisposableImage.split("/")[1]}
					help.push(optionDTO)
				this.setState({ messages: help });
				})
				.catch((err) => {
					console.log(err);
				});	


			}
				
				

			});


		})
		.catch((err) => {
			console.log(err);
		});	

		

	}
	handleGetAlbumImages=(id)=>{
		Axios.get(BASE_URL + "/api/feedPosts/api/feedAlbum/images/"+id,  {  headers: { Authorization: getAuthHeader() } })
		.then((res) => {

		this.setState({ albumImages: res.data });
		})
		.catch((err) => {
			console.log(err);
		});	
	}


	
    componentDidMount() {
		let sender = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length-1)
		this.setState({ user1: sender });
		this.amIFollowingThisUser();
		this.isChatDeleted();
		this.handleGetChat();

			};

	render() {
		return (
			<React.Fragment>
    		<TopBar />
			<Header />
			
                               
					 <table className="table" style={{ width: "100%",  maxHeight: "200px",
    overflowY: "scroll", marginTop: "10rem" , marginLeft:"15%", marginRight:"15%"}}>
                            <tbody hidden={this.state.deletedChat && this.state.deletedForUser == this.state.user1}>
                                {this.state.messages.map((message) => (
							<tr id={message.id} key={message.id}  style={{ width: "100%",marginTop:"2rem" }}>

<tr style={{ width: "100%" }}>
							<td>{message.time}</td>
							<td><label>   &nbsp; {message.username}:  &nbsp; </label>  <label><b>{message.text}</b></label></td>
							<div hidden ={message.feedPost=="000000000000000000000000"} style={{margin:"5%" }}>
							<label style={{color:"blueviolet" }}>{message.feedUser}</label><br/>
							<img
                                className="img-fluid"
                                src={"http://localhost:80/api/feedPosts/api/feed/fileMessage/"+message.feedPost+"/"+this.state.user1}
                                width="500px"
								height="500px"
                                alt="You can not see this picture. This account is private."
                              />
							</div>
							<div hidden ={message.storyPost=="000000000000000000000000"} style={{margin:"5%" }}>
							<label style={{color:"blueviolet" }}>{message.feedUser}</label><br/>
							<img
                                className="img-fluid"
                                src={"http://localhost:80/api/storyPosts/api/story/fileMessage/"+message.storyPost+"/"+this.state.user1}
                                width="500px"
								height="500px"
                                alt="You can not see this picture. This account is private or this story is expired."
                              />
							</div>


							<div hidden ={message.media==="000000000000000000000000" || message.media ===undefined}>
							<Button
							hidden ={message.openedDisposable || message.senderId == this.state.user1}
							style={{marginTop:"2%" }}
									 onClick={() =>  this.handleOpenDisposable(message.media, message.disposableImage)}
									className="btn btn-primary btn-md"
									id="sendMessageButton"
									type="button"
								>
									Open
									</Button>
									<Button
									style={{marginTop:"2%" }}

							hidden ={!message.openedDisposable || message.senderId == this.state.user1}
									className="btn btn-secondary btn-md"
									id="sendMessageButton"
									type="button"
								>
									Opened
									</Button>


									
									<Button
									style={{marginTop:"2%" }}

							hidden ={!message.senderId == this.state.user1}
									className="btn btn-secondary btn-md"
									id="sendMessageButton"
									type="button"
								>
									Sent
									</Button>
									
							</div>



							  <div hidden ={message.albumPost ==="000000000000000000000000"} style={{margin:"5%" }}>
							  <label style={{color:"blueviolet" }}>{message.feedUser}</label><br/>
							  {this.state.albumImages.map((message) => (
							<tr id={message.id} key={message.id}  style={{ width: "100%" }}>

							<tr style={{ width: "100%" }}>
							
							  <img
                                className="img-fluid"
                                src={"http://localhost:80/api/feedPosts/api/feedAlbum/files/"+message.Id}
                                width="500px"
								height="500px"
                                alt="description"
                              />
							
							
							</tr>
</tr>




                                ))}
							
							</div>
							</tr>
</tr>




                                ))}
								
                              
                            </tbody>
                        </table>
                  


			<table className="table" style={{ width: "80%", marginTop: "2rem" , marginRight:"15%", marginLeft:"15%"}}>
                            <tbody>
								<tr style={{ width: "100%"}} >
								<td colSpan="3">
								<div className="form-group controls mb-0 pb-2" style={{ color: "#6c757d", opacity: 1 }}>
									<input
										placeholder="New message"
										className="form-control"
										id="comment"
										type="text"
										onChange={this.handleMessageChange}
										value={this.state.message}
									/>
								</div>
								</td>
								<td colSpan="2">	
									<Button
									style={{ background: "#1977cc"}}
									onClick={() => this.handleSendMessage(this.state.message)}
									className="btn btn-primary btn-md"
									id="sendMessageButton"
									type="button"
								>
									Send
									</Button>
									</td>
								</tr>
								<tr style={{ width: "100%"}}>
									<td colSpan="3">
                            <label>
                                Upload a file:  &nbsp;      </label>

                                <input type="file" className="btn btn-outline-secondary btn-sm" name="file" onChange={this.onChangeHandler} />
								</td>
								<td colspan="2">
								<Button
									style={{ background: "#1977cc", width: "50%" }}
									onClick={() => this.handleSendDisposable(this.state.message)}
									className="btn btn-primary btn-md"
									id="sendMessageButton"
									type="button"
								>
									<FiSend/>
									</Button>
							</td>
							</tr>
							<tr>
							<Button
									hidden ={this.state.following || this.state.deletedChat}
									style={{marginTop:"2%", marginLeft :"40%" }}
									onClick={() => this.handleDeleteChat()}
									className="btn btn-danger btn-md"
									id="sendMessageButton"
									type="button"
								>
									Delete chat
									</Button>
							</tr>
                            </tbody>
                        </table>
              
                        <DisposableImageModal
					  show={this.state.showDisposable}
					  onCloseModal={this.handleDisposableClose}
					  header=""
					  media = {this.state.disposableForOpenMedia}
					

					  
					  />
			</React.Fragment>
		);
	}
}

export default Chat;

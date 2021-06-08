import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import Stories from 'react-insta-stories';
import { BiRightArrow } from "react-icons/bi";
import { BiLeftArrow } from "react-icons/bi";

import ConvertImage from "react-convert-image";
class StoriesModal extends Component {


	state = {
		count : 0,	ss: [],
		photos: [],
		peopleLikes: [],
		peopleDislikes: [],
		peopleComments: [],
		albums: [],
		showLikesModal: false,
		showDislikesModal: false,
		showCommentsModal: false,

		showStories: false,
		showWriteCommentModal: false,
		showAddPostToCollection: false,
		selectedPostId: -1,
		collections: [],
		showWriteCommentModalAlbum: false,
		users: [],
		pics: [],
		image: [],
		converted: undefined,
		help: [],
		ubiucse: "",
		pictures: [],
		bla: [1, 2],
		imageUrl: "",
		helpImage: "",
		hid: true,
		ready: false,
		stories: [],
		convertedImage: "",
		count: 0,
		userIsLogged: false,
		ssAlbums: [],
		usern: "",
		brojac : 0,
		br: 0,
		myCollectionAlbums : [],
		showAddAlbumToCollectionAlbum : false,
		userIsLoggedIn : true,
		stoori: [],
	}

	close = () => {

		this.setState({ count: 0 });



	}
   onClick = () =>{
	
	   let a = this.state.count;
	   if(a === this.props.brojac-1){
	   }else{
		a = a+1;
		this.setState({ count: a });
	   }
   }

   componentDidMount(){
	this.setState({ brojac :0 });
	
	
	   console.log(this.props.stories)
   }
   onClickBack = () =>{
	let a = this.state.count;
	if(a-1 === -1){
	}else{
	 a = a-1;
	 this.setState({ count: a });
	}
   
}
handleConvertedImage = (converted, username) => {
		
	/*var hh = this.state.stories;
	this.setState({ br: this.state.br +1 });
	if (this.state.usern === "") {

		this.setState({
			usern: username.username,
		});

		let st = { id: this.state.br, stories: [] }
		let storiji = {
			url: converted, header: {
				heading: username.username,
				subheading: 'CLOSE FRIENDS',

			},
		}
		st.stories.push(storiji)
		hh.push(st)
		this.setState({
			stories: hh,
		});




		if (this.state.brojac === hh.length) {
			this.setState({
				ready: true,
			});
		}
	}

	else if (this.state.usern === username.username) {

		this.state.stories.forEach(l => {
			l.stories.forEach(ll => {
				console.log(ll)
				if (ll.header.heading === username.username) {
					
					console.log(ll)
					let storiji = {
						url: converted, header: {
							heading: username.username,
							subheading: 'CLOSE FRIENDS',

						},
					}
					
					l.stories.push(storiji)
					var pom =l
					hh.pop(l)
					hh.push(pom)

				}

			
				this.setState({
					stories: hh,
				});

			})
		})





		if (this.state.brojac === hh.length) {
			this.setState({
				ready: true,
			});
		}
	}
	else{
		this.setState({
			usern: username.username,
		});

		let st = { id:  this.state.br, stories: [] }
		let storiji = {
			url: converted, header: {
				heading: username.username,
				subheading: 'CLOSE FRIENDS',

			},
		}
		st.stories.push(storiji)
		hh.push(st)
		this.setState({
			stories: hh,
		});




		if (this.state.brojac === hh.length) {
			this.setState({
				ready: true,
			});
		}
		console.log(hh)
	}*/
	this.setState({ stories :[] });
	var bro = this.state.brojac;

	var hh = [];
	let st = [] 
		let storiji = {
			url: converted, header: {
				heading: username,
				subheading: 'CLOSE FRIENDS',

			},
		}
	st.push(storiji)
	hh.push(storiji)
	console.log(hh)
	this.setState({stories: hh})


	/*if(bro === this.props.stories.length-1){
		alert("EVO ME")
		this.setState({ready : true})
	}*/
	this.setState({ready : true})
	bro = bro +1 
	this.setState({brojac : bro})


}
	render() {
		return (
			
			<Modal
			
				opacity= {0.5}
				show={this.props.show}
				
				transparent={true}
				aria-labelledby="contained-modal-title-vcenter"
				centered
				onHide={this.props.onCloseModal}
			>
				<Modal.Header closeButton onClick = {this.close}  style={{ alignItems: 'center', justifyContent: 'center', backgroundColor: 'rgba(0,0,0,0.5)'}}>
					
				</Modal.Header>
				<Modal.Body   style={{flex: 1, alignItems: 'center', justifyContent: 'center', backgroundColor: 'rgba(0,0,0,0.5)'}}>
			
					<div hidden={this.state.hid}>
						
								<ConvertImage
									image={this.props.stt.s}
									onConversion={e => this.handleConvertedImage(e, this.props.stt.username)}

								/>
							</div>

			



				<table style={{ width: "100%" }}>
				<tbody>
				
						{this.state.ready && <div>
							
								<td>
							<Stories
							stories={this.state.stories}
							defaultInterval={1500}
							width={432}
							height={768}
						/> 
						</td>


					
						
						</div>
						}
						
					
					
						
						</tbody>
				</table>
	
	
                              
                           
						
				</Modal.Body>
				
			</Modal>
		);
	}
}

export default StoriesModal;

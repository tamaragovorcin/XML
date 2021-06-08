import React, { Component } from "react";
import { Button, Modal } from "react-bootstrap";
import Stories from 'react-insta-stories';
import { Carousel } from 'react-responsive-carousel';
import { BiRightArrow } from "react-icons/bi";
import { BiLeftArrow } from "react-icons/bi";

class StoriesModal extends Component {


	state = {
		count : 0,
	}

	close = () => {

		this.setState({ count: 0 });



	}
   onClick = () =>{
	
	   let a = this.state.count;
	   if(a === this.props.brojac-1){
		alert("Preko je")
	   }else{
		alert("Povecavam")
		a = a+1;
		this.setState({ count: a });
	   }
	   alert(a)
   }

   componentDidMount(){
	alert(this.state.count)
   }
   onClickBack = () =>{
	alert(this.state.count)
	let a = this.state.count;
	if(a-1 === -1){
		alert("Minus jedan je")
	}else{
		alert("Smanjujem" + a)
	 a = a-1;
	 this.setState({ count: a });
	}
   
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
				
				<table style={{ width: "100%" }}>
				<tbody>
				<td><BiLeftArrow onClick = {this.onClickBack} /></td>
						{this.props.ready && <div>
							
								<td>
							<Stories
							stories={this.props.stories[this.state.count].stories}
							defaultInterval={1500}
							width={432}
							height={768}
						/> 
						</td>
						</div>
						}
							<td><BiRightArrow onClick = {this.onClick} /></td>
					
					
						
						</tbody>
				</table>
	
	
                              
                           
						
				</Modal.Body>
				
			</Modal>
		);
	}
}

export default StoriesModal;

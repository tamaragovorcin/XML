
import React from "react";
import { Tabs, Tab } from 'react-bootstrap';
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import TopBar from "../components/TopBar";
import Header from "../components/Header";
import Axios from "axios";
import { BASE_URL } from "../constants.js";
import { Carousel } from 'react-responsive-carousel';
import { Button } from "react-bootstrap";
class Tokens extends React.Component {
    state = {
        token: "",
    }




    componentDidMount() {

        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)
        Axios.get(BASE_URL + "/api/users/api/token/" + id)
				.then((res) => {
					if (res.status === 401) {
						this.setState({ errorHeader: "Bad credentials!", errorMessage: "Wrong username or password.", hiddenErrorAlert: false });
					} else if (res.status === 500) {
						this.setState({ errorHeader: "Internal server error!", errorMessage: "Server error.", hiddenErrorAlert: false });
					} else {
						this.setState({
							
							token : res.data, 
						});
					}
				})
				.catch ((err) => {
			console.log(err);
		});




    }

    handleNew = () => {
        let id = localStorage.getItem("userId").substring(1, localStorage.getItem('userId').length - 1)
        Axios.get(BASE_URL + "/api/users/api/generateToken/" + id)
				.then((res) => {
					if (res.status === 401) {
						this.setState({ errorHeader: "Bad credentials!", errorMessage: "Wrong username or password.", hiddenErrorAlert: false });
					} else if (res.status === 500) {
						this.setState({ errorHeader: "Internal server error!", errorMessage: "Server error.", hiddenErrorAlert: false });
					} else {
						this.setState({
							
							token : res.data, 
						});
					}
				})
				.catch ((err) => {
			console.log(err);
		});



    }


    render() {
        return (
            <React.Fragment>
                <TopBar />
                <Header />
                <div className="container" style={{ marginTop: "10%" }}>
                    <h5 className=" text-center mb-0 mt-2 text-uppercase">My current token</h5>

                </div>
                <div className="container" style={{ marginLeft: "5rem", marginRight: "5rem"}}  ><h5  className=" text-center mb-0 mt-2 text-uppercase">{this.state.token}</h5></div>
                <div style={{ marginTop: "10%" }}><Button
                    style={{ background: "#1977cc", marginTop: "15px", marginLeft: "40%", width: "20%" }}
                    onClick={this.handleNew}
                    className="btn btn-primary btn-xl"
                    id="sendMessageButton"
                    type="button"
                >
                    Generate new token
                </Button></div>



            </React.Fragment>

        );
    }

}
export default Tokens;
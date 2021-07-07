
import React from "react";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import TopBar from "../components/TopBar";
import Header from "../components/Header";
import Axios from "axios";
import { BASE_URL } from "../constants.js";
import ModalDialog from "../components/ModalDialog";
import getAuthHeader from "../GetHeader";

class BestInfluencers extends React.Component {
    state = {
        bestInfluencers : []
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
        this.handleGetBestInfluencers()
    }
    
    handleGetBestInfluencers = ()=> {
      
		Axios.get(BASE_URL + "/api/campaign/bestPromoters/", {  headers: { Authorization: getAuthHeader() } } )
			.then((res) => {
				this.setState({ bestInfluencers: res.data });
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
         
            <div className="container">
                        <div className="container" style={{ marginTop: "10rem", marginRight: "10rem" }}>
                            <h3>
                            <h5 className=" text-center mb-0 mt-2 text-uppercase">Best influencers</h5>
                            </h3>
                            <table className="table" style={{ width: "100%" }}>
                                        <tbody>
                                        {this.state.bestInfluencers.map((user) => (
                                                <tr id={user.UserId} key={user.UserId}>
                                                    
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Username: </b> {user.Username}
                                                        </div>
                                                    </td>
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Number of partnerships: </b> {user.NumberOfPartnerships}
                                                        </div>
                                                    </td>
                                                    <td >
                                                        <div style={{ marginTop: "1rem"}}>
                                                            <b>Number of clicks on link from post: </b> {user.NumberOfClicks}
                                                        </div>
                                                    </td>
                                                   
                                                </tr>

                                            ))}
                                            <tr>
                                                <td></td>
                                                <td></td>
                                                <td></td>
                                            </tr>
                                        </tbody>
                                    </table>
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
export default BestInfluencers;
import React from "react";
import Header from "../components/Header";
import TopBar from "../components/TopBar";
import { Link } from "react-router-dom";

import { CgProfile } from 'react-icons/cg';
class HomePage extends React.Component {
	
	render() {
		return (
			<React.Fragment>
				<TopBar />
				<Header />

				<section className="d-flex align-items-top">
					<div className="container"  style={{ marginTop: "9rem", marginRight: "10rem"}}>
						<h1><CgProfile  size={200}  />  </h1>
					</div>
					
				</section>
				<section style={{ marginTop: "0rem", marginLeft: "50rem"}}>
					<div className="container" >
						<h3>USERNAME</h3>
					</div>
				</section>
			</React.Fragment>
		);
	}
}

export default HomePage;
import React from "react";
import Header from "./Header";
import TopBar from "./TopBar";
import { Link } from "react-router-dom";

class HomePage extends React.Component {
	
	render() {
		return (
			<React.Fragment>
				<TopBar />
				<Header />

				<section id="hero" className="d-flex align-items-top">
					<div className="container">
						<h1>Welcome</h1>
                        <Link  to="/registration" className="btn-get-started scrollto">
							Register
						</Link>
					</div>
					
				</section>
			</React.Fragment>
		);
	}
}

export default HomePage;
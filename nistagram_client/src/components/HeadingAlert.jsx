import React, { Component } from "react";

class HeadingAlert extends Component {
	render() {
		return (
			<div className="alert alert-danger alert-dismissible fade show" hidden={this.props.hidden} role="alert">
				<h4 className="alert-heading">{this.props.header}</h4>
				<hr />
				<p className="mb-0">{this.props.message}</p>
				<button type="button" className="close" onClick={this.props.handleCloseAlert}>
					<span aria-hidden="true">&times;</span>
				</button>
			</div>
		);
	}
}

export default HeadingAlert;

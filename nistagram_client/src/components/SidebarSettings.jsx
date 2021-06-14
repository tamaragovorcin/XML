import React, { Component } from "react";
import { ProSidebar, Menu, MenuItem, SubMenu } from 'react-pro-sidebar';
import { Link } from "react-router-dom";

class SidebarSettings extends Component{
    render() {

		return (
			<React.Fragment>
			
				<div>
					 
                    <div class="sidenav shadow p-3 bg-white rounded">
					<Link to="/editProfile">Edit profile</Link>
					<Link to="/passwordChange">Change password</Link>
                    <Link to="/settings">Notification settings</Link>
                    </div>
				</div>
				
			</React.Fragment>
		);
	}
}
export default SidebarSettings
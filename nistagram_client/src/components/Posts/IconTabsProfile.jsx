import React from "react";
import {Tabs, Tab} from 'react-bootstrap';



class IconTabsProfile extends React.Component {
  constructor(props){
    super(props);
    this.state = {
        key: 1 | props.activeKey
    }
    this.handleSelect = this.handleSelect.bind(this);
}

handleSelect (key) {
    console.log("selected " + key);
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
            Tab 1 content
            </Tab>
            <Tab eventKey={2} title="Albums">
            Tab 2 content
            </Tab>
            <Tab eventKey={3} title="All stories">
            Tab 3 content
            </Tab>
            <Tab eventKey={4} title="Highlights">
            Tab 4 content
            </Tab>
            <Tab eventKey={5} title="Saved">
            Tab 5 content
            </Tab>
        </Tabs>
    );

	}
}
export default IconTabsProfile;
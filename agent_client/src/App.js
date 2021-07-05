import './App.css';
import { HashRouter as Router, Link, Switch, Route  } from "react-router-dom";

import Login from './pages/Login'
import Registration from './pages/Registration'
import Unauthorized from './components/Unauthorized'
import PasswordChange from './pages/PasswordChange'
import Orders from './pages/Orders'
import All from './pages/All'
import AllAgents from './pages/AllAgents'
import Homepage from './pages/Homepage'
import NewCampaigns from './pages/NewCampaigns'
import Ordered from './pages/Ordered'
function App() {
  return (
    <Router>
			<Switch>
				<Link exact to="/login" path="/login" component={Login} />
				
				<Link exact to="/" path="/" component={All} />
				<Link exact to="/registration" path="/registration" component={Registration} />
				<Link exact to="/unauthorized" path="/unauthorized" component={Unauthorized} />
				<Link exact to="/allAgents" path="/allAgents" component={AllAgents} />
				<Link exact to="/passwordChange" path="/passwordChange" component={PasswordChange} />
				<Link exact to="/all" path="/all" component={All} />
				<Link exact to="/orders" path="/orders" component={Orders} />
				<Link exact to="/ordered" path="/ordered" component={Ordered} />
				<Link exact to="/newCampaigns" path="/newCampaigns" component={NewCampaigns} />

			</Switch>
	</Router>
  );
}

export default App;

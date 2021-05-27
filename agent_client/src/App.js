import logo from './logo.svg';
import './App.css';

import { HashRouter as Router, Link, Switch } from "react-router-dom";
import HomePage from './components/HomePage'
import Login from './components/Login'
import Registration from './components/Registration'
import ProfileInfo from './components/Users/ProfileInfo'
import Unauthorized from './components/Unauthorized'


function App() {
  return (
    <Router>
			<Switch>
				<Link exact to="/" path="/" component={HomePage} />
				<Link exact to="/login" path="/login" component={Login} />
				<Link exact to="/registration" path="/registration" component={Registration} />
				<Link exact to="/profile-info" path="/profile-info" component={ProfileInfo} />
				<Link exact to="/unauthorized" path="/unauthorized" component={Unauthorized} />
			</Switch>
	</Router>
  );
}


export default App;

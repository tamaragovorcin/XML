import './App.css';
import { HashRouter as Router, Link, Switch } from "react-router-dom";
import HomePage from './pages/HomePage'
import Login from './pages/Login'
import Registration from './pages/Registration'
import ProfileInfo from './components/Users/ProfileInfo'
import Unauthorized from './components/Unauthorized'
import UserProfilePage from './pages/UserProfilePage'
import ProfilePage from './pages/ProfilePage'
function App() {
  return (
    <Router>
			<Switch>
				<Link exact to="/" path="/" component={HomePage} />
				<Link exact to="/login" path="/login" component={Login} />
				<Link exact to="/registration" path="/registration" component={Registration} />
				<Link exact to="/profile-info" path="/profile-info" component={ProfileInfo} />
				<Link exact to="/unauthorized" path="/unauthorized" component={Unauthorized} />
				<Link exact to="/userChangeProfile" path="/userChangeProfile" component={UserProfilePage} />
				<Link exact to="/profilePage" path="/profilePage" component={ProfilePage} />
			</Switch>
	</Router>
  );
}

export default App;

import './App.css';
import { HashRouter as Router, Link, Switch, Route  } from "react-router-dom";
import HomePage from './pages/HomePage'
import Login from './pages/Login'
import Registration from './pages/Registration'
import Unauthorized from './components/Unauthorized'
import UserProfilePage from './pages/UserProfilePage'
import ProfilePage from './pages/ProfilePage'
import EditProfile from './pages/EditProfile';
import PasswordChange from './pages/PasswordChange'
import Favorites from './pages/Favorites';
import PharmacyProfilePage from './pages/FollowerProfilePage';
import Search from './pages/Search';

function App() {
  return (
    <Router>
			<Switch>
				<Link exact to="/" path="/" component={HomePage} />
				<Link exact to="/login" path="/login" component={Login} />
				<Link exact to="/registration" path="/registration" component={Registration} />
				<Link exact to="/unauthorized" path="/unauthorized" component={Unauthorized} />
				<Link exact to="/userChangeProfile" path="/userChangeProfile" component={UserProfilePage} />
				<Link exact to="/profilePage" path="/profilePage" component={ProfilePage} />
				<Link exact to="/favorites" path="/favorites" component={Favorites} />

				<Link exact to="/passwordChange" path="/passwordChange" component={PasswordChange} />
				<Link exact to="/settings" path="/settings" component={EditProfile} />
				<Route path="/followerProfilePage/:id" children={<PharmacyProfilePage />} />
				<Link exact to="/seacrh" path="/seacrh" component={Search} />

			</Switch>
	</Router>
  );
}

export default App;

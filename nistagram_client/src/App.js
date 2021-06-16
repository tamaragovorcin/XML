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
import CloseFriends from './pages/CloseFriends';
import FollowRequest from './pages/FollowRequest';
import Settings from './pages/PrivacySettings';
import NotificationSettings from './pages/NotificationSettings'
import Notifications from './pages/Notifications'
import LikedAndDislikedPosts from './pages/LikedAndDislikedPosts';
import VerifyRequest from './pages/VerifyRequest';



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
				<Link exact to="/editProfile" path="/editProfile" component={EditProfile} />
				<Route path="/followerProfilePage/:id" children={<PharmacyProfilePage />} />
				<Link exact to="/seacrh" path="/seacrh" component={Search} />
				<Link exact to="/closeFriends" path="/closeFriends" component={CloseFriends} />
				<Link exact to="/followRequest" path="/followRequest" component={FollowRequest} />
				<Link exact to="/settings" path="/settings" component={Settings} />
				<Link exact to="/notificationSettings" path="/notificationSettings" component={NotificationSettings} />
				<Link exact to="/likedAndDisliked" path="/likedAndDisliked" component={LikedAndDislikedPosts} />
				<Link exact to="/verifyRequest" path="/verifyRequest" component={VerifyRequest} />
				<Link exact to="/notifications" path="/notifications" component={Notifications} />


			</Switch>
	</Router>
  );
}

export default App;

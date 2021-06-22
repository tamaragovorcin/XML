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
import LikedPosts from './pages/LikedPosts';
import DislikedPosts from './pages/DislikedPosts';
import VerifyRequest from './pages/VerifyRequest';
import ReportedPosts from './pages/ReportedPosts';
import AgentsRequests from './pages/AgentsRequests';
import RegisterNewAgent from './pages/RegisterNewAgent';
import PartnershipRequests from './pages/PartnershipRequests';
import BestInfluencers from './pages/BestInfluencers';
import ChatEnginePreview  from './chat/ChatEngine'

import './App.css';
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
				<Link exact to="/verifyRequest" path="/verifyRequest" component={VerifyRequest} />
				<Link exact to="/notifications" path="/notifications" component={Notifications} />

				<Link exact to="/likedPosts" path="/likedPosts" component={LikedPosts} />
				<Link exact to="/dislikedPosts" path="/dislikedPosts" component={DislikedPosts} />
				<Link exact to="/verifyRequest" path="/verifyRequest" component={VerifyRequest} />
				<Link exact to="/reportedPosts" path="/reportedPosts" component={ReportedPosts} />
				<Link exact to="/agentsR" path="/agentsR" component={AgentsRequests} />
				<Link exact to="/registerNewAgent" path="/registerNewAgent" component={RegisterNewAgent} />

				<Link exact to="/partnershipRequests" path="/partnershipRequests" component={PartnershipRequests} />
				<Link exact to="/bestInfluencers" path="/bestInfluencers" component={BestInfluencers} />
				<Link exact to="/chatEngine" path="/chatEngine" component={ChatEnginePreview} />

			</Switch>
	</Router>
  );
}

export default App;

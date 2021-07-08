package main

import (
	"github.com/gorilla/mux"

)


func (app *application) routes() *mux.Router {
	// Register handler functions.



	r := mux.NewRouter()

	r.HandleFunc("/api/proba/{token}", app.proba).Methods("GET")
	r.HandleFunc("/api/token/{userId}", app.getToken).Methods("GET")
	r.HandleFunc("/api/generateToken/{userId}", app.generateNewToken).Methods("GET")
	r.HandleFunc("/api/", app.getAllUsers).Methods("GET")
	r.HandleFunc("/api/all/{userId}", app.getAllUsersWithoutLogged).Methods("GET")

	r.HandleFunc("/api/{id}", app.findUserByID).Methods("GET")
	r.HandleFunc("/api/user/update/",  app.updateUser).Methods("POST")

	r.HandleFunc("/api/", app.insertUser).Methods("POST")
	r.HandleFunc("/admin/", app.insertAdmin).Methods("POST")

	r.HandleFunc("/api/search/{name}", app.search).Methods("GET")
	r.HandleFunc("/api/login", app.loginUser).Methods("POST")
	r.HandleFunc("/api/user/privacy/{userId}", app.findUserPrivacy).Methods("GET")
	r.HandleFunc("/api/user/username/{userId}", app.findUserUsername).Methods("GET")
	r.HandleFunc("/api/user/closeFriends/{userId}", app.findUserCloseFriends).Methods("GET")
	r.HandleFunc("/api/user/username/category/{userId}", app.findUserUsernameIfInfluencer).Methods("GET")
	r.HandleFunc("/api/user/genderOk/{userId}/{gender}", IsAuthorized(app.findIfGenderIsOk)).Methods("GET")
	r.HandleFunc("/api/user/dateOfBirthOk/{userId}/{dateOne}/{dateTwo}", IsAuthorized(app.findIfDateOfBirthIsOk)).Methods("GET")

	r.HandleFunc("/api/user/userId/{token}", IsAuthorized(app.findUserIdIfTokenExists)).Methods("GET")


	r.HandleFunc("/api/user/addToCloseFriends/", IsAuthorized(app.addUserToCloseFriends)).Methods("POST")
	r.HandleFunc("/api/user/removeFromCloseFriends/", IsAuthorized(app.removeUserFromCloseFriends)).Methods("POST")

	//r.HandleFunc("/api/getLoggedIn", app.getLoggedIn).Methods("GET")

	r.HandleFunc("/profileInformation/", IsAuthorized(app.getAllProfileInformation)).Methods("GET")
	r.HandleFunc("/profileInformation/{id}", IsAuthorized(app.findProfileInformationByID)).Methods("GET")
	r.HandleFunc("/profileInformation/", IsAuthorized(app.insertProfileInformation)).Methods("POST")
	r.HandleFunc("/profileInformation/{id}", IsAuthorized(app.deleteProfileInformation)).Methods("DELETE")


	r.HandleFunc("/api/role/", IsAuthorized(app.getAllRoles)).Methods("GET")
	r.HandleFunc("/api/role/{id}", IsAuthorized(app.findRoleByID)).Methods("GET")
	r.HandleFunc("/api/role/", IsAuthorized(app.insertRole)).Methods("POST")
	r.HandleFunc("/api/role/{id}", IsAuthorized(app.deleteRole)).Methods("DELETE")

	r.HandleFunc("/agent/", IsAuthorized(app.getAllAgents)).Methods("GET")
	r.HandleFunc("/agentRequests/", IsAuthorized(app.getAllAgentsRequests)).Methods("GET")
	r.HandleFunc("/agent/{id}", IsAuthorized(app.findAgentByID)).Methods("GET")
	r.HandleFunc("/agent", app.insertAgent).Methods("POST")
	r.HandleFunc("/agent/{id}", IsAuthorized(app.deleteAgent)).Methods("DELETE")

	r.HandleFunc("/api/user/profileImage/{userId}", IsAuthorized(app.saveImage)).Methods("POST")
	r.HandleFunc("/api/user/profileImage/{userId}", app.getUsersProfileImage).Methods("GET")

	r.HandleFunc("/notification/", app.getAllNotification).Methods("GET")
	r.HandleFunc("/notification/", app.insertNotification).Methods("POST")
	r.HandleFunc("/notification/{id}", app.deleteNotification).Methods("DELETE")

	r.HandleFunc("/api/user/privacySettings/", app.changePrivacySettings).Methods("POST")
	r.HandleFunc("/api/user/privacySettings/{userId}",app.getUsersPrivacySettings).Methods("GET")
	r.HandleFunc("/api/user/notificationSettings/{userId}",app.getUsersNotificationSettings).Methods("GET")

	r.HandleFunc("/api/mute/",app.muteUser).Methods("POST")
	r.HandleFunc("/api/block/",app.blockUser).Methods("POST")
	r.HandleFunc("/api/unmute/",app.unmuteUser).Methods("POST")
	r.HandleFunc("/api/unblock/",app.unblockUser).Methods("POST")
	r.HandleFunc("/api/user/blockedUsers/{userId}", app.findBlockedUsers).Methods("GET")
	r.HandleFunc("/api/user/mutedUsers/{userId}", app.findMutedUsers).Methods("GET")
	r.HandleFunc("/api/checkIfMuted/{subjectId}/{objectId}", app.checkIfUserIsMuted).Methods("GET")
	r.HandleFunc("/api/checkIfBlocked/{subjectId}/{objectId}", app.checkIfUserIsBlocked).Methods("GET")
	r.HandleFunc("/api/user/allowTags/{userId}", app.checkIfUserAllowsTags).Methods("GET")

	r.HandleFunc("/remove/{id}", app.deleteUser).Methods("DELETE")

	r.HandleFunc("/verificationRequestAll", app.getAllRequestVerification).Methods("GET")
	r.HandleFunc("/api/verificationRequest", app.insertVerification).Methods("POST")
	r.HandleFunc("/api/image/{userIdd}/{verificationId}", app.saveImageVerification).Methods("POST")
	r.HandleFunc("/verificationRequest/{id}", app.deleteVerification).Methods("DELETE")
	r.HandleFunc("/verificationRequest/accept/", app.acceptVerificationRequest).Methods("POST")

	r.HandleFunc("/agents/accept/",app.acceptAgentsRequest).Methods("POST")
	r.HandleFunc("/agent/byAdmin/", app.insertAgentByAdmin).Methods("POST")

	r.HandleFunc("/api/turnOnNotifications", app.turnOnNotificationsForUser).Methods("POST")
	r.HandleFunc("/api/addSettings/{userId}", app.addSettings).Methods("GET")
	r.HandleFunc("/api/updateNotifications", app.updateNotifications).Methods("POST")
	r.HandleFunc("/api/sendNotificationPost/{postType}/{userId}", app.sendNotificationPost).Methods("GET")
	r.HandleFunc("/api/sendNotificationComment/{writer}/{user}/{content}", app.sendNotificationComment).Methods("GET")
	r.HandleFunc("/api/getPostNotifications/{userId}", app.getPostNotifications).Methods("GET")
	r.HandleFunc("/api/getCommentNotification/{userId}", app.getCommentNotifications).Methods("GET")


	return r
}
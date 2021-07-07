package main

import (
	"github.com/gorilla/mux"

)


func (app *application) routes() *mux.Router {
	// Register handler functions.



	r := mux.NewRouter()

	r.HandleFunc("/api/proba/{token}", IsAuthorized(app.proba)).Methods("GET")
	r.HandleFunc("/api/token/{userId}", IsAuthorized(app.getToken)).Methods("GET")
	r.HandleFunc("/api/generateToken/{userId}", IsAuthorized(app.generateNewToken)).Methods("GET")
	r.HandleFunc("/api/", IsAuthorized(app.getAllUsers)).Methods("GET")
	r.HandleFunc("/api/all/{userId}", IsAuthorized(app.getAllUsersWithoutLogged)).Methods("GET")

	r.HandleFunc("/api/{id}", IsAuthorized(app.findUserByID)).Methods("GET")
	r.HandleFunc("/api/user/update/",  IsAuthorized(app.updateUser)).Methods("POST")

	r.HandleFunc("/api/", IsAuthorized(app.insertUser)).Methods("POST")
	r.HandleFunc("/admin/", IsAuthorized(app.insertAdmin)).Methods("POST")

	r.HandleFunc("/api/search/{name}", IsAuthorized(app.search)).Methods("GET")
	r.HandleFunc("/api/login", IsAuthorized(app.loginUser)).Methods("POST")
	r.HandleFunc("/api/user/privacy/{userId}", IsAuthorized(app.findUserPrivacy)).Methods("GET")
	r.HandleFunc("/api/user/username/{userId}", IsAuthorized(app.findUserUsername)).Methods("GET")
	r.HandleFunc("/api/user/closeFriends/{userId}", IsAuthorized(app.findUserCloseFriends)).Methods("GET")
	r.HandleFunc("/api/user/username/category/{userId}", IsAuthorized(app.findUserUsernameIfInfluencer)).Methods("GET")
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
	r.HandleFunc("/agent", IsAuthorized(app.insertAgent)).Methods("POST")
	r.HandleFunc("/agent/{id}", IsAuthorized(app.deleteAgent)).Methods("DELETE")

	r.HandleFunc("/api/user/profileImage/{userId}", IsAuthorized(app.saveImage)).Methods("POST")
	r.HandleFunc("/api/user/profileImage/{userId}", IsAuthorized(app.getUsersProfileImage)).Methods("GET")

	r.HandleFunc("/notification/", IsAuthorized(app.getAllNotification)).Methods("GET")
	r.HandleFunc("/notification/", IsAuthorized(app.insertNotification)).Methods("POST")
	r.HandleFunc("/notification/{id}", IsAuthorized(app.deleteNotification)).Methods("DELETE")

	r.HandleFunc("/api/user/privacySettings/", IsAuthorized(app.changePrivacySettings)).Methods("POST")
	r.HandleFunc("/api/user/privacySettings/{userId}",IsAuthorized(app.getUsersPrivacySettings)).Methods("GET")
	r.HandleFunc("/api/user/notificationSettings/{userId}",IsAuthorized(app.getUsersNotificationSettings)).Methods("GET")

	r.HandleFunc("/api/mute/",IsAuthorized(app.muteUser)).Methods("POST")
	r.HandleFunc("/api/block/",IsAuthorized(app.blockUser)).Methods("POST")
	r.HandleFunc("/api/unmute/",IsAuthorized(app.unmuteUser)).Methods("POST")
	r.HandleFunc("/api/unblock/",IsAuthorized(app.unblockUser)).Methods("POST")
	r.HandleFunc("/api/user/blockedUsers/{userId}", IsAuthorized(app.findBlockedUsers)).Methods("GET")
	r.HandleFunc("/api/user/mutedUsers/{userId}", IsAuthorized(app.findMutedUsers)).Methods("GET")
	r.HandleFunc("/api/checkIfMuted/{subjectId}/{objectId}", IsAuthorized(app.checkIfUserIsMuted)).Methods("GET")
	r.HandleFunc("/api/checkIfBlocked/{subjectId}/{objectId}", IsAuthorized(app.checkIfUserIsBlocked)).Methods("GET")
	r.HandleFunc("/api/user/allowTags/{userId}", IsAuthorized(app.checkIfUserAllowsTags)).Methods("GET")

	r.HandleFunc("/remove/{id}", IsAuthorized(app.deleteUser)).Methods("DELETE")

	r.HandleFunc("/verificationRequestAll", IsAuthorized(app.getAllRequestVerification)).Methods("GET")
	r.HandleFunc("/api/verificationRequest", IsAuthorized(app.insertVerification)).Methods("POST")
	r.HandleFunc("/api/image/{userIdd}/{verificationId}", IsAuthorized(app.saveImageVerification)).Methods("POST")
	r.HandleFunc("/verificationRequest/{id}", IsAuthorized(app.deleteVerification)).Methods("DELETE")
	r.HandleFunc("/verificationRequest/accept/", IsAuthorized(app.acceptVerificationRequest)).Methods("POST")

	r.HandleFunc("/agents/accept/", IsAuthorized(app.acceptAgentsRequest)).Methods("POST")
	r.HandleFunc("/agent/byAdmin/", IsAuthorized(app.insertAgentByAdmin)).Methods("POST")

 	r.HandleFunc("/api/turnOnNotifications", IsAuthorized(app.turnOnNotificationsForUser)).Methods("POST")
	r.HandleFunc("/api/addSettings/{userId}", IsAuthorized(app.addSettings)).Methods("GET")
	r.HandleFunc("/api/updateNotifications", IsAuthorized(app.updateNotifications)).Methods("POST")
	r.HandleFunc("/api/sendNotificationPost/{postType}/{userId}", IsAuthorized(app.sendNotificationPost)).Methods("GET")
	r.HandleFunc("/api/sendNotificationComment/{writer}/{user}/{content}", IsAuthorized(app.sendNotificationComment)).Methods("GET")
	r.HandleFunc("/api/getPostNotifications/{userId}", IsAuthorized(app.getPostNotifications)).Methods("GET")
	r.HandleFunc("/api/getCommentNotification/{userId}", IsAuthorized(app.getCommentNotifications)).Methods("GET")


	return r
}
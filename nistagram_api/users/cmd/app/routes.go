package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()

	r.HandleFunc("/api/", app.getAllUsers).Methods("GET")
	r.HandleFunc("/api/all/{userId}", app.getAllUsersWithoutLogged).Methods("GET")

	r.HandleFunc("/api/{id}", app.findUserByID).Methods("GET")
	r.HandleFunc("/api/user/update/", app.updateUser).Methods("POST")
	r.HandleFunc("/api/", app.insertUser).Methods("POST")
	r.HandleFunc("/admin/", app.insertAdmin).Methods("POST")

	r.HandleFunc("/api/search/{name}", app.search).Methods("GET")
	r.HandleFunc("/api/login", app.loginUser).Methods("POST")
	r.HandleFunc("/api/user/privacy/{userId}", app.findUserPrivacy).Methods("GET")
	r.HandleFunc("/api/user/username/{userId}", app.findUserUsername).Methods("GET")
	r.HandleFunc("/api/user/closeFriends/{userId}", app.findUserCloseFriends).Methods("GET")

	r.HandleFunc("/api/user/addToCloseFriends/", app.addUserToCloseFriends).Methods("POST")
	r.HandleFunc("/api/user/removeFromCloseFriends/", app.removeUserFromCloseFriends).Methods("POST")

	//r.HandleFunc("/api/getLoggedIn", app.getLoggedIn).Methods("GET")

	r.HandleFunc("/profileInformation/", app.getAllProfileInformation).Methods("GET")
	r.HandleFunc("/profileInformation/{id}", app.findProfileInformationByID).Methods("GET")
	r.HandleFunc("/profileInformation/", app.insertProfileInformation).Methods("POST")
	r.HandleFunc("/profileInformation/{id}", app.deleteProfileInformation).Methods("DELETE")

	r.HandleFunc("/verification/", app.getAllVerifications).Methods("GET")
	r.HandleFunc("/verification/{id}", app.findVerificationByID).Methods("GET")
	r.HandleFunc("/verification/", app.insertVerification).Methods("POST")
	r.HandleFunc("/verification/{id}", app.deleteVerification).Methods("DELETE")

	r.HandleFunc("/api/role/", app.getAllRoles).Methods("GET")
	r.HandleFunc("/api/role/{id}", app.findRoleByID).Methods("GET")
	r.HandleFunc("/api/role/", app.insertRole).Methods("POST")
	r.HandleFunc("/api/role/{id}", app.deleteRole).Methods("DELETE")

	r.HandleFunc("/agent/", app.getAllAgents).Methods("GET")
	r.HandleFunc("/agent/{id}", app.findAgentByID).Methods("GET")
	r.HandleFunc("/agent/", app.insertAgent).Methods("POST")
	r.HandleFunc("/agent/{id}", app.deleteAgent).Methods("DELETE")

	r.HandleFunc("/api/user/profileImage/{userId}", app.saveImage).Methods("POST")
	r.HandleFunc("/api/user/profileImage/{userId}", app.getUsersProfileImage).Methods("GET")

	r.HandleFunc("/notification/", app.getAllNotification).Methods("GET")
	r.HandleFunc("/notification/{id}", app.findByIDNotification).Methods("GET")
	r.HandleFunc("/notification/", app.insertNotification).Methods("POST")
	r.HandleFunc("/notification/{id}", app.deleteNotification).Methods("DELETE")

	r.HandleFunc("/api/user/privacySettings/", app.changePrivacySettings).Methods("POST")
	r.HandleFunc("/api/user/privacySettings/{userId}",app.getUsersPrivacySettings).Methods("GET")
	r.HandleFunc("/api/mute/",app.muteUser).Methods("POST")
	r.HandleFunc("/api/block/",app.blockUser).Methods("POST")
	r.HandleFunc("/api/unmute/",app.unmuteUser).Methods("POST")
	r.HandleFunc("/api/unblock/",app.unblockUser).Methods("POST")
	r.HandleFunc("/api/user/blockedUsers/{userId}", app.findBlockedUsers).Methods("GET")
	r.HandleFunc("/api/user/mutedUsers/{userId}", app.findMutedUsers).Methods("GET")
	r.HandleFunc("/api/checkIfMuted/{subjectId}/{objectId}", app.checkIfUserIsMuted).Methods("GET")
	r.HandleFunc("/api/checkIfBlocked/{subjectId}/{objectId}", app.checkIfUserIsBlocked).Methods("GET")
	r.HandleFunc("/api/user/allowTags/{userId}", app.checkIfUserAllowsTags).Methods("GET")

	r.HandleFunc("/user/remove/{id}", app.deleteUser).Methods("DELETE")

	//r.HandleFunc("api/allVerificationRequest/", app.getAllRequestVerification).Methods("GET")

	return r
}
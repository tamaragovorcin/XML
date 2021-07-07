package main

import (
	//"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"strconv"
	"strings"
	"users/saga"
	//"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"users/pkg/dtos"
	"users/pkg/models"
)


func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			fmt.Println("No Token Found")

			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized")
			return
		}


		authStringHeader := r.Header.Get("Authorization")
		if authStringHeader == "" {
			fmt.Errorf("Neki eror za auth")
		}
		authHeader := strings.Split(authStringHeader, "Bearer ")
		jwtToken := authHeader[1]

		token, err := jwt.Parse(jwtToken, func (token *jwt.Token) (interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("luna") , nil
		})

		if err != nil {
			fmt.Println("Your Token has been expired.")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}



		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			rolesString, _ := claims["roles"].(string)
			fmt.Println(rolesString)
			var tokenRoles []models.Role

			if err := json.Unmarshal([]byte(rolesString), &tokenRoles); err != nil {
				fmt.Println("Usercccc.")
			}



		} else{
			fmt.Println("User authorize fail.")
		}
	}


}


func equalPasswords(hashedPwd string, passwordRequest string) bool {

	byteHash := []byte(hashedPwd)
	plainPwd := []byte(passwordRequest)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}


func (app *application) loginUser(w http.ResponseWriter, r *http.Request)  {

	var loginRequest dtos.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		app.serverError(w, err)
	}
	//ctx = c.Request().Context()
	user, err := app.users.FindByUsername(loginRequest.Username)
	if user == nil {
		app.infoLog.Println("User not found")
	}

	if err != nil {
		app.infoLog.Println("Invalid email")
	}
	if user.ApprovedAgent== "wait" {
		app.infoLog.Println("Agent is not accepted")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
	}


	token := ""

	if(user.Token == ""){
		token, _ = generateToken(user)

		user.Token = token
		_, _ = app.users.Update(*user)
	} else {
		token = user.Token
	}









	//rolesString, _ := json.Marshal(user.ProfileInformation.Roles[0].Name)
	//b, err := json.Marshal(user)
	if err != nil {
		app.serverError(w, err)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	expireTime := time.Now().Add(time.Hour).Unix() * 1000
	userToken := dtos.UserTokenState{ AccessToken: token, Roles: user.ProfileInformation.Roles[0].Name, UserId: user.Id, ExpiresIn: expireTime,

	}
	bb, err := json.Marshal(userToken)
	w.Write(bb)
}

func generateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	rolesString, _ := json.Marshal(user.ProfileInformation.Roles)
	expireTime := time.Now().Add(time.Hour).Unix() * 1000
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.ProfileInformation.Email
	claims["name"] = user.ProfileInformation.Name
	claims["surname"] = user.ProfileInformation.LastName
	claims["username"] = user.ProfileInformation.Username
	claims["roles"] = string(rolesString)
	claims["id"] = user.Id
	claims["exp"] = strconv.FormatInt(expireTime, 10)

	return token.SignedString([]byte("luna"))
}


func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	users, err := app.users.GetAll()

	if err != nil {
		app.serverError(w, err)
	}
	usersAll := getUsersWithoutAdmin(users)
	b, err := json.Marshal(usersAll)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Users have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) getAllAgentsRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	users, err := app.users.GetAll()

	if err != nil {
		app.serverError(w, err)
	}
	usersList :=[]models.User{}
	for _, oneUser := range users {
		if(strings.Contains(oneUser.ApprovedAgent, "wait")){
			usersList = append(usersList, oneUser)
		}
	}
	usersAll := usersList
	b, err := json.Marshal(usersAll)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Users have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func getUsersWithoutAdmin(users []models.User) interface{} {
	usersList :=[]models.User{}
	for _, oneUser := range users {
		if oneUser.ProfileInformation.Roles[0].Name!="ADMIN" {
			usersList = append(usersList, oneUser)
		}
	}
	return usersList
}


func (app *application) proba(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)

	token := vars["token"]
	users, _ := app.users.GetAll()
	for _, oneUser := range users {

		if oneUser.Token==token {
			app.infoLog.Println("Token je okej")
		} else{
			app.infoLog.Println("Nije okej")
		}
	}



}

func (app *application) getToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)

	user := vars["userId"]
	users, err := app.users.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	token := ""
	for _, oneUser := range users {

		if oneUser.Id.Hex()==user {
			token = oneUser.Token
		}
	}

	b, err := json.Marshal(token)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Found token")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}



func (app *application) generateNewToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)

	user := vars["userId"]
	users, err := app.users.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	token := ""
	for _, oneUser := range users {

		if oneUser.Id.Hex()==user {

				token, _ = generateToken(&oneUser)

				oneUser.Token = token
				_, _ = app.users.Update(oneUser)

		}
	}

	b, err := json.Marshal(token)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Found token")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}


func (app *application) getAllUsersWithoutLogged(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)

	user := vars["userId"]
	users, err := app.users.GetAll()
	if err != nil {
		app.serverError(w, err)
	}
	userWithoutLogged := []models.User{}
	for _, oneUser := range users {

		if oneUser.Id.Hex()!=user {
			if oneUser.ProfileInformation.Roles[0].Name!="ADMIN" {
				userWithoutLogged = append(userWithoutLogged, oneUser)
			}
		}
	}

	b, err := json.Marshal(userWithoutLogged)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Users have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	name := vars["name"]

	users, err := app.users.GetAll()
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("User not found")
			return
		}
		app.serverError(w, err)
	}

	var u = []models.User{}

	for i, s := range users {
		fmt.Println(i, s)
		if(strings.Contains(s.ProfileInformation.Username, name)){
			u = append(u, s)
		}
	}

	b, err := json.Marshal(u)
	if err != nil {
		app.serverError(w, err)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id := vars["id"]
	//intVar, err := strconv.Atoi(id)
	intVar, err := primitive.ObjectIDFromHex(id)
	m, err := app.users.FindByID(intVar)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("User not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) acceptAgentsRequest(w http.ResponseWriter, r *http.Request) {

	var mm dtos.AgentsReactionDTO
	err := json.NewDecoder(r.Body).Decode(&mm)
	if err != nil {
		app.serverError(w, err)
	}

	user, err := app.users.FindByID(mm.UserId)
	if user == nil {
		app.infoLog.Println("User not found")
	}

	var userUpdate = models.User{
		Id: user.Id,
		ProfileInformation: user.ProfileInformation,
		Biography: user.Biography,
		Private: user.Private,
		Verified: false,
		Website: user.Website,
		Category : user.Category,
		ApprovedAgent: "true",
	}

	insertResult, err := app.users.Update(userUpdate)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

func (app *application) findUserUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)

	userId := vars["userId"]
	intVar, err := primitive.ObjectIDFromHex(userId)
	m, err := app.users.FindByID(intVar)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("User not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m.ProfileInformation.Username)
	fmt.Println(m.ProfileInformation.Username)
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}


func (app *application) findUserUsernameIfInfluencer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	intVar, err := primitive.ObjectIDFromHex(userId)
	m, err := app.users.FindByID(intVar)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("User not found")
			return
		}
		app.serverError(w, err)
	}
	forMarshal := ""
	if m.Category=="INFLUENCER" {
		forMarshal = m.ProfileInformation.Username
	} else {
		forMarshal = "not"
	}
	b, err := json.Marshal(forMarshal)
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) findUserPrivacy(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	intVar, err := primitive.ObjectIDFromHex(userId)
	m, err := app.users.FindByID(intVar)
	if err != nil {
	if err.Error() == "ErrNoDocuments" {
		app.infoLog.Println("User not found")
		return
	}
	app.serverError(w, err)
	}
	writing := ""
	if m.Private==true {
		writing="private"
	}else if m.Private==false {
		writing = "public"
	}
	b, err := json.Marshal(writing)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findIfGenderIsOk(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	gender := vars["gender"]

	intVar, err := primitive.ObjectIDFromHex(userId)
	m, err := app.users.FindByID(intVar)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("User not found")
			return
		}
		app.serverError(w, err)
	}
	writing := ""
	if strings.ToLower(m.ProfileInformation.Gender)== strings.ToLower(gender){
		writing="sameGender"
	}else {
		writing = "notSame"
	}
	b, err := json.Marshal(writing)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) findUserIdIfTokenExists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	token := vars["token"]
	users, _ := app.users.GetAll()
	userIdString := ""
	for _, oneUser := range users {

		if strings.ToLower(oneUser.Token) == strings.ToLower(token) {
			userIdString = oneUser.Id.Hex()
			app.infoLog.Println("Token je okej")
		} else{
			app.infoLog.Println("Nije okej")
		}
	}
	forMarshal:=""
	fmt.Println(userIdString)
	if userIdString=="" {
		forMarshal = "not"
	}else {
		forMarshal = userIdString

	}
	b, err := json.Marshal(forMarshal)
	if err != nil {
		app.serverError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
func (app *application) findIfDateOfBirthIsOk(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	userId := vars["userId"]
	dateOne := vars["dateOne"]
	dateTwo := vars["dateTwo"]

	intVar, err := primitive.ObjectIDFromHex(userId)
	m, err := app.users.FindByID(intVar)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("User not found")
			return
		}
		app.serverError(w, err)
	}
	writing := ""
	if dateOfBirthBetweenTwoDates(m.ProfileInformation.DateOfBirth, dateOne,dateTwo){
		writing="dateOfBirthOk"
	}else {
		writing = "notOk"
	}
	b, err := json.Marshal(writing)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a user")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func dateOfBirthBetweenTwoDates(birth string, one string, two string) bool {
	layout := "2006-01-02T15:04:05.000Z"
	stringBirth := birth+"T11:45:26.371Z"
	stringDateOne := one+"T11:45:26.371Z"
	stringDateTwo := two+"T11:45:26.371Z"

	timeBirth, err := time.Parse(layout, stringBirth)
	timeDateOne, err := time.Parse(layout, stringDateOne)
	timeDateTwo, err := time.Parse(layout, stringDateTwo)

	if err != nil {
		fmt.Println(err)
	}
	if timeBirth.After(timeDateOne) && timeBirth.Before(timeDateTwo) {
		return true
	}
	return false
}
func HashAndSaltPasswordIfStrong(password string) (string, error) {


	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash), err
}

func (app *application) insertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")



	var m dtos.UserRequest
	var able = true
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}


	users, err := app.users.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	for i, s := range users {
		fmt.Println(i, s)
		if (s.ProfileInformation.Username == m.Username) {
			app.infoLog.Printf("This username is already taken")

			able = false

			var e = errors.New("This username is already taken")
			b, err := json.Marshal(e)
			if err != nil {
				app.serverError(w, err)
			}

			w.Header().Set("Content-Type", "text")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(b)
			app.serverError(w, errors.New("This username is already taken"))
			break

		}
	}
		/*if( s.ProfileInformation.Email == m.Email){
			app.infoLog.Printf("User with this email already exists")
			able = false

			var e = errors.New("User with this email already exists")
			b, err := json.Marshal(e)
			if err != nil {
				app.serverError(w, err)
			}



			w.Header().Set("Content-Type", "text")
			w.WriteHeader(http.StatusConflict)
			w.Write(b)
			app.serverError(w, errors.New("User with this email already exists"))
			break

		}

	}*/

	if able  {
		hashAndSalt, err := HashAndSaltPasswordIfStrong(m.Password)

		var profileInformation = models.ProfileInformation{
			Name: m.Name, LastName: m.LastName,
			Email:       m.Email,
			Username:    m.Username,
			Password:    hashAndSalt,
			Roles:       []models.Role{{Name: "USER"}},
			PhoneNumber: m.PhoneNumber,
			Gender:      m.Gender, //models.Gender(m.Gender),
			DateOfBirth: m.DateOfBirth,
		}
		intId := primitive.NewObjectID()
		var user = models.User{
			Id: intId,
			Status: 0,
			ProfileInformation: profileInformation,
			Biography:          m.Biography,
			Private:            m.Private,
			Verified:           false,
			Website: m.Website,
			ApprovedAgent: "false",
		}
		idd := intId
		users, err := app.users.GetAll()
		for _, s := range users {

			if s.ProfileInformation.Email == m.Email {

				idd = s.Id
			}
		}
		insertResult, err := app.users.Insert(user)


		if err != nil {
			app.serverError(w, err)
		}



		m := saga.Message{Service: saga.ServiceInteraction, SenderService: saga.ServiceUser, Action: saga.ActionStart, User: idd.String(), User2: intId.String()}
		fmt.Println(m)
		saga.NewOrchestrator().Next(saga.InteractionChannel, saga.ServiceInteraction, m)


		time.Sleep(2 * time.Second)

		u,_ := app.users.FindByID(intId)

		fmt.Println(u.Status)
		if(u.Status == 1){


			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			idMarshaled, _ := json.Marshal(insertResult.InsertedID)

			w.Write(idMarshaled)


		}

		if(u.Status == 2){




			_, _ = app.users.DeleteId(intId)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusConflict)
			idMarshaled, _ := json.Marshal(insertResult.InsertedID)

			w.Write(idMarshaled)


		}





	}
}

func (app *application) insertAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var m dtos.UserRequest
	var able = true
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}


	users, err := app.users.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	for i, s := range users {
		fmt.Println(i, s)
		if(s.ProfileInformation.Username == m.Username){
			app.infoLog.Printf("This username is already taken")

			able = false


			var e = errors.New("This username is already taken")
			b, err := json.Marshal(e)
			if err != nil {
				app.serverError(w, err)
			}



			w.Header().Set("Content-Type", "text")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(b)
			app.serverError(w, errors.New("This username is already taken"))
			break

		}

		if( s.ProfileInformation.Email == m.Email){
			app.infoLog.Printf("User with this email already exists")
			able = false

			var e = errors.New("User with this email already exists")
			b, err := json.Marshal(e)
			if err != nil {
				app.serverError(w, err)
			}



			w.Header().Set("Content-Type", "text")
			w.WriteHeader(http.StatusConflict)
			w.Write(b)
			app.serverError(w, errors.New("User with this email already exists"))
			break

		}

	}

	if able  {
		hashAndSalt, err := HashAndSaltPasswordIfStrong(m.Password)

		var profileInformation = models.ProfileInformation{
			Name: m.Name, LastName: m.LastName,
			Email:       m.Email,
			Username:    m.Username,
			Password:    hashAndSalt,
			Roles:       []models.Role{{Name: "ADMIN"}},
			PhoneNumber: m.PhoneNumber,
			Gender:      m.Gender, //models.Gender(m.Gender),
			DateOfBirth: m.DateOfBirth,

		}

		var user = models.User{
			ProfileInformation: profileInformation,
			Biography:          m.Biography,
			Private:            m.Private,
			Verified:           false,
			Website: m.Website,
			ApprovedAgent : "false",
		}

		insertResult, err := app.users.Insert(user)
		if err != nil {
			app.serverError(w, err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		idMarshaled, err := json.Marshal(insertResult.InsertedID)

		w.Write(idMarshaled)
	}
}
func (app *application) insertAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var m dtos.UserRequest
	var able = true
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}


	users, err := app.users.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	for i, s := range users {
		fmt.Println(i, s)
		if(s.ProfileInformation.Username == m.Username){
			app.infoLog.Printf("This username is already taken")

			able = false


			var e = errors.New("This username is already taken")
			b, err := json.Marshal(e)
			if err != nil {
				app.serverError(w, err)
			}



			w.Header().Set("Content-Type", "text")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(b)
			app.serverError(w, errors.New("This username is already taken"))
			break

		}

		if( s.ProfileInformation.Email == m.Email){
			app.infoLog.Printf("User with this email already exists")
			able = false

			var e = errors.New("User with this email already exists")
			b, err := json.Marshal(e)
			if err != nil {
				app.serverError(w, err)
			}



			w.Header().Set("Content-Type", "text")
			w.WriteHeader(http.StatusConflict)
			w.Write(b)
			app.serverError(w, errors.New("User with this email already exists"))
			break

		}

	}

	if able  {
		hashAndSalt, err := HashAndSaltPasswordIfStrong(m.Password)

		var profileInformation = models.ProfileInformation{
			Name: m.Name, LastName: m.LastName,
			Email:       m.Email,
			Username:    m.Username,
			Password:    hashAndSalt,
			Roles:       []models.Role{{Name: "AGENT"}},
			PhoneNumber: m.PhoneNumber,
			Gender:      m.Gender, //models.Gender(m.Gender),
			DateOfBirth: m.DateOfBirth,
		}

		var user = models.User{
			ProfileInformation: profileInformation,
			Biography:          m.Biography,
			Private:            m.Private,
			Verified:           false,
			Website: m.Website,
			ApprovedAgent : "wait",
		}

		insertResult, err := app.users.Insert(user)
		if err != nil {
			app.serverError(w, err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		idMarshaled, err := json.Marshal(insertResult.InsertedID)

		w.Write(idMarshaled)
	}
}
func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id := vars["id"]
	intId, err := primitive.ObjectIDFromHex(id)

	deleteResult, err := app.users.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}
	removeFromCloseFriends(intId,app)
	removeFromBlocked(intId,app)
	removeFromMuted(intId,app)

	app.infoLog.Printf("Have been eliminated %d users(s)", deleteResult.DeletedCount)
}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {

		var m dtos.UserUpdateRequest
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			app.serverError(w, err)
		}
		sb := m.Id
		sb = sb[1:]
		sb = sb[:len(sb)-1]
	intId, err := primitive.ObjectIDFromHex(sb)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		uss, err := app.users.FindByID(intId)

	if uss == nil {
			app.infoLog.Println("User not found")
		}

		if err != nil {
			app.infoLog.Println("Invalid email")
		}
		var profileInformation = models.ProfileInformation{
			Name: m.Name,
			LastName: m.LastName,
			Email:       m.Email,
			Username:    m.Username,
			PhoneNumber: m.PhoneNumber,
			Gender:  m.Gender,//models.Gender(m.Gender),
			DateOfBirth: m.DateOfBirth,
		}

		var user = models.User{
			Id: intId,
			ProfileInformation: profileInformation,
			Biography: m.Biography,
			Private: m.Private,
			Verified: uss.Verified,
			Website: m.Website,
			ApprovedAgent : "false",
		}

		insertResult, err := app.users.Update(user)

	if err != nil {
			app.serverError(w, err)
		}

	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}
func (app *application) deleteAgent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.users.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d Agent(s)", deleteResult.DeletedCount)
}

func (app *application) insertAgentByAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var m dtos.UserRequest
	var able = true
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}


	users, err := app.users.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	for i, s := range users {
		fmt.Println(i, s)
		if(s.ProfileInformation.Username == m.Username){
			app.infoLog.Printf("This username is already taken")

			able = false


			var e = errors.New("This username is already taken")
			b, err := json.Marshal(e)
			if err != nil {
				app.serverError(w, err)
			}



			w.Header().Set("Content-Type", "text")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(b)
			app.serverError(w, errors.New("This username is already taken"))
			break

		}

		if( s.ProfileInformation.Email == m.Email){
			app.infoLog.Printf("User with this email already exists")
			able = false

			var e = errors.New("User with this email already exists")
			b, err := json.Marshal(e)
			if err != nil {
				app.serverError(w, err)
			}



			w.Header().Set("Content-Type", "text")
			w.WriteHeader(http.StatusConflict)
			w.Write(b)
			app.serverError(w, errors.New("User with this email already exists"))
			break

		}

	}

	if able  {
		hashAndSalt, err := HashAndSaltPasswordIfStrong(m.Password)

		var profileInformation = models.ProfileInformation{
			Name: m.Name, LastName: m.LastName,
			Email:       m.Email,
			Username:    m.Username,
			Password:    hashAndSalt,
			Roles:       []models.Role{{Name: "AGENT"}},
			PhoneNumber: m.PhoneNumber,
			Gender:      m.Gender, //models.Gender(m.Gender),
			DateOfBirth: m.DateOfBirth,
		}

		var user = models.User{
			ProfileInformation: profileInformation,
			Biography:          m.Biography,
			Private:            m.Private,
			Verified:           false,
			Website: m.Website,
			ApprovedAgent : "true",
		}

		insertResult, err := app.users.Insert(user)
		if err != nil {
			app.serverError(w, err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		idMarshaled, err := json.Marshal(insertResult.InsertedID)

		w.Write(idMarshaled)
	}
}
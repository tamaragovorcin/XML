package main

import (
	//"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"strings"

	"errors"
	//"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"users/pkg/dtos"
	"users/pkg/models"
)

type UserHandlers struct {

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

	token, err := generateToken(user)


	//rolesString, _ := json.Marshal(user.ProfileInformation.Roles[0].Name)
	//b, err := json.Marshal(user)
	if err != nil {
		app.serverError(w, err)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)


	userToken := dtos.UserTokenState{ AccessToken: token, Roles: user.ProfileInformation.Roles[0].Name, UserId: user.Id,

	}
	bb, err := json.Marshal(userToken)
	w.Write(bb)
}

func generateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	rolesString, _ := json.Marshal(user.ProfileInformation.Roles)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.ProfileInformation.Email
	claims["name"] = user.ProfileInformation.Name
	claims["surname"] = user.ProfileInformation.LastName
	claims["username"] = user.ProfileInformation.Username
	claims["roles"] = string(rolesString)
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	return  token.SignedString([]byte("luna"))
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

func getUsersWithoutAdmin(users []models.User) interface{} {
	usersList :=[]models.User{}
	for _, oneUser := range users {
		if oneUser.ProfileInformation.Roles[0].Name!="ADMIN" {
			usersList = append(usersList, oneUser)
		}
	}
	return usersList
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
			Roles:       []models.Role{{Name: "USER"}},
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
		}

		insertResult, err := app.users.Update(user)

	if err != nil {
			app.serverError(w, err)
		}

	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)
}

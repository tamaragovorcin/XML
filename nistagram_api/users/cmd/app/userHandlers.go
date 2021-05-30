package main

import (
	//"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"os"
	"strconv"

	//"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
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


	rolesString, _ := json.Marshal(user.ProfileInformation.Roles)

	//b, err := json.Marshal(user)
	if err != nil {
		app.serverError(w, err)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)


	userToken := dtos.UserTokenState{ AccessToken: token, Roles: string(rolesString), UserId: user.Id,

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
	chats, err := app.users.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(chats)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Users have been listed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}


func (app *application) findUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id := vars["id"]
	intVar, err := strconv.Atoi(id)
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
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	hashAndSalt, err := HashAndSaltPasswordIfStrong(m.Password)
	var profileInformation = models.ProfileInformation{
		Id: 1,
		Name: m.Name, LastName: m.LastName,
		Email:       m.Email,
		Username:    m.Username,
		Password:    hashAndSalt,
		Roles:       []models.Role{{ Id: 5, Name: "USER"}},
		PhoneNumber: m.PhoneNumber,
		Gender:  m.Gender,//models.Gender(m.Gender),
		DateOfBirth: m.DateOfBirth,
	}


	var user = models.User{Id: 5,
		ProfileInformation: profileInformation,
		Biography: m.Biography,
		Private: m.Private,
		Verified: false,
	}


	insertResult, err := app.users.Insert(user)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New user have been created, id=%s", insertResult.InsertedID)
}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.users.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d users(s)", deleteResult.DeletedCount)
}
func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	app.infoLog.Printf("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")


		var m dtos.UserUpdateRequest
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			app.serverError(w, err)
		}
		intId, err := strconv.Atoi(m.Id)
		if err != nil {
			// handle error
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
		app.infoLog.Printf("USERNAMEEE, %s",uss.ProfileInformation.Username)
		var profileInformation = models.ProfileInformation{
			Id: uss.Id,
			Name: m.Name,
			LastName: m.LastName,
			Email:       m.Email,
			Username:    m.Username,
			PhoneNumber: m.PhoneNumber,
			Gender:  m.Gender,//models.Gender(m.Gender),
			DateOfBirth: m.DateOfBirth,
		}


		var user = models.User{
			Id: uss.Id,
			ProfileInformation: profileInformation,
			Biography: m.Biography,
			Private: m.Private,
			Verified: false,
		}


		insertResult, err := app.users.Update(user)
		if err != nil {
			app.serverError(w, err)
		}

	app.infoLog.Printf("New user have been created, id=%s", insertResult.UpsertedID)

}

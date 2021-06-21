package main

import (
	"AgentApp/pkg/dtos"
	"AgentApp/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
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
			Roles:       m.Role,
			PhoneNumber: m.PhoneNumber,
			Gender:      m.Gender, //models.Gender(m.Gender),
			DateOfBirth: m.DateOfBirth,
		}

		var user = models.User{
			ProfileInformation: profileInformation,
			Website: m.Website,
		}

		insertResult, err := app.users.Insert(user)
		if err != nil {
			app.serverError(w, err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		idMarshaled, err := json.Marshal(insertResult.InsertedID)
		fmt.Println(idMarshaled)

		w.Write(idMarshaled)
	}
}

func HashAndSaltPasswordIfStrong(password string) (string, error) {


	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash), err
}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.users.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d users(s)", deleteResult.DeletedCount)
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
	//rolesString, _ := json.Marshal(user.ProfileInformation.Roles)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.ProfileInformation.Email
	claims["name"] = user.ProfileInformation.Name
	claims["surname"] = user.ProfileInformation.LastName
	claims["username"] = user.ProfileInformation.Username
	claims["roles"] = user.ProfileInformation.Roles
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	return  token.SignedString([]byte("luna"))
}


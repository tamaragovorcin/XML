package main

import (
	//"context"
	"encoding/json"
	//"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
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




/*func (app *application) loginUser(w http.ResponseWriter, r *http.Request) error  {

c echo.Context,ctx context.Context

	loginRequest := &dtos.LoginRequest{}
	if err := c.Bind(loginRequest); err != nil {
		return err
	}

	ctx = c.Request().Context()
	user, err := app.users.FindByUsername(loginRequest.Email)
	if err != nil {
		return errors.New("invalid email address")
	}

	if err != nil && user != nil {
		return c.JSON(http.StatusForbidden, map[string]string{
			"userId" : string(user.Id),
		})
	}

	token, err := generateToken(user)


	rolesString, _ := json.Marshal(user.ProfileInformation.Roles)
	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": token,
		"roles" : string(rolesString),
	})
}*/

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

	m, err := app.users.FindByID(id)
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
	app.infoLog.Println("Users ssssssssssssssss")
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
		Roles:       []models.Role{{ Id: 26, Name: "USER"}},
		PhoneNumber: m.PhoneNumber,
		Gender:  m.Gender,//models.Gender(m.Gender),
		DateOfBirth: m.DateOfBirth,
	}


	var user = models.User{Id: 7,
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

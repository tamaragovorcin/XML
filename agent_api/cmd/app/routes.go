package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()

	r.HandleFunc("/api/user/", IsAuthorized(app.getAllUsers)).Methods("GET")
	r.HandleFunc("/api/user/{id}", IsAuthorized(app.findUserByID)).Methods("GET")
	r.HandleFunc("/api/user/", IsAuthorized(app.insertUser)).Methods("POST")
	r.HandleFunc("/api/user/{id}", IsAuthorized(app.deleteUser)).Methods("DELETE")
	r.HandleFunc("/api/login", IsAuthorized(app.loginUser)).Methods("POST")
	r.HandleFunc("/location/", IsAuthorized(app.getAllLocations)).Methods("GET")
	r.HandleFunc("/location/{id}", IsAuthorized(app.findLocationByID)).Methods("GET")
	r.HandleFunc("/location/", IsAuthorized(app.insertLocation)).Methods("POST")
	r.HandleFunc("/location/{id}", IsAuthorized(app.deleteLocation)).Methods("DELETE")

	r.HandleFunc("/api/image/{userIdd}", IsAuthorized(app.saveImage)).Methods("POST")
	r.HandleFunc("/api/removeImage", IsAuthorized(app.deleteImage)).Methods("POST")
	r.HandleFunc("/api/addImages", IsAuthorized(app.addImages)).Methods("POST")


	r.HandleFunc("/api/feedAlbum/edit/{userIdd}", IsAuthorized(app.edit)).Methods("POST")
	r.HandleFunc("/product/", IsAuthorized(app.getAllProducts)).Methods("GET")
	r.HandleFunc("/product/{id}", IsAuthorized(app.findProductByID)).Methods("GET")
	r.HandleFunc("/api/product/remove/{id}", IsAuthorized(app.deleteProduct)).Methods("GET")
	r.HandleFunc("/api/product/{userId}", IsAuthorized(app.insertProduct)).Methods("POST")

	r.HandleFunc("/api/feedAlbum/usersAlbums/{userIdd}", IsAuthorized(app.getUsersFeedAlbums)).Methods("GET")
	r.HandleFunc("/api/feedAlbum/all/{userIdd}", IsAuthorized(app.getPosts)).Methods("GET")



	r.HandleFunc("/api/addToCart", IsAuthorized(app.addToCart)).Methods("POST")
	r.HandleFunc("/api/getAllCart/{userIdd}", IsAuthorized(app.getCart)).Methods("GET")
	r.HandleFunc("/api/getOrder/{userIdd}", IsAuthorized(app.getOrder)).Methods("GET")
	r.HandleFunc("/api/cart/remove/{id}", IsAuthorized(app.deleteCart)).Methods("GET")
	r.HandleFunc("/api/removeCart/{id}", IsAuthorized(app.removeCart)).Methods("GET")

	r.HandleFunc("/chosenProduct/", IsAuthorized(app.getAllChosenProducts)).Methods("GET")
	r.HandleFunc("/chosenProduct/{id}", IsAuthorized(app.findCHosenProductByID)).Methods("GET")
	r.HandleFunc("/chosenProduct/", IsAuthorized(app.insertChosenProduct)).Methods("POST")
	r.HandleFunc("/chosenProduct/{id}", IsAuthorized(app.deleteChosenProduct)).Methods("DELETE")


	r.HandleFunc("/api/purchase/{id}", IsAuthorized(app.getAllPurchases)).Methods("GET")
	r.HandleFunc("/purchase/{id}", IsAuthorized(app.findPurchaseByID)).Methods("GET")
	r.HandleFunc("/api/purchase", IsAuthorized(app.insertPurchase)).Methods("POST")
	r.HandleFunc("/purchase/{id}", IsAuthorized(app.deletePurchase)).Methods("DELETE")


	r.HandleFunc("/content/", IsAuthorized(app.getAllContents)).Methods("GET")
	r.HandleFunc("/content/{id}", IsAuthorized(app.findContentByID)).Methods("GET")
	r.HandleFunc("/content/", IsAuthorized(app.insertContent)).Methods("POST")
	r.HandleFunc("/content/{id}", IsAuthorized(app.deleteContent)).Methods("DELETE")

	r.HandleFunc("/api/bestCampaigns/{token}", IsAuthorized(app.getCampaignMonitoring)).Methods("GET")

	return r
}

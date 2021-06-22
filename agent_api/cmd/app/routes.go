package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/api/user/", app.getAllUsers).Methods("GET")
	r.HandleFunc("/api/user/{id}", app.findUserByID).Methods("GET")
	r.HandleFunc("/api/user/", app.insertUser).Methods("POST")
	r.HandleFunc("/api/user/{id}", app.deleteUser).Methods("DELETE")
	r.HandleFunc("/api/login", app.loginUser).Methods("POST")
	r.HandleFunc("/location/", app.getAllLocations).Methods("GET")
	r.HandleFunc("/location/{id}", app.findLocationByID).Methods("GET")
	r.HandleFunc("/location/", app.insertLocation).Methods("POST")
	r.HandleFunc("/location/{id}", app.deleteLocation).Methods("DELETE")

	r.HandleFunc("/api/image/{userIdd}", app.saveImage).Methods("POST")
	r.HandleFunc("/api/removeImage", app.deleteImage).Methods("POST")
	r.HandleFunc("/api/addImages", app.addImages).Methods("POST")


	r.HandleFunc("/api/feedAlbum/edit/{userIdd}", app.edit).Methods("POST")
	r.HandleFunc("/product/", app.getAllProducts).Methods("GET")
	r.HandleFunc("/product/{id}", app.findProductByID).Methods("GET")
	r.HandleFunc("/api/product/remove/{id}", app.deleteProduct).Methods("GET")
	r.HandleFunc("/api/product/{userId}", app.insertProduct).Methods("POST")

	r.HandleFunc("/api/feedAlbum/usersAlbums/{userIdd}", app.getUsersFeedAlbums).Methods("GET")
	r.HandleFunc("/api/feedAlbum/all/{userIdd}", app.getPosts).Methods("GET")



	r.HandleFunc("/api/addToCart", app.addToCart).Methods("POST")
	r.HandleFunc("/api/getAllCart/{userIdd}", app.getCart).Methods("GET")
	r.HandleFunc("/api/getOrder/{userIdd}", app.getOrder).Methods("GET")
	r.HandleFunc("/api/cart/remove/{id}", app.deleteCart).Methods("GET")
	r.HandleFunc("/api/removeCart/{id}", app.removeCart).Methods("GET")

	r.HandleFunc("/chosenProduct/", app.getAllChosenProducts).Methods("GET")
	r.HandleFunc("/chosenProduct/{id}", app.findCHosenProductByID).Methods("GET")
	r.HandleFunc("/chosenProduct/", app.insertChosenProduct).Methods("POST")
	r.HandleFunc("/chosenProduct/{id}", app.deleteChosenProduct).Methods("DELETE")


	r.HandleFunc("/api/purchase/{id}", app.getAllPurchases).Methods("GET")
	r.HandleFunc("/purchase/{id}", app.findPurchaseByID).Methods("GET")
	r.HandleFunc("/api/purchase", app.insertPurchase).Methods("POST")
	r.HandleFunc("/purchase/{id}", app.deletePurchase).Methods("DELETE")


	r.HandleFunc("/content/", app.getAllContents).Methods("GET")
	r.HandleFunc("/content/{id}", app.findContentByID).Methods("GET")
	r.HandleFunc("/content/", app.insertContent).Methods("POST")
	r.HandleFunc("/content/{id}", app.deleteContent).Methods("DELETE")
	return r
}

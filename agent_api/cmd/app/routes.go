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

	r.HandleFunc("/location/", app.getAllLocations).Methods("GET")
	r.HandleFunc("/location/{id}", app.findLocationByID).Methods("GET")
	r.HandleFunc("/location/", app.insertLocation).Methods("POST")
	r.HandleFunc("/location/{id}", app.deleteLocation).Methods("DELETE")

	r.HandleFunc("/product/", app.getAllProducts).Methods("GET")
	r.HandleFunc("/product/{id}", app.findProductByID).Methods("GET")
	r.HandleFunc("/product/", app.insertProduct).Methods("POST")
	r.HandleFunc("/product/{id}", app.deleteProduct).Methods("DELETE")

	r.HandleFunc("/chosenProduct/", app.getAllChosenProducts).Methods("GET")
	r.HandleFunc("/chosenProduct/{id}", app.findCHosenProductByID).Methods("GET")
	r.HandleFunc("/chosenProduct/", app.insertChosenProduct).Methods("POST")
	r.HandleFunc("/chosenProduct/{id}", app.deleteChosenProduct).Methods("DELETE")


	r.HandleFunc("/purchase/", app.getAllPurchases).Methods("GET")
	r.HandleFunc("/purchase/{id}", app.findPurchaseByID).Methods("GET")
	r.HandleFunc("/purchase/", app.insertPurchase).Methods("POST")
	r.HandleFunc("/purchase/{id}", app.deletePurchase).Methods("DELETE")


	r.HandleFunc("/content/", app.getAllContents).Methods("GET")
	r.HandleFunc("/content/{id}", app.findContentByID).Methods("GET")
	r.HandleFunc("/content/", app.insertContent).Methods("POST")
	r.HandleFunc("/content/{id}", app.deleteContent).Methods("DELETE")
	return r
}

package main

import (
	"AgentApp/pkg/dtos"
	"AgentApp/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func (app *application) getAllProducts(w http.ResponseWriter, r *http.Request) {
	// Get all bookings stored
	bookings, err := app.products.GetAll()
	if err != nil {
		app.serverError(w, err)
	}

	// Convert booking list into json encoding
	b, err := json.Marshal(bookings)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("products have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findProductByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find booking by id
	m, err := app.products.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("products not found")
			return
		}
		// Any other error will send an internal server error
		app.serverError(w, err)
	}

	// Convert booking to json encoding
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Println("Have been found a products")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}


func (app *application) saveImage(w http.ResponseWriter, r *http.Request)  {

	vars := mux.Vars(r)
	userId := vars["userIdd"]
	r.ParseMultipartForm(32 << 20)
	file, hander, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())

	}
	res1 := strings.HasPrefix(userId, "\"")
	if res1 == true {
		userId = userId[1:]
		userId = userId[:len(userId)-1]
	}

	defer file.Close()
	var path = "cmd/app/images/"+hander.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 777)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	io.Copy(f, file)



}


func (app *application) addImages(w http.ResponseWriter, r *http.Request)  {

	var m dtos.AddImagesDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}

	var pom = models.Product{}
	app.infoLog.Printf("",m)
	products, err := app.products.GetAll()
	for _, prod := range products {
		if(prod.Id==m.PostId){
			for _, med := range m.Media {
				prod.Media = append(prod.Media, med)
				pom = prod
				_, _ = app.products.Delete(prod.Id)
			}
		}
	}

	_, _ = app.products.Insert(pom)

	app.infoLog.Printf("Success")


}


func (app *application) insertProduct(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userId := vars["userId"]
	var m dtos.ProductDTO
	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	var post = models.Product{
		 User : userIdPrimitive,
		DateTime : time.Now(),
		Price: m.Price,
		Quantity: m.Quantity,
		Name : m.Name,
		Media: m.Media,
	}

	insertResult, err := app.products.Insert(post)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.InsertedID)
	w.Write(idMarshaled)
}


func (app *application) deleteProduct(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	userId := vars["id"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	// Delete booking by id
	deleteResult, err := app.products.Delete(userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d products(s)", deleteResult.DeletedCount)
}



func (app *application) deleteImage(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url

	var m dtos.DeleteImageDTO
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	products, err := app.products.GetAll()
	al := models.Product{}
	app.infoLog.Printf("Have bdddd", m.AlbumId, m.Image)
	for _, prod := range products {

		if(prod.Id == m.AlbumId){

			for _, media := range prod.Media {
				if len(prod.Media)==1{

					prod.Media = []string{}
				}else {
					if (media == m.Image) {
						prod.Media = append(prod.Media[:1], prod.Media[2:]...)

					}
				}

			}
			al = prod
			_, _ = app.products.Delete(prod.Id)
		}

	}

	_, err = app.products.Insert(al)




	// Delete booking by id

	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated products(s)")
}
func findFeedAlbumsByUserId(albums []models.Product, idPrimitive primitive.ObjectID) ([]models.Product, error){
	feedAlbumsUser := []models.Product{}

	for _, album := range albums {
		if	album.User.String()==idPrimitive.String() {
			feedAlbumsUser = append(feedAlbumsUser, album)
		}
	}
	return feedAlbumsUser, nil
}

func findRest(albums []models.Product, idPrimitive primitive.ObjectID) ([]models.Product, error){
	feedAlbumsUser := []models.Product{}

	for _, album := range albums {
		if	album.User.String()!=idPrimitive.String() {
			feedAlbumsUser = append(feedAlbumsUser, album)
		}
	}
	return feedAlbumsUser, nil
}

func findMyCart(albums []models.Cart, idPrimitive primitive.ObjectID) ([]models.Cart, error){
	feedAlbumsUser := []models.Cart{}

	for _, album := range albums {
		if	album.Buyer.String()==idPrimitive.String() {
			feedAlbumsUser = append(feedAlbumsUser, album)
		}
	}
	return feedAlbumsUser, nil
}

func findMyPurchase(albums []models.Purchase, idPrimitive primitive.ObjectID) ([]models.Purchase, error){
	feedAlbumsUser := []models.Purchase{}

	for _, album := range albums {
		if	album.Buyer.String()==idPrimitive.String() {
			feedAlbumsUser = append(feedAlbumsUser, album)
		}
	}
	return feedAlbumsUser, nil
}

func (app *application) getUsersFeedAlbums(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userIdd"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	allAlbums, _ := app.products.GetAll()
	usersFeedAlbums, err := findFeedAlbumsByUserId(allAlbums, userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}

	feedAlbumResponse := []dtos.ProductResponseDTO{}

	for _, album := range usersFeedAlbums {

		images := album.Media

		if err != nil {
			app.serverError(w, err)
		}

		feedAlbumResponse = append(feedAlbumResponse, toResponseAlbum(album, images))

	}

	imagesMarshaled, err := json.Marshal(feedAlbumResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}
func toResponseAlbum(feedAlbum models.Product, imageList []string) dtos.ProductResponseDTO {
	imagesBuffered := [][]byte{}




	if(len(imageList) == 0){

	}else {

		for _, image2 := range imageList {

			f, _ := os.Open("cmd/app/images/" + image2)

			defer f.Close()
			image, _, _ := image.Decode(f)
			buffer := new(bytes.Buffer)

			if err := jpeg.Encode(buffer, image, nil); err != nil {
				log.Println("unable to encode image.")
			}
			imageBuffered := buffer.Bytes()
			imagesBuffered = append(imagesBuffered, imageBuffered)
		}
	}

	return dtos.ProductResponseDTO{
		Id: feedAlbum.Id,
		DateTime : strings.Split(feedAlbum.DateTime.String(), " ")[0],
		Media : imagesBuffered,
		User  : feedAlbum.User,
		Price : feedAlbum.Price,
		Quantity   : feedAlbum.Quantity,
		Name    : feedAlbum.Name,
		MediaOrig: feedAlbum.Media,


	}
}




func (app *application) getPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userIdd"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	allAlbums, _ := app.products.GetAll()
	usersFeedAlbums, err := findRest(allAlbums, userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}
	feedAlbumResponse := []dtos.ProductResponseDTO{}

	for _, album := range usersFeedAlbums {

		images := album.Media

		if err != nil {
			app.serverError(w, err)
		}

		feedAlbumResponse = append(feedAlbumResponse, toResponseAlbum(album, images))

	}

	imagesMarshaled, err := json.Marshal(feedAlbumResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}












func (app *application) edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userIdd"]
	m := []dtos.ProductResponseDTO{}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)
	log.Println("Obrisano...", m)

	allAlbums, _ := app.products.GetAll()
	//usersFeedAlbums, err := findFeedAlbumsByUserId(allAlbums, userIdPrimitive)

	for _, album := range allAlbums {
		if	album.User.String()==userIdPrimitive.String() {
		app.products.Delete(album.Id)
		}
	}

	for _, a := range m {
		var post = models.Product{
			User : a.User,
			DateTime : time.Now(),
			Price: a.Price,
			Quantity: a.Quantity,
			Name : a.Name,
			Media: a.MediaOrig,
		}

		_, err := app.products.Insert(post)
		if err != nil {
			return
		}
	}

}



func (app *application) addToCart(w http.ResponseWriter, req *http.Request) {

	var m dtos.CartDTO
	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	var item = models.Cart{
		Buyer : m.User,
		ChosenProducts : m.Product,
		Quantity: m.Quantity,
	}

	insertResult, err := app.cart.Insert(item)
	if err != nil {
		app.serverError(w, err)
	}
	allAlbums, _ := app.products.GetAll()
	var pom = models.Product{}
	for _, prod := range allAlbums {


		if prod.Id == m.Product{
			pom = prod
			help , _ := strconv.Atoi(prod.Quantity)
			h2 , _ := strconv.Atoi(m.Quantity)
			h3 := help - h2
			pom.Quantity = strconv.Itoa(h3)
			_, _ = app.products.Delete(prod.Id)
		}

	}

	_, _ = app.products.Insert(pom)


	app.infoLog.Printf("New content have been created, id=%s", insertResult.InsertedID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	idMarshaled, err := json.Marshal(insertResult.InsertedID)
	w.Write(idMarshaled)
}

func (app *application) getCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userIdd"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	allAlbums, _ := app.cart.GetAll()
	allProducts, _ := app.products.GetAll()
	usersFeedAlbums, err := findMyCart(allAlbums, userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}

	feedAlbumResponse := []dtos.CartFrontDTO{}

	for _, album := range usersFeedAlbums {
		for _, product := range allProducts {
			if(album.ChosenProducts == product.Id){

				images:= product.Media
				feedAlbumResponse = append(feedAlbumResponse, toResponseCart(album, images, product))

			}
		}


		if err != nil {
			app.serverError(w, err)
		}



	}

	imagesMarshaled, err := json.Marshal(feedAlbumResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}



func toResponseCart(feedAlbum models.Cart, imageList []string, product models.Product) dtos.CartFrontDTO {
	imagesBuffered := [][]byte{}


	for _, image2 := range imageList {

		f, _ := os.Open("cmd/app/images/" + image2)

		defer f.Close()
		image, _, _ := image.Decode(f)
		buffer := new(bytes.Buffer)

		if err := jpeg.Encode(buffer, image, nil); err != nil {
			log.Println("unable to encode image.")
		}
		imageBuffered :=buffer.Bytes()
		imagesBuffered= append(imagesBuffered, imageBuffered)
	}

	productt := dtos.ProductResponseDTO{

		Price : product.Price,
		Name    : product.Name,
		MediaOrig: product.Media,


	}
	return dtos.CartFrontDTO{
		Id: feedAlbum.Id,
		Media : imagesBuffered,
		User  : feedAlbum.Buyer,
		Quantity : feedAlbum.Quantity,
		Product    : productt,


	}
}

func (app *application) deleteCart(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	userId := vars["id"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	// Delete booking by id
	deleteResult, err := app.cart.Delete(userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated %d products(s)", deleteResult.DeletedCount)
}


func (app *application) removeCart(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	userId := vars["id"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	// Delete booking by id
	// Delete booking by id
	carts, err := app.cart.GetAll()

	for _, cart := range carts {

		if(cart.Buyer == userIdPrimitive){
			_, _ = app.cart.Delete(cart.Id)
		}
	}


	if err != nil {
		app.serverError(w, err)
	}

	app.infoLog.Printf("Have been eliminated products(s)")
}


func (app *application) getOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userIdd"]
	userIdPrimitive, _ := primitive.ObjectIDFromHex(userId)

	purch, _ := app.purchases.GetAll()
	usersFeedAlbums, err := findMyPurchase(purch, userIdPrimitive)
	if err != nil {
		app.serverError(w, err)
	}

	feedAlbumResponse := []dtos.PurchaseResponseDTO{}

	for _, album := range usersFeedAlbums {

		log.Println("mjhjbkgkjhgk.")
			feedAlbumResponse = append(feedAlbumResponse, toResponseOrder(album))



	}
	log.Println("blbbalbla" , feedAlbumResponse)
	imagesMarshaled, err := json.Marshal(feedAlbumResponse)
	if err != nil {
		app.serverError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(imagesMarshaled)
}



func toResponseOrder(purchase models.Purchase) dtos.PurchaseResponseDTO {
	imagesBuffered := [][]byte{}
	response := []dtos.PurchDTO{}

		for _, prod := range purchase.Products {
			images := prod.Product.MediaOrig

			for _, image2 := range images {

				f, _ := os.Open("cmd/app/images/" + image2)

				defer f.Close()
				image, _, _ := image.Decode(f)
				buffer := new(bytes.Buffer)

				if err := jpeg.Encode(buffer, image, nil); err != nil {
					log.Println("unable to encode image.")
				}
				imageBuffered :=buffer.Bytes()
				imagesBuffered= append(imagesBuffered, imageBuffered)
			}

			r := dtos.PurchDTO{Quantity: prod.Quantity, Price: prod.Product.Price, Name: prod.Product.Name, Media: imagesBuffered, MediaOrig: images,
			}

			response = append(response, r)

		}


		return dtos.PurchaseResponseDTO{Product: response, User: purchase.Buyer, Id: purchase.Id, Location: purchase.Location}


}
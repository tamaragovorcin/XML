package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"users/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"users/pkg/models/mongodb"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	users    *mongodb.UserModel
	roles    *mongodb.RoleModel
	verifications *mongodb.VerificationModel
	profileInformation *mongodb.ProfileInformationModal
	agents *mongodb.AgentModel
}

func main() {

	fmt.Printf("Found multiple documents (array of pointers): %+v\n")

	// Define command-line flags
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4000, "HTTP server network port")
	mongoURI := flag.String("mongoURI", "mongodb://localhost:27017", "Database hostname url")
	mongoDatabse := flag.String("mongoDatabse", "users", "Database name")
	enableCredentials := flag.Bool("enableCredentials", false, "Enable the use of credentials for mongo connection")
	flag.Parse()

	// Create logger for writing information and error messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Create mongo client configuration
	co := options.Client().ApplyURI(*mongoURI)
	if *enableCredentials {
		co.Auth = &options.Credential{
			Username: os.Getenv("MONGODB_USERNAME"),
			Password: os.Getenv("MONGODB_PASSWORD"),
		}
	}


	// Establish database connection
	client, err := mongo.NewClient(co)
	if err != nil {
		errLog.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		errLog.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	infoLog.Printf("Database connection established")

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		infoLog:  infoLog,
		errorLog: errLog,
		users: &mongodb.UserModel{
			C: client.Database(*mongoDatabse).Collection("users"),
		},
		roles: &mongodb.RoleModel{
			C: client.Database(*mongoDatabse).Collection("roles"),
		},
		verifications: &mongodb.VerificationModel{
			C: client.Database(*mongoDatabse).Collection("verifications"),
		},

		profileInformation: &mongodb.ProfileInformationModal{
			C: client.Database(*mongoDatabse).Collection("profileInformation"),
		},
	}

	// Initialize a new http.Server struct.
	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", serverURI)
	err = srv.ListenAndServe()
	errLog.Fatal(err)

	collection := client.Database("test").Collection("roles")

	// Some dummy data to add to the Database
	admin := models.Role{"1", "ADMIN"}
	user := models.Role{"2", "USER"}

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), admin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Insert multiple documents
	trainers := []interface{}{admin, user}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)


}

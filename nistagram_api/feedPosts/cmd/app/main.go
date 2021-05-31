package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"feedPosts/pkg/models/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	comments    *mongodb.CommentModel
	feedPosts   *mongodb.FeedPostModel
	posts       *mongodb.PostsModel
	locations   *mongodb.LocationModel
	albumFeeds   *mongodb.AlbumFeedModel
	collections   *mongodb.CollectionModel
	images   *mongodb.ImageModel
}

func main() {

	fmt.Printf("Found multiple documents (array of pointers): %+v\n")

	// Define command-line flags
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4001, "HTTP server network port")
	mongoURI := flag.String("mongoURI", "mongodb://localhost:27017", "Database hostname url")
	mongoDatabse := flag.String("mongoDatabse", "feedPosts", "Database name")
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
		comments: &mongodb.CommentModel{
			C: client.Database(*mongoDatabse).Collection("comments"),
		},
		locations: &mongodb.LocationModel{
			C: client.Database(*mongoDatabse).Collection("locations"),
		},
		posts: &mongodb.PostsModel{
			C: client.Database(*mongoDatabse).Collection("posts"),
		},
		feedPosts: &mongodb.FeedPostModel{
			C: client.Database(*mongoDatabse).Collection("feedPosts"),
		},
		albumFeeds: &mongodb.AlbumFeedModel{
			C: client.Database(*mongoDatabse).Collection("albumFeed"),
		},
		collections: &mongodb.CollectionModel{
			C: client.Database(*mongoDatabse).Collection("collections"),
		},
		images: &mongodb.ImageModel{
			C: client.Database(*mongoDatabse).Collection("images"),
		},
	}

	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	router := app.routes();
	http.ListenAndServe(serverURI, setHeaders(router))

}
func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//anyone can make a CORS request (not recommended in production)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//only allow GET, POST, and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, FETCH")
		//Since I was building a REST API that returned JSON, I set the content type to JSON here.
		w.Header().Set("Content-Type", "application/json")
		//Allow requests to have the following headers
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, cache-control")
		//if it's just an OPTIONS request, nothing other than the headers in the response is needed.
		//This is essential because you don't need to handle the OPTIONS requests in your handlers now
		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

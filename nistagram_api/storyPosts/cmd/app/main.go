package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"storyPosts/pkg/models/mongodb"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	storyPosts *mongodb.StoryPostModel
	highlights *mongodb.HighlightModel
	albumStories   *mongodb.AlbumStoryModel
	posts       *mongodb.PostsModel
	locations   *mongodb.LocationModel
	images   *mongodb.ImageModel
	highlightsAlbum   *mongodb.HighlightAlbumModel


}

func main() {
	fmt.Printf("Found multiple documents (array of pointers): %+v\n")

	// Define command-line flags
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4004, "HTTP server network port")
	//mongoURI := flag.String("mongoURI", "mongodb://localhost:27017", "Database hostname url")
	mongoURI := flag.String("mongoURI", "mongodb://db_users:27017", "Database hostname url")

	mongoDatabse := flag.String("mongoDatabse", "storyPosts", "Database name")
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
		locations: &mongodb.LocationModel{
			C: client.Database(*mongoDatabse).Collection("locations"),
		},
		posts: &mongodb.PostsModel{
			C: client.Database(*mongoDatabse).Collection("posts"),
		},
		storyPosts: &mongodb.StoryPostModel{
			C: client.Database(*mongoDatabse).Collection("storyPosts"),
		},
		highlights: &mongodb.HighlightModel{
			C: client.Database(*mongoDatabse).Collection("highlights"),
		},
		albumStories: &mongodb.AlbumStoryModel{
			C: client.Database(*mongoDatabse).Collection("albumStories"),
		},
		images: &mongodb.ImageModel{
			C: client.Database(*mongoDatabse).Collection("images"),
		},
		highlightsAlbum: &mongodb.HighlightAlbumModel{
			C: client.Database(*mongoDatabse).Collection("highlightsAlbum"),
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
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, FETCH,DELETE")
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
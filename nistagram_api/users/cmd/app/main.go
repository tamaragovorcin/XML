package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"users/pkg/models"
	"users/pkg/models/mongodb"
	"users/saga"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	users    *mongodb.UserModel
	roles    *mongodb.RoleModel
	verifications *mongodb.VerificationModel
	profileInformation *mongodb.ProfileInformationModal
	agents *mongodb.AgentModel
	profileImage *mongodb.ProfileImageModel
	notification   *mongodb.NotificationModel
	settings  *mongodb.SettingsModel
	notificationForUser *mongodb.NotificationForUserModel
	notificationContent *mongodb.NotificationContentModel
	images   *mongodb.ImageModel
	orchestrator *saga.Orchestrator
	tracer opentracing.Tracer

}
func main() {

	fmt.Printf("Found multiple documents (array of pointers): %+v\n")

	// Define command-line flags
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4006, "HTTP server network port")
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
		profileImage:  &mongodb.ProfileImageModel{
			C: client.Database(*mongoDatabse).Collection("profileImages"),
		},
		notification: &mongodb.NotificationModel{
			C: client.Database(*mongoDatabse).Collection("notifications"),
		},
		settings: &mongodb.SettingsModel{
			C: client.Database(*mongoDatabse).Collection("settings"),
		},
		notificationForUser: &mongodb.NotificationForUserModel{
			C: client.Database(*mongoDatabse).Collection("notificationForUser"),
		},
		notificationContent: &mongodb.NotificationContentModel{
			C: client.Database(*mongoDatabse).Collection("notificationContent"),
		},
		agents: &mongodb.AgentModel{
			C: client.Database(*mongoDatabse).Collection("agents"),
		},
		images: &mongodb.ImageModel{
			C: client.Database(*mongoDatabse).Collection("images"),
		},





	}


	go saga.NewOrchestrator().Start()
	go app.RedisConnection()
	// Initialize a new http.Server struct.
	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)

	router := app.routes()
	http.ListenAndServe(serverURI, setHeaders(router))

}


func (app *application) RedisConnection() {
	// create client and ping redis
	var err error
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	if _, err = client.Ping().Result(); err != nil {
		log.Fatalf("error creating redis client %s", err)
	}

	// subscribe to the required channels
	pubsub := client.Subscribe(saga.UserChannel, saga.ReplyChannel)
	if _, err = pubsub.Receive(); err != nil {
		log.Fatalf("error subscribing %s", err)
	}
	defer func() { _ = pubsub.Close() }()
	ch := pubsub.Channel()

	log.Println("starting the order service")
	for {
		select {
		case msg := <-ch:
			m := saga.Message{}
			err := json.Unmarshal([]byte(msg.Payload), &m)
			if err != nil {
				log.Println(err)
				continue
			}

			switch msg.Channel {
			case saga.UserChannel:

				// Happy Flow
				if m.Action == saga.ActionStart {



					if m.SenderService == saga.ServiceInteraction {
						if m.Ok {
							user := m.User2
							iddd := strings.Split(user, "\"")
							id, _ := primitive.ObjectIDFromHex(iddd[1])
							res, err := app.users.UpdateS(id, models.FINISHED)
							log.Println(res)
							if err != nil {
								return
							}

							log.Println("FINISHED")

						}} else {
						user := m.User2
						iddd := strings.Split(user, "\"")
						id, _ := primitive.ObjectIDFromHex(iddd[1])
						res, err := app.users.UpdateS(id, models.CANCELLED)
						log.Println(res)
						if err != nil {
							return
						}

						log.Println("CANCEL")
					}


				}

				if m.Action == saga.ActionRollback {
					user := m.User2
					iddd := strings.Split(user, "\"")

					id, _ := primitive.ObjectIDFromHex(iddd[1])
					log.Println("lala", id)
					res, err := app.users.UpdateS(id, models.CANCELLED)
					log.Println(res)
					if err != nil {
						return
					}

					u,_ := app.users.FindByID(id)

					log.Println("CANCELLED %d", u.Status)


				/*	w := http.ResponseWriter


					w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(http.StatusOK)
						idMarshaled, _ := json.Marshal("ubacili smo")

						w.Write(idMarshaled)*/

				}
			}
		}
	}
}



func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//anyone can make a CORS request (not recommended in production)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//only allow GET, POST, and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, DELETE")
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

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)


func routes() *mux.Router {
	// Register handler functions.

	configuration := parseConfiguration()
	driver, err := configuration.newDriver()
	if err != nil {
		log.Fatal(err)
	}
	defer unsafeClose(driver)
	r := mux.NewRouter()
	r.HandleFunc("/api/followRequest", CreateFollow(driver, configuration.Database)).Methods("POST")
	r.HandleFunc("/movie/vote/{id}", voteInMovieHandlerFunc(driver, configuration.Database)).Methods("GET")


	return r
}
func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//anyone can make a CORS request (not recommended in production)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//only allow GET, POST, and OPTIONS
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
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
func main(){
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4006, "HTTP server network port")

	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	router := routes()
	http.ListenAndServe(serverURI, setHeaders(router))

}

func parseConfiguration() *Neo4jConfiguration {
	database := lookupEnvOrGetDefault("NEO4J_DATABASE", "no-waiter-userInteraction")
	if !strings.HasPrefix(lookupEnvOrGetDefault("NEO4J_VERSION", "4"), "4") {
		database = ""
	}
	return &Neo4jConfiguration{
		Url:      lookupEnvOrGetDefault("NEO4J_URI", "bolt://localhost:7687"),
		Username: lookupEnvOrGetDefault("NEO4J_USER", "neo4j"),
		Password: lookupEnvOrGetDefault("NEO4J_PASSWORD", "neo4jdb"),
		Database: database,
	}
}

func lookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}

func unsafeClose(closeable io.Closer) {
	if err := closeable.Close(); err != nil {
		log.Fatal(fmt.Errorf("could not close resource: %w", err))
	}
}
type FollowRequest struct {
	Id uuid.UUID  `json:"_id,omitempty"`
	Following  uuid.UUID `json:"following,omitempty"`
	Follower   uuid.UUID `json:"follower,omitempty"`
	Approved  bool `json:"approved,omitempty"`
	DateTime time.Time `json:"dateTime,omitempty"`
}

type Report struct {
	Id uuid.UUID `json:"_id,omitempty"`
	ComplainingUser uuid.UUID `json:"complainingUser,omitempty"`
	ReportedUser uuid.UUID `json:"reportedUser,omitempty"`
	FeedPost uuid.UUID `json:"feedPost,omitempty"`
	StoryPost uuid.UUID `json:"storyPost,omitempty"`
}

type VoteResult struct {
	Updates int `json:"updates"`
}

type Neo4jConfiguration struct {
	Url      string
	Username string
	Password string
	Database string
}

func (nc *Neo4jConfiguration) newDriver() (neo4j.Driver, error) {
	return neo4j.NewDriver(nc.Url, neo4j.BasicAuth(nc.Username, nc.Password, ""))
}



func CreateFollow(driver neo4j.Driver, database string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("POGODIIOOOO JE")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var m FollowRequest
		err := json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			fmt.Println("Error")
		}
		session := driver.NewSession(neo4j.SessionConfig{
			AccessMode:   neo4j.AccessModeWrite,
			DatabaseName: database,
		})
		defer unsafeClose(session)

		voteResult, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			result, err := tx.Run(
				"MATCH (follower:User), (following:User) WHERE follower.id = $followerId AND following.id = $followingId CREATE (follower)-[:FOLLOW]->(following)",
				map[string]interface{}{"follower": m.Follower,
					"following": m.Following,
					})
			if err != nil {
				return nil, err
			}
			var summary, _ = result.Consume()
			var voteResult VoteResult
			voteResult.Updates = summary.Counters().PropertiesSet()

			return voteResult, nil
		})
		if err != nil {
			log.Println("error voting for movie:", err)
			return
		}
		err = json.NewEncoder(w).Encode(voteResult)
		if err != nil {
			log.Println("error writing volte result response:", err)
		}
	}
}

func voteInMovieHandlerFunc(driver neo4j.Driver, database string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("POGOOOODI")
		w.Header().Set("Content-Type", "application/json")
		title, _ := url.QueryUnescape(req.URL.Path[len("/movie/vote/"):])

		session := driver.NewSession(neo4j.SessionConfig{
			AccessMode:   neo4j.AccessModeWrite,
			DatabaseName: database,
		})
		defer unsafeClose(session)

		voteResult, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			result, err := tx.Run(
				`MATCH (m:Movie {title: $title}) 
				WITH m, (CASE WHEN exists(m.votes) THEN m.votes ELSE 0 END) AS currentVotes
				SET m.votes = currentVotes + 1;`,
				map[string]interface{}{"title": title})
			if err != nil {
				return nil, err
			}
			var summary, _ = result.Consume()
			var voteResult VoteResult
			voteResult.Updates = summary.Counters().PropertiesSet()

			return voteResult, nil
		})
		if err != nil {
			log.Println("error voting for movie:", err)
			return
		}
		err = json.NewEncoder(w).Encode(voteResult)
		if err != nil {
			log.Println("error writing volte result response:", err)
		}
	}
}

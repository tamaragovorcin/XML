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
	r := mux.NewRouter()

	r.HandleFunc("/api/followRequest", CreateFollowRequest(driver, configuration.Database)).Methods("POST")
	r.HandleFunc("/api/followApproved", AcceptFollowRequest(driver, configuration.Database)).Methods("POST")
	r.HandleFunc("/api/createUser", CreateUser(driver, configuration.Database)).Methods("POST")
	r.HandleFunc("/api/allFollowRequest", ReturnFollowRequests(driver, configuration.Database)).Methods("POST")

	r.HandleFunc("/api/followRequest", CreateFollow(driver, configuration.Database)).Methods("POST")
	r.HandleFunc("/api/createUser", CreateUser(driver, configuration.Database)).Methods("POST")

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
	configuration := parseConfiguration()
	driver, err := configuration.newDriver()
	if err != nil {
		log.Fatal(err)
	}

	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 4005, "HTTP server network port")

	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	router := routes()
	http.ListenAndServe(serverURI, setHeaders(router))

	defer unsafeClose(driver)

}

func parseConfiguration() *Neo4jConfiguration {


	if !strings.HasPrefix(lookupEnvOrGetDefault("NEO4J_VERSION", "4"), "4") {
		//database = ""
	}

	return &Neo4jConfiguration{
		Url:      lookupEnvOrGetDefault("NEO4J_URI", "bolt://localhost:7687"),
		Username: lookupEnvOrGetDefault("NEO4J_USER", "neo4j"),
		Password: lookupEnvOrGetDefault("NEO4J_PASSWORD", "root"),

		//Database: database,

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
type Follow struct {
	Id uuid.UUID  `json:"_id,omitempty"`
	Following  User`json:"following,omitempty"`
	Follower   User`json:"follower,omitempty"`
	Approved  bool `json:"approved,omitempty"`
	DateTime time.Time `json:"dateTime,omitempty"`
}
type FollowRequest struct {
	Follower string `json:"follower"`
	Following string `json:"following"`
}

type FollowRequestDTO struct {
	Follower string `json:"follower"`
	Following string `json:"following"`
}
type FollowDTO struct {
	FollowerId string `json:"FollowerId"`
	FollowingId string `json:"FollowingId"`
}

type User struct {
	Id string `json:"id"`
}

type Users struct {
	Users []string `json:"users"`
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



func CreateFollowRequest(driver neo4j.Driver, database string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("POGODIIOOOO JE")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var m FollowRequestDTO

		fmt.Println(m.Follower)
		fmt.Println(m.Follower)


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
				"MATCH (follower:User), (following:User) WHERE follower.id = $followerId AND following.id = $followingId CREATE (follower)-[:FOLLOWREQUEST]->(following)",

				map[string]interface{}{"followerId": m.Follower,
					"followingId": m.Following,
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
func AcceptFollowRequest(driver neo4j.Driver, database string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("POGODIIOOOO JE")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var m FollowRequestDTO
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
				"MATCH (following:User)<-[f:FOLLOWREQUEST]-(follower:User) WHERE following.id = followingId AND follower.id = followerId DELETE f",

				map[string]interface{}{"followerId": m.Follower,
					"followingId": m.Following,
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

func CreateFollow(driver neo4j.Driver, database string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("POGODIIOOOO JE")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var m FollowDTO
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

				map[string]interface{}{"followerId": m.FollowerId,
					"followingId": m.FollowingId,
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
func ReturnFollowRequests(driver neo4j.Driver, database string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("POGODIIOOOO JE")
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		var m User
		err := json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			fmt.Println("Error")
		}
		session := driver.NewSession(neo4j.SessionConfig{
			AccessMode:   neo4j.AccessModeWrite,
			DatabaseName: database,
		})


		defer unsafeClose(session)
		result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			query := "MATCH (following:User)<-[f:FOLLOWREQUEST]-(follower:User) WHERE following.id = $followingId return follower.id as id"
			parameters := map[string]interface{}{
				"followingId": m.Id,
			}
			records, err := tx.Run(query, parameters)
			if err != nil {
				return nil, err
			}
			users := Users{}
			for records.Next() {
				record := records.Record()
				id, _ := record.Get("id")
				users.Users = append(users.Users, id.(string))
			}
			return users, nil
		})
		if err != nil {
			log.Println("error querying graph:", err)
			return
		}
		log.Println(result)

	}
}


func CreateUser(driver neo4j.Driver, database string) func(w http.ResponseWriter,r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("POGODIIOOOO")
		var m User
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
				"CREATE (:User{id:$uId})",
				map[string]interface{}{
					"uId": m.Id,
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
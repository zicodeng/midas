package main
import (
	"os"
	"fmt"
	"net/http"	
	"log"
	"midas/server/handlers"
	"github.com/go-redis/redis"
	"time"	
	"midas/server/sessions"
	"midas/server/users"
	mgo "gopkg.in/mgo.v2"
	
)
//main is the main entry point for the server
func main() {
	sessKey := os.Getenv("SESSIONKEY")
	redisaddr := os.Getenv("REDISADDR")
	if len(redisaddr) == 0 {
		redisaddr = "127.0.0.1:6379"
	}

	mongoDBaddr := os.Getenv("DBADDR")
	if len(mongoDBaddr) == 0 {
		mongoDBaddr = "127.0.0.1:27017"
	}	
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}
	println(mongoDBaddr)
	sess, err := mgo.Dial(mongoDBaddr)
	if err != nil {
		log.Fatal("problem connecting to the database")
	}	
	userStore := users.NewMongoStore(sess, "demo", "users")
	client := redis.NewClient(&redis.Options{
		Addr: redisaddr,
	})

	sessionStore := sessions.NewRedisStore(client, time.Hour)

	ctx := handlers.NewHandlerContext(sessionStore,userStore, sessKey)

	tlskey := os.Getenv("TLSKEY")
	tlscert := os.Getenv("TLSCERT")
	if len(tlskey) == 0 || len(tlscert) == 0 {
		log.Fatal("please set TLSKEY and TLSCERT")
	}

	mux := http.NewServeMux()
	// /v1/users: UsersHandler
	// /v1/users/me: UsersMeHandler
	// /v1/sessions: SessionsHandler
	// /v1/sessions/mine: SessionsMineHandler
	mux.HandleFunc("/v1/test", ctx.TestHandler)
	
	mux.HandleFunc("/v1/users", ctx.UsersHandler)
	mux.HandleFunc("/v1/users/me", ctx.UsersMeHandler)
	mux.HandleFunc("/v1/sessions", ctx.SessionsHandler)
	mux.HandleFunc("/v1/sessions/mine", ctx.SessionsMineHandler)
	
	
	
	
	fmt.Printf("server is listening at %s \n", addr)
	corsHandler := handlers.NewCORSHandler(mux)
	
	
	/* TODO: add code to do the following
	- Read the ADDR environment variable to get the address
	  the server should listen on. If empty, default to ":80"
	- Create a new mux for the web server.
	- Tell the mux to call your handlers.SummaryHandler function
	  when the "/v1/summary" URL path is requested.
	- Start a web server listening on the address you read from
	  the environment variable, using the mux you created as
	  the root handler. Use log.Fatal() to report any errors
	  that occur when trying to start the web server.
	*/
	log.Fatal(http.ListenAndServeTLS(addr, tlscert, tlskey, corsHandler))   
	
}
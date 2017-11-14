package handlers
import (
	"time"
	"midas/server/users"
)

//TODO: define a session state struct for this web server
//see the assignment description for the fields you should include
//remember that other packages can only see exported fields!

type SessionState struct {
	Time time.Time
	User users.User
}
// The sessions package you implemented in the last assignment saves any 
// sort of sessions state the caller wants to use. It's now time to define 
// the session state struct for our web server.

// In the /servers/gateway/handlers/sessionstate.go file, define a struct 
// with fields capable of tracking the following information:

// The time at which this session began
// The authenticated users.User who started the session
// That's all we need for now, but you can add other fields to this struct 
// in the future to capture and track other data about the session.
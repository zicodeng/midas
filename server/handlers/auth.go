package handlers
import (
	"time"
	"net/http"
	"encoding/json"
	"fmt"
	"midas/server/users"	
	"midas/server/sessions"	
)

const headerContentType = "Content-Type"

const contentTypeJSON = "application/json"
func (ctx *Context) TestHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "POST":	
		nu := &users.NewUser{}
		if err := json.NewDecoder(r.Body).Decode(nu); err != nil{
			http.Error(w, fmt.Sprintf("error encoding request body: %v", err), http.StatusBadRequest)	
			return		
		}
		w.WriteHeader(http.StatusCreated)
		respond(w, nu)
	default:
		http.Error(w, "method must be POST", http.StatusMethodNotAllowed)
		return
	}
}
	
func (ctx *Context) UsersHandler(w http.ResponseWriter, r *http.Request){
	
	switch r.Method {
	case "POST":
		nu := &users.NewUser{}
		if err := json.NewDecoder(r.Body).Decode(nu); err != nil{
			http.Error(w, fmt.Sprintf("error encoding request body: %v", err), http.StatusBadRequest)	
			return		
		}
		if err := nu.Validate(); err != nil{
			http.Error(w, fmt.Sprintf("error invalid new user: %v", err), http.StatusUnprocessableEntity)	
			return
		} 
		if user, err := ctx.usersStore.GetByEmail(nu.Email); err != users.ErrUserNotFound && user != nil{
			println(user.Email)
			http.Error(w, fmt.Sprintf("error user email already in use: %v", err), http.StatusUnprocessableEntity)	
			return
		}
		

		user, err := ctx.usersStore.Insert(nu) 
		if err != nil {
			http.Error(w, fmt.Sprintf("error adding user: %v", err), http.StatusInternalServerError)
			return
		}

		sessionState := &SessionState{
			Time: time.Now(),
			User: *user,
		}
		if _, err = sessions.BeginSession(ctx.signingKey, ctx.sessionsStore, sessionState, w); err != nil {
			http.Error(w, fmt.Sprintf("error begining session: %v", err), http.StatusInternalServerError)
			return
		} 
		w.WriteHeader(http.StatusCreated)
		respond(w, user)
	default:
		http.Error(w, "method must be POST", http.StatusMethodNotAllowed)
		return
	}
}

func (ctx *Context) UsersMeHandler(w http.ResponseWriter, r *http.Request){

	sessionState := &SessionState{}
	sessionID, err := sessions.GetState(r, ctx.signingKey, ctx.sessionsStore, sessionState)
	if err != nil {
		http.Error(w, fmt.Sprintf("error authenticating: %v", err), http.StatusUnauthorized)	
		return
	}
	switch r.Method {
	case "GET":		
		respond(w, sessionState.User)
	case "PATCH":
		updates := &users.Updates{}
		if err := json.NewDecoder(r.Body).Decode(updates); err != nil{
			http.Error(w, fmt.Sprintf("error encoding request body: %v", err), http.StatusBadRequest)	
			return		
		}
		if err := ctx.usersStore.Update(sessionState.User.ID, updates); err != nil {
			http.Error(w, fmt.Sprintf("error updating user: %v", err), http.StatusBadRequest)	
			return	
		}
		sessionState.User.FirstName = updates.FirstName
		sessionState.User.LastName = updates.LastName		
		if err = ctx.sessionsStore.Save(sessionID, sessionState); err != nil {
			http.Error(w, fmt.Sprintf("error saving session data: %v", err), http.StatusInternalServerError)
			return
		}
		respond(w, sessionState.User)		
	default:
		http.Error(w, "method must be GET or PATCH", http.StatusMethodNotAllowed)
		return
	}
}
func (ctx *Context) SessionsHandler(w http.ResponseWriter, r *http.Request){

	switch r.Method {
	case "POST":
		creds := &users.Credentials{}
		if err := json.NewDecoder(r.Body).Decode(creds); err != nil{
			http.Error(w, fmt.Sprintf("error encoding request body: %v", err), http.StatusBadRequest)	
			return		
		}
		user, err := ctx.usersStore.GetByEmail(creds.Email)
		
		if err != nil {
			
			http.Error(w, fmt.Sprintf("invalid credentials: %v", err), http.StatusBadRequest)	
			return	
		}
		if err = user.Authenticate(creds.Password); err != nil {
			http.Error(w, fmt.Sprintf("invalid credentials: %v", err), http.StatusBadRequest)	
			return	
		}
		sessionState := &SessionState{
			Time: time.Now(),
			User: *user,
		}
		if _, err = sessions.BeginSession(ctx.signingKey, ctx.sessionsStore, sessionState, w); err != nil {
			http.Error(w, fmt.Sprintf("error begining session: %v", err), http.StatusInternalServerError)
			return
		} 
		
		respond(w, sessionState.User)		
	default:
		http.Error(w, "method must be POST", http.StatusMethodNotAllowed)
		return
	}
}

func (ctx *Context) SessionsMineHandler(w http.ResponseWriter, r *http.Request){
	// This function handles requests for the "current session" 
	// resource, and allows clients to end that session. The HTTP 
	// method must be DELETE. For any other HTTP method, respond 
	// with an http.StatusMethodNotAllowed error. If there is an 
	// error getting the session state, respond with an http.StatusUnauthorized error.
	
	// If the method is DELETE, follow these steps:
	
	// End the current session
	// Respond with the string "signed out"
	switch r.Method {
	case "DELETE":
		if _, err := sessions.EndSession(r, ctx.signingKey, ctx.sessionsStore); err != nil {
			http.Error(w, fmt.Sprintf("error authenticating: %v", err), http.StatusUnauthorized)	
			return
		}	
		w.Write([]byte("signed out"))
	default:
		http.Error(w, "method must be DELETE", http.StatusMethodNotAllowed)
		return
	}
}

func respond(w http.ResponseWriter, value interface{}) {
	w.Header().Add(headerContentType, contentTypeJSON)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		http.Error(w, fmt.Sprintf("error encoding response value to JSON: %v", err), http.StatusInternalServerError)
	}
}
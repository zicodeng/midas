package handlers
import (
	"midas/server/users"
	"midas/server/sessions"	
)

type Context struct {
	sessionsStore sessions.Store
	usersStore users.Store
	signingKey string
}

func NewHandlerContext(sessionsStore sessions.Store, usersStore users.Store, signingKey string) *Context{
	return &Context{
		sessionsStore: sessionsStore,
		usersStore: usersStore,
		signingKey: signingKey,
	}
}

func (ctx *Context) GetUserStore()users.Store{
	return ctx.usersStore
}

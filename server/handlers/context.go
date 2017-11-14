package handlers
import (
	"github.com/info344-a17/challenges-cjjaeger/servers/gateway/models/users"
	"github.com/info344-a17/challenges-cjjaeger/servers/gateway/sessions"		
	"github.com/info344-a17/challenges-cjjaeger/servers/gateway/indexes"
	
)

type Context struct {
	sessionsStore sessions.Store
	usersStore users.Store
	signingKey string
	trie	indexes.Trie
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

func (ctx *Context) SetTrie(trie indexes.Trie){
	ctx.trie = trie	
}

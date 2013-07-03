package blogfe

import (
	"net/http"

	"appengine"
	"appengine/user"
)

type AppEngineAuth struct{}

func NewAuthenticator() *AppEngineAuth {
	return new(AppEngineAuth)
}

func (a *AppEngineAuth) CheckAuth(r *http.Request) bool {
	return user.IsAdmin(appengine.NewContext(r))
}

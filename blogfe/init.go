package blogfe

import (
	"net/http"

	"appengine"
	"appengine/user"

	blog "github.com/arunjitsingh/goappengine/blog"
	jsonrpc "github.com/arunjitsingh/rpc/v2/json2"
	rpc "github.com/gorilla/rpc/v2"
)

const (
	allowedCORSOrigins = "*.ajsd.in,*.arunjit-test.appspot.com"
	allowedCORSHeaders = "X-Requested-With,Content-Type"
	devCORSOrigins     = "*"
)

type AppEngineAuth struct{}

func (a *AppEngineAuth) CheckAuth(r *http.Request) bool {
	return user.IsAdmin(appengine.NewContext(r))
}

type server struct {
	h http.Handler
}

func newServer(h http.Handler) *server {
	return &server{h}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	corsOrigins := allowedCORSOrigins
	if appengine.IsDevAppServer() {
		corsOrigins = devCORSOrigins
	}
	w.Header().Set("Access-Control-Allow-Origin", corsOrigins)
	w.Header().Set("Access-Control-Allow-Headers", allowedCORSHeaders)
	s.h.ServeHTTP(w, r)
}

func init() {
	s := rpc.NewServer()
	s.RegisterCodec(jsonrpc.NewCodec(), "application/json")
	s.RegisterService(blog.NewAuthenticatedService(&AppEngineAuth{}, false), "blog")
	http.Handle("/rpc", newServer(s))
}

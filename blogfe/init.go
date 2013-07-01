package blogfe

import (
	"net/http"
	"strings"

	"appengine"
	"appengine/user"

	auth "github.com/arunjitsingh/go/auth"
	blog "github.com/arunjitsingh/goblogae"
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

// TODO(arunjitsingh): Move to some appengine utils.
func hostname(r *http.Request) string {
	appID := appengine.AppID(appengine.NewContext(r))
	parts := strings.Split(appID, ":")
	if len(parts) == 1 {
		return parts[0] + ".appspot.com" // "appid" => appid.appspot.com
	}
	return parts[1] + "." + parts[0] // "example.com:appid" => appid.example.com
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
	} else {
		corsOrigins += "," + hostname(r)
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

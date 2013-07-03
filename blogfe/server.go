package blogfe

import (
	"net/http"

	"appengine"
)

const (
	allowedCORSOrigins = "*.ajsd.in,*.arunjit-test.appspot.com"
	allowedCORSHeaders = "X-Requested-With,Content-Type"
	devCORSOrigins     = "*"
)

type Server struct {
	h http.Handler
}

func NewServer(h http.Handler) *Server {
	return &Server{h}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	corsOrigins := allowedCORSOrigins
	if appengine.IsDevAppServer() {
		corsOrigins = devCORSOrigins
	}
	w.Header().Set("Access-Control-Allow-Origin", corsOrigins)
	w.Header().Set("Access-Control-Allow-Headers", allowedCORSHeaders)
	if r.Method == "OPTIONS" {
		return
	}
	s.h.ServeHTTP(w, r)
}

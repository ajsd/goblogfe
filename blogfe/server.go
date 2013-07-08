package blogfe

import (
	"net/http"

	"appengine"
)

const (
	allowedCORSOrigins     = "http://ajsd.github.io"
	allowedCORSHeaders     = "X-Requested-With,Content-Type"
	allowedCORSCredentials = "true"
	devCORSOrigins         = "http://localhost:9000"
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
	w.Header().Set("Access-Control-Allow-Credentials", allowedCORSCredentials)

	if r.Method == "OPTIONS" {
		return
	}
	s.h.ServeHTTP(w, r)
}

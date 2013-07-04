package blogfe

import (
	"net/http"

	blog "github.com/ajsd/goblogae"
	jsonrpc "github.com/arunjitsingh/rpc/v2/json2"
	rpc "github.com/gorilla/rpc/v2"
)

func init() {
	s := rpc.NewServer()
	s.RegisterCodec(jsonrpc.NewCodec(), "application/json")
	s.RegisterService(blog.NewAuthenticatedService(NewAuthenticator(), false), "blog")
	http.Handle("/rpc", NewServer(s))

	http.HandleFunc(loginURL, Login)
	http.HandleFunc(logoutURL, Logout)
	http.HandleFunc(redirectURL, Redirect)
}

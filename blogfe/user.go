package blogfe

import (
	"net/http"
	"net/url"

	"appengine"
	"appengine/user"
)

const (
	nextURLParam = "next"
	loginURL     = "/login"
	logoutURL    = "/logout"
	redirectURL  = "/redirect"
)

func createRedirect(u string) string {
	return redirectURL + "?next=" + url.QueryEscape(u)
}

func Login(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if url, err := user.LoginURL(c, createRedirect(r.FormValue(nextURLParam))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if url, err := user.LoginURL(c, createRedirect(r.FormValue(nextURLParam))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, r.FormValue(nextURLParam), http.StatusFound)
}

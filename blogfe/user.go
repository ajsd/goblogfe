package blogfe

import (
	"net/http"

	"appengine"
	"appengine/user"
)

func Login(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if url, err := user.LoginURL(c, r.Header.Get("Referer")); err != nil {
		http.Error(w, "Couldn't log in", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if url, err := user.LogoutURL(c, r.Header.Get("Referer")); err != nil {
		http.Error(w, "Couldn't log out", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, url, http.StatusFound)
	}
}

package blogfe

import (
	"net/http"

	"appengine"
	"appengine/user"
)

const nextURLParam = "next"

func Login(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if url, err := user.LoginURL(c, r.FormValue(nextURLParam)); err != nil {
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
	if url, err := user.LoginURL(c, r.FormValue(nextURLParam)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, url, http.StatusFound)
	}
}

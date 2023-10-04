package src

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HandleBasic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var redirection Redirect
	for _, red := range Cfg.Redirections {
		if id == red.Id {
			redirection = red
		}
	}
	if redirection.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	redirect := vars["redirect"]
	loc := Location{redirection.Path.generateOrigin(), redirect}
	http.Redirect(w, r, loc.generateUrl(), http.StatusSeeOther)
}

func HandleSlug(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["redirect"]
	redirect := Redirection{}
	DB.First(&redirect, "slug = ?", slug)
	if redirect.Slug != slug {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	http.Redirect(w, r, redirect.Destination, http.StatusSeeOther)
}

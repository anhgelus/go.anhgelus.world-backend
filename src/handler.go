package src

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HandleBasic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var redirection *Redirect
	for _, red := range cfg.Redirections {
		if id == red.Id {
			redirection = &red
		}
	}
	if redirection == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	redirect := vars["redirect"]
	loc := Location{redirect, redirection.Path.generateOrigin()}
	http.Redirect(w, r, loc.generateUrl(), http.StatusSeeOther)
}

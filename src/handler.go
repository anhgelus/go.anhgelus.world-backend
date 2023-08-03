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
		if redirection != nil {
			continue
		}
		if id == red.Id {
			println("id", red.Path)
			redirection = &red
		}
	}
	if redirection == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	println("path", redirection.Path)
	for i, u := range vars {
		println(i, u)
	}
	redirect := vars["redirect"]
	loc := Location{redirection.Path.generateOrigin(), redirect}
	http.Redirect(w, r, loc.generateUrl(), http.StatusSeeOther)
}

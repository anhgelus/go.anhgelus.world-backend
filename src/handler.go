package src

import (
	"net/http"
)

func HandleBasic(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	http.Redirect(w, r, "", http.StatusSeeOther)
}

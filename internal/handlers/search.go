package handlers

import (
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("research")
	if data == "" {
		errorResponse(w, http.StatusBadRequest)
	}
	
}

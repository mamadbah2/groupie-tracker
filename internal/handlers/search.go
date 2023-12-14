package handlers

import (
	"bim/internal/models"
	"bim/internal/pkg"
	"net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("research")
	if data == "" {
		errorResponse(w, http.StatusBadRequest)
	}
	allArtist := getArtists()
	allRelation := getRelations()
	Finded := pkg.FindedArtist(data, allArtist, allRelation)
	renderTemplates(w, "search", &models.TemplateData{AllArtist: Finded})
}

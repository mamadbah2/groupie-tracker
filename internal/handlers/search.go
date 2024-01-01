package handlers

import (
	"bim/internal/models"
	"bim/internal/pkg"
	"net/http"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("research")
	if data == "" {
		errorResponse(w, http.StatusBadRequest)
		return
	}
	data = strings.ToLower(data)
	tabValueSearch := strings.Split(data, " - ")
	var valueSearch string
	var Finded []models.Artists
	allArtist := getArtists()
	allRelation, err := getRelations()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError)
		return
	}
	valueSearch = tabValueSearch[0]
	if len(tabValueSearch) == 2 {
		Finded = pkg.FindedArtist(valueSearch, allArtist, allRelation)
	} else {
		Finded = pkg.FindedArtistGlob(valueSearch, allArtist, allRelation)
	}

	renderTemplates(w, "search", &models.TemplateData{AllArtist: Finded})
}

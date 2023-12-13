package handlers

import (
	"bim/internal/models"
	"net/http"
	"strconv"
)

// Gestionnaire de la page d'accueil
func Home(w http.ResponseWriter, r *http.Request) {
	// Vérification de la méthode
	if r.Method == http.MethodGet {
		// Contrôle des URLs
		if r.URL.Path != "/" {
			errorResponse(w, http.StatusNotFound)
			return
		}

		renderTemplates(w, "home", &models.TemplateData{AllArtist: getArtists(), AllRelation: getRelations()})
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Contrôle des URLs
		if r.URL.Path != "/detail" {
			errorResponse(w, http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(r.URL.Query().Get("Id"))
		if err != nil {
			errorResponse(w, http.StatusBadRequest)
			return
		}
		renderTemplates(w, "detail", &models.TemplateData{Artist: getArtist(id), Relation: getRelation(id)})
	}
}

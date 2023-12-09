package handlers

import (
	"bim/internal/models"
	"net/http"
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

		renderTemplates(w, "home", &models.TemplateData{AllArtist: getArtists()})
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	
}

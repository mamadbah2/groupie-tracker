package main

import (
	"bim/config"
	"bim/internal/handlers"
	"fmt"
	"net/http"
	"os"
)

func Run(tab []string) {
	if len(tab) == 0 {
		var appConfig config.Config // Appel des éléments de configuration

		templateCache, err := handlers.CreateTemplateCache() // Création de cache pour switcher entre les pages
		if err != nil {                                      // Retour en cas d'erreur
			panic(err)
		}

		appConfig.TemplateCache = templateCache // Création du cache de la vue
		appConfig.Port = ":1111"                // Configuration du port
		appConfig.StaticDir = "assets"          // Configuration des fichiers statiques
		handlers.CreateTemplates(&appConfig)    // Création de la vue
		router := http.NewServeMux()            // Appel du serveur Mux pour gérer URLs

		// Passage des fichiers statiques
		statics := http.FileServer(http.Dir(appConfig.StaticDir))
		router.Handle("/static/", http.StripPrefix("/static/", statics))

		// Passages du gestionnaires des templates
		router.HandleFunc("/", handlers.Home)
		router.HandleFunc("/detail", handlers.Detail)

		// Lancement du serveur d'application
		fmt.Println("Server started on port http://0.0.0.0" + appConfig.Port)
		http.ListenAndServe(appConfig.Port, router)
		os.Exit(0) // Sortie du programme
	}
	fmt.Println("---<< Too many arguments >>---")
}

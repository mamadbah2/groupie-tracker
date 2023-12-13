package handlers

import (
	"bim/internal/models"
	"encoding/json"
	"net/http"
)

const apiURL = "https://groupietrackers.herokuapp.com/api/"

func get(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		return err
	}
	return nil
}

func getArtists() []models.Artists {
	var artists []models.Artists
	err := get(apiURL+"artists", &artists)
	if err != nil {
		return nil
	}
	return artists
}

func getArtist(id int) models.Artists {

	return getArtists()[id-1]
}

func getRelations() models.Relation {
	var relations models.Relation
	err := get(apiURL+"relation", &relations)
	if err != nil {
		return models.Relation{}
	}
	return relations
}

func getRelation(id int) interface{} {
	return getRelations().Index[id-1]
}

func getCordonnate(latitude *float64, longitude *float64) {

}

package handlers

import (
	"bim/internal/models"
	"encoding/json"
	"net/http"
)

const KEYMAP = "LxqL2Zy8AMyJhz1aVcLPd43rMTiaLydz"
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

func getRelations() (models.Relation, error) {
	var relations models.Relation
	err := get(apiURL+"relation", &relations)

	return relations, err
}

func getRelation(id int) interface{} {
	getIt, _ := getRelations()
	return getIt.Index[id-1]
}

func getCordonnate(place string, latitude *float64, longitude *float64) error {
	respMap, err := http.Get("https://www.mapquestapi.com/geocoding/v1/address?key=" + KEYMAP + "&location=" + place)
	var locate models.APIResponse
	err = json.NewDecoder(respMap.Body).Decode(&locate)
	*latitude = locate.Results[0].Locations[0].LatLng.Lat
	*longitude = locate.Results[0].Locations[0].LatLng.Lng
	return err
}

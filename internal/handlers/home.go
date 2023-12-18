package handlers

import (
	"bim/internal/models"
	"bim/internal/pkg"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Gestionnaire de la page d'accueil
func Home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path != "/" {
			errorResponse(w, http.StatusNotFound)
			return
		}
		all1 := getArtists()
		all2 := getRelations()
		fString := pkg.GetUniqueLocation(all2)
		renderTemplates(w, "home", &models.TemplateData{AppInfos: models.App{AppName: appName, PageTitle: "List of Artists", PageName: "Home"}, AllArtist: all1, FilterLocations: fString})

	case http.MethodPost:
		var Arts []models.Artists
		var nums []string
		all1 := getArtists()
		all2 := getRelations()
		yearStart := pkg.Atoi(r.FormValue("yearStart"))
		yearEnd := pkg.Atoi(r.FormValue("yearEnd"))
		faStart, err4 := time.Parse("02-01-2006", pkg.BestFormat(r.FormValue("faStart")))
		faEnd, err5 := time.Parse("02-01-2006", pkg.BestFormat(r.FormValue("faEnd")))
		for i := '1'; i <= '8'; i++ {
			if r.FormValue("members"+string(i)) != "" {
				nums = append(nums, r.FormValue("members"+string(i)))
			}
		}
		members, err3 := pkg.TabToTab(nums)
		// members = append(members, num)GetLocations
		fString := pkg.GetUniqueLocation(all2)
		location := r.FormValue("location")
		if err3 != nil || err4 != nil || err5 != nil {
			errorResponse(w, http.StatusBadRequest)
			return
		}
		Arts = pkg.FilterArtists(all1, all2, faStart, faEnd, yearStart, yearEnd, members, location)

		renderTemplates(w, "home", &models.TemplateData{AllArtist: Arts, AllRelation: getRelations(), FilterLocations: fString})
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
		if err != nil || (id < 1 || id > 52) {
			errorResponse(w, http.StatusBadRequest)
			return
		}
		var coordinate models.AllGeoCoordinate
		for place := range getRelations().Index[id-1].DatesLocations {
			var x, y float64
			err = getCordonnate(place, &x, &y)
			coordinate.Lat = append(coordinate.Lat, x)
			coordinate.Lng = append(coordinate.Lng, y)
		}
		fmt.Println(coordinate)
		renderTemplates(w, "detail", &models.TemplateData{Artist: getArtist(id), Relation: getRelation(id), Coordinate: coordinate})
	}
}

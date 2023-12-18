package models

type AllGeoCoordinate struct {
	Lat []float64
	Lng []float64
}

// Pour le mapQuest

type APIResponse struct {
	Results []struct {
		Locations []struct {
			LatLng struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"latLng"`
		} `json:"locations"`
	} `json:"results"`
}

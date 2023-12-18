package pkg

import (
	"bim/internal/models"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetUniqueLocation(locations models.Relation) []string {
	var results []string
	var tmp []string
	var slices []string
	for _, l := range locations.Index {
		for locate := range l.DatesLocations {
			tmp = append(tmp, locate)
		}
	}
	for _, t := range tmp {
		slices = append(slices, t)
	}
	results = append(results, slices[0])
	slices = slices[1:]
	for _, s := range slices {
		found := false
		for _, r := range results {
			if s == r {
				found = true
				break
			}
		}
		if !found {
			results = append(results, s)
		}
	}
	return results
}

func FilterArtists(artists []models.Artists, locations models.Relation, startDate time.Time, endDate time.Time, startYear int, endYear int, tabMember []int, locationQuery string) []models.Artists {
	var filtered []models.Artists
	for _, artist := range artists {
		// Filter by date range
		if !artistBelongsToDate(artist, startDate, endDate) {
			continue
		}
		// Filter by year range
		if !artistBelongsToYear(artist, startYear, endYear) {
			continue
		}
		// Filter by member range
		if !artistBelongsToMember(artist, tabMember) {
			continue
		}
		// Filter by location
		if !artistBelongsToLocation(artist, locationQuery, locations) && locationQuery != "none" {
			continue
		}
		filtered = append(filtered, artist)
	}
	return filtered
}

// artistBelongsToDateRange checks if the artist's first album date is within the specified range
func artistBelongsToDate(artist models.Artists, start, end time.Time) bool {
	date, err := time.Parse("02-01-2006", artist.FirstAlbum)
	if err != nil {
		fmt.Println("Error parsing date", err)
		return false
	}
	return date.After(start) && date.Before(end)
}

// artistBelongsToYearRange checks if the artist's creation date is within the specified range
func artistBelongsToYear(artist models.Artists, start, end int) bool {
	return artist.CreationDate >= start && artist.CreationDate <= end
}

// artistBelongsToMemberRange checks if the artist's number of members is within the specified range
func artistBelongsToMember(artist models.Artists, tab []int) bool {
	if len(tab) < 0 || len(tab) > 8 {
		return false
	}
	sort.Ints(tab)
	if len(tab) == 1 {
		return len(artist.Members) == tab[0]
	}
	if len(tab) == 0 {
		return true
	}
	/*for _, v := range tab {
		if len(artist.Members) == v {
			return true
		}
	}*/
	// A revoir
	return len(artist.Members) >= tab[0] && len(artist.Members) <= tab[len(tab)-1]
	// return false
}

// artistBelongsToLocation checks if the artist's location matches the specified query
func artistBelongsToLocation(artist models.Artists, query string, locations models.Relation) bool {
	var IDs []int
	for _, index := range locations.Index {
		for location := range index.DatesLocations {
			if location == query {
				IDs = append(IDs, index.Id)
			}
		}
	}
	for _, id := range IDs {
		if artist.Id == id {
			return true
		}
	}
	return false
}

func Atoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error parsing string to int", err)
		return 0
	}
	return result
}

func BestFormat(str string) string {
	tmp := strings.Split(str, "-")
	var s string
	var tab []string
	for i := len(tmp) - 1; i >= 0; i-- {
		tab = append(tab, tmp[i])
	}
	for i, t := range tab {
		if i > 0 && i <= len(tab)-1 {
			s += "-"
		}
		s += t
	}
	return s
}

func TabToTab(tab []string) ([]int, error) {
	var result []int
	for _, t := range tab {
		val, err := strconv.Atoi(t)
		if err != nil {
			return nil, err
		}
		result = append(result, val)
	}
	return result, nil
}

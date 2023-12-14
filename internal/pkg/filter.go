package pkg

import (
	"bim/internal/models"
	"fmt"
	"sort"
	"strconv"
	"time"
)

// DatesBelongsTo returns a slice of artists by date range
// func DatesBelongsTo(start, end time.Time, artists []models.Artists) []models.Artists {
// 	var filtered []models.Artists
// 	for _, artist := range artists {
// 		date, err := time.Parse("02-01-2006", artist.FirstAlbum)
// 		if err != nil {
// 			fmt.Println("Error parsing date", err)
// 			continue
// 		}
// 		if date.After(start) && date.Before(end) {
// 			filtered = append(filtered, artist)
// 			continue
// 		}
// 	}
// 	return filtered
// }

// YearsBelongsTo returns a slice of artists by year range
// func YearsBelongsTo(start, end int, artists []models.Artists) []models.Artists {
// 	var filtered []models.Artists
// 	for _, artist := range artists {
// 		if artist.CreationDate >= start && artist.CreationDate <= end {
// 			filtered = append(filtered, artist)
// 			continue
// 		}
// 	}
// 	return filtered
// }

// ArtistsByMembers returns a slice of artists by number of members
func ArtistByMembers(tab []int, artists []models.Artists) []models.Artists {
	if len(tab) <= 0 || len(tab) > 8 {
		return nil
	}
	sort.Ints(tab)
	var result []models.Artists
	for _, artist := range artists {
		for _, v := range tab {
			if len(artist.Members) == v {
				result = append(result, artist)
				continue
			}
		}
	}
	return result
}

// ArtistsByLocation returns a slice of artists by location
func ArtistByLocation(query string, artists []models.Artists, relation models.Relation) []models.Artists {
	var result []models.Artists
	var IDs []int
	for _, index := range relation.Index {
		if _, location := index.DatesLocations[query]; location {
			IDs = append(IDs, index.Id)
		}
	}
	for _, artist := range artists {
		for _, id := range IDs {
			if artist.Id == id {
				result = append(result, artist)
			}
		}
	}
	return result
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

// FilterArtists returns a slice of artists
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
		if !artistBelongsToLocation(artist, locationQuery, locations) {
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
	if len(tab) <= 0 || len(tab) > 8 {
		return false
	}
	sort.Ints(tab)
	if len(tab) == 1 {
		return len(artist.Members) == tab[0]
	}
	return len(artist.Members) >= tab[0] && len(artist.Members) <= tab[len(tab)-1]
}

// artistBelongsToLocation checks if the artist's location matches the specified query
func artistBelongsToLocation(artist models.Artists, query string, relation models.Relation) bool {
	var IDs []int
	for _, index := range relation.Index {
		if _, location := index.DatesLocations[query]; location {
			IDs = append(IDs, index.Id)
		}
	}
	for _, id := range IDs {
		if artist.Id == id {
			return true
		}
	}
	return false
}

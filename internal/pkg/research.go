package pkg

import (
	"bim/internal/models"
	"strconv"
)

func doublon(table []int, data int) bool {
	for _, id := range table {
		if id == data {
			return true
		}
	}
	return false
}

func unionTable(table1, table2 []int) []int {
	var table []int
	for _, t := range table1 {
		if !doublon(table, t) {
			table = append(table, t)
		}
	}
	for _, t := range table2 {
		if !doublon(table, t) {
			table = append(table, t)
		}
	}
	return table
}

func matchIdAndArtist(tableId []int, allArtist []models.Artists) []models.Artists {
	var artistFinded []models.Artists
	for _, id := range tableId {
		artistFinded = append(artistFinded, allArtist[id-1])
	}
	return artistFinded
}

func researchInArtist(data string, tableArtists []models.Artists) []int {
	var tabId []int
	for _, artist := range tableArtists {
		if artist.Name == data || artist.FirstAlbum == data {
			tabId = append(tabId, artist.Id)
		}
		annee, err := strconv.ParseInt(data, 10, 64)
		if err == nil {
			if annee == int64(artist.CreationDate) {
				tabId = append(tabId, artist.Id)
			}
		}
		for _, member := range artist.Members {
			if member == data {
				tabId = append(tabId, artist.Id)

			}
		}
	}
	return tabId
}

func researchInRelation(data string, tableRelations models.Relation) []int {
	var tabId []int
	for _, relation := range tableRelations.Index {
		if _, location := relation.DatesLocations[data]; location {
			tabId = append(tabId, relation.Id)
		}
	}
	return tabId
}

func FindedArtist(data string, allArtist []models.Artists, allRelation models.Relation) []models.Artists {
	return matchIdAndArtist(unionTable(researchInArtist(data, allArtist), researchInRelation(data, allRelation)), allArtist)
}

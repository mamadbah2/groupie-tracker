package pkg

import (
	"bim/internal/models"
	"strconv"
	"strings"
)

func doublon(table []int, data int) bool {
	for _, id := range table {
		if id == data {
			return true
		}
	}
	return false
}

// Prend 2 tableaux et unie les tableaux tout en eliminant les doublons
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

// Cette fonction prend la donnee qu'on veut rechercher et tout les artists
// Apres avoir rechercher ça... Il retourne tous les Id qui correspondent a la
// recherche
func researchInArtist(data string, tableArtists []models.Artists) []int {
	var tabId []int
	for _, artist := range tableArtists {
		if strings.ToLower(artist.Name) == data || strings.ToLower(artist.FirstAlbum) == data {
			tabId = append(tabId, artist.Id)
		}
		annee, err := strconv.ParseInt(data, 10, 64)
		if err == nil {
			if annee == int64(artist.CreationDate) {
				tabId = append(tabId, artist.Id)
			}
		}
		for _, member := range artist.Members {
			if strings.ToLower(member) == data {
				tabId = append(tabId, artist.Id)

			}
		}
	}
	return tabId
}

func researchInArtistGlob(data string, tableArtists []models.Artists) []int {
	var tabId []int
	for _, artist := range tableArtists {
		if strings.Contains(strings.ToLower(artist.Name), data) || strings.Contains(artist.FirstAlbum, data) {
			tabId = append(tabId, artist.Id)
		}
		if strings.Contains(strconv.Itoa(artist.CreationDate), data) {
			tabId = append(tabId, artist.Id)
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), data) {
				tabId = append(tabId, artist.Id)

			}
		}
	}
	return tabId
}

// Cette fonction prend la donnee qu'on veut rechercher et toutes les relations
// Apres avoir rechercher ça... Il retourne tous les Id qui correspondent a la
// recherche
func researchInRelation(data string, tableRelations models.Relation) []int {
	var tabId []int
	for _, relation := range tableRelations.Index {
		for key := range relation.DatesLocations {
			if strings.ToLower(key) == data {
				tabId = append(tabId, relation.Id)
			}
		}
	}
	return tabId
}

func researchInRelationGlob(data string, tableRelations models.Relation) []int {
	var tabId []int
	for _, relation := range tableRelations.Index {
		for key := range relation.DatesLocations {
			if strings.Contains(strings.ToLower(key), data) {
				tabId = append(tabId, relation.Id)
			}
		}
	}
	return tabId
}

// Cette fonction fait correspondre les id du tableau d'id avec l'artist correspondant
// et retourne les artists trouvés

func matchIdAndArtist(tableId []int, allArtist []models.Artists) []models.Artists {
	var artistFinded []models.Artists
	for _, id := range tableId {
		artistFinded = append(artistFinded, allArtist[id-1])
	}
	return artistFinded
}

func FindedArtist(data string, allArtist []models.Artists, allRelation models.Relation) []models.Artists {
	return matchIdAndArtist(unionTable(researchInArtist(data, allArtist), researchInRelation(data, allRelation)), allArtist)
}

func FindedArtistGlob(data string, allArtist []models.Artists, allRelation models.Relation) []models.Artists {
	return matchIdAndArtist(unionTable(researchInArtistGlob(data, allArtist), researchInRelationGlob(data, allRelation)), allArtist)
}
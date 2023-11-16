package search

import (
	"encoding/json"
	"groupie-tracker/internal/models"
	"groupie-tracker/internal/pkg"
	"strconv"
)

var ArtistURL string = "https://groupietrackers.herokuapp.com/api/artists"

func SearchByTxt(searchText string, body []byte) (map[string]interface{}, error) {
	suggestions := make([]string, 0)
	artists := make(map[int]string)
	var groups models.Groups
	json.Unmarshal(body, &groups.Groups)
	for _, artist := range groups.Groups {
		if pkg.Search(searchText, artist.Name, "") {
			suggestions = append(suggestions, artist.Name)
			artists[artist.Id] = artist.Name
		}
		for _, members := range artist.Members {
			if pkg.Search(searchText, members, "") {
				suggestions = append(suggestions, members+" - a member of "+artist.Name)
				artists[artist.Id] = artist.Name
			}
		}
		creationDate := strconv.Itoa(artist.CreationDate)
		if pkg.Search(searchText, creationDate, "") {
			suggestions = append(suggestions, creationDate+" - creation date of "+artist.Name)
			artists[artist.Id] = artist.Name
		}
		if pkg.Search(searchText, artist.FirstAlbum, "album") {
			suggestions = append(suggestions, artist.FirstAlbum+" - first album date of "+artist.Name)
			artists[artist.Id] = artist.Name
		}
	}
	responseMap := map[string]interface{}{
		"suggestions": suggestions,
		"artists":     artists,
	}

	return responseMap, nil
}

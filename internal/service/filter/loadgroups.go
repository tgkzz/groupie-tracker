package filter

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/internal/models"
	"groupie-tracker/internal/service/api"
	"strconv"
)

var (
	ArtistURL   string = "https://groupietrackers.herokuapp.com/api/artists"
	LocationURL string = "https://groupietrackers.herokuapp.com/api/locations/"
	RelationURL string = "https://groupietrackers.herokuapp.com/api/relation/"
)

func LoadGroups() ([]models.ResultGroup, error) {
	var result []models.ResultGroup

	for i := 1; i <= 52; i++ {
		var Group models.ResultGroup

		id := strconv.Itoa(i)

		fmt.Println(id)

		// marshaling artist url
		bodyArtist, err := api.ReturnResponseBody(ArtistURL + "/" + id)
		if err != nil {
			return result, err
		}
		var group models.Group
		json.Unmarshal(bodyArtist, &group)

		// marshalling locations url
		bodyLocation, err := api.ReturnResponseBody(LocationURL + id)
		if err != nil {
			return result, err
		}
		var locations models.Locations
		json.Unmarshal(bodyLocation, &locations)

		// marshaling relations url
		bodyRelation, err := api.ReturnResponseBody(RelationURL + id)
		if err != nil {
			return result, err
		}
		var dateLocation models.DateLocation
		json.Unmarshal(bodyRelation, &dateLocation)

		Group.Id = group.Id
		Group.Name = group.Name
		Group.Image = group.Image
		Group.CreationDate = group.CreationDate
		Group.Members = group.Members
		Group.FirstAlbum = group.FirstAlbum

		for _, loc := range locations.Locate {
			Group.ConcertData = append(Group.ConcertData, models.Concerts{Location: loc, Dates: dateLocation.DateLoc[loc]})
		}

		result = append(result, Group)
	}

	return result, nil
}

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

func LoadGroupsGoroutine() ([]models.ResultGroup, error) {
	var result []models.ResultGroup
	numGroups := 52

	groupCh := make(chan models.ResultGroup, numGroups)
	errCh := make(chan error, numGroups)

	for i := 1; i <= numGroups; i++ {
		go func(id int) {
			var Group models.ResultGroup

			fmt.Println(id)

			// marshaling artist url
			bodyArtist, err := api.ReturnResponseBody(ArtistURL + "/" + strconv.Itoa(id))
			if err != nil {
				errCh <- err
				return
			}
			var group models.Group
			json.Unmarshal(bodyArtist, &group)

			// marshalling locations url
			bodyLocation, err := api.ReturnResponseBody(LocationURL + strconv.Itoa(id))
			if err != nil {
				errCh <- err
				return
			}
			var locations models.Locations
			json.Unmarshal(bodyLocation, &locations)

			// marshaling relations url
			bodyRelation, err := api.ReturnResponseBody(RelationURL + strconv.Itoa(id))
			if err != nil {
				errCh <- err
				return
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

			groupCh <- Group

		}(i)
	}

	for i := 0; i < numGroups; i++ {
		select {
		case group := <-groupCh:
			result = append(result, group)
		case err := <-errCh:
			return nil, err
		}
	}

	return result, nil
}

func LoadGroupsDefault() ([]models.ResultGroup, error) {
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

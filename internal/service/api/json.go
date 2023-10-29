package api

import (
	"encoding/json"
	"groupie-tracker/internal/models"
)

func GroupsJsonMarshalling(url string) (models.Groups, error) {
	emptyStruct := models.Groups{}

	var groups models.Groups

	body, err := ReturnResponseBody(url)
	if err != nil {
		return emptyStruct, err
	}

	if err = json.Unmarshal(body, &groups.Groups); err != nil {
		return emptyStruct, err
	}

	return groups, err
}

func GroupJsonMarshalling(url string) (models.Group, error) {
	emptyStruct := models.Group{}

	var group models.Group

	bodyArtist, err := ReturnResponseBody(url)
	if err != nil {
		return emptyStruct, err
	}

	if err = json.Unmarshal(bodyArtist, &group); err != nil {
		return emptyStruct, err
	}

	return group, err
}

func LocationJsonMarshalling(url string) (models.Locations, error) {
	emptyStruct := models.Locations{}

	var locations models.Locations

	bodyLocation, err := ReturnResponseBody(url)
	if err != nil {
		return emptyStruct, err
	}

	err = json.Unmarshal(bodyLocation, &locations)
	if err != nil {
		return emptyStruct, err
	}

	return locations, err
}

func RelationJsonMarshalling(url string) (models.DateLocation, error) {
	emptyStruct := models.DateLocation{}

	var dateLocation models.DateLocation

	bodyRelation, err := ReturnResponseBody(url)
	if err != nil {
		return emptyStruct, err
	}

	if err := json.Unmarshal(bodyRelation, &dateLocation); err != nil {
		return emptyStruct, err
	}

	return dateLocation, err
}

func FormResultGroup(group models.Group, locations models.Locations, dateLocation models.DateLocation) models.ResultGroup {
	var ResultGroup models.ResultGroup

	ResultGroup.Id = group.Id
	ResultGroup.Image = group.Image
	ResultGroup.Name = group.Name
	ResultGroup.CreationDate = group.CreationDate
	ResultGroup.Members = group.Members
	ResultGroup.FirstAlbum = group.FirstAlbum

	for _, loc := range locations.Locate {
		ResultGroup.ConcertData = append(ResultGroup.ConcertData, models.Concerts{Location: loc, Dates: dateLocation.DateLoc[loc]})
	}

	return ResultGroup
}

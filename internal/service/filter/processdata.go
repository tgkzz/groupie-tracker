package filter

import (
	"groupie-tracker/internal/models"
	"groupie-tracker/internal/pkg"
)

func ProcessData(filter models.Filter) ([]models.ResultGroup, error) {
	emptyStruct := []models.ResultGroup{}

	var Result []models.ResultGroup

	allGroups, err := LoadGroups()
	if err != nil {
		return emptyStruct, err
	}

	for _, group := range allGroups {
		if group.CreationDate > filter.CreationDateFrom &&
			group.CreationDate < filter.CreationDateTo &&
			pkg.InTheSlice(len(filter.Members), filter.Members) &&
			filter.FirstAlbum < pkg.TakeYearFromData(group.FirstAlbum) {
			Result = append(Result, group)
		}
	}

	return Result, nil
}

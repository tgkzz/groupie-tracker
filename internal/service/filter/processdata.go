package filter

import (
	"fmt"
	"groupie-tracker/internal/models"
	"groupie-tracker/internal/pkg"
)

// TODO : implement bad request (filter.CreationDateFrom > filter.CreationDateTo and etc.)
func ProcessData(filter models.Filter) ([]models.ResultGroup, error) {
	emptyStruct := []models.ResultGroup{}

	var Result []models.ResultGroup

	allGroups, err := LoadGroups()
	if err != nil {
		return emptyStruct, err
	}

	fmt.Println("Applied Filters:")
	fmt.Printf("Creation date (from): %#v\n", filter.CreationDateFrom)
	fmt.Printf("Creation date (To): %#v\n", filter.CreationDateTo)
	fmt.Printf("FirstAlbum: %#v\n", filter.FirstAlbum)
	fmt.Printf("Possible member num: %#v\n", filter.Members)

	for _, group := range allGroups {
		if group.CreationDate >= filter.CreationDateFrom &&
			group.CreationDate <= filter.CreationDateTo &&
			pkg.InTheSlice(len(group.Members), filter.Members) &&
			filter.FirstAlbum <= pkg.TakeYearFromData(group.FirstAlbum) {
			Result = append(Result, group)
		}
	}

	return Result, nil
}

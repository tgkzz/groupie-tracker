package filter

import (
	"fmt"
	"groupie-tracker/internal/models"
	"groupie-tracker/internal/pkg"
)

// TODO : add location filter
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
	fmt.Printf("FirstAlbum (from): %#v\n", filter.FirstAlbumFrom)
	fmt.Printf("FirstAlbum (To): %#v\n", filter.FirstAlbumTo)
	fmt.Printf("Possible member num: %#v\n", filter.Members)

	for _, group := range allGroups {
		if group.CreationDate >= filter.CreationDateFrom && group.CreationDate <= filter.CreationDateTo &&
			filter.FirstAlbumFrom <= pkg.TakeYearFromData(group.FirstAlbum) && filter.FirstAlbumTo >= pkg.TakeYearFromData(group.FirstAlbum) &&
			pkg.InTheSlice(len(group.Members), filter.Members) {
			Result = append(Result, group)
		}

	}

	return Result, nil
}

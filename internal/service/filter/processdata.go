package filter

import (
	"fmt"
	"groupie-tracker/internal/models"
	"groupie-tracker/internal/pkg"
	"time"
)

// TODO : add location filter
// TODO : think of how to get locations []string
func ProcessData(filter models.Filter) ([]models.ResultGroup, error) {
	emptyStruct := []models.ResultGroup{}

	var Result []models.ResultGroup

	start := time.Now()
	allGroups, err := LoadGroupsGoroutine()
	if err != nil {
		return emptyStruct, err
	}
	end := time.Since(start)
	fmt.Println(end)

	fmt.Println("Applied Filters:")
	fmt.Printf("Creation date (from): %#v\n", filter.CreationDateFrom)
	fmt.Printf("Creation date (To): %#v\n", filter.CreationDateTo)
	fmt.Printf("FirstAlbum (from): %#v\n", filter.FirstAlbumFrom)
	fmt.Printf("FirstAlbum (To): %#v\n", filter.FirstAlbumTo)
	fmt.Printf("Possible member num: %#v\n", filter.Members)
	fmt.Printf("Locations: %s\n", filter.Location)

	for _, group := range allGroups {

		allLocation := []string{}

		for _, concert := range group.ConcertData {
			allLocation = append(allLocation, concert.Location)
		}

		if group.CreationDate >= filter.CreationDateFrom && group.CreationDate <= filter.CreationDateTo &&
			filter.FirstAlbumFrom <= pkg.FetchYearFromData(group.FirstAlbum) && filter.FirstAlbumTo >= pkg.FetchYearFromData(group.FirstAlbum) &&
			pkg.InTheSliceInt(len(group.Members), filter.Members) &&
			pkg.InTheSliceString(filter.Location, allLocation) {
			Result = append(Result, group)
		}

	}

	return Result, nil
}

package filter

import (
	"groupie-tracker/internal/models"
	"strconv"
)

func DataHandling(from string, to string, firstAlbumFrom string, firstAlbumTo string, members []string) (models.Filter, error) {
	var err error

	emptyStruct := models.Filter{}

	var result models.Filter

	if from == "" {
		from = "0"
	}

	if to == "" {
		to = "922337203685477580"
	}

	if firstAlbumFrom == "" {
		firstAlbumFrom = "0"
	}

	if firstAlbumTo == "" {
		firstAlbumFrom = "922337203685477580"
	}

	result.CreationDateFrom, err = strconv.Atoi(from)
	if err != nil {
		return emptyStruct, err
	}

	result.CreationDateTo, err = strconv.Atoi(to)
	if err != nil {
		return emptyStruct, err
	}

	result.FirstAlbumTo, err = strconv.Atoi(firstAlbumTo)
	if err != nil {
		return emptyStruct, err
	}

	result.FirstAlbumFrom, err = strconv.Atoi(firstAlbumFrom)
	if err != nil {
		return emptyStruct, err
	}

	if members == nil {
		for i := 1; i <= 10; i++ {
			result.Members = append(result.Members, i)
		}
	} else {
		for _, num := range members {
			tmp, err := strconv.Atoi(num)
			if err != nil {
				return emptyStruct, err
			}
			result.Members = append(result.Members, tmp)
		}
	}

	return result, nil
}

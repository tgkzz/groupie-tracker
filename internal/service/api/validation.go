package api

import (
	"strconv"
	"strings"
)

func PathValidation(path string) error {
	id := strings.TrimPrefix(path, "/groups/")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	if id == "" || idNum > 52 {
		return err
	}

	return nil
}

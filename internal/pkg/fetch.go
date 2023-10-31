package pkg

import (
	"strconv"
	"strings"
)

func FetchYearFromData(str string) int {
	parts := strings.Split(str, "-")

	result, _ := strconv.Atoi(parts[2])

	return result
}

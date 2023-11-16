package pkg

import "strings"

func Search(text, term, is string) bool {
	term = strings.ToLower(term)
	arr := Split(term, is)
	res := 0
	res2 := false
	text = strings.ToLower(text)
	for i := range arr {
		if len(arr[i]) == 0 {
			continue
		}
		for j := 0; j < len(text); j++ {
			if len(arr[i]) > j {
				if text[j] != arr[i][j] {
					res++
				}
			}
		}
		if res == 0 {
			return true
		}
		res = 0
	}
	return res2
}
func Split(text, term string) []string {
	switch term {
	case "album":
		return strings.Split(text, "-")
	default:
		return strings.Split(text, " ")
	}
}

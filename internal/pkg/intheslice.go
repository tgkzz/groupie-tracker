package pkg

func InTheSliceInt(a int, all []int) bool {
	for _, num := range all {
		if num == a {
			return true
		}
	}

	return false
}

func InTheSliceString(a string, all []string) bool {
	if a == "" {
		return true
	}

	for _, str := range all {
		if str == a {
			return true
		}
	}

	return false
}

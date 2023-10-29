package pkg

func InTheSlice(a int, all []int) bool {
	for _, num := range all {
		if num == a {
			return true
		}
	}

	return false
}

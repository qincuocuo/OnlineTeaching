package utils

func IsContainInSlice(target string, items []string) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

func IsContainIntInSlice(target int, items []int) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

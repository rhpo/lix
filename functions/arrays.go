package functions

func InArray[T comparable](array []T, item T) bool {
	// for loop
	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
}

func Remove[T comparable](array []T, item T) []T {
	// for loop
	var result []T
	for _, v := range array {
		if v != item {
			result = append(result, v)
		}
	}
	return result
}

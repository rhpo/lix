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

func InsertAt(str string, item string, pos int) string {
	if pos < 0 || pos > len(str) {
		return str
	}
	var result []byte
	for i := 0; i < len(str); i++ {
		if i == pos {
			result = append(result, item...)
		}
		result = append(result, str[i])
	}
	if pos == len(str) {
		result = append(result, item...)
	}
	return string(result)
}

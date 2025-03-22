package functions

import "strings"

func Join(A []string) string {
	str := ""
	for i := 0; i < len(A); i++ {
		str += A[i]
	}
	return str
}

func Split(str string) []string {
	res := make([]string, 0)
	for i := 0; i < len(str); i++ {
		res = append(res, string(str[i]))
	}
	return res
}

func Trim(str string) string {
	return strings.TrimSpace(str)
}

func Reverse(str string) string {
	result := Split(str)
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return Join(result)
}

func Sort(A []string) []string {
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
	return A
}

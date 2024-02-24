package utils

/**
 * Limit number of items
*/
func Limit[T any](array []T, limit int) []T {
	var arr []T
	var l int

	if limit > len(array) {
		l = len(array)
	} else {
		l = limit
	}

	for i := 0; i < l; i++ {
		arr = append(arr, array[i])
	}
	return arr
}

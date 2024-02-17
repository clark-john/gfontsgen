package utils

import "strings"

/**
 * This is only used for key-value pairs with equals between them
 */
func ToMap(arr []string) map[string]string {
	m := make(map[string]string)
	for _, k := range arr {
		j := strings.Split(k, "=")
		m[j[0]] = j[1]
	}

	return m
}

/**
 * check if a string is in the array
*/
func IsIn(array []string, value string) bool {
	for _, val := range array {
		if strings.EqualFold(val, strings.ToLower(value)) {
			return true
		}
	}
	return false
}

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

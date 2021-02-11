package main

import "strconv"

// for []int{1, 2, 3}, this function returns "1, 2, 3"
func marshalValues(values []int) string {
	result := ""
	for _, value := range values {
		if result != "" {
			result = result + ", "
		}
		result = result + strconv.Itoa(value)
	}

	return result
}

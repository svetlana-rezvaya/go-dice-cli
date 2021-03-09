package main

// for []int{1, 2, 3}, this function returns "1, 2, 3"
func marshalValues(values []int, useUnicode bool) string {
	result := ""
	for _, value := range values {
		if result != "" {
			result = result + ", "
		}
		result = result + marshalValue(value, useUnicode)
	}

	return result
}

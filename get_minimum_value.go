package main

func getMinimumValue(values []int) int {
	minimum := 0 // in this case there cannot be 0 in values
	for _, value := range values {
		if minimum == 0 || value < minimum {
			minimum = value
		}
	}

	return minimum
}

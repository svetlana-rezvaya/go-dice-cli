package main

func getMaximumValue(values []int) int {
	maximum := 0 // in this case there cannot be 0 in values
	for _, value := range values {
		if maximum == 0 || value > maximum {
			maximum = value
		}
	}

	return maximum
}

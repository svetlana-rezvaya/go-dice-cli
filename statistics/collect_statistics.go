package statistics

// CollectStatistics ...
func CollectStatistics(values []int) Statistics {
	minimum := getMinimumValue(values)
	maximum := getMaximumValue(values)
	sum := getSumOfValues(values)
	return Statistics{Minimum: minimum, Maximum: maximum, Sum: sum}
}

func getMinimumValue(values []int) int {
	minimum := 0 // in this case there cannot be 0 in values
	for _, value := range values {
		if minimum == 0 || value < minimum {
			minimum = value
		}
	}

	return minimum
}

func getMaximumValue(values []int) int {
	maximum := 0 // in this case there cannot be 0 in values
	for _, value := range values {
		if maximum == 0 || value > maximum {
			maximum = value
		}
	}

	return maximum
}

func getSumOfValues(values []int) int {
	sum := 0
	for _, value := range values {
		sum = sum + value
	}

	return sum
}

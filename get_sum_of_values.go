package main

func getSumOfValues(values []int) int {
	sum := 0
	for _, value := range values {
		sum = sum + value
	}

	return sum
}

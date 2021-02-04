package main

func collectStatistics(values []int) statistics {
	minimum := getMinimumValue(values)
	maximum := getMaximumValue(values)
	sum := getSumOfValues(values)
	return statistics{minimum: minimum, maximum: maximum, sum: sum}
}

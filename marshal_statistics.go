package main

import "strconv"

func marshalStatistics(stats statistics) string {
	return "minimum: " + strconv.Itoa(stats.minimum) + "\n" +
		"maximum: " + strconv.Itoa(stats.maximum) + "\n" +
		"sum: " + strconv.Itoa(stats.sum) + "\n"
}

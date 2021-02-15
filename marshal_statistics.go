package main

import "fmt"

func marshalStatistics(stats statistics) string {
	return fmt.Sprintf(
		"minimum: %d\nmaximum: %d\nsum: %d\n",
		stats.minimum,
		stats.maximum,
		stats.sum,
	)
}

package main

import "fmt"

func marshalStatistics(stats statistics, useUnicode bool) string {
	return fmt.Sprintf(
		"minimum: %s\nmaximum: %s\nsum: %s\n",
		marshalValue(stats.minimum, useUnicode),
		marshalValue(stats.maximum, useUnicode),
		marshalValue(stats.sum, useUnicode),
	)
}

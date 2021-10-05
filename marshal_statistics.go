package dice

import "fmt"

// MarshalStatistics ...
func MarshalStatistics(stats Statistics, useUnicode bool) string {
	return fmt.Sprintf(
		"minimum: %s\nmaximum: %s\nsum: %s\n",
		marshalValue(stats.Minimum, useUnicode),
		marshalValue(stats.Maximum, useUnicode),
		marshalValue(stats.Sum, useUnicode),
	)
}

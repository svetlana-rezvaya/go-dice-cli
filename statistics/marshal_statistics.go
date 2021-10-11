package statistics

import (
	"fmt"

	dice "github.com/svetlana-rezvaya/go-dice-cli"
)

// MarshalStatistics ...
func MarshalStatistics(stats Statistics, useUnicode bool) string {
	return fmt.Sprintf(
		"minimum: %s\nmaximum: %s\nsum: %s\n",
		dice.MarshalValue(stats.Minimum, useUnicode),
		dice.MarshalValue(stats.Maximum, useUnicode),
		dice.MarshalValue(stats.Sum, useUnicode),
	)
}

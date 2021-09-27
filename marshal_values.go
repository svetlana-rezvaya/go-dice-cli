package main

import (
	"strconv"
	"strings"
)

// MarshalValues ...
// For []int{1, 2, 3}, this function returns "1, 2, 3".
func MarshalValues(values []int, useUnicode bool) string {
	result := ""
	for _, value := range values {
		if result != "" {
			result = result + ", "
		}
		result = result + marshalValue(value, useUnicode)
	}

	return result
}

func marshalValue(value int, useUnicode bool) string {
	if useUnicode {
		return toUnicodeDices(value)
	}

	return strconv.Itoa(value)
}

func toUnicodeDices(value int) string {
	if value == 0 {
		return ""
	}
	if value <= 6 {
		return string(rune('\u2680' + value - 1))
	}
	if value < 10 {
		return "\u2685" + toUnicodeDices(value-6)
	}

	order := value / 10
	rest := value % 10
	result := strings.Repeat("\u2684\u2684 ", order) + toUnicodeDices(rest)
	return strings.TrimSpace(result)
}

package main

import (
	"strconv"
	"strings"
)

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

package main

func toUnicodeDices(value int) string {
	if value <= 6 {
		return string('\u2680' + value - 1)
	}
	if value < 10 {
		return "\u2685" + toUnicodeDices(value-6)
	}

	panic("toUnicodeDices: not yet implemented")
}

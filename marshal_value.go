package main

func toUnicodeDices(value int) string {
	if value <= 6 {
		return string('\u2680' + value - 1)
	}

	panic("toUnicodeDices: not yet implemented")
}

package dice

import "testing"

func TestMarshalStatistics_withoutUnicode(test *testing.T) {
	stats := Statistics{Minimum: 7, Maximum: 9, Sum: 24}
	result := MarshalStatistics(stats, false)

	if result != "minimum: 7\nmaximum: 9\nsum: 24\n" {
		test.Fail()
	}
}

func TestMarshalStatistics_withUnicode(test *testing.T) {
	stats := Statistics{Minimum: 7, Maximum: 9, Sum: 24}
	result := MarshalStatistics(stats, true)

	if result != "minimum: ⚅⚀\nmaximum: ⚅⚂\nsum: ⚄⚄ ⚄⚄ ⚃\n" {
		test.Fail()
	}
}

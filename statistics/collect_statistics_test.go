package statistics

import "testing"

func TestCollectStatistics(test *testing.T) {
	stats := CollectStatistics([]int{9, 7, 8})

	wantedStats := Statistics{Minimum: 7, Maximum: 9, Sum: 24}
	if stats != wantedStats {
		test.Fail()
	}
}

func Test_getMinimumValue(test *testing.T) {
	minimum := getMinimumValue([]int{9, 7, 8})

	if minimum != 7 {
		test.Fail()
	}
}

func Test_getMaximumValue(test *testing.T) {
	maximum := getMaximumValue([]int{7, 9, 8})

	if maximum != 9 {
		test.Fail()
	}
}

func Test_getSumOfValues(test *testing.T) {
	sum := getSumOfValues([]int{9, 7, 8})

	if sum != 24 {
		test.Fail()
	}
}

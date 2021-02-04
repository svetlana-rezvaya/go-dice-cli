package main

import "testing"

func Test_collectStatistics(test *testing.T) {
	stats := collectStatistics([]int{9, 7, 8})

	wantedStats := statistics{minimum: 7, maximum: 9, sum: 24}
	if stats != wantedStats {
		test.Fail()
	}
}

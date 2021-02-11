package main

import "testing"

func Test_marshalStatistics(test *testing.T) {
	stats := statistics{minimum: 7, maximum: 9, sum: 24}
	result := marshalStatistics(stats)

	if result != "minimum: 7\nmaximum: 9\nsum: 24\n" {
		test.Fail()
	}
}

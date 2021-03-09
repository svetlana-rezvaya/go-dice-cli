package main

import "testing"

func Test_marshalStatistics_withoutUnicode(test *testing.T) {
	stats := statistics{minimum: 7, maximum: 9, sum: 24}
	result := marshalStatistics(stats, false)

	if result != "minimum: 7\nmaximum: 9\nsum: 24\n" {
		test.Fail()
	}
}

func Test_marshalStatistics_withUnicode(test *testing.T) {
	stats := statistics{minimum: 7, maximum: 9, sum: 24}
	result := marshalStatistics(stats, true)

	if result != "minimum: ⚅⚀\nmaximum: ⚅⚂\nsum: ⚄⚄ ⚄⚄ ⚃\n" {
		test.Fail()
	}
}

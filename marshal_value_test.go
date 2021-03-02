package main

import (
	"testing"
)

func Test_toUnicodeDices(test *testing.T) {
	type data struct {
		value        int
		wantedResult string
	}

	tests := []data{
		data{value: 1, wantedResult: "⚀"},
		data{value: 2, wantedResult: "⚁"},
		data{value: 3, wantedResult: "⚂"},
		data{value: 4, wantedResult: "⚃"},
		data{value: 5, wantedResult: "⚄"},
		data{value: 6, wantedResult: "⚅"},
		data{value: 7, wantedResult: "⚅⚀"},
		data{value: 8, wantedResult: "⚅⚁"},
		data{value: 9, wantedResult: "⚅⚂"},
	}
	for _, testData := range tests {
		result := toUnicodeDices(testData.value)

		if result != testData.wantedResult {
			test.Logf("failed: %d", testData.value)
			test.Fail()
		}
	}
}

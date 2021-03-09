package main

import "testing"

func Test_marshalValues_withoutUnicode(test *testing.T) {
	result := marshalValues([]int{1, 2, 3}, false)

	if result != "1, 2, 3" {
		test.Fail()
	}
}

func Test_marshalValues_withUnicode(test *testing.T) {
	result := marshalValues([]int{1, 2, 3}, true)

	if result != "⚀, ⚁, ⚂" {
		test.Fail()
	}
}

package main

import "testing"

func Test_getMinimumValue(test *testing.T) {
	minimum := getMinimumValue([]int{9, 7, 8})

	if minimum != 7 {
		test.Fail()
	}
}

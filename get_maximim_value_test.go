package main

import "testing"

func Test_getMaximumValue(test *testing.T) {
	maximum := getMaximumValue([]int{7, 9, 8})

	if maximum != 9 {
		test.Fail()
	}
}

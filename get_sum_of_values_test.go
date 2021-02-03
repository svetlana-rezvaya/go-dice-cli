package main

import "testing"

func Test_getSumOfValues(test *testing.T) {
	sum := getSumOfValues([]int{9, 7, 8})

	if sum != 24 {
		test.Fail()
	}
}

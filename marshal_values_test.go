package main

import "testing"

func Test_marshalValues(test *testing.T) {
	result := marshalValues([]int{1, 2, 3})

	if result != "1, 2, 3" {
		test.Fail()
	}
}

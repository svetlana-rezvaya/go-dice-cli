package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_generateDiceThrows(test *testing.T) {
	rand.Seed(1)

	diceThrows := generateDiceThrows(3, 8)

	wantedDiceThrows := []int{2, 8, 8}
	if !reflect.DeepEqual(diceThrows, wantedDiceThrows) {
		test.Fail()
	}
}

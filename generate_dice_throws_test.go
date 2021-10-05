package dice

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerateDiceThrows(test *testing.T) {
	rand.Seed(1)

	diceThrows := GenerateDiceThrows(3, 8)

	wantedDiceThrows := []int{2, 8, 8}
	if !reflect.DeepEqual(diceThrows, wantedDiceThrows) {
		test.Fail()
	}
}

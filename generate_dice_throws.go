package dice

import "math/rand"

// GenerateDiceThrows ...
func GenerateDiceThrows(throwCount int, faceCount int) []int {
	diceThrows := []int{}
	for i := 0; i < throwCount; i = i + 1 {
		value := rand.Intn(faceCount) + 1
		diceThrows = append(diceThrows, value)
	}

	return diceThrows
}

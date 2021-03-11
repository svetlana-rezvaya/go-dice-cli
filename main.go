package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	options, err := parseFlags(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	diceThrows := generateDiceThrows(options.throwCount, options.faceCount)
	fmt.Println(marshalValues(diceThrows, options.useUnicode))

	stats := collectStatistics(diceThrows)
	fmt.Print(marshalStatistics(stats, options.useUnicode))
}

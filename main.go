package main

import (
	"flag"
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
		if err == flag.ErrHelp {
			fmt.Print(helpMessage)
			os.Exit(0)
		}
		log.Fatal(err)
	}

	diceThrows := generateDiceThrows(options.throwCount, options.faceCount)
	fmt.Println(marshalValues(diceThrows, options.useUnicode))

	stats := collectStatistics(diceThrows)
	fmt.Print(marshalStatistics(stats, options.useUnicode))
}

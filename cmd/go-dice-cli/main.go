package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	dice "github.com/svetlana-rezvaya/go-dice-cli"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	options, err := dice.ParseFlags(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			fmt.Print(dice.HelpMessage)
			os.Exit(0)
		}
		log.Fatal(err)
	}

	diceThrows := dice.GenerateDiceThrows(options.ThrowCount, options.FaceCount)
	fmt.Println(dice.MarshalValues(diceThrows, options.UseUnicode))

	stats := dice.CollectStatistics(diceThrows)
	fmt.Print(dice.MarshalStatistics(stats, options.UseUnicode))
}

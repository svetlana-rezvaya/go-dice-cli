package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const helpMessage = "" +
	"Usage:\n" +
	"  go-dice-cli -h | -help | --help\n" +
	"  go-dice-cli [-unicode] <dice>\n" +
	"  go-dice-cli [-unicode] -dice STRING\n" +
	"  go-dice-cli [-unicode] -throws INTEGER -faces INTEGER\n" +
	"\n" +
	"Options:\n" +
	"  -h, -help, --help  — show the help message and exit;\n" +
	"  -dice STRING       — number of throws and dice faces in the dice notation;\n" +
	"  -throws INTEGER    — number of throws;\n" +
	"  -faces INTEGER     — number of dice faces;\n" +
	"  -unicode           — use Unicode to show dices.\n" +
	"\n" +
	"Arguments:\n" +
	"  <dice>  — number of throws and dice faces in the dice notation.\n"

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

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	dice "github.com/svetlana-rezvaya/go-dice-cli"
	"github.com/svetlana-rezvaya/go-dice-cli/internal/cli"
	"github.com/svetlana-rezvaya/go-dice-cli/statistics"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	options, err := cli.ParseFlags(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			fmt.Print(cli.HelpMessage)
			os.Exit(0)
		}
		log.Fatal(err)
	}

	diceThrows := dice.GenerateDiceThrows(options.ThrowCount, options.FaceCount)
	fmt.Println(dice.MarshalValues(diceThrows, options.UseUnicode))

	stats := statistics.CollectStatistics(diceThrows)
	fmt.Print(statistics.MarshalStatistics(stats, options.UseUnicode))
}

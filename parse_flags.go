package dice

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
)

// HelpMessage ...
const HelpMessage = "" +
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

// Options ...
type Options struct {
	ThrowCount int
	FaceCount  int
	UseUnicode bool
}

// ParseFlags ...
func ParseFlags(arguments []string) (Options, error) {
	flags := flag.NewFlagSet("go-dice-cli", flag.ContinueOnError)
	flags.SetOutput(ioutil.Discard)

	diceFlag := flags.String("dice", "", "")
	throwsFlag := flags.Int("throws", 0, "")
	facesFlag := flags.Int("faces", 0, "")
	unicodeFlag := flags.Bool("unicode", false, "")
	err := flags.Parse(arguments)
	if err != nil {
		if err == flag.ErrHelp {
			return Options{}, err
		}
		return Options{}, fmt.Errorf("unable to parse the flags: %s", err)
	}

	// 1. as a string in a positional argument
	diceNotation := flags.Arg(0)
	if diceNotation != "" {
		throwCount, faceCount, err := parseDiceNotation(diceNotation)
		if err != nil {
			return Options{},
				fmt.Errorf("unable to parse the positional argument: %s", err)
		}

		options := Options{
			ThrowCount: throwCount,
			FaceCount:  faceCount,
			UseUnicode: *unicodeFlag,
		}
		return options, nil
	}

	// 2. as a string in the 'dice' flag
	diceNotation = *diceFlag
	if diceNotation != "" {
		throwCount, faceCount, err := parseDiceNotation(diceNotation)
		if err != nil {
			return Options{}, fmt.Errorf("unable to parse the 'dice' flag: %s", err)
		}

		options := Options{
			ThrowCount: throwCount,
			FaceCount:  faceCount,
			UseUnicode: *unicodeFlag,
		}
		return options, nil
	}

	// 3. as numbers in the 'throws' and 'faces' flags
	throwCount := *throwsFlag
	if throwCount == 0 {
		return Options{}, errors.New("the 'throws' flag is required")
	}

	faceCount := *facesFlag
	if faceCount == 0 {
		return Options{}, errors.New("the 'faces' flag is required")
	}

	options := Options{
		ThrowCount: throwCount,
		FaceCount:  faceCount,
		UseUnicode: *unicodeFlag,
	}
	return options, nil
}

package main

import (
	"errors"
	"flag"
	"io/ioutil"
)

type options struct {
	throwCount int
	faceCount  int
	useUnicode bool
}

func parseFlags(arguments []string) (options, error) {
	flags := flag.NewFlagSet("go-dice-cli", flag.ContinueOnError)
	flags.SetOutput(ioutil.Discard)

	diceFlag := flags.String("dice", "", "")
	throwsFlag := flags.Int("throws", 0, "")
	facesFlag := flags.Int("faces", 0, "")
	unicodeFlag := flags.Bool("unicode", false, "")
	err := flags.Parse(arguments)
	if err != nil {
		if err == flag.ErrHelp {
			return options{}, err
		}
		return options{}, errors.New("unable to parse the flags: " + err.Error())
	}

	// 1. as a string in a positional argument
	diceNotation := flags.Arg(0)
	if diceNotation != "" {
		throwCount, faceCount, err := parseDiceNotation(diceNotation)
		if err != nil {
			return options{},
				errors.New("unable to parse the positional argument: " + err.Error())
		}

		options := options{
			throwCount: throwCount,
			faceCount:  faceCount,
			useUnicode: *unicodeFlag,
		}
		return options, nil
	}

	// 2. as a string in the 'dice' flag
	diceNotation = *diceFlag
	if diceNotation != "" {
		throwCount, faceCount, err := parseDiceNotation(diceNotation)
		if err != nil {
			return options{},
				errors.New("unable to parse the 'dice' flag: " + err.Error())
		}

		options := options{
			throwCount: throwCount,
			faceCount:  faceCount,
			useUnicode: *unicodeFlag,
		}
		return options, nil
	}

	// 3. as numbers in the 'throws' and 'faces' flags
	throwCount := *throwsFlag
	if throwCount == 0 {
		return options{}, errors.New("the 'throws' flag is required")
	}

	faceCount := *facesFlag
	if faceCount == 0 {
		return options{}, errors.New("the 'faces' flag is required")
	}

	options := options{
		throwCount: throwCount,
		faceCount:  faceCount,
		useUnicode: *unicodeFlag,
	}
	return options, nil
}

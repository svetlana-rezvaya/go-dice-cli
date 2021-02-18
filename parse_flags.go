package main

import (
	"errors"
	"flag"
)

func parseFlags(arguments []string) (throwCount int, faceCount int, err error) {
	flags := flag.NewFlagSet("go-dice-cli", flag.ContinueOnError)
	diceFlag := flags.String("dice", "",
		"number of throws and dice faces in the dice notation; "+
			"this flag can be used as a positional argument")
	throwsFlag := flags.Int("throws", 0, "number of throws")
	facesFlag := flags.Int("faces", 0, "number of dice faces")
	err = flags.Parse(arguments)
	if err != nil {
		if err == flag.ErrHelp {
			return 0, 0, err
		}
		return 0, 0, errors.New("unable to parse the flags: " + err.Error())
	}

	// 1. as a string in a positional argument
	diceNotation := flags.Arg(0)
	if diceNotation != "" {
		throwCount, faceCount, err = parseDiceNotation(diceNotation)
		if err != nil {
			return 0, 0,
				errors.New("unable to parse the positional argument: " + err.Error())
		}
	}

	// 2. as a string in the 'dice' flag
	diceNotation = *diceFlag
	if diceNotation != "" {
		throwCount, faceCount, err = parseDiceNotation(diceNotation)
		if err != nil {
			return 0, 0, errors.New("unable to parse the 'dice' flag: " + err.Error())
		}
	}

	// 3. as numbers in the 'throws' and 'faces' flags
	throwCount = *throwsFlag
	if throwCount == 0 {
		return 0, 0, errors.New("the 'throws' flag is required")
	}

	faceCount = *facesFlag
	if faceCount == 0 {
		return 0, 0, errors.New("the 'faces' flag is required")
	}

	return throwCount, faceCount, nil
}

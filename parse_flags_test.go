package main

import (
	"flag"
	"testing"
)

func Test_parseFlags_successWithPositionalArgument(test *testing.T) {
	throwCount, faceCount, err := parseFlags([]string{"2d20"})

	if throwCount != 2 {
		test.Fail()
	}

	if faceCount != 20 {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_parseFlags_errorWithPositionalArgument(test *testing.T) {
	throwCount, faceCount, err := parseFlags([]string{"#d20"})

	if throwCount != 0 {
		test.Fail()
	}

	if faceCount != 0 {
		test.Fail()
	}

	wantedErrStr := "unable to parse the positional argument: " +
		"unable to parse the throw count: " +
		"strconv.Atoi: parsing \"#\": invalid syntax"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_parseFlags_successWithDiceFlag(test *testing.T) {
	throwCount, faceCount, err := parseFlags([]string{"-dice", "2d20"})

	if throwCount != 2 {
		test.Fail()
	}

	if faceCount != 20 {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_parseFlags_errorWithDiceFlag(test *testing.T) {
	throwCount, faceCount, err := parseFlags([]string{"-dice", "#d20"})

	if throwCount != 0 {
		test.Fail()
	}

	if faceCount != 0 {
		test.Fail()
	}

	wantedErrStr := "unable to parse the 'dice' flag: " +
		"unable to parse the throw count: " +
		"strconv.Atoi: parsing \"#\": invalid syntax"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_parseFlags_successWithThrowsAndFacesFlags(test *testing.T) {
	throwCount, faceCount, err :=
		parseFlags([]string{"-throws", "2", "-faces", "20"})

	if throwCount != 2 {
		test.Fail()
	}

	if faceCount != 20 {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_parseFlags_errorWithoutThrowsFlag(test *testing.T) {
	throwCount, faceCount, err := parseFlags([]string{"-faces", "20"})

	if throwCount != 0 {
		test.Fail()
	}

	if faceCount != 0 {
		test.Fail()
	}

	if err == nil || err.Error() != "the 'throws' flag is required" {
		test.Fail()
	}
}

func Test_parseFlags_errorWithoutFacesFlag(test *testing.T) {
	throwCount, faceCount, err := parseFlags([]string{"-throws", "2"})

	if throwCount != 0 {
		test.Fail()
	}

	if faceCount != 0 {
		test.Fail()
	}

	if err == nil || err.Error() != "the 'faces' flag is required" {
		test.Fail()
	}
}

func Test_parseFlags_errorWithIncorrectFlag(test *testing.T) {
	throwCount, faceCount, err := parseFlags([]string{"-incorrect"})

	if throwCount != 0 {
		test.Fail()
	}

	if faceCount != 0 {
		test.Fail()
	}

	wantedErrStr := "unable to parse the flags: " +
		"flag provided but not defined: -incorrect"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_parseFlags_errorWithHelpFlag(test *testing.T) {
	throwCount, faceCount, err := parseFlags([]string{"-help"})

	if throwCount != 0 {
		test.Fail()
	}

	if faceCount != 0 {
		test.Fail()
	}

	if err != flag.ErrHelp {
		test.Fail()
	}
}

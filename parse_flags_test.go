package main

import (
	"flag"
	"testing"
)

func Test_parseFlags_successWithPositionalArgument(test *testing.T) {
	receivedOptions, err := parseFlags([]string{"2d20"})

	wantedOptions := options{throwCount: 2, faceCount: 20, useUnicode: false}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_parseFlags_successWithPositionalArgumentAndUnicode(test *testing.T) {
	receivedOptions, err := parseFlags([]string{"-unicode", "2d20"})

	wantedOptions := options{throwCount: 2, faceCount: 20, useUnicode: true}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_parseFlags_errorWithPositionalArgument(test *testing.T) {
	receivedOptions, err := parseFlags([]string{"#d20"})

	wantedOptions := options{throwCount: 0, faceCount: 0, useUnicode: false}
	if receivedOptions != wantedOptions {
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
	receivedOptions, err := parseFlags([]string{"-dice", "2d20"})

	wantedOptions := options{throwCount: 2, faceCount: 20}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_parseFlags_errorWithDiceFlag(test *testing.T) {
	receivedOptions, err := parseFlags([]string{"-dice", "#d20"})

	wantedOptions := options{throwCount: 0, faceCount: 0}
	if receivedOptions != wantedOptions {
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
	receivedOptions, err := parseFlags([]string{"-throws", "2", "-faces", "20"})

	wantedOptions := options{throwCount: 2, faceCount: 20}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func Test_parseFlags_errorWithoutThrowsFlag(test *testing.T) {
	receivedOptions, err := parseFlags([]string{"-faces", "20"})

	wantedOptions := options{throwCount: 0, faceCount: 0}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err == nil || err.Error() != "the 'throws' flag is required" {
		test.Fail()
	}
}

func Test_parseFlags_errorWithoutFacesFlag(test *testing.T) {
	receivedOptions, err := parseFlags([]string{"-throws", "2"})

	wantedOptions := options{throwCount: 0, faceCount: 0}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err == nil || err.Error() != "the 'faces' flag is required" {
		test.Fail()
	}
}

func Test_parseFlags_errorWithIncorrectFlag(test *testing.T) {
	receivedOptions, err := parseFlags([]string{"-incorrect"})

	wantedOptions := options{throwCount: 0, faceCount: 0}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	wantedErrStr := "unable to parse the flags: " +
		"flag provided but not defined: -incorrect"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_parseFlags_errorWithHelpFlag(test *testing.T) {
	receivedOptions, err := parseFlags([]string{"-help"})

	wantedOptions := options{throwCount: 0, faceCount: 0}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err != flag.ErrHelp {
		test.Fail()
	}
}

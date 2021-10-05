package dice

import (
	"flag"
	"testing"
)

func TestParseFlags_successWithPositionalArgument(test *testing.T) {
	receivedOptions, err := ParseFlags([]string{"2d20"})

	wantedOptions := Options{ThrowCount: 2, FaceCount: 20, UseUnicode: false}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestParseFlags_successWithPositionalArgumentAndUnicode(test *testing.T) {
	receivedOptions, err := ParseFlags([]string{"-unicode", "2d20"})

	wantedOptions := Options{ThrowCount: 2, FaceCount: 20, UseUnicode: true}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestParseFlags_errorWithPositionalArgument(test *testing.T) {
	receivedOptions, err := ParseFlags([]string{"#d20"})

	wantedOptions := Options{ThrowCount: 0, FaceCount: 0, UseUnicode: false}
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

func TestParseFlags_successWithDiceFlag(test *testing.T) {
	receivedOptions, err := ParseFlags([]string{"-dice", "2d20"})

	wantedOptions := Options{ThrowCount: 2, FaceCount: 20}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestParseFlags_errorWithDiceFlag(test *testing.T) {
	receivedOptions, err := ParseFlags([]string{"-dice", "#d20"})

	wantedOptions := Options{ThrowCount: 0, FaceCount: 0}
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

func TestParseFlags_successWithThrowsAndFacesFlags(test *testing.T) {
	receivedOptions, err := ParseFlags([]string{"-throws", "2", "-faces", "20"})

	wantedOptions := Options{ThrowCount: 2, FaceCount: 20}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err != nil {
		test.Fail()
	}
}

func TestParseFlags_errorWithoutThrowsFlag(test *testing.T) {
	receivedOptions, err := ParseFlags([]string{"-faces", "20"})

	wantedOptions := Options{ThrowCount: 0, FaceCount: 0}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err == nil || err.Error() != "the 'throws' flag is required" {
		test.Fail()
	}
}

func TestParseFlags_errorWithoutFacesFlag(test *testing.T) {
	receivedOptions, err := ParseFlags([]string{"-throws", "2"})

	wantedOptions := Options{ThrowCount: 0, FaceCount: 0}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err == nil || err.Error() != "the 'faces' flag is required" {
		test.Fail()
	}
}

func TestParseFlags_errorWithIncorrectFlag(test *testing.T) {
	receivedOptions, err := ParseFlags([]string{"-incorrect"})

	wantedOptions := Options{ThrowCount: 0, FaceCount: 0}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	wantedErrStr := "unable to parse the flags: " +
		"flag provided but not defined: -incorrect"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func TestParseFlags_errorWithHelpFlag(test *testing.T) {
	receivedOptions, err := ParseFlags([]string{"-help"})

	wantedOptions := Options{ThrowCount: 0, FaceCount: 0}
	if receivedOptions != wantedOptions {
		test.Fail()
	}

	if err != flag.ErrHelp {
		test.Fail()
	}
}

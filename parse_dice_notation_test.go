package main

import "testing"

func Test_parseDiceNotation_successWithFullData(test *testing.T) {
	throwCount, faceCount, err := parseDiceNotation("2d20")

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

func Test_parseDiceNotation_successWithShortData(test *testing.T) {
	throwCount, faceCount, err := parseDiceNotation("d20")

	if throwCount != 1 {
		test.Fail()
	}
	if faceCount != 20 {
		test.Fail()
	}
	if err != nil {
		test.Fail()
	}
}

func Test_parseDiceNotation_errorWithLessParts(test *testing.T) {
	throwCount, faceCount, err := parseDiceNotation("2")

	if throwCount != 0 {
		test.Fail()
	}
	if faceCount != 0 {
		test.Fail()
	}
	if err == nil || err.Error() != "invalid count of parts" {
		test.Fail()
	}
}

func Test_parseDiceNotation_errorWithMoreParts(test *testing.T) {
	throwCount, faceCount, err := parseDiceNotation("2d20d200")

	if throwCount != 0 {
		test.Fail()
	}
	if faceCount != 0 {
		test.Fail()
	}
	if err == nil || err.Error() != "invalid count of parts" {
		test.Fail()
	}
}

func Test_parseDiceNotation_errorWithIncorrectThrowCount(test *testing.T) {
	throwCount, faceCount, err := parseDiceNotation("#d20")

	if throwCount != 0 {
		test.Fail()
	}

	if faceCount != 0 {
		test.Fail()
	}

	wantedErrStr := "unable to parse the throw count: " +
		"strconv.Atoi: parsing \"#\": invalid syntax"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

func Test_parseDiceNotation_errorWithIncorrectFaceCount(test *testing.T) {
	throwCount, faceCount, err := parseDiceNotation("2d##")

	if throwCount != 0 {
		test.Fail()
	}

	if faceCount != 0 {
		test.Fail()
	}

	wantedErrStr := "unable to parse the face count: " +
		"strconv.Atoi: parsing \"##\": invalid syntax"
	if err == nil || err.Error() != wantedErrStr {
		test.Fail()
	}
}

package dice

import "testing"

func TestParseDiceNotation_successWithFullData(test *testing.T) {
	throwCount, faceCount, err := ParseDiceNotation("2d20")

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

func TestParseDiceNotation_successWithShortData(test *testing.T) {
	throwCount, faceCount, err := ParseDiceNotation("d20")

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

func TestParseDiceNotation_errorWithLessParts(test *testing.T) {
	throwCount, faceCount, err := ParseDiceNotation("2")

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

func TestParseDiceNotation_errorWithMoreParts(test *testing.T) {
	throwCount, faceCount, err := ParseDiceNotation("2d20d200")

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

func TestParseDiceNotation_errorWithIncorrectThrowCount(test *testing.T) {
	throwCount, faceCount, err := ParseDiceNotation("#d20")

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

func TestParseDiceNotation_errorWithIncorrectFaceCount(test *testing.T) {
	throwCount, faceCount, err := ParseDiceNotation("2d##")

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

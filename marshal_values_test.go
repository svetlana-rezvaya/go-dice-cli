package dice

import "testing"

func TestMarshalValues_withoutUnicode(test *testing.T) {
	result := MarshalValues([]int{1, 2, 3}, false)

	if result != "1, 2, 3" {
		test.Fail()
	}
}

func TestMarshalValues_withUnicode(test *testing.T) {
	result := MarshalValues([]int{1, 2, 3}, true)

	if result != "⚀, ⚁, ⚂" {
		test.Fail()
	}
}

func TestMarshalValue(test *testing.T) {
	type data struct {
		value        int
		useUnicode   bool
		wantedResult string
	}

	tests := []data{
		data{value: 28, useUnicode: false, wantedResult: "28"},
		data{value: 28, useUnicode: true, wantedResult: "⚄⚄ ⚄⚄ ⚅⚁"},
	}
	for _, testData := range tests {
		result := MarshalValue(testData.value, testData.useUnicode)

		if result != testData.wantedResult {
			test.Logf("failed: %d %t", testData.value, testData.useUnicode)
			test.Fail()
		}
	}
}

func Test_toUnicodeDices(test *testing.T) {
	type data struct {
		value        int
		wantedResult string
	}

	tests := []data{
		data{value: 1, wantedResult: "⚀"},
		data{value: 2, wantedResult: "⚁"},
		data{value: 3, wantedResult: "⚂"},
		data{value: 4, wantedResult: "⚃"},
		data{value: 5, wantedResult: "⚄"},
		data{value: 6, wantedResult: "⚅"},
		data{value: 7, wantedResult: "⚅⚀"},
		data{value: 8, wantedResult: "⚅⚁"},
		data{value: 9, wantedResult: "⚅⚂"},
		data{value: 10, wantedResult: "⚄⚄"},
		data{value: 12, wantedResult: "⚄⚄ ⚁"},
		data{value: 18, wantedResult: "⚄⚄ ⚅⚁"},
		data{value: 20, wantedResult: "⚄⚄ ⚄⚄"},
		data{value: 22, wantedResult: "⚄⚄ ⚄⚄ ⚁"},
		data{value: 28, wantedResult: "⚄⚄ ⚄⚄ ⚅⚁"},
	}
	for _, testData := range tests {
		result := toUnicodeDices(testData.value)

		if result != testData.wantedResult {
			test.Logf("failed: %d", testData.value)
			test.Fail()
		}
	}
}

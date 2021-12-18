package customerimporter

import (
	"reflect"
	"testing"
)

type TestCases struct {
	inputFile string
	expected  map[string]int
}

var testCases = []TestCases{
	{inputFile: "testdata/test-1.csv", expected: map[string]int{"github.io": 1}},
}

func TestReadFile(t *testing.T) {
	for _, test := range testCases {
		result, err := ImportCustomers(test.inputFile)
		if err != nil {
			t.Fatal("Expected Result not given")
		}

		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Expected Domains \n%v\n did not equal actual Domains \n%v\n", result, test.expected)
		}

	}
}

package customerimporter

import (
	"reflect"
	"testing"
)

type TestCases struct {
	inputFile     string
	expected      map[string]int
	expectedError bool
}

var testCases = []TestCases{
	{inputFile: "testdata/single-customer.csv", expected: map[string]int{"github.io": 1}},
	{inputFile: "testdata/multiple-customers.csv", expected: map[string]int{"github.io": 1, "cyberchimps.com": 1, "hubpages.com": 1}},
	{inputFile: "testdata/multiple-customers-single-domain.csv", expected: map[string]int{"github.io": 3}},

	// Unhappy Path
	{inputFile: "testdata/this-file-does-not-exist.csv", expectedError: true},
}

func TestReadFile(t *testing.T) {
	for _, test := range testCases {
		result, err := ImportCustomers(test.inputFile)
		if test.expectedError && err == nil {
			t.Fatal("Expected Error to be returned")
		}
		if err != nil {
			t.Fatalf("Did not expect to receive error: \n%v\n", err)
		}

		if !reflect.DeepEqual(result, test.expected) {
			t.Fatalf("Expected Domains \n%v\n did not equal actual Domains \n%v\n", result, test.expected)
		}

	}
}

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
	{inputFile: "testdata/single-customer.csv", expected: map[string]int{"github.io": 1}},
	{inputFile: "testdata/multiple-customers.csv", expected: map[string]int{"github.io": 1, "cyberchimps.com": 1, "hubpages.com": 1}},
	{inputFile: "testdata/multiple-customers-single-domain.csv", expected: map[string]int{"github.io": 3}},
	{inputFile: "testdata/multiple-customers-single-domain.csv", expected: map[string]int{"github.io": 3}},
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

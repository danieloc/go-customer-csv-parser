package main

import (
	"fmt"

	"github.com/danieloc/go-customer-csv/customerimporter"
)

func main() {
	domains, err := customerimporter.ImportCustomers("customers.csv")
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}
	fmt.Println(domains)
}

package main

import (
	"fmt"

	"github.com/danieloc/go-customer-csv/customerimporter"
)

func main() {
	domains := customerimporter.ImportCustomers()
	fmt.Println(domains)
}

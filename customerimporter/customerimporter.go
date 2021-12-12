// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Customer struct {
	first_name string
	last_name  string
	email      string
	gender     string
	ip_address string
}

const (
	first_name = iota
	last_name
	email
	gender
	ip_address
)

func ImportCustomers() {
	file, err := os.Open("customers.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(file))
	//Discard the header
	_, headerErr := reader.Read()
	if err != nil {
		log.Fatal(headerErr)
	}
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(Customer{
			first_name: row[first_name],
			last_name:  row[last_name],
			email:      row[email],
			gender:     row[gender],
			ip_address: row[ip_address],
		})
	}
}

// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Customer struct {
	email string
}

const (
	first_name = iota
	last_name
	email
	gender
	ip_address
)

func getDomain(email string) (string, error) {
	components := strings.Split(email, "@")
	if len(components) != 2 {
		return "", errors.New("Missing domain")
	}
	_, domain := components[0], components[1]
	return domain, nil
}

func recordDomain(domain string, allDomains map[string]int) {
	allDomains[domain]++

}

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

	allDomains := make(map[string]int)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		cust := Customer{
			email: row[email],
		}

		customerDomain, err := getDomain(cust.email)
		if err != nil {
			fmt.Println("Missing domain in email", cust.email)
			continue
		}

		recordDomain(customerDomain, allDomains)
	}
}

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Change(s string, prog string, version string) string {
	// Parse the data
	data := strings.Split(s, "\n")

	// Get the Version
	ver := data[5][9:]

	// Get the phone
	phone := data[3][7:]

	verMatched, _ := regexp.MatchString("^[0-9]+\\.[0-9]+$", ver)
	phoneMatched, _ := regexp.MatchString("^\\+1-[0-9]{3}-[0-9]{3}-[0-9]{4}$", phone)

	if ! verMatched || ! phoneMatched {
		return "ERROR: VERSION or PHONE"
	} else {
		if ver != "2.0" {
			ver = version
		}

		return fmt.Sprintf("Program: %v Author: g964 Phone: +1-503-555-0090 Date: 2019-01-01 Version: %v", prog, ver)
	}
}

func main() {
	s1 := "Program title: Primes\nAuthor: Kern\nCorporation: Gold\nPhone: +1-503-555-0091\nDate: Tues April 9, 2005\nVersion: 6.7\nLevel: Alpha"
	s11 := "Program title: Hammer\nAuthor: Tolkien\nCorporation: IB\nPhone: +1-503-555-0090\nDate: Tues March 29, 2017\nVersion: 2.0\nLevel: Release"
	s12 := "Program title: Primes\nAuthor: Kern\nCorporation: Gold\nPhone: +1-503-555-009\nDate: Tues April 9, 2005\nVersion: 6.7\nLevel: Alpha"

	fmt.Println(Change(s1, "Ladder", "1.1"))
	fmt.Println(Change(s11, "Balance", "1.5.6"))
	fmt.Println(Change(s12, "Ladder", "1.1"))
}

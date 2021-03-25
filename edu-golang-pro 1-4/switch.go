package main

import (
	"fmt"
	"regexp"
)

func switchRegExp() {
	asString := "mike@g.com"

	var negative = regexp.MustCompile(`-`)
	var floatingPoint = regexp.MustCompile(`\d?\.\d`)
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)

	switch {
	case negative.MatchString(asString):
		fmt.Println("Negative number")
	case floatingPoint.MatchString(asString):
		fmt.Println("Floating point!")
	case email.MatchString(asString):
		fmt.Println("It is an email!")
		fallthrough
	default:
		fmt.Println("Something else!")
	}
}

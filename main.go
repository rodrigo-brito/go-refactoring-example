package main

import (
	"example/lib"
	"fmt"
	"time"
)

type Validator struct {}

func (*Validator) IsValidNote(note float64) bool {
	return note >= 0 && note <= 100
}

func IsValidYear(year int) bool {
	return year > 1900 && year <=time.Now().Year()
}

func main() {
	validator := new(Validator)
	lvalidator := new(lib.Validator)
	fmt.Println("Valid 10 = ", lvalidator.IsValidMoney(10))
	fmt.Println("Valid -1 = ", lvalidator.IsValidMoney(-1))
	fmt.Println("Valid 0 = ", lvalidator.IsValidMoney(0))
	fmt.Println("Valid tomorrow = ", lib.IsValidDate(time.Now().AddDate(0,0,1)))
	fmt.Println("Valid yesterday = ", lib.IsValidDate(time.Now().AddDate(0,0,-1)))
	fmt.Println("Valid year = ", IsValidYear(2020))
	fmt.Println("Valid note = ", validator.IsValidNote(99.9))
}

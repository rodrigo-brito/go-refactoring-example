package main

import (
	"example/lib"
	"fmt"
	"time"
)

func main() {
	validator := new(lib.Validator)
	fmt.Println("Valid 10 = ", validator.IsValidCash(10))
	fmt.Println("Valid -1 = ", validator.IsValidCash(-1))
	fmt.Println("Valid 0 = ", validator.IsValidCash(0))
	fmt.Println("Valid tomorrow = ", validator.IsValidDate(time.Now().AddDate(0,0,1)))
	fmt.Println("Valid yesterday = ", validator.IsValidDate(time.Now().AddDate(0,0,-1)))
	fmt.Println("Valid year = ", validator.IsValidYear(2020))
	fmt.Println("Valid note = ", validator.IsValidNotes(99.9))
}

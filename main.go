package main

import (
	"example/lib"
	"fmt"
	"time"
)

func main() {
	validator := new(lib.Validator)
	fmt.Println("Valid 10 = ", validator.IsValidMoney(10))
	fmt.Println("Valid -1 = ", validator.IsValidMoney(-1))
	fmt.Println("Valid 0 = ", validator.IsValidMoney(0))
	fmt.Println("Valid tomorrow = ", lib.IsValidDate(time.Now().AddDate(0,0,1)))
	fmt.Println("Valid yesterday = ", lib.IsValidDate(time.Now().AddDate(0,0,-1)))
}

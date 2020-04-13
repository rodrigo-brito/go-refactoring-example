package main

import (
	"example/lib"
	"fmt"
	"time"
)

func IsValidDate(date time.Time) bool {
	today := time.Now()
	return date.After(today)
}

func main() {
	fmt.Println("Valid 10 = ", lib.IsValidMoney(10))
	fmt.Println("Valid -1 = ", lib.IsValidMoney(-1))
	fmt.Println("Valid 0 = ", lib.IsValidMoney(0))
	fmt.Println("Valid tomorrow = ", IsValidDate(time.Now().AddDate(0,0,1)))
	fmt.Println("Valid yesterday = ", IsValidDate(time.Now().AddDate(0,0,-1)))
}

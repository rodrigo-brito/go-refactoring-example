package lib

import (
	"fmt"
	"time"
)

type validator struct {}

type Validator interface {
	IsValidDate(date time.Time) bool
	IsValidNotes(note float64) bool
	IsValidCash(value float64) bool
	IsValidYear(year int) bool
}

func NewValidator() Validator {
	return new(validator)
}

type TimeValidate interface {
	IsValidDate(date time.Time) bool
}

func (validator) IsValidDate(date time.Time) bool {
	today := time.Now()
	return date.After(today)
}

func (validator) IsValidNotes(note float64) bool {
	return note >= 0 && note <= 100
}

func (validator) IsValidCash(value float64) bool {
	return value > 0
}

func  (validator) IsValidYear(year int) bool {
	return year > 1900 && year <=time.Now().Year()
}

func PrintType(value int) {
	if (value % 2 == 0) {
		fmt.Println("Odd")
	} else {
		fmt.Println("even")
	}
}
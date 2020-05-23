package lib

import (
	"fmt"
	"time"
)

type validator struct{}

type Validator interface {
	IsValidDate(date time.Time) bool
}

func NewValidator() Validator {
	return new(validator)
}

type TimeValidate interface {
	IsValidDate(date time.Time) bool
}

func IsValidYear(year int) bool {
	return year > 1900 && year <= time.Now().Year()
}

func (validator) IsValidDate(date time.Time) bool {
	today := time.Now()
	return date.After(today)
}

func isEven(value int) {
	if value%2 == 0 {
		fmt.Println("Odd")
	} else {
		fmt.Println("even")
	}
}

func PrintType(value int) {
	isEven(value)
}

package lib

import "time"

type Validator struct {}

type TimeValidate interface {
	IsValidDate(date time.Time) bool
}

func (Validator) IsValidDate(date time.Time) bool {
	today := time.Now()
	return date.After(today)
}

func (Validator) IsValidNotes(note float64) bool {
	return note >= 0 && note <= 100
}

func (Validator) IsValidCash(value float64) bool {
	return value > 0
}

func  (Validator) IsValidYear(year int) bool {
	return year > 1900 && year <=time.Now().Year()
}
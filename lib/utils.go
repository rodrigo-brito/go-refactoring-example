package lib

import "time"

type Validator struct {}

func (v *Validator) IsValidMoney(value float64) bool {
	return value > 0
}

func IsValidDate(date time.Time) bool {
	today := time.Now()
	return date.After(today)
}
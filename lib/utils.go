package lib

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func ValidateCPF(cpf string) error {
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)
	if len(cpf) != 11 {
		return errors.New("CPF inválido")
	}
	var eq bool
	var dig string
	for _, val := range cpf {
		if len(dig) == 0 {
			dig = string(val)
		}
		if string(val) == dig {
			eq = true
			continue
		}
		eq = false
		break
	}
	if eq {
		return errors.New("CPF inválido")
	}

	i := 10
	sum := 0
	for index := 0; index < len(cpf)-2; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}

	prod := sum * 10
	mod := prod % 11
	if mod == 10 {
		mod = 0
	}
	digit1, _ := strconv.Atoi(string(cpf[9]))
	if mod != digit1 {
		return errors.New("CPF inválido")
	}
	i = 11
	sum = 0
	for index := 0; index < len(cpf)-1; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}
	prod = sum * 10
	mod = prod % 11
	if mod == 10 {
		mod = 0
	}
	digit2, _ := strconv.Atoi(string(cpf[10]))
	if mod != digit2 {
		return errors.New("CPF inválido")
	}

	return nil
}

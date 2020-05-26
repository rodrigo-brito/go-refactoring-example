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

func validateDigits(cpf string) error {
	var eq bool
	var digits string
	for _, val := range cpf {
		if len(digits) == 0 {
			digits = string(val)
		}
		if string(val) == digits {
			eq = true
			continue
		}
		eq = false
		break
	}
	if eq {
		return errors.New("CPF inválido")
	}
	return nil
}

func ValidateCPF(cpf string) error {
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)
	if len(cpf) != 11 {
		return errors.New("CPF inválido")
	}

	err := validateDigits(cpf)
	if err != nil {
		return err
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

func ValidateCNPJ(cnpj string) error {
	cnpj = strings.Replace(cnpj, ".", "", -1)
	cnpj = strings.Replace(cnpj, "-", "", -1)
	cnpj = strings.Replace(cnpj, "/", "", -1)
	if len(cnpj) != 14 {
		return errors.New("invalid CNPJ")
	}

	algs := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var algProdCpfDig1 = make([]int, 12, 12)
	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(cnpj[key]))
		tmp := val * intParsed
		algProdCpfDig1[key] = tmp
	}
	sum := 0
	for _, val := range algProdCpfDig1 {
		sum += val
	}
	digit1 := sum % 11
	if digit1 < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - digit1
	}

	char12, _ := strconv.Atoi(string(cnpj[12]))
	if char12 != digit1 {
		return errors.New("invalid CPNJ")
	}
	algs = append([]int{6}, algs...)

	var algProdCpfDig2 = make([]int, 13, 13)
	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(cnpj[key]))

		tmp := val * intParsed
		algProdCpfDig2[key] = tmp
	}
	sum = 0
	for _, val := range algProdCpfDig2 {
		sum += val
	}

	digit2 := sum % 11
	if digit2 < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - digit2
	}

	char13, _ := strconv.Atoi(string(cnpj[13]))
	if char13 != digit2 {
		return errors.New("invalid CNPJ")
	}

	return nil
}

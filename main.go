package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"example/lib"
)

func main() {
	validator := lib.NewValidator()
	fmt.Println("Valid tomorrow = ", validator.IsValidDate(time.Now().AddDate(0, 0, 1)))
	fmt.Println("Valid yesterday = ", validator.IsValidDate(time.Now().AddDate(0, 0, -1)))
}

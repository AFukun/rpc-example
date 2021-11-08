package calculator

import (
	"errors"
	"strconv"
	"strings"
)

func parseMessage(message string) (c []string) {
	eString := string(message[:])
	return strings.Fields(eString)
}

func parseArgs(c []string) (float64, float64, error) {
	num1, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	num2, err := strconv.ParseFloat(c[2], 64)
	if err != nil {
		return 0.0, 0.0, err
	}
	return num1, num2, nil
}

func Calc(equation string) (float64, error) {
	result := 0.0
	e := parseMessage(equation)
	num1, num2, err := parseArgs(e)
	if err != nil {
		return 0.0, err
	}
	switch e[1] {
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0.0 {
			return 0.0, errors.New("error: you tried to divide by zero")
		}
		result = num1 / num2
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	}
	return result, nil
}

// Package calculator provides a library for simple
// calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding
// them together.
func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (result float64, err error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("Cannot take square root of a negative number: %f", num)
	}
	return math.Sqrt(num), nil
}

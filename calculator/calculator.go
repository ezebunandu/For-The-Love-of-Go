// Package calculator does simple calculations
package calculator

import (
	"errors"
	"math"
)

// Add takes two integers and returns the sum of them
func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("cannot take sqrt of a negative number")
	}
	return math.Sqrt(a), nil
}

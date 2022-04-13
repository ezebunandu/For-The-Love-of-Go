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
func AddMany(inputs ...float64) float64 {
	var sum float64 = 0
	for _, input := range inputs {
		sum += input
	}
	return sum
}

func SubtractMany(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	result := inputs[0]
	for _, input := range inputs[1:] {
		result -= input
	}
	return result

}

func MultiplyMany(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	var multiple float64 = 1
	for _, input := range inputs {
		multiple *= input
	}
	return multiple
}

func DivideMany(inputs ...float64) (float64, error) {
	if len(inputs) == 0 {
		return 0, nil
	}
	result := inputs[0]
	for _, input := range inputs[1:] {
		if input == 0 {
			return 0, errors.New("cannot divide by 0")
		}
		result /= input
	}
	return result, nil
}

func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("Cannot take square root of a negative number: %f", num)
	}
	return math.Sqrt(num), nil
}

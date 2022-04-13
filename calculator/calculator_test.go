package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAddMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs   []float64
		expected float64
	}
	testCases := []testCase{
		{inputs: []float64{}, expected: 0},
		{inputs: []float64{2}, expected: 2},
		{inputs: []float64{1, 2}, expected: 3},
		{inputs: []float64{1, 2, 3}, expected: 6},
	}
	for _, testcase := range testCases {
		got := calculator.AddMany(testcase.inputs...)
		if testcase.expected != got {
			t.Errorf("Add(%v): expected %f, got %f", testcase.inputs, testcase.expected, got)
		}
	}
}

func TestSubtractMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs   []float64
		expected float64
	}
	testCases := []testCase{
		{inputs: []float64{}, expected: 0},
		{inputs: []float64{2}, expected: 2},
		{inputs: []float64{2, 2}, expected: 0},
		{inputs: []float64{1, 2, 3}, expected: -4},
	}
	for _, testcase := range testCases {
		got := calculator.SubtractMany(testcase.inputs...)
		if testcase.expected != got {
			t.Errorf("Subtract(%v): expected %f, got %f", testcase.inputs, testcase.expected, got)
		}
	}
}

func TestMultiplyMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs   []float64
		expected float64
	}
	testCases := []testCase{
		{inputs: []float64{}, expected: 0},
		{inputs: []float64{2}, expected: 2},
		{inputs: []float64{1, 2}, expected: 2},
		{inputs: []float64{2, 3, 4}, expected: 24},
		{inputs: []float64{-1, 2, -1}, expected: 2},
	}
	for _, testcase := range testCases {
		got := calculator.MultiplyMany(testcase.inputs...)
		if testcase.expected != got {
			t.Errorf("Multiply(%v): expected %f, got %f", testcase.inputs, testcase.expected, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b, expected float64
	}
	testCases := []testCase{
		{a: 4, b: 2, expected: 2},
		{a: 0, b: 5, expected: 0},
		{a: 5, b: -1, expected: -5},
		{a: -2, b: -1, expected: 2},
	}
	for _, testcase := range testCases {
		got, err := calculator.Divide(testcase.a, testcase.b)
		if err != nil {
			t.Fatalf("Divide(%f, %f): want no error for valid input, got %v", testcase.a, testcase.b, err)
		}
		if testcase.expected != got {
			t.Errorf("Add(%f, %f): expected %f, got %f", testcase.a, testcase.b, testcase.expected, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("Divide(1, 0): want error for invalid input, got nil")
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestSqrt(t *testing.T) {
	type sqrtTestCase struct {
		num, squareRoot float64
	}
	t.Parallel()
	testCases := []sqrtTestCase{
		{num: 4, squareRoot: 2},
		{num: 1, squareRoot: 1},
		{num: 0, squareRoot: 0},
		{num: 2, squareRoot: 1.4142135},
	}
	for _, testcase := range testCases {
		got, err := calculator.Sqrt(testcase.num)
		if err != nil {
			t.Fatalf("Sqrt(%f): want no error for valid input, got %v", testcase.num, err)
		}
		if !closeEnough(got, testcase.squareRoot, 0.0001) {
			t.Errorf("Sqrt(%f): expected %f, got %f", testcase.num, testcase.squareRoot, got)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Sqrt(-1)
	if err == nil {
		t.Fatalf("Sqrt(-1): want error for squareroot of a negative number, got nil")
	}
}

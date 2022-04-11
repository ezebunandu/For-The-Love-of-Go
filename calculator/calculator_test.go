package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type testCase struct {
	a, b, expected float64
}

type sqrtTestCase struct {
	num, squareRoot float64
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, expected: 4},
		{a: 2, b: -3, expected: -1},
		{a: 5, b: 0, expected: 5},
	}
	for _, testcase := range testCases {
		got := calculator.Add(testcase.a, testcase.b)
		if testcase.expected != got {
			t.Errorf("Add(%f, %f): expected %f, got %f", testcase.a, testcase.b, testcase.expected, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 4, b: 2, expected: 2},
		{a: 1, b: 1, expected: 0},
		{a: 2, b: 4, expected: -2},
	}
	for _, testcase := range testCases {
		got := calculator.Subtract(testcase.a, testcase.b)
		if testcase.expected != got {
			t.Errorf("Subtract(%f, %f): expected %f, got %f", testcase.a, testcase.b, testcase.expected, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 4, b: 2, expected: 8},
		{a: 2, b: 0, expected: 0},
		{a: 2, b: -1, expected: -2},
	}
	for _, testcase := range testCases {
		got := calculator.Multiply(testcase.a, testcase.b)
		if testcase.expected != got {
			t.Errorf("Multiply(%f, %f): expected %f, got %f", testcase.a, testcase.b, testcase.expected, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
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

package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 1, b: 0, want: 1},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("wanted: %f, got: %f", tc.want, got)

		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 2, b: 1, want: 1},
		{a: 2, b: 0, want: 2},
		{a: 1, b: -1, want: 2},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("wanted: %f, got: %f", tc.want, got)
		}
	}
}
func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 0, b: 1, want: 0},
		{a: 2, b: 1, want: 2},
		{a: 1, b: -1, want: -1},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("wanted: %f, got: %f", tc.want, got)
		}
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 5, b: 2, want: 2.5},
		{a: 1, b: -1, want: -1},
		{a: -1, b: -1, want: 1},
		{a: 1, b: 3, want: 0.333333},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f, %f): want %f, got: %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(5, 0)
	if err == nil {
		t.Error("wanted err but got none")
	}

}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    float64
		want float64
	}
	testCases := []testCase{
		{a: 25, want: 5},
		{a: 0, want: 0},
		{a: 5, want: 2.230607},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if !closeEnough(tc.want, got, 0.01) {
			t.Errorf("Sqrt(%f): want %f, got: %f", tc.a, tc.want, got)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Sqrt(-5)
	if err == nil {
		t.Error("wanted err but got none")
	}

}

package mathlib

import (
	"fmt"
	"math"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero, HAHAHA")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("cannot calculate square root of a negative number")
	}
	return math.Sqrt(a), nil
}

func Square(a int) int {
	return a * a
}

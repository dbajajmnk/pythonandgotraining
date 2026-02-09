package calculator

import "errors"

var ErrDivideByZero = errors.New("divide by zero")

func Add(a, b int) int { return a + b }

/**
2+2 
0,2
-2,3 
-2,-3

**/

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

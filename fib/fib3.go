package fib

import (
	"errors"
)

var memo = map[int]int{0: 0, 1: 1}

func Fib3(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("Fib2 function needs non negative input.")
	} else {
		fibVal, ok := memo[n]
		if !ok {
			minus1, _ := Fib3(n - 1)
			minus2, _ := Fib3(n - 2)
			memo[n] = minus1 + minus2
			return memo[n], nil
		} else {
			return fibVal, nil
		}
	}
}

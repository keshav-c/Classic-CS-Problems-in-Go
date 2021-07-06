package fib

import "errors"

func Fib3(n int) (int, int, error) {
	var nCalls int = 0
	var memo = map[int]int{0: 0, 1: 1}
	var fib func(n int) (int, error)
	fib = func(n int) (int, error) {
		nCalls++
		if n < 0 {
			return -1, errors.New("Fib3 function needs non negative input.")
		} else {
			fibVal, ok := memo[n]
			if !ok {
				minus1, _ := fib(n - 1)
				minus2, _ := fib(n - 2)
				memo[n] = minus1 + minus2
				return memo[n], nil
			} else {
				return fibVal, nil
			}
		}
	}
	fibVal, err := fib(n)
	return fibVal, nCalls, err
}

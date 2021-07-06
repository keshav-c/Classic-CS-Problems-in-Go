package fib

import "errors"

func Fib2(n int) (int, int, error) {
	var nCalls int = 0
	var fib func(n int) (int, error)
	fib = func(n int) (int, error) {
		nCalls++
		if n < 0 {
			return -1, errors.New("Fib2 function needs non negative input.")
		} else if n == 0 || n == 1 {
			return n, nil
		} else {
			minus1, _ := fib(n - 1)
			minus2, _ := fib(n - 2)
			return minus1 + minus2, nil
		}
	}
	fibVal, err := fib(n)
	return fibVal, nCalls, err
}

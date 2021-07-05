package fib

import "errors"

func Fib2(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("Fib2 function needs non negative input.")
	} else if n == 0 || n == 1 {
		return n, nil
	} else {
		minus1, _ := Fib2(n - 1)
		minus2, _ := Fib2(n - 2)
		return minus1 + minus2, nil
	}
}

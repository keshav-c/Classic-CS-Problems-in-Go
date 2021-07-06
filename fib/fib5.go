package fib

import "errors"

func Fib5(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("Fib5 function needs non negative input.")
	} else if n == 0 || n == 1 {
		return n, nil
	} else {
		last := 0
		next := 1
		for i := 1; i < n; i++ {
			last, next = next, last+next
		}
		return next, nil
	}
}

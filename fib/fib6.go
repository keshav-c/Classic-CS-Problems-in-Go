package fib

import "errors"

func Fib6(n int) (<-chan int, error) {
	if n < 0 {
		return nil, errors.New("Fib6 generator function needs non negative input.")
	} else {
		fibOut := make(chan int)
		go func() {
			fibOut <- 0
			if n > 0 {
				fibOut <- 1
				last := 0
				next := 1
				for i := 1; i < n; i++ {
					last, next = next, last+next
					fibOut <- next
				}
				close(fibOut)
			}
		}()
		return fibOut, nil
	}
}

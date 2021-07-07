package main

import (
	"fmt"
	"github.com/keshav-c/classicCS/fib"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fibChan, err := fib.Fib6(n)
	if err != nil {
		log.Fatal(err)
	}
	for i := range fibChan {
		fmt.Println(i)
	}
}

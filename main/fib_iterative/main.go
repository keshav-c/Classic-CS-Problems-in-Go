package main

import (
	"fmt"
	"github.com/keshav-c/classicCS/fib"
	"github.com/keshav-c/classicCS/util"
	"log"
	"os"
	"strconv"
	"time"
)

func fibCaller(n int) int {
	defer util.Timer(time.Now())
	fibn, err := fib.Fib5(n)
	if err != nil {
		log.Fatal(err)
	}
	return fibn
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fibVal := fibCaller(n)
	fmt.Printf("Fib %d is %d\n", n, fibVal)
}

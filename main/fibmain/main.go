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

func fibCaller(n int) (int, int) {
	defer util.Timer(time.Now())
	//fibn, nCalls, err := fib.Fib2(n)
	fibn, nCalls, err := fib.Fib3(n)
	if err != nil {
		log.Fatal(err)
	}
	return fibn, nCalls
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fibn, nCalls := fibCaller(n)
	//fibn, err := fib.Fib5(n)
	fmt.Printf("Fib %d is %d\n", n, fibn)
	log.Printf("%d calls were made to fib.", nCalls)
}

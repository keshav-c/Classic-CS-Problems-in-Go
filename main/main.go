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

func main() {
	defer util.Timer(time.Now())
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fibn, err := fib.Fib2(n)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Fib %d is %d\n", n, fibn)
	}
}

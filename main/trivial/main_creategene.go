package main

import (
	"github.com/keshav-c/classicCS/trivial"
	"github.com/keshav-c/classicCS/util"
	"log"
	"os"
	"strconv"
	"time"
)

func createGeneData(fileName string, length int) error {
	defer util.Timer(time.Now())
	return trivial.GenerateGeneSequence(fileName, length) // 30 MegaBytes = 31457280 Bytes
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fileName := os.Args[1]
	length, err := strconv.Atoi(os.Args[2])
	check(err)
	err = createGeneData(fileName, length)
	check(err)
}

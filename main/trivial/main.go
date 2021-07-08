package main

import (
	"github.com/keshav-c/classicCS/trivial"
	"github.com/keshav-c/classicCS/util"
	"time"
)

func main() {
	defer util.Timer(time.Now())
	trivial.GenerateGeneSequence("s.txt", 30000000)
}

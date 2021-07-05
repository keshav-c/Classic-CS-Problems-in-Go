package util

import (
	"log"
	"time"
)

func Timer(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("Execution time: %s", elapsed)
}

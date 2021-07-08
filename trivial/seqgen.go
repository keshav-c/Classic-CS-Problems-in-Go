package trivial

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GenerateGeneSequence(fileName string, size int) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	geneSequence := getGeneSequence(size)
	n, err := file.WriteString(geneSequence)
	if err != nil {
		return err
	}
	log.Printf("%d bytes were written to file: %s.\n", n, fileName)
	if err = file.Close(); err != nil {
		return err
	}
	return nil
}

func getGene() string {
	genes := map[int]string{0: "A", 1: "C", 2: "T", 3: "G"}
	return genes[rand.Intn(4)]
}

func getGeneSequence(size int) string {
	rand.Seed(time.Now().UnixNano())
	var b strings.Builder
	for i := 0; i < size; i++ {
		fmt.Fprintf(&b, "%s", getGene())
	}
	return b.String()
}


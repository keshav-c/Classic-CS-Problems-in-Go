package main

import (
	"fmt"
	"github.com/keshav-c/classicCS/trivial"
	"github.com/keshav-c/classicCS/util"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func createGeneData(fileName string, length int) {
	defer util.Timer(time.Now())
	trivial.GenerateGeneSequence(fileName, length) // 30 MegaBytes = 31457280 Bytes
}

func getGeneData(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func compress(gene string) []byte {
	code := map[rune]byte{'A': 0b00, 'C': 0b01, 'G': 0b10, 'T': 0b11}
	geneSlice := []rune(strings.ToUpper(gene))
	cmprsnLen := len(geneSlice)/4 + 2
	cmprsdGene := make([]byte, 0, cmprsnLen)
	nucRead := 0
	var cByte byte = 0
	for _, nuc := range geneSlice {
		cByte <<= 2
		cByte |= code[nuc]
		nucRead++
		if nucRead == 4 {
			cmprsdGene = append(cmprsdGene, cByte)
			nucRead = 0
			cByte = 0
		}
	}
	if nucRead > 0 {
		cmprsdGene = append(cmprsdGene, cByte)
	}
	cmprsdGene = append(cmprsdGene, byte(nucRead))
	return cmprsdGene
}

func main() {
	fileName := os.Args[1]
	length, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	createGeneData(fileName, length)
	geneData, err := getGeneData(fileName)
	fmt.Println(geneData)
	if err != nil {
		log.Fatal(err)
	}
	compressedData := compress(geneData)
	fmt.Println(compressedData)
}

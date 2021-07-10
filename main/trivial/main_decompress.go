package main

import (
	"github.com/keshav-c/classicCS/trivial"
	"github.com/keshav-c/classicCS/util"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func getData(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func decompressGene(data []byte) (string, error) {
	defer util.Timer(time.Now())
	return trivial.Decompress(data)
}

func writeDecmprsdData(dcmprsd string, outFileName string) error {
	file, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	n, err := file.WriteString(dcmprsd)
	if err != nil {
		return err
	}
	log.Printf("%d decompressed nucleotides written to file: %s\n", n, outFileName)
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	inFileName := os.Args[1]
	outFileName := os.Args[2]
	binData, err := getData(inFileName)
	check(err)
	dcmprsdGene, err := decompressGene(binData)
	check(err)
	err = writeDecmprsdData(dcmprsdGene, outFileName)
	check(err)
}

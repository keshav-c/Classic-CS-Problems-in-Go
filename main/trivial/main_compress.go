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

func compressGene(data string) ([]byte, error) {
	defer util.Timer(time.Now())
	return trivial.Compress(data)
}

func writeCmprsdData(cmprsdBytes []byte, outFileName string) error {
	file, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	n, err := file.Write(cmprsdBytes)
	if err != nil {
		return err
	}
	log.Printf("%d compressed bytes were written to file: %s\n", n, outFileName)
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
	geneFileName := os.Args[1]
	outFileName := os.Args[2]
	data, err := getData(geneFileName)
	check(err)
	gene := string(data)
	cmprsd, err := compressGene(gene)
	check(err)
	err = writeCmprsdData(cmprsd, outFileName)
	check(err)
}

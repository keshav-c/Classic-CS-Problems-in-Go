package main

import (
	"github.com/keshav-c/classicCS/onetimepad"
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

func encrypt(data []byte) ([]byte, []byte, error) {
	defer util.Timer(time.Now())
	return onetimepad.Encrypt(data)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	plainFileName := os.Args[1]
	plainBytes, err := getData(plainFileName)
	check(err)

	cypher, key, err := encrypt(plainBytes)
	check(err)

	cypherFile, err := os.Create("cypher")
	check(err)
	_, err = cypherFile.Write(cypher)
	check(err)
	err = cypherFile.Close()
	check(err)

	keyFile, err := os.Create("key")
	check(err)
	_, err = keyFile.Write(key)
	check(err)
	err = keyFile.Close()
	check(err)
}

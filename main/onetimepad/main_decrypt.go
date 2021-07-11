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

func decrypt(cypher, key []byte) ([]byte, error) {
	defer util.Timer(time.Now())
	return onetimepad.Decrypt(cypher, key)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cypherFileName := os.Args[1]
	cypher, err := getData(cypherFileName)
	check(err)
	keyFileName := os.Args[2]
	key, err := getData(keyFileName)
	check(err)

	plainBytes, err := decrypt(cypher, key)
	check(err)

	decryptedFileName := os.Args[3]
	decryptedFile, err := os.Create(decryptedFileName)
	check(err)
	_, err = decryptedFile.Write(plainBytes)
	check(err)
	err = decryptedFile.Close()
	check(err)
}

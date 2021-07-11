package main

import (
	"fmt"
	"github.com/keshav-c/classicCS/onetimepad"
	"log"
)

func main() {
	plain := "abc"
	plainBytes := []byte(plain)
	fmt.Println(plain, "in bytes", printBinSlice(plainBytes))
	cypherBytes, key, err := onetimepad.Encrypt(plainBytes)
	if err != nil {
		log.Fatal(err)
	}
	cypher := string(cypherBytes)
	fmt.Println(cypher, "in bytes", printBinSlice(cypherBytes))
	keyStr := string(key)
	fmt.Println(keyStr, "in bytes", printBinSlice(key))
	decryptedBytes, err := onetimepad.Decrypt(cypherBytes, key)
	if err != nil {
		log.Fatal(err)
	}
	decrypted := string(decryptedBytes)
	fmt.Println(decrypted, "in bytes", printBinSlice(decryptedBytes))
}

func printBinSlice(bs []byte) string {
	a := ""
	for _, n := range bs {
		a += fmt.Sprintf("%08b", n)
	}
	return a
}

package onetimepad

import (
	"crypto/rand"
	"errors"
)

func Encrypt(plain []byte) ([]byte, []byte, error) {
	l := len(plain)
	dummy, err := randNBytes(l)
	if err != nil {
		return nil, nil, err
	}
	cypher := make([]byte, l)
	for i, pi := range plain {
		cypher[i] = pi ^ dummy[i]
	}
	return cypher, dummy, nil
}

func Decrypt(cypher, key []byte) ([]byte, error) {
	l := len(cypher)
	if l != len(key) {
		return nil, errors.New("Cyphertext and key are of different lengths.")
	}
	plain := make([]byte, l)
	for i, ci := range cypher {
		plain[i] = ci ^ key[i]
	}
	return plain, nil
}

func randNBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

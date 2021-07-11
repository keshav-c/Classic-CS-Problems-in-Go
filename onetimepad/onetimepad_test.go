package onetimepad

import (
	"bytes"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	plain := "One Time Pad!"
	plainBytes := []byte(plain)
	enc, key, _ := Encrypt(plainBytes)
	dec, _ := Decrypt(enc, key)
	if !bytes.Equal(plainBytes, dec) {
		t.Error("One time pad encryption/decryption failed.")
	}
}

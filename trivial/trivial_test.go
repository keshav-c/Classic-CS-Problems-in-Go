package trivial

import (
	"testing"
)

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ai := range a {
		if ai != b[i] {
			return false
		}
	}
	return true
}

func TestCompress4xNucs(t *testing.T) {
	gene := "GTAGAGATGCAT"
	expected := []byte{0xb2, 0x23, 0x93, 0x04}
	compressed, err := Compress(gene)
	if !equal(expected, compressed) || err != nil {
		t.Error("Failed to compress sequence of 4x nucleotides.")
	}
}

func TestCompressLt4Nucs(t *testing.T) {
	gene := "GT"
	expected := []byte{0x0b, 0x02}
	compressed, err := Compress(gene)
	if !equal(expected, compressed) || err != nil {
		t.Error("Failed to compress sequence of less than 4 nucleotides.")
	}
}

func TestCompress4xpNucs(t *testing.T) {
	gene := "AATGCCCAAGA"
	expected := []byte{0x0e, 0x54, 0x08, 0x03}
	compressed, err := Compress(gene)
	if !equal(expected, compressed) || err != nil {
		t.Error("Failed to compress sequence of 4x + n nucleotides.")
	}
}

func TestCompressInvalidNucs(t *testing.T) {
	gene := "AYCT"
	compressed, err := Compress(gene)
	if compressed != nil || err.Error() != "Invalid Nucleotide in sequence." {
		t.Error("Failed to handle sequence with invalid nucleotides.")
	}
}

func TestDecompress2Bytes(t *testing.T) {
	compressed := []byte{0x0b, 0x02}
	expected := "GT"
	decompressed, err := Decompress(compressed)
	if expected != decompressed || err != nil {
		t.Error("Failed to decompress 2 bytes.")
	}
}

func TestDecompressMt2Bytes(t *testing.T) {
	compressed := []byte{0x0e, 0x54, 0x08, 0x03}
	expected := "AATGCCCAAGA"
	decompressed, err := Decompress(compressed)
	if expected != decompressed || err != nil {
		t.Error("Failed to decompress more than 2 bytes.")
	}
}

func TestDecompress1Byte(t *testing.T) {
	compressed := []byte{0x0e}
	expected := ""
	decompressed, err := Decompress(compressed)
	if expected != decompressed || err.Error() != "Invalid compressed gene (min 2 bytes)" {
		t.Error("Failed to handle compressed data with 1 byte.")
	}
}

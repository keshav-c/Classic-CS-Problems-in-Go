package trivial

import (
	"errors"
	"fmt"
	"strings"
)

var code = map[rune]byte{
	'A': 0b00,
	'C': 0b01,
	'G': 0b10,
	'T': 0b11,
}

var decode = map[byte]rune{
	0b00: 'A',
	0b01: 'C',
	0b10: 'G',
	0b11: 'T',
}

func Compress(gene string) ([]byte, error) {
	geneSlice := []rune(strings.ToUpper(gene))
	cmprsnLen := len(geneSlice)/4 + 2
	cmprsdGene := make([]byte, 0, cmprsnLen)
	var cByte, nucRead byte
	for _, nuc := range geneSlice {
		cByte <<= 2
		encodedNuc, ok := code[nuc]
		if !ok {
			return nil, errors.New("Invalid Nucleotide in sequence.")
		}
		cByte |= encodedNuc
		nucRead++
		if nucRead == 4 {
			cmprsdGene = append(cmprsdGene, cByte)
			nucRead, cByte = 0, 0
		}
	}
	if nucRead > 0 {
		cmprsdGene = append(cmprsdGene, cByte, nucRead)
	} else {
		cmprsdGene = append(cmprsdGene, 4)
	}
	return cmprsdGene, nil
}

func Decompress(cmprsdByts []byte) (string, error) {
	var b strings.Builder
	l := len(cmprsdByts)
	b.Grow(l * 4) // Size of decompressed gene
	if l < 2 {
		return "", errors.New("Invalid compressed gene (min 2 bytes)")
	}
	for _, cByte := range cmprsdByts[:l-2] {
		fmt.Fprintf(&b, "%s", decompressByte(cByte, 4))
	}
	fmt.Fprintf(&b, "%s", decompressByte(cmprsdByts[l-2], int(cmprsdByts[l-1])))
	return b.String(), nil
}

func decompressByte(cByte byte, nNucs int) string {
	var nucs [4]rune
	for nucsRead := 0; nucsRead < 4; nucsRead++ {
		encodedBits := cByte & byte(0b11)
		nucs[3-nucsRead] = decode[encodedBits]
		cByte >>= 2
	}
	return string(nucs[4-nNucs:])
}

#! /bin/bash

gene="gene.txt"
cmprsd="comp.bin"
dcmprsd="decomp.txt"

echo "Removing old files..."
rm -f $gene $cmprsd $dcmprsd
echo "Generating new gene..."
./seqgen $gene 31457280
echo "Compressing new gene..."
./compress $gene $cmprsd
echo "Decoding compressed gene..."
./decompress $cmprsd $dcmprsd

if cmp -s $gene $dcmprsd; then
	printf "The file %s is the same as %s\n" $dcmprsd $gene
else
	printf "The file %s is NOT the same as %s\n" $dcmprsd $gene
fi

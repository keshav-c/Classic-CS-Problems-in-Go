#! /bin/bash

gene="gene.txt"
cmprsd="comp.bin"
dcmprsd="decomp.txt"

echo "Removing old files..."
rm -f $gene $cmprsd $dcmprsd
echo "Generating new gene..."
go run main_creategene.go $gene 31457280
echo "Compressing new gene..."
go run main_compress.go $gene $cmprsd
echo "Decoding compressed gene..."
go run main_decompress.go $cmprsd $dcmprsd

if cmp -s $gene $dcmprsd; then
	printf "The file %s is the same as %s\n" $dcmprsd $gene
else
	printf "The file %s is NOT the same as %s\n" $dcmprsd $gene
fi

#! /bin/bash

plain="alice.txt"
url="https://www.gutenberg.org/files/11/11-0.txt"
cypher="cypher"
key="key"
decrypted="decrypted.txt"

echo "Downloading plain text file ..."
curl $url --output $plain
echo "Encrypting plain text file: alice.txt"
./otp_encrypt $plain
echo "Decrypting encoded file: $cypher; with key: $key ..."
./otp_decrypt $cypher $key $decrypted

echo "Compare $plain and $decrypted ..."
if cmp -s $plain $decrypted; then
	echo "$plain and $decrypted are the same."
else
	echo "$plain and $decrypted are NOT the same."
fi

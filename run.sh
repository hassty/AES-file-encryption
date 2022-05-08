#!/bin/sh

bin=./aes
keyfile=aes.key
message='tihomirov timofey mihaylovich'
iv=deadfa11feedc0de

printf "cafebabedeadbeef" >"$keyfile"

printf "%10s %s\n" "key:" "$(cat $keyfile)"
printf "%10s %s\n" "iv:" "$iv"

encrypted=$(echo "$message" | $bin --keyfile $keyfile --iv $iv encrypt)
printf "encrypted: %s\n" "$encrypted"

decrypted=$(echo "$encrypted" | $bin --keyfile $keyfile --iv $iv decrypt)
printf "decrypted: %s\n" "$decrypted"

#!/bin/sh

bin=./aes
keyfile=aes.key
message='tihomirov timofey mihaylovich'

printf "cafebabedeadbeef" >"$keyfile"

printf "%10s %s\n" "key:" "$(cat $keyfile)"

encrypted=$(echo "$message" | $bin --keyfile $keyfile encrypt)
printf "encrypted: %s\n" "$encrypted"

decrypted=$(echo "$encrypted" | $bin --keyfile $keyfile decrypt)
printf "decrypted: %s\n" "$decrypted"

rm "$keyfile"

package main

import (
	"log"
	"os"

	rhdiff "github.com/faagerholm/rhdiff/src"
)

func Signature(inStr, outStr string, chunkLen, strongLen int) {

	file, err := os.Open(inStr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sig, err := os.OpenFile(outStr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0600))
	if err != nil {
		log.Fatal(err)
	}
	defer sig.Close()

	_, err = rhdiff.Signature(file, sig, uint32(chunkLen), uint32(strongLen))
	if err != nil {
		log.Fatal(err)
	}
}

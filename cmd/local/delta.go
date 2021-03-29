package main

import (
	"log"
	"os"

	rhdiff "github.com/faagerholm/rhdiff/src"
)

func Delta(sigPtr, inPtr, outPtr string, chunkPtr int) {
	// Compare the binary difference between old file and new file
	// With the help of signature and new file => Output: delta

	signature, err := os.Open(sigPtr)
	if err != nil {
		log.Fatal(err)
	}
	defer signature.Close()

	in, err := os.Open(inPtr)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	delta, err := os.OpenFile(outPtr, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0600))
	if err != nil {
		log.Fatal(err)
	}
	defer delta.Close()

	sig, err := rhdiff.LoadSigFile(signature, uint32(chunkPtr))

	if err := rhdiff.Delta(sig, in, delta); err != nil {
		log.Fatal(err)
	}

}

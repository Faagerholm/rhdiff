package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// main.go signature file.txt signature.bin
	sigCommand := flag.NewFlagSet("signature", flag.ExitOnError)
	// main.go delta signature.bin new_file.txt delta.bin
	deltaCommand := flag.NewFlagSet("delta", flag.ExitOnError)

	// signature subcommand flag pointers
	sigInputPtr := sigCommand.String("i", "", "Input file. (Required)")
	sigOutputPtr := sigCommand.String("s", "signature.bin", "Output signature file.")
	sigChunkPtr := sigCommand.Int("chunk", 16, "Chunk size.")
	sigStrongPtr := sigCommand.Int("strong", 32, "Chunk size.")

	// delta subcommand flag pointers
	deltaSigPtr := deltaCommand.String("s", "signature.bin", "Signature file. (Required)")
	deltaInputPtr := deltaCommand.String("i", "", "Input file. (Required)")
	deltaOutputPtr := deltaCommand.String("d", "delta.bin", "Output delta file.")
	deltaStrongPtr := deltaCommand.Int("strong", 32, "Strong length.")

	if len(os.Args) < 2 {
		fmt.Println("signature or delta command required")
		os.Exit(1)
	}

	// Parse subcommands
	switch os.Args[1] {
	case "signature":
		sigCommand.Parse(os.Args[2:])
	case "delta":
		deltaCommand.Parse(os.Args[2:])
	default:
		//TODO: FIX DEFAULTS
		flag.PrintDefaults()
		os.Exit(1)
	}

	if sigCommand.Parsed() {
		if *sigInputPtr == "" {
			sigCommand.PrintDefaults()
			os.Exit(1)
		}

		Signature(*sigInputPtr, *sigOutputPtr, *sigStrongPtr, *sigChunkPtr)
		fmt.Printf("Signature created: %s", *sigOutputPtr)
		os.Exit(1)
	}

	if deltaCommand.Parsed() {
		if *deltaInputPtr == "" {
			deltaCommand.PrintDefaults()
			os.Exit(1)
		}
		if *deltaSigPtr == "" {
			deltaCommand.PrintDefaults()
			os.Exit(1)
		}

		Delta(*deltaSigPtr, *deltaInputPtr, *deltaOutputPtr, *deltaStrongPtr)
		fmt.Printf("Delta file created: %s", *deltaOutputPtr)
		os.Exit(1)
	}
}

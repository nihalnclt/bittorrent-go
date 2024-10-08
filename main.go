package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nihalnclt/bittorrent-go/torrentfile"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <inputPath> <outputPath>")
		os.Exit(1)
	}

	// Retrieve input and output paths from command-line arguments
	inPath := os.Args[1]
	outPath := os.Args[2]

	if inPath == "" || outPath == "" {
		fmt.Println("Error: inputPath or outputPath is empty")
		os.Exit(1)
	}

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("InfoHash: %x\n", tf.InfoHash)
	fmt.Printf("Announce: %s\n", tf.Announce)

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second)
}

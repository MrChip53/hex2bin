package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
)

var inputFile string
var outputFile string

func init() {
	flag.StringVar(&inputFile, "i", "", "input file")
	flag.StringVar(&outputFile, "o", "out.bin", "output file")
	flag.Parse()
}

func main() {
	fileBytes, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	fileString := strings.ReplaceAll(string(fileBytes), "\r\n", "")
	fileString = strings.ReplaceAll(fileString, "\n", "")

	bytes, err := hex.DecodeString(fileString)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(outputFile, bytes, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Binary output to file at " + outputFile + ".")
}

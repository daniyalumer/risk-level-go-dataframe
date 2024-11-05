package main

import (
	"flag"
	"log"
	"os"

	"github.com/daniyalumer/csvdataframe/risklevel"
)

func main() {
	inputFile := flag.String("input", "", "Input CSV file path (required)")
	outputFile := flag.String("output", "", "Output CSV file path (required)")
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	log.Printf("Looking for CSV file: %s\n", *inputFile)

	if _, err := os.Stat(*inputFile); err != nil {
		if os.IsNotExist(err) {
			log.Printf("Error: CSV file not found %s\n", *inputFile)
		} else {
			log.Printf("Error accessing file %s\n", *inputFile)
		}
		return
	}

	risklevel.RiskLevelAssessment(*inputFile, *outputFile)

	log.Printf("Processing complete. Output saved to %s\n", *outputFile)

}

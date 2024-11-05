package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/daniyalumer/csvdataframe/risklevel"
)

func main() {
	inputFile := flag.String("input", "", "Input CSV file path (required)")
	outputPath := flag.String("output", "", "Output directory path (required)")
	flag.Parse()

	if *inputFile == "" || *outputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	path, fileName := filepath.Split(*inputFile)

	log.Printf("Looking for CSV file: %s in directory: %s\n", fileName, path)

	if _, err := os.Stat(*inputFile); err != nil {
		if os.IsNotExist(err) {
			log.Printf("Error: CSV file not found: %s\n", fileName)
		} else {
			log.Printf("Error accessing file %s: %v\n", fileName, err)
		}
		return
	}

	risklevel.RiskLevelAssessment(path, fileName, *outputPath)

	log.Printf("Processing complete. Output saved to %s\n", *outputPath)

}

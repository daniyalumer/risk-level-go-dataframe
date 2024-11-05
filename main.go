package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/daniyalumer/csvdataframe/risklevel"
)

func main() {
	path, fileName := filepath.Split(os.Args[1])

	outputPath := os.Args[2]

	fmt.Printf("Looking for CSV file: %s in directory: %s\n", fileName, path)

	if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
		fmt.Printf("Error: CSV file not found: %s\n", fileName)
		return
	}

	risklevel.RiskLevelAssessment(path, fileName, outputPath)

	fmt.Println("Processing complete. Output saved to %s\n", outputPath)

}

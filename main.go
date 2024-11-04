package main

import (
	"fmt"
	"os"

	"github.com/daniyalumer/csvdataframe/risklevel"
)

func main() {
	path := "."

	csv_name := "241015-page-stats-aggregator.csv"

	fmt.Printf("Looking for CSV file: %s in directory: %s\n", csv_name, path)

	if _, err := os.Stat(path + "/" + csv_name); os.IsNotExist(err) {
		fmt.Printf("Error: CSV file not found: %s\n", csv_name)
		return
	}

	risklevel.RiskLevelAssessment(path, csv_name)

	fmt.Println("Processing complete. Output saved to: page-stats-aggregator.csv")

}

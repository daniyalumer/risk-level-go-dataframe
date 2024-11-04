package main

import (
	"fmt"
	"os"

	dataframe "github.com/datumbrain/go-dataframe"
)

func main() {
	path := "."

	csv_name := "241015-page-stats-aggregator.csv"

	fmt.Printf("Looking for CSV file: %s in directory: %s\n", csv_name, path)

	if _, err := os.Stat(path + "/" + csv_name); os.IsNotExist(err) {
		fmt.Printf("Error: CSV file not found: %s\n", csv_name)
		return
	}

	df := dataframe.CreateDataFrame(path, csv_name)
	fmt.Println("Processing data...")

	df.NewField("Risk Level")
	for _, row := range df.FrameRecords {
		if row.ConvertToInt("Score", df.Headers) <= 4 && row.ConvertToInt("Score", df.Headers) >= 0 {
			row.Update("Risk Level", "High", df.Headers)
		} else if row.ConvertToInt("Score", df.Headers) <= 7 && row.ConvertToInt("Score", df.Headers) >= 4 {
			row.Update("Risk Level", "Medium", df.Headers)
		} else if row.ConvertToInt("Score", df.Headers) <= 9 && row.ConvertToInt("Score", df.Headers) >= 8 {
			row.Update("Risk Level", "Low", df.Headers)
		} else {
			row.Update("Risk Level", "No Risk", df.Headers)
		}
	}

	df.Sort("Risk Level")
	df.SaveDataFrame(path, "page-stats-aggregator.csv")
	fmt.Println("Processing complete. Output saved to: page-stats-aggregator.csv")

}

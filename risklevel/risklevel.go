package risklevel

import (
	"fmt"

	dataframe "github.com/datumbrain/go-dataframe"
)

func RiskLevelAssessment(path string, csvName string, outputPath string) {
	df := dataframe.CreateDataFrame(path, csvName)
	fmt.Println("Processing data...")

	df.NewField("Risk Level")
	for _, row := range df.FrameRecords {
		switch {
		case row.ConvertToInt("Score", df.Headers) >= 0 && row.ConvertToInt("Score", df.Headers) < 4:
			row.Update("Risk Level", "High", df.Headers)
		case row.ConvertToInt("Score", df.Headers) >= 4 && row.ConvertToInt("Score", df.Headers) < 7:
			row.Update("Risk Level", "Medium", df.Headers)
		case row.ConvertToInt("Score", df.Headers) >= 7 && row.ConvertToInt("Score", df.Headers) < 9:
			row.Update("Risk Level", "Low", df.Headers)
		case row.ConvertToInt("Score", df.Headers) == 10:
			row.Update("Risk Level", "No Risk", df.Headers)
		}
	}

	df.Sort("Risk Level")
	if !df.SaveDataFrame(outputPath, "page-stats-aggregator") {
		fmt.Printf("Error: failed to save CSV to output path: %s\n", outputPath)
		return
	}
}

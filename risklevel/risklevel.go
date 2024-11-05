package risklevel

import (
	"log"

	dataframe "github.com/datumbrain/go-dataframe"
)

func RiskLevelAssessment(path string, csvName string, outputPath string) {
	df := dataframe.CreateDataFrame(path, csvName)
	log.Println("Processing data...")

	df.NewField("Risk Level")
	for _, row := range df.FrameRecords {
		score := row.ConvertToInt("Score", df.Headers)
		switch {
		case score >= 0 && score < 4:
			row.Update("Risk Level", "High", df.Headers)
		case score >= 4 && score < 7:
			row.Update("Risk Level", "Medium", df.Headers)
		case score >= 7 && score < 9:
			row.Update("Risk Level", "Low", df.Headers)
		case score == 10:
			row.Update("Risk Level", "No Risk", df.Headers)
		}
	}

	df.Sort("Risk Level")
	if !df.SaveDataFrame(outputPath, "page-stats-aggregator") {
		log.Printf("Error: failed to save CSV to output path: %s\n", outputPath)
		return
	}
}

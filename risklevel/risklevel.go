package risklevel

import (
	"fmt"

	dataframe "github.com/datumbrain/go-dataframe"
)

func RiskLevelAssessment(path string, csv_name string) {
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
	df.SaveDataFrame("", "page-stats-aggregator")
}

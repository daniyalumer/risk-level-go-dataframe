package risklevel

import (
	"log"
	"os"
	"path/filepath"

	dataframe "github.com/datumbrain/go-dataframe"
)

func RiskLevelAssessment(inputPath string, outputFile string) {
	inputPath, inputFileName := filepath.Split(inputPath)

	df := dataframe.CreateDataFrame(inputPath, inputFileName)
	log.Println("Processing data...")

	df.NewField("Risk Level")
	for _, row := range df.FrameRecords {
		score := row.ConvertToInt("Score", df.Headers)
		switch {
		case score >= 0 && score <= 4:
			row.Update("Risk Level", "High", df.Headers)
		case score >= 4 && score <= 7:
			row.Update("Risk Level", "Medium", df.Headers)
		case score >= 7 && score <= 9:
			row.Update("Risk Level", "Low", df.Headers)
		case score == 10:
			row.Update("Risk Level", "No Risk", df.Headers)
		}
	}

	df.Sort("Risk Level")

	outputPath, outputFileName := filepath.Split(outputFile)
	if _, err := os.Stat(outputPath); err != nil {
		if os.IsNotExist(err) {
			log.Printf("Error: Output directory does not exist: %s\n", outputPath)
			return
		}
		log.Printf("Error accessing output directory: %s\n", outputPath)
		return
	}

	log.Printf("Saving output to %s\n", outputPath)
	if !df.SaveDataFrame(outputPath, outputFileName) {
		log.Printf("Error: failed to save CSV to output path: %s\n", outputPath)
		return
	}
}

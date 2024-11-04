# Page Stats Aggregator

This Go program reads data from a CSV file, assigns a "Risk Level" based on page view scores, sorts the data by the new column, and saves the updated data back to a CSV file.

## Table of Contents
- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [Input Format](#input-format)
- [Output](#output)
- [Dependencies](#dependencies)

## Overview

This program utilizes the `go-dataframe` package to process CSV data. It adds a "Risk Level" column based on the "Score" field using these rules:

- **Score 0-4:** `High`
- **Score 5-7:** `Medium`
- **Score 8-9:** `Low`
- **Score 10:** `No Risk`

## Installation

1. Ensure Go is installed on your system
2. Install the required package:
   ```bash
   go get github.com/datumbrain/go-dataframe
   ```
3. Clone this repository or download the source code

## Usage

1. Place your input CSV file in the same directory as the program
2. Run the program:
   ```bash
   go run main.go
   ```
3. The processed data will be saved in the root directory

## Input Format

The input CSV file should contain at least a "Score" column with numerical values from 0-10.

Example input format:
```csv
Page,Score,OtherColumns...
/home,8,...
/about,4,...
```

## Output

The program will:
1. Add a new "Risk Level" column
2. Sort data by risk level
3. Save the results to `page-stats-aggregator.csv`

Example output format:
```csv
Page,Score,Risk Level
/about,4,High
/home,8,Low
```

## Dependencies

- Go 1.x
- [go-dataframe](https://github.com/datumbrain/go-dataframe) package
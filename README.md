# Page Stats Aggregator

This Go program reads data from a CSV file, assigns a "Risk Level" based on page view scores, sorts the data by the new column, and saves the updated data back to a CSV file.

## Table of Contents
- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [File Structure](#file-structure)
- [Example](#example)

## Overview

This program utilizes the `go-dataframe` package to load and process data from a CSV file named `241015-page-stats-aggregator.csv`. It adds a new column called "Risk Level" based on conditions applied to the "Score" field:

- **Score 0-4:** `High`
- **Score 5-7:** `Medium`
- **Score 8-9:** `Low`
- **Score 10:** `No Risk`

The program then sorts the data by the "Risk Level" column and saves the updated data to a new file named `page-stats-aggregator.csv`.

## Installation

To use this program, you must have Go installed on your system. Additionally, install the `go-dataframe` package by running:

```bash
go get github.com/datumbrain/go-dataframe

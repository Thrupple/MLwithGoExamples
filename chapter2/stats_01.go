// Usage:
//  go run stats_01.go iris.csv
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat"
)

///////////////////////////////////////////////////////////////////////////////

var (
	flagSkipHeader = flag.Bool("skip_header", true, "Skip CSV header row")
)

var (
	ErrEmpty = fmt.Errorf("Empty string")
)

///////////////////////////////////////////////////////////////////////////////

func ParseFloat(string_value string) (float64, error) {
	// Check for empty value
	if len(strings.TrimSpace(string_value)) == 0 {
		return 0.0, ErrEmpty
	}
	// Convert to integer
	if float_value, err := strconv.ParseFloat(string_value, 64); err != nil {
		return 0.0, err
	} else {
		return float64(float_value), nil
	}
}

func AnalyseSepalLength(rows [][]string) error {
	sepal_length := make([]float64, 0)
	for line_number, row := range rows {
		if len(row) != 5 {
			return fmt.Errorf("Line %v: Expected 5 values", line_number+1)
		}
		// Skip header
		if *flagSkipHeader && line_number == 0 {
			continue
		}
		// Retrieve sepal_length
		
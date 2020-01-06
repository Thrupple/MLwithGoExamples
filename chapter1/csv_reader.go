
// Usage:
//  go run csv_reader.go data.csv
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////

var (
	flagSkipHeader = flag.Bool("skip_header", true, "Skip CSV header row")
)

var (
	ErrEmpty = fmt.Errorf("Empty string")
)

///////////////////////////////////////////////////////////////////////////////

func ParseInteger(string_value string) (int, error) {
	// Check for empty value
	if len(strings.TrimSpace(string_value)) == 0 {
		return 0, ErrEmpty
	}
	// Convert to integer
	if int_value, err := strconv.ParseInt(string_value, 10, 32); err != nil {
		return 0, err
	} else {
		return int(int_value), nil
	}
}

func MaxInt(a, b int) int {
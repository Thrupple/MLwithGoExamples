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
		
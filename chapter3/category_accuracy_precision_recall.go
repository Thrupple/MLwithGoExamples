// Usage:
//  go run chapter3/category_accuracy.go chapter3/labeled.csv
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	// Utilities for reading data
	"github.com/djthorpe/MachineLearning/util"
)

///////////////////////////////////////////////////////////////////////////////

func RunMain() int {
	if flag.NArg() != 1 {
		log.Println("Expected file argument")
		return -1
	}

	table, _ := util.NewTable()
	filename := flag.Arg(0)
	if err := table.ReadCSV(filename, false, true, true); err != nil {
		log.Println("Unable to read CSV:", err)
		return -1
	}

	// Calculate the mean absolut
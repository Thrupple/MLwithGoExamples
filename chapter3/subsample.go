// Usage:
//  go run chapter3/subsample.go chapter3/time_series.csv
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

	// One in every four is the testing set
	training_rows := make([]int, 0)
	testing_rows := make([]int, 0)
	for row := 0; row < len(table.Rows); row++ {
		if row%4 == 0 {
			testing_rows = append(testing_rows, row)
		} else {
			training_rows = append
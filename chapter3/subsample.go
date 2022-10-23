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

	
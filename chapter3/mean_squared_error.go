
// Usage:
//  go get -u gonum.org/v1/gonum/stat
//  go run chapter3/mean_squared_error.go chapter3/time_series.csv
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"

	// Utilities for reading data
	"github.com/djthorpe/MachineLearning/util"
	"gonum.org/v1/gonum/stat"
)

///////////////////////////////////////////////////////////////////////////////

func RunMain() int {
	if flag.NArg() != 1 {
		log.Println("Expected file argument")
		return -1
	}

	table, _ := util.NewTable()
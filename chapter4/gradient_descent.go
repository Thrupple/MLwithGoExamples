
// Usage:
//  go run chapter4/gradient_descent.go chapter4/advertising.csv
package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"path"

	// Frameworks
	"github.com/djthorpe/MachineLearning/util"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	LEARNING_RATE = 0.001
)

///////////////////////////////////////////////////////////////////////////////

func RunMain() int {
	if flag.NArg() < 1 {
		log.Println("Expected file argument", flag.NArg())
		return -1
	}

	table, _ := util.NewTable()
	filename := flag.Arg(0)
	if err := table.ReadCSV(filename, false, true, true); err != nil {
		log.Println("Unable to read CSV:", err)
		return -1
	}

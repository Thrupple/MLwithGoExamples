// Usage:
//  go get -u gonum.org/v1/plot/...
//  go run plot_01.go iris.csv
//  open iris.csv_hist.png
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path"

	"github.com/djthorpe/MachineLearning/util"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

///////////////////////////////////////////////////////////////////////////////

var (
	ErrEmpty = fmt.Errorf("Empty string")
)

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
	if petal_length, err
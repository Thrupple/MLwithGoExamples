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
	ErrEmpty = fmt.Errorf("Empty
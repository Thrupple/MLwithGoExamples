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
	"gonum
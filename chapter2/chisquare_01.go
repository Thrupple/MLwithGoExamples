// Usage:
//  go get -u gonum.org/v1/gonum/stat
//  go run chisquare_01.go
package main

import (
	"flag"
	"fmt"
	"os"

	"gonum.org/v1/gonum/stat"
)

///////////////////////////////////////////////////////////////////////////////

func RunMain() int {
	// Define observed and expected values. Most
	// of the time these will come from your
	// data (website visits, etc.).
	observed := []float64{48, 52}
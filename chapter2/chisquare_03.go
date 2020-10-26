// Usage:
//  go get -u gonum.org/v1/gonum/stat
//  go get -u golang.org/x/exp/rand
//  go run chapter2/chisquare_03.go
package main

import (
	"flag"
	"fmt"
	"os"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

///////////////////////////////////////////////////////////////////////////////

func RunMain() int {
	// Define the observed frequencies
	observed := []float64{
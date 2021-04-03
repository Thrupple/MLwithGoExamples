// Usage:
//  go get -u gonum.org/v1/gonum/...
//  go run matrix_01.go
package main

import (
	"flag"
	"fmt"
	"os"

	"gonum.org/v1/gonum/mat"
)

///////////////////////////////////////////////////////////////////////////////
// The core packages of the gonum suite are written in pure Go with
// some assembly. Installation is done using go get:
//
//   go get -u gonum.org/v1/gonum/...
//

// RunMain runs the main program
func RunMain() int {

	// Create a flat representation of our matrix.
	components := []float64{1.2, -5.7, -2.4, 7.3}

	// Form our matrix (the first argument is the number of
	// rows and the second argument is the number of columns).
	a := 
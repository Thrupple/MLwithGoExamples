// Usage:
//  go get -u gonum.org/v1/gonum/...
//  go run matrix_02.go
package main

import (
	"flag"
	"fmt"
	"math"
	"os"

	"gonum.org/v1/gonum/mat"
)

///////////////////////////////////////////////////////////////////////////////
// The core packages of the gonum suite are written in pure Go with
// some assembly. Installation is done using go get:
//
//   go get -u gonum.org/v1/gonum/...
//

// RunMain runs the main program. Demonstrates adding, multiplying, powers,
// applying function to elements
func RunMain() int {
	// Create two matrices of the same size, a and b.
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 
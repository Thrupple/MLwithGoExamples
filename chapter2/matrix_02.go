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
//   go get -u gonum.org/v1/gonum/.
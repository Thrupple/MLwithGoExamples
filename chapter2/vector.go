// Usage:
//  go get -u gonum.org/v1/gonum/...
//  go run vector.go
package main

import (
	"flag"
	"fmt"
	"os"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
)

///////////////////////////////////////////////////////////////////////////////
// The core packages of the gonum suite are written in pure Go with
// some assembly. Installation is done using go get:
//
//   go get -u gonum.org/v1/gonum/...

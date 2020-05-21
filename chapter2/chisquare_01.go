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
	// Define observed and exp
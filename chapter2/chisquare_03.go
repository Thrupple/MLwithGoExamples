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
		260.0, // This number is the number of observed with no regular exercise.
		135.0, // This number is the number of observed with sporatic exercise.
		105.0, // This number is the number of observed with regular exercise.
	}
	// Define the total observed.
	totalObserved := 500.0

	// Calculate the expected frequencies (again assuming the null Hypothesis).
	expected := []float64{
		totalObserved * 0.60,
		totalObserved * 0.25,
		totalObserved * 0.15,
	}

	// Calculate the ChiSquare test statistic.
	chiSquare := stat.ChiSquare(observed, expected)

	// Output the test statistic to standard out.
	fmt.Println("Observed", observed)
	fmt.Println("Expected", expected)
	fmt.Println("chiSquare", chiSquare)

	// Create a Chi-squared distribution with K degrees of freedom.
	// In this case 
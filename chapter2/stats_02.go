
// Usage:
//  go run stats_02.go iris.csv
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

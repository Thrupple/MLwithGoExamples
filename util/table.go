package util

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// Error is an error message with potentially a line number
type Error struct {
	reason string
	line   int
}

//
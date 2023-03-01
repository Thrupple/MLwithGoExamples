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

// Value is a single value in the data
type Value struct {
	Str      string
	_Uint64  *uint64
	_Int64   *int64
	_Float64 *float64
}

// Table is the table of values with optional column headers
type Table struct {
	Columns []string
	colmap  map[string]int
	Rows    [][]*Value
}

var (
	ErrDuplicateColumn = &Error{reason: "Duplicate or invalid column name"}
	ErrDimensionError  = &Error{reason: "Too many values for row"}
	ErrOutOfRange      = &Error{reason: "Index out of range"}
	ErrNotFound        = &Error{reason: "Column Not Found"}
)

// NewTable creates a new table with specified columns
func NewTable(columns ...string) (*Table, error) {
	this := new(Table)
	if err := this.SetColumns(columns...); err != nil {
		return nil, err
	}
	return this, nil
}

// Subsample creates a new table from an existing table with
// the specified rows
func (this *Table) Subsample(rows []int) (*Table, error) {
	that := new(Table)
	if err := that.SetColumns(this.Columns...); err != nil {
		return nil, err
	} else {
		that.Rows = make([][]*Value, 0, len(rows))
		for _, row := range rows {
			if row < 0 || row >= len(this.Rows) {
				return nil, ErrOutOfRange
			}
			that.Rows = append(that.Rows, this.Rows[row])
		}
	}
	return that, nil
}

// SetColumns sets the columns f
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

// SetColumns sets the columns for the table
func (this *Table) SetColumns(columns ...string) error {
	this.Columns = make([]string, 0, len(columns))
	this.colmap = make(map[string]int, len(columns))
	if err := this.AppendColumns(columns...); err != nil {
		return err
	}
	return nil
}

// NumberOfColumns returns the number of columns for the
// table
func (this *Table) NumberOfColumns() int {
	return len(this.Columns)
}

// AppendColumns appends columns onto the table
func (this *Table) AppendColumns(columns ...string) error {
	// Update columns and colmap
	for i, column := range columns {
		if _, exists := this.colmap[column]; exists {
			return ErrDuplicateColumn
		}
		this.colmap[column] = i
		this.Columns = append(this.Columns, column)
	}
	return nil
}

// TypeForColumn returns uint, int or float as a string depending
// on whether a column is all uint, int or float. It can also
// return empty string if indeterminate (empty data, for example)
func (this *Table) TypeForColumn(c string) (string, error) {
	if n, exists := this.colmap[c]; exists == false {
		return "", ErrNotFound
	} else {
		var not_float, not_uint, not_int, any bool
		for _, values := range this.Rows {
			if n >= len(values) || values[n] == nil {
				continue
			}
			// We have seen a value
			any = true
			// Check for int first
			if not_int == false {
				// check for int
				if _, err := values[n].Int64(); err != nil {
					not_int = true
				}
			}
			// Then check for uint
			if not_uint == false {
				// check for uint
				if _, err := values[n].Uint64(); err != nil {
					not_uint = true
				}
			}
			// Finally check for float
			if not_float == false {
				// check for float
				if _, err := values[n].Float64(); err != nil {
					not_float = true
				}
			}
		}
		if any == false {
			return "", ErrOutOfRange
		} else if not_int == true && not_uint == true && not_float == true {
			return "", nil
		} else if not_int == true && not_uint == true {
			return "float", nil
		} else if not_uint == true {
			return "int", nil
		} else {
			return "uint", nil
		}
	}
}

// AppendStringRow appends a row of string values onto the table
// and will return an error if the length of the string exceeds
// the number of columns. If you 
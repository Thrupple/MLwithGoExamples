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
// the number of columns. If you set treat_empty_as_nil to true
// then any string value which is only whitespace or of zero length
// is treated as nil
func (this *Table) AppendStringRow(values []string, treat_empty_as_nil bool) error {
	if len(values) > len(this.Columns) {
		return ErrDimensionError
	}
	// Create a row of values
	row := make([]*Value, len(this.Columns))
	for i := 0; i < len(values); i++ {
		if i >= len(values) {
			continue
		} else if treat_empty_as_nil && (values[i] == "" || strings.TrimSpace(values[i]) == "") {
			continue
		} else {
			row[i] = &Value{Str: values[i]}
		}
	}

	// Append row
	if this.Rows == nil {
		this.Rows = make([][]*Value, 0, 1)
	}
	this.Rows = append(this.Rows, row)

	// Return success
	return nil
}

// StringRow returns a row as an array of string values for row index n. If
// any values are nil then the nil_string is used
func (this *Table) StringRow(n int, nil_string string) ([]string, error) {
	if n < 0 || n >= len(this.Rows) {
		return nil, ErrOutOfRange
	}
	values := this.Rows[n]
	row := make([]string, len(this.Columns))
	for i := range row {
		if i >= len(values) || values[i] == nil {
			row[i] = nil_string
		} else {
			row[i] = values[i].Str
		}
	}
	return row, nil
}

// StringColumn returns all values in a specific named column, c. If
// any values are nil then the nil_string is used
func (this *Table) StringColumn(c string, nil_string string) ([]string, error) {
	if n, exists := this.colmap[c]; exists == false {
		return nil, ErrNotFound
	} else {
		column := make([]string, len(this.Rows))
		for i, values := range this.Rows {
			if n >= len(values) || values[n] == nil {
				column[i] = nil_string
			} else {
				column[i] = values[n].Str
			}
		}
		return column, nil
	}
}

// FloatColumn returns all values in a specific named column, c as float64 values. If
// any values are nil then the nil_value is used (usually 0.0). If any value cannot be
// converted to a float, then an error is returned
func (this *Table) FloatColumn(c string, nil_value float64) ([]float64, error) {
	if n, exists := this.colmap[c]; exists == false {
		return nil, ErrNotFound
	} else {
		column := make([]float64, len(this.Rows))
		for i, values := range this.Rows {
			if n >= len(values) || values[n] == nil {
				column[i] = nil_value
			} else {
				if float, err := values[n].Float64(); err != nil {
					return nil, err
				} else {
					column[i] = float
				}
			}
		}
		return column, nil
	}
}

// UintColumn returns all values in a specific named column, c as uint values. If
// any values are nil then the nil_value is used (usually 0). If any value cannot be
// converted to a uint, then an error is returned
func (this *Table) UintColumn(c string, nil_value uint) ([]uint, error) {
	if n, exists := this.colmap[c]; exists == false {
		return nil, ErrNotFound
	} else {
		column := make([]uint, len(this.Rows))
		for i, values := range this.Rows {
			if n >= len(values) || values[n] == nil {
				column[i] = nil_value
			} else {
				if value, err := values[n].Uint64(); err != nil {
					return nil, err
				} else {
					column[i] = uint(value)
				}
			}
		}
		return column, nil
	}
}

// UintPointerColumn returns all values in a specific named column, c as
// *uint values. If any values are nil then the pointer is nil. If any value cannot be
// converted to a uint, then an error is returned
func (this *Table) UintPointerColumn(c string) ([]*uint, error) {
	if n, exists := this.colmap[c]; exists == false {
		return nil, ErrNotFound
	} else {
		column := make([]*uint, len(this.Rows))
		for i, values := range this.Rows {
			if n >= len(values) || values[n] == nil {
				column[i] = nil
			} else {
				if value, err := values[n].Uint64(); err != nil {
					return nil, err
				} else {
					v := uint(value)
					column[i] = &v
				}
			}
		}
		return column, nil
	}
}

// UintValues returns unique values in a specific names column, c
// or will return error on failure
func (this *Table) UintValues(c string) ([]ui
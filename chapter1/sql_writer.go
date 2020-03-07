
// Usage:
//  go run sql_writer.go -db test.sql
package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

///////////////////////////////////////////////////////////////////////////////
// STRUCTURES

type CitibikeStationData struct {
	LastUpdated *mytime     `json:"last_updated"`
	TTL         *myduration `json:"ttl"`
	Data        *mydata     `json:"data"`
}

// We need to interpret the JSON into a time.Time structure
type mytime struct {
	t time.Time
}

// We need to interpret the JSON duration into a time.Duration structure
type myduration struct {
	d time.Duration
}

// Data in the JSON only contains an array of stations
type mydata struct {
	Stations []*station `json:"stations"`
}

// Data structure which represents a Citibike station
type station struct {
	StationId      string  `json:"station_id"`
	IsInstalled    uint    `json:"is_installed"`
	IsRenting      uint    `json:"is_renting"`
	IsReturning    uint    `json:"is_returning"`
	DocksAvailable uint    `json:"num_docks_available"`
	DocksDisabled  uint    `json:"num_docks_disabled"`
	BikesAvailable uint    `json:"num_bikes_available"`
	BikesDisabled  uint    `json:"num_bikes_disabled"`
	LastReported   *mytime `json:"last_reported"`
}

///////////////////////////////////////////////////////////////////////////////

const (
	// The URL which contains the citibike information
	CITIBIKE_URL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

	// The table name
	SQL_TABLENAME = "station"
)

var (
	FlagDatabasePath = flag.String("db", "", "Path to the database")

	SQL_COLUMNS = []string{
		"id integer not null primary key",
		"last_reported string not null",
		"is_installed bool",
		"is_renting bool",
		"docks_available integer",
		"docks_disabled integer",
		"bikes_available integer",
		"bikes_disabled integer",
	}
)

///////////////////////////////////////////////////////////////////////////////
// PARSERS AND STRINGIFY

// Unmarshall a unixtime into a time.Time structure
func (t *mytime) UnmarshalJSON(j []byte) error {
	if unixtime, err := strconv.ParseInt(string(j), 10, 64); err != nil {
		return err
	} else {
		t.t = time.Unix(unixtime, 0)
		return nil
	}
}

// Unmarshall a pure number into a time.Duration structure, where the
// number represents a second
func (d *myduration) UnmarshalJSON(j []byte) error {
	if seconds, err := strconv.ParseInt(string(j), 10, 64); err != nil {
		return err

// Usage:
//  go run json_reader.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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

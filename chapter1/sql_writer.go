
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
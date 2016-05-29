package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var ColumnName string
var ColumnNames string
var HeadersOnly bool
var ColumnStats bool

func init() {
	flag.StringVar(&ColumnName, "column-name", "", "print values for a specific column name")
	flag.StringVar(&ColumnNames, "column-names", "",
		"print values for specific column names, separated by commas")
	flag.BoolVar(&HeadersOnly, "headers-only", false, "print headers only, one per line")
	flag.BoolVar(&ColumnStats, "column-stats", false,
		`print column value length statistics: "Column Name" min mean max`)
}

func readHeaders(r *csv.Reader) []string {
	headers, err := r.Read()
	if err != nil {
		log.Fatalf("Could not read CSV header row: %s", err)
	}
	return headers
}

func printColumnValues(targetColumnNames []string, headers []string, r *csv.Reader) {
	targetColumnIdxs := make([]int, len(targetColumnNames))
	for i := range targetColumnIdxs {
		targetColumnIdxs[i] = -1
	}

	for targetIdx, targetColumnName := range targetColumnNames {
		for idx, columnName := range headers {
			if columnName == targetColumnName {
				if targetColumnIdxs[targetIdx] != -1 {
					log.Fatalf(`Duplicate column "%s"`, quote(targetColumnName))
				}
				targetColumnIdxs[targetIdx] = idx
			}
		}
		if targetColumnIdxs[targetIdx] == -1 {
			log.Fatalf(`Column "%s" not found in CSV header`, quote(targetColumnName))
		}
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		values := []string{}
		for _, idx := range targetColumnIdxs {
			values = append(values, record[idx])
		}
		fmt.Println(strings.Join(values, ","))
	}
}

func printColumnStats(headers []string, r *csv.Reader) {
	maxLengths := make([]int, len(headers))
	minLengths := make([]int, len(headers))
	for i := range minLengths {
		minLengths[i] = -1
	}
	sumLengths := make([]int, len(headers))
	n := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		for idx, val := range record {
			if len(val) > maxLengths[idx] {
				maxLengths[idx] = len(val)
			}
			if minLengths[idx] == -1 || len(val) < minLengths[idx] {
				minLengths[idx] = len(val)
			}
			sumLengths[idx] += len(val)
		}
		n++
	}

	for idx, columnName := range headers {
		fmt.Printf("\"%s\" %d %d %d\n", quote(columnName),
			minLengths[idx], sumLengths[idx]/n, maxLengths[idx])
	}
}

func main() {
	flag.Parse()
	csvFile, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvFile)
	headers := readHeaders(r)
	if HeadersOnly {
		for _, columnName := range headers {
			fmt.Println(columnName)
		}
		return
	} else if ColumnStats {
		printColumnStats(headers, r)
	} else if ColumnName != "" {
		printColumnValues([]string{ColumnName}, headers, r)
	} else if ColumnNames != "" {
		printColumnValues(strings.Split(ColumnNames, ","), headers, r)
	}
}

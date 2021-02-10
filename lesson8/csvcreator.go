package main

import (
	"encoding/csv"
	"errors"
	"os"
	"sync"
)

var mu sync.Mutex

// CreateCsvFile desc
func CreateCsvFile(data [][]string) error {
	mu.Lock()
	defer mu.Unlock()
	file, err := os.Create("csv_report.csv"); if err != nil {
		return errors.New("cannot create file")
	}
	defer file.Close()
	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	for _, rec := range data {
		if err = csvWriter.Write(rec); err != nil {
			return errors.New("cannot write date to csv file")
		}
	}
	return nil
}

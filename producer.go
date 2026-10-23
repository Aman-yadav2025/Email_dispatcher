package main

import (
	"encoding/csv"
	"os"
)

func loadRecipient(filePath string,ch chan Recipient) error {
	f, err := os.Open(filePath)

	defer close(ch)

	if err != nil {
		return err
	}

	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil{
		return nil
	}

	for _, record := range records[1:] {
		ch<-Recipient{
			Name: record[1],
			Email: record[3],
		}
	}
	return nil
}

package utils

import (
	"encoding/csv"
	"io"
)

func EncodeRecordsToCSV(recs [][]string, w io.Writer) error {
	writer := csv.NewWriter(w)

	// We might want to change this to flushing every intervals to save memory. That could be done by determining a flush limit,
	// get the size of one record and make an estimation of when we should flush.
	if err := writer.WriteAll(recs); err != nil {
		return err
	}
	writer.Flush()

	return nil
}

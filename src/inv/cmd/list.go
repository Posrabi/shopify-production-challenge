package cmd

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Posrabi/shopify-backend-project/src/common/utils"
)

func listHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	items, err := invService.ListItems(ctx)
	if err != nil {
		return err
	}

	if r.Header.Get("Content-Type") == "text/csv" {
		records := [][]string{items[0].ToCSVHeader()}
		for _, i := range items {
			records = append(records, i.ToCSVRecord())
		}
		w.Header().Set("Content-Type", "text/csv")
		return utils.EncodeRecordsToCSV(records, w)
	}

	if err := json.NewEncoder(w).Encode(items); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}

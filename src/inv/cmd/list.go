package cmd

import (
	"context"
	"encoding/json"
	"net/http"
)

func listHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := context.Background()
	items, err := invService.ListItems(ctx)
	if err != nil {
		return err
	}

	if err := json.NewEncoder(w).Encode(items); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}

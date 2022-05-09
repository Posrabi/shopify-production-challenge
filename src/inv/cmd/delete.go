package cmd

import (
	"context"
	"net/http"
)

func deleteHandler(w http.ResponseWriter, r *http.Request) error {
	itemIDs := r.URL.Query()["id"]

	ctx := context.Background()
	for _, id := range itemIDs {
		if err := invService.DeleteItem(ctx, id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	}

	return nil
}

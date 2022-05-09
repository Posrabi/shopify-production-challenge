package cmd

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Posrabi/shopify-backend-project/src/inv/pkg/entity"
)

func editHandler(w http.ResponseWriter, r *http.Request) error {
	item := entity.Item{}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	ctx := context.Background()
	if err := invService.EditItem(ctx, &item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return nil
}

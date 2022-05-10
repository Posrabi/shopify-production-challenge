package cmd

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Posrabi/shopify-backend-project/src/inv/pkg/entity"
)

func createHandler(w http.ResponseWriter, r *http.Request) error {
	request := entity.CreateItemRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil && err != io.EOF {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	ctx := context.Background()
	if err := invService.CreateItem(ctx, &request.AssignCityToItem().Item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return nil
}

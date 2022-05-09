package cmd

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Posrabi/shopify-backend-project/src/inv/pkg/entity"
)

func shipHandler(w http.ResponseWriter, r *http.Request) error {
	shipment := entity.Shipment{}
	if err := json.NewDecoder(r.Body).Decode(&shipment); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	ctx := context.Background()
	if err := invService.ShipItem(ctx, &shipment); err != nil {
		if strings.HasPrefix(err.Error(), "Not enough inventory") {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return err
	}

	return nil
}

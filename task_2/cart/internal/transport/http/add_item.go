package cartTransportHTTP

import (
	"encoding/json"
	"net/http"
)

type AddItemRequest struct {
	Count uint32 `json:"count"`
}

func (h *CartHTTPHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := parsePathValueToInt64(r, "user_id")
	if err != nil {
		sendError(w, err, "failed to parse user_id from path", http.StatusBadRequest)
		return
	}

	skuID, err := parsePathValueToInt64(r, "sku_id")
	if err != nil {
		sendError(w, err, "failed to parse sku_id from path", http.StatusBadRequest)
		return
	}

	var req AddItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, err, "failed to decode request", http.StatusBadRequest)
		return
	}

	if err := h.cartService.AddItem(ctx, userID, skuID, req.Count); err != nil {
		sendError(w, err, "failed to add item", noneStatusCode)
		return
	}

	w.WriteHeader(http.StatusOK)
}

package cartTransportHTTP

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *CartHTTPHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := parsePathValueToInt64(r, "user_id")
	if err != nil {
		sendError(w, err, "failed to parse user_id from path", http.StatusBadRequest)
		return
	}

	cart, err := h.cartService.GetCart(ctx, userID)
	if err != nil {
		sendError(w, err, "failed to get cart from service", noneStatusCode)
		return
	}

	cartDTO := cartDTOFromDomain(cart)

	body, err := json.Marshal(cartDTO)
	if err != nil {
		sendError(w, err, "failed to marshal cart DTO", noneStatusCode)
		return
	}

	if err := responseSuccessJSON(w, http.StatusOK, body); err != nil {
		log.Printf("failed to write get_cart response: %v", err)
		return
	}
}

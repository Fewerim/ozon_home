//go:build api

package test

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func newUserID() int64 {
	return rand.Int64N(1_000_000_000_000_000_000) + 1
}

func postCartItem(t *testing.T, client *http.Client, userID, sku int64, count uint32) int {
	t.Helper()

	url := fmt.Sprintf("http://127.0.0.1:8080/user/%d/cart/%d", userID, sku)
	body := fmt.Sprintf(`{"count":%d}`, count)
	request, err := http.NewRequestWithContext(
		t.Context(), http.MethodPost, url, strings.NewReader(body),
	)
	require.NoError(t, err)

	response, err := client.Do(request)
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())
	return response.StatusCode
}

func TestCartAddItemAPI(t *testing.T) {
	client := &http.Client{Timeout: 5 * time.Second}
	userID := newUserID()

	statusCode := postCartItem(t, client, userID, testSKU, 2)
	require.Equal(t, http.StatusOK, statusCode)
}

func TestCartGetCartAPI(t *testing.T) {
	client := &http.Client{Timeout: 5 * time.Second}
	userID := newUserID()
	require.Equal(t, http.StatusOK, postCartItem(t, client, userID, testSKU, 2))

	url := fmt.Sprintf("http://127.0.0.1:8080/user/%d/cart", userID)
	request, err := http.NewRequestWithContext(t.Context(), http.MethodGet, url, nil)
	require.NoError(t, err)

	response, err := client.Do(request)
	require.NoError(t, err)
	defer response.Body.Close()
	require.Equal(t, http.StatusOK, response.StatusCode)

	var cart struct {
		Items []struct {
			Sku   int64  `json:"sku"`
			Price uint32 `json:"price"`
			Count uint32 `json:"count"`
		} `json:"items"`
		TotalPrice uint32 `json:"total_price"`
	}
	require.NoError(t, json.NewDecoder(response.Body).Decode(&cart))
	require.Len(t, cart.Items, 1)
	require.Equal(t, testSKU, cart.Items[0].Sku)
	require.Equal(t, uint32(2), cart.Items[0].Count)
	require.Equal(t, cart.Items[0].Price*cart.Items[0].Count, cart.TotalPrice)
}

func TestCartMissingProductAPI(t *testing.T) {
	client := &http.Client{Timeout: 5 * time.Second}
	userID := newUserID()

	statusCode := postCartItem(t, client, userID, missingSKU, 1)
	require.Equal(t, http.StatusPreconditionFailed, statusCode)
}

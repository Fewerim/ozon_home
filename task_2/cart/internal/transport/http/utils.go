package cartTransportHTTP

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Fewerim/ozon_home/task_2/cart/internal/domain"
)

const noneStatusCode = 0

func parsePathValueToInt64(r *http.Request, name string) (int64, error) {
	return strconv.ParseInt(r.PathValue(name), 10, 64)
}

func responseSuccessJSON(w http.ResponseWriter, statusCode int, data []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(data); err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	return nil
}

func sendError(w http.ResponseWriter, err error, msg string, statusCode int) {
	var errStatusCode int
	if statusCode == noneStatusCode {
		errStatusCode = mappingErrorToStatusCode(err)
	} else {
		errStatusCode = statusCode
	}

	errMsg := fmt.Errorf("%s %w", msg, err).Error()

	http.Error(w, errMsg, errStatusCode)
}

func mappingErrorToStatusCode(err error) int {
	if errors.Is(err, domain.ErrInvalidArgument) {
		return http.StatusBadRequest
	}

	if errors.Is(err, domain.ErrNotFound) {
		return http.StatusNotFound
	}

	if errors.Is(err, domain.ErrPreconditionFailed) {
		return http.StatusPreconditionFailed
	}

	return http.StatusInternalServerError
}

package utils

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type HTTPError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Cause   string `json:"cause,omitempty"`
}

func NewHTTPError(message string, code int, err error) *HTTPError {
	return &HTTPError{
		Message: message,
		Code:    code,
		Cause:   err.Error(),
	}
}

// HandleHttpError writes a json HTTPError to the response writer
func HandleHttpError(
	ctx context.Context,
	w http.ResponseWriter,
	message string,
	code int,
	err error,
) {
	httpErr := NewHTTPError(message, code, err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpErr.Code)
	bytes, err := json.Marshal(httpErr)
	if err != nil {
		log.Printf("ERROR: Failed to Marshal ERROR JSON: %v", err)
	}
	w.Write(bytes)
}

func (e *HTTPError) Error() string {
	return e.Message
}

func HttpJsonResponse(
	ctx context.Context,
	w http.ResponseWriter,
	code int,
	data any,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	bytes, err := json.Marshal(data)
	if err != nil {
		HandleHttpError(ctx, w, "Failed to Marshal JSON", http.StatusInternalServerError, err)
		return
	}
	w.Write(bytes)
}

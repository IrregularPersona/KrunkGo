package krunkgo

import (
	"fmt"
)

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error (%d): %s", e.StatusCode, e.Message)
}

type RateLimitError struct {
	RetryAfter uint64
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("Rate limit exceeded. Retry after %d seconds", e.RetryAfter)
}

type DecodeError struct {
	Message string
	Body    string
	Field   string
}

func (e *DecodeError) Error() string {
	fieldStr := e.Field
	if fieldStr == "" {
		fieldStr = "unknown"
	}
	return fmt.Sprintf("Decode error (key: %s): %s | Body: %s", fieldStr, e.Message, e.Body)
}

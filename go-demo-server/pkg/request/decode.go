package request

import (
	"encoding/json"
	"io"
)

func decodeJSON[T any](ioBody io.ReadCloser) (*T, error) {
	var payload T
	err := json.NewDecoder(ioBody).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

package request

import (
	"errors"
	"net/http"
)

func Json[T any](r *http.Request) (*T, error) {
	body, err := decodeJSON[T](r.Body)
	if err != nil {
		return nil, errors.New("invalid json body")
	}
	err = validateStruct(body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

package request

import "github.com/go-playground/validator/v10"

func validateStruct[T any](payload T) error {
	validate := validator.New()
	return validate.Struct(payload)
}

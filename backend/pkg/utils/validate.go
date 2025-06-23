package utils

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidateStruct[T any](payload T) error {
	return validate.Struct(payload)
}

package internalerrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok || len(validationErrors) == 0 {
			return err
		}
		validationError := validationErrors[0]

		field := strings.ToLower(validationError.StructField())

		switch validationError.Tag() {
		case "required":
			return errors.New(field + " is required")
		case "min":
			return errors.New(field + " must be at least " + validationError.Param() + " characters long")
		case "max":
			return errors.New(field + " must be no more than " + validationError.Param() + " characters long")
		case "email":
			return errors.New(field + " must be a valid email")
		default:
			return errors.New("invalid field")
		}
	}
	return nil
}
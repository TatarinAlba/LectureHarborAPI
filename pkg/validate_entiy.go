package pkg

import (
	validationLib "github.com/go-playground/validator/v10"
)

func ValidateEntity(entityToValidate interface{}) error {
	validator := validationLib.New(validationLib.WithRequiredStructEnabled())
	err := validator.Struct(entityToValidate)
	if err != nil {
		return err
	}
	return nil
}

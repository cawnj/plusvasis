package templates

import (
	"github.com/go-playground/validator/v10"
)

func Validate(job NomadJob) error {
	validate := validator.New()
	err := validate.RegisterValidation("kvPairs", validateKvPairs)
	if err != nil {
		return err
	}
	return validate.Struct(job)
}

func validateKvPairs(fl validator.FieldLevel) bool {
	env := fl.Field().Interface().([][]string)
	for _, keyValue := range env {
		// check we have a key and value
		if len(keyValue) != 2 {
			return false
		}
		// check key and value are not empty
		if keyValue[0] == "" || keyValue[1] == "" {
			return false
		}
	}
	return true
}

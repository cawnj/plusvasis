package templates

import (
	"github.com/go-playground/validator/v10"
)

func Validate(job NomadJob) error {
	validate := validator.New()
	return validate.Struct(job)
}

package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var V = validator.New()

func ValidateStruct(s interface{}) error {
	return V.Struct(s)
}

func RegisterCustomRules() {
	V.RegisterValidation("pan", PanValidation)
	V.RegisterValidation("phone", PhoneValidation)
	V.RegisterValidation("email", EmailValidation)
}

func PanValidation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[A-Z]{5}[0-9]{4}[A-Z]$`)
	return re.MatchString(fl.Field().String())
}

func PhoneValidation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[6-9][0-9]{9}$`)
	return re.MatchString(fl.Field().String())
}

func EmailValidation(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(fl.Field().String())
}

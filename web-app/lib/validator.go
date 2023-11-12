package lib

import (
	"unicode"

	"github.com/golodash/galidator"
)

var g = galidator.New()

func GetValidator(i any) galidator.Validator {
	return g.Validator(i)
}

func GetCustomValidator(i any, cv galidator.Validators) galidator.Validator {
	return g.CustomValidators(cv).Validator(i)
}

func ValidateStrongPass(field any) bool {
	str, ok := field.(string)

	if !ok {
		return false
	}

	letters := 0
	number := false
	upper := false
	special := false

	for _, c := range str {
		letters++

		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c):
			special = true
		default:
		}
	}

	if letters > 7 && number && upper && special {
		return true
	}
	return false
}

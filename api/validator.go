package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/ulugbek0217/simple_bank/util"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsValidCurrency(currency)
	}
	return false
}

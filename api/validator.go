package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/krysopath/fluteapi/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}

var validEmailDomain validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if email, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsResolvableEmail(email)
	}
	return false
}

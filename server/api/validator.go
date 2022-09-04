package api

import (
	util "github.com/OktarianTB/stock-trading-simulator-golang/utils"
	"github.com/go-playground/validator/v10"
)

var validFrequency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if frequency, ok := fieldLevel.Field().Interface().(string); ok {
		// check frequency is supported
		return util.IsValidFrequency(frequency)
	}
	return false
}

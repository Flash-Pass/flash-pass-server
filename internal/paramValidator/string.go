package paramValidator

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
)

const (
	MobilePattern         = "^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$"
	StringIsMobileTag     = "stringIsMobile"
	StringNotBothEmptyTag = "stringNotBothEmpty"
)

func stringIsMobile(field validator.FieldLevel) bool {
	mobile := field.Field().String()
	ok, err := regexp.MatchString(MobilePattern, mobile)
	if err != nil {
		return false
	}
	return ok
}

func notBothEmpty(field validator.FieldLevel) bool {
	curr := field.Field().String()
	kind := field.Field().Kind()

	if kind != reflect.String {
		return false
	}

	if curr != "" {
		return true
	}

	related, relatedKind, ok := field.GetStructFieldOK()
	if !ok || relatedKind != kind {
		return false
	}

	return related.String() != ""
}

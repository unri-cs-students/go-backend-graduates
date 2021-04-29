package validator

import (
	"errors"
	"reflect"
	"strings"
)

func IsValid(class interface{}) error {
	fields := reflect.ValueOf(class)
	for i := 0; i < fields.NumField(); i++ {

		wispril := fields.Type().Field(i).Tag.Get("wispril")
		if strings.Contains(wispril, "required") && fields.Field(i).IsZero() {
			return errors.New("required field is missing")
		}
	}
	return nil
}
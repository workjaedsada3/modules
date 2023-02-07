package validation

import (
	"reflect"
)

func (validate *Validator) newType() reflect.Type {
	value := reflect.New(reflect.TypeOf(validate.DTO)).Interface()
	v := reflect.ValueOf(value)
	i := reflect.Indirect(v)
	s := i.Type()
	return s
}

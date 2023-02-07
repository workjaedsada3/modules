package validation

import (
	"fmt"
	"regexp"
)

func (validate *Validator) checkString(value interface{}) bool {
	xType := fmt.Sprintf("%T", value)
	if value != nil {
		if xType != "string" {
			return false
		}
	}
	return true
}
func (validate *Validator) checkNumber(value interface{}) bool {
	xType := fmt.Sprintf("%T", value)
	return xType == "number"
}

// /
func (validate *Validator) checkRequired(value interface{}) bool {
	value = fmt.Sprintf("%v", value)
	return value != ""
}

func (validate *Validator) checkRequiredArray(value interface{}) bool {
	return true
}

func (validate *Validator) checkMax(value interface{}, length int) bool {
	// value =
	return len([]rune(fmt.Sprintf("%s", value))) <= length
}

func (validate *Validator) checkPattern(value interface{}, pattern string) bool {
	var IsLetter = regexp.MustCompile(pattern).MatchString
	return IsLetter(fmt.Sprintf("%s", value))
}

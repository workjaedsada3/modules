package validation

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func (validate *Validator) Validator() *Validator {
	value := validate.newType().Elem()
	numFields := value.NumField()
	for i := 0; i < numFields; i++ {
		// string1 := strings.SplitN(value.Field(i).Tag.Get("validator"), ",", 2)[0]
		string2 := strings.Split(value.Field(i).Tag.Get("validator"), "|")
		if (string2)[0] != "" {
			for j := 0; j < len(string2); j++ {
				validate.checkValidate(value.Field(i).Tag.Get("json"), string2[j], value.Field(i))
			}
		}
	}
	return validate
}

func (validate *Validator) ArrayValidator() *Validator {
	// value := validate.newType().Elem()
	va := reflect.New(reflect.TypeOf(validate.ArrDTO)).Interface()
	v := reflect.ValueOf(va)
	i := reflect.Indirect(v)
	value := i.Type()

	numFields := value.NumField()
	for i := 0; i < numFields; i++ {
		// string1 := strings.SplitN(value.Field(i).Tag.Get("validator"), ",", 2)[0]
		string2 := strings.Split(value.Field(i).Tag.Get("validator"), "|")
		if (string2)[0] != "" {
			for j := 0; j < len(string2); j++ {
				validate.checkValidate(value.Field(i).Tag.Get("json"), string2[j], value.Field(i))
			}
		}
	}
	return validate
}

///

func (validate *Validator) checkValidate(name string, values interface{}, field reflect.StructField) {
	message1 := strings.SplitN(field.Tag.Get("validator_message"), ",", 2)[0]
	message := strings.Split(message1, "|")[0]
	message2 := strings.Split(message1, "|")
	datas := convertStructToMap(validate.DTO)
	length := 0
	pattern := ""

	if strings.Split(fmt.Sprintf("%s", values), "=")[0] == "max" {
		intVar, _ := strconv.Atoi(strings.Split(fmt.Sprintf("%s", values), "=")[1])
		length = intVar
		values = strings.Split(fmt.Sprintf("%s", values), "=")[0]
	} else if strings.Split(fmt.Sprintf("%s", values), "=")[0] == "min" {
		intVar, _ := strconv.Atoi(strings.Split(fmt.Sprintf("%s", values), "=")[1])
		length = intVar
		values = strings.Split(fmt.Sprintf("%s", values), "=")[0]
	} else if strings.Split(fmt.Sprintf("%s", values), "=")[0] == "pattern" {
		pattern = strings.Split(fmt.Sprintf("%s", values), "=")[1]
		values = strings.Split(fmt.Sprintf("%s", values), "=")[0]
	}

	switch values {
	case "string":
		error := validate.checkString(datas[name])
		msg := validate.customMessage(values, message, message2, fmt.Sprintln("must be string"))

		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
	case "max": // max
		error := validate.checkMax(datas[name], length)
		msg := validate.customMessage(values, message, message2, fmt.Sprintf("maximum length %d", length))
		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
	case "letter": // max
		error := validate.checkPattern(datas[name], `^[a-zA-Z]+$`)
		msg := validate.customMessage(values, message, message2, "must be letter only.")
		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
	case "alpha_num": // max
		error := validate.checkPattern(datas[name], `^[a-zA-Z0-9]*$`)
		msg := validate.customMessage(values, message, message2, "must be letter or digit only")
		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
	case "pattern": // pattern
		error := validate.checkPattern(datas[name], pattern)
		msg := validate.customMessage(values, message, message2, "field invalid pattern.")
		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
	case "datetime":
		error := validate.checkPattern(datas[name], `(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2})`)
		msg := validate.customMessage(values, message, message2, "invalid datetime format (YYYY-dd-mm HH:ii).")
		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
	case "date":
		error := validate.checkPattern(datas[name], `(\d{4})-(\d{2})-(\d{2})$`)
		msg := validate.customMessage(values, message, message2, "invalid date format (YYYY-dd-mm).")
		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
	case "required":
		error := validate.checkRequired(datas[name])
		msg := validate.customMessage(values, message, message2, "field is required.")
		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
	case "required_array":
		error := validate.checkRequiredArray(datas[name])
		msg := validate.customMessage(values, message, message2)
		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
	case "number":
		error := validate.checkNumber(datas[name])
		msg := validate.customMessage(values, message, message2)

		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}
	case "email":
		error := validate.checkPattern(datas[name], `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
		msg := validate.customMessage(values, message, message2, "invalid email format.")
		if !error {
			validate.Error = append(validate.Error, map[string]interface{}{
				"input":   name,
				"message": msg,
			})
		}

	}
}

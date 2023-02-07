package validation

import (
	"fmt"
	"reflect"
	"strings"
)

func (validate *Validator) customMessage(values interface{}, message string, message2 []string, custom_message ...string) string {
	var msg string
	if len(message2) > 0 {
		for i := 0; i < len(message2); i++ {
			check2 := strings.Contains(message2[i], fmt.Sprintf("%s", values))
			if check2 {
				msg = strings.Split(message2[i], fmt.Sprintf("%s:", values))[1]
			}
		}
	}
	if msg == "" {
		msg = fmt.Sprintf(" %s", values)
	}
	if len(custom_message) > 0 {
		msg = custom_message[0]
	}
	return msg
}

func convertStructToMap(st interface{}) map[string]interface{} {

	reqRules := make(map[string]interface{})

	v := reflect.ValueOf(st)
	t := reflect.TypeOf(st)
	// log.Println(t)
	// reqRules[key] = v.FieldByName(t.Field())
	for i := 0; i < v.Elem().NumField(); i++ {
		key := strings.ToLower(t.Elem().Field(i).Tag.Get("json"))
		// typ := v.FieldByName(t.Elem().Field(i).Name).Kind().String()
		// structTag := t.Elem().Field(i).Tag.Get("json")
		// jsonName := strings.TrimSpace(strings.Split(structTag, ",")[0])
		value := v.Elem().FieldByName(t.Elem().Field(i).Name)
		// log.Println(structTag)
		// log.Println(jsonName)
		// log.Println(typ)
		// if jsonName != "" && jsonName != "-" {
		// 	key = jsonName
		// }
		reqRules[key] = value

		// if typ == "string" {
		// 	if !(value.String() == "" && strings.Contains(structTag, "omitempty")) {
		// 		fmt.Println(key, value)
		// 		fmt.Println(key, value.String())
		// 	}
		// } else if typ == "int" {
		// 	reqRules[key] = value.Int()
		// } else {
		// 	reqRules[key] = value.Interface()
		// }

	}
	// for i := 0; i < v2.NumField(); i++ {
	// 	key := strings.ToLower(t.Field(i).Name)
	// 	typ := v.FieldByName(t.Field(i).Name).Kind().String()
	// 	structTag := t.Field(i).Tag.Get("json")
	// 	jsonName := strings.TrimSpace(strings.Split(structTag, ",")[0])
	// 	value := v.FieldByName(t.Field(i).Name)

	// 	// if jsonName is not empty use it for the key
	// 	if jsonName != "" && jsonName != "-" {
	// 		key = jsonName
	// 	}

	// 	if typ == "string" {
	// 		if !(value.String() == "" && strings.Contains(structTag, "omitempty")) {
	// 			fmt.Println(key, value)
	// 			fmt.Println(key, value.String())
	// 			reqRules[key] = value.String()
	// 		}
	// 	} else if typ == "int" {
	// 		reqRules[key] = value.Int()
	// 	} else {
	// 		reqRules[key] = value.Interface()
	// 	}

	// }
	return reqRules
}

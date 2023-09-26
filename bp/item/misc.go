package item

import (
	"fmt"
	"reflect"
	"strings"
)

func convertMapToStruct(m map[string]interface{}, s interface{}) error {
	structValue := reflect.ValueOf(s).Elem()
	structType := structValue.Type()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := structType.Field(i)
		fieldName := strings.Split(fieldType.Tag.Get("json"), ",")[0]

		if val, ok := m[fieldName]; ok {
			// Check if the field is a pointer and the map value is a pointer
			if field.Kind() == reflect.Ptr && reflect.TypeOf(val).Kind() == reflect.Ptr {
				// Create a new instance of the field's underlying type
				field.Set(reflect.New(fieldType.Type.Elem()))

				// Set the value of the field to the dereferenced map value
				field.Elem().Set(reflect.ValueOf(val).Elem())
			} else {
				// Set the value of the field directly
				switch fieldType.Type.Kind() {
				case reflect.Int:
					if v, ok := val.(float64); ok {
						field.SetInt(int64(v))
					}
				default:
					field.Set(reflect.ValueOf(val))
				}
				// field.Set(reflect.ValueOf(val))
			}
		}
	}

	return nil
}

func PascalToLower(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] >= 'A' && s[i] <= 'Z' {
			result += "_" + string(s[i]+32)
		} else {
			result += string(s[i])
		}
	}
	return strings.ToLower(result)
}

func GetComponentName(component interface{}) string {
	// Check if component is a string
	if reflect.TypeOf(component).Kind() == reflect.String {
		return component.(string)
	}
	t := reflect.TypeOf(component)
	var name string
	switch t.Kind() {
	case reflect.Ptr:
		name = t.Elem().Name()
	case reflect.Struct:
		name = t.Name()
	default:
		panic(fmt.Sprintf("invalid component type: %s", t.Kind()))
	}
	name = strings.Replace(name, "_", ".", -1)
	name = PascalToLower(name)
	name = "minecraft:" + name
	name = strings.Replace(name, "._", ".", -1)
	return name
}
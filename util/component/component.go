package component

import (
	"fmt"
	"reflect"
	"strings"
)

func ConvertMapToStruct(m map[string]interface{}, s interface{}) error {
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
			} else if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Int {
				if v, ok := val.(float64); ok {
					intV := int(v)
					field.Set(reflect.ValueOf(&intV))
				}
			} else {
				// Set the value of the field directly
				switch fieldType.Type.Kind() {
				case reflect.Int:
					if v, ok := val.(float64); ok {
						field.SetInt(int64(v))
					}
				case reflect.Slice:
					if elements, ok := val.([]interface{}); ok {
						sliceType := fieldType.Type.Elem()
						slice := reflect.MakeSlice(fieldType.Type, len(elements), len(elements))

						for j, elem := range elements {
							switch elem.(type) {
							case map[string]interface{}:
								nestedStruct := reflect.New(sliceType).Interface()
								err := ConvertMapToStruct(elem.(map[string]interface{}), nestedStruct)
								if err != nil {
									return err
								}
								slice.Index(j).Set(reflect.ValueOf(nestedStruct).Elem())
							default:
								slice.Index(j).Set(reflect.ValueOf(elem))
							}
						}

						field.Set(slice)
					} else {
						return fmt.Errorf("map value for slice field %s is not an array", fieldName)
					}
				case reflect.Ptr:
					if reflect.TypeOf(val).ConvertibleTo(fieldType.Type) {
						field.Set(reflect.ValueOf(val).Convert(fieldType.Type))
					} else {
						return fmt.Errorf("map value for field %s is not convertible to type %s", fieldName, fieldType.Type)
					}
				default:
					field.Set(reflect.ValueOf(val))
				}
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
	t := reflect.TypeOf(component)
	if t.Kind() == reflect.String && t.Name() == "string" {
		return component.(string)
	}
	var name string
	switch t.Kind() {
	case reflect.Ptr:
		name = t.Elem().Name()
	default:
		name = t.Name()
	}
	name = strings.Replace(name, "_", ".", -1)
	name = PascalToLower(name)
	name = "minecraft:" + name
	name = strings.Replace(name, "._", ".", -1)
	return name
}

func Get(old interface{}, new interface{}) (interface{}, error) {
	oldType := reflect.TypeOf(old)
	newType := reflect.TypeOf(new)
	// We have 4 cases:
	// 1. old is a pointer and new is not a pointer
	// 2. old is not a pointer and new is a pointer
	// 3. old is a pointer and new is a pointer
	// 4. old is not a pointer and new is not a pointer
	// If the type is match, return it
	if newType == oldType {
		return new, nil
	}
	// Convert map to struct
	err := ConvertMapToStruct(old.(map[string]interface{}), new)
	if err != nil {
		return nil, err
	}
	// Assign new to the struct
	old = new
	return new, nil
}

package entityExtractor

import (
	"gorm.io/gorm"
	"reflect"
	"strings"
)

// ConvertEntityToMap converts any entity struct to a map[string]interface{}
func ConvertEntityToMap(entity interface{}) map[string]interface{} {
	v := reflect.ValueOf(entity)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return nil
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return nil
	}

	result := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldValue := v.Field(i)

		// Ignore fields from gorm.Model
		if field.Anonymous && field.Type == reflect.TypeOf(gorm.Model{}) {
			continue
		}

		// Always include boolean fields
		if fieldValue.Kind() == reflect.Bool {
			tag := strings.Replace(field.Tag.Get("gorm"), "column:", "", -1)
			if tag != "" {
				result[tag] = fieldValue.Interface()
			} else {
				result[field.Name] = fieldValue.Interface()
			}
			continue
		}

		// Ignore zero-value fields, including nil values
		if fieldValue.IsZero() {
			continue
		}

		// Use the gorm tag if available, otherwise use the field name
		tag := strings.Replace(field.Tag.Get("gorm"), "column:", "", -1)
		if tag != "" {
			result[tag] = fieldValue.Interface()
		} else {
			result[field.Name] = fieldValue.Interface()
		}
	}
	return result
}

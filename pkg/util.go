package pkg

import (
	"fmt"
	"reflect"
)

// MapToStruct converts a map[string]interface{} to a struct.
// The map keys must match the struct field names.
func MapToStruct(m map[string]interface{}, s interface{}) error {
	// Get the value of the interface as a reflect.Value
	v := reflect.ValueOf(s)
	// Check that the argument is a pointer to a struct
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("argument must be a pointer to a struct, got %T", s)
	}
	// Get the element that the pointer points to
	v = v.Elem()
	// Get the type of the element
	t := v.Type()
	// Iterate over the fields in the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// Check if the field has a "db" tag
		name := field.Tag.Get("db")
		// If no "db" tag is present, use the field name as the map key
		if name == "" {
			name = field.Name
		}
		// Check if the map has a value for the key
		if value, ok := m[name]; ok {
			// Set the field value to the value in the map
			v.Field(i).Set(reflect.ValueOf(value))
		}
	}
	return nil
}

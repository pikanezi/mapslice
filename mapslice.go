// mapslice maps a specific field in a slice of structs to a slice of native types.
package mapslice

import (
	"fmt"
	"github.com/gocarina/structs"
	"reflect"
)

var (
	// ErrNotSlice happens when the value passed is not a slice.
	ErrNotSlice = fmt.Errorf("not slice")

	// ErrNotString happens when the value of the field is not a string.
	ErrNotString = fmt.Errorf("not string")

	// ErrNotInt happens when the value of the field is not an int.
	ErrNotInt = fmt.Errorf("not int")

	// ErrNotFloat happens when the value of the field is not a float64.
	ErrNotFloat = fmt.Errorf("not float64")

	// ErrNotBool happens when the value of the field is not a bool.
	ErrNotBool = fmt.Errorf("not bool")
)

// MapSliceToString maps a field to a slice of string.
func MapSliceToString(slice interface{}, fieldName string) (s []string, err error) {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(slice)
		for i := 0; i < val.Len(); i++ {
			fields := structs.Fields(val.Index(i).Interface())
			for _, f := range fields {
				if f.IsExported() && f.Name() == fieldName {
					e := f.Value()
					switch e.(type) {
					case string:
						s = append(s, e.(string))
					default:
						return nil, ErrNotString
					}
				}
			}
		}
	default:
		return nil, ErrNotSlice
	}
	return
}

// MapSliceToInt maps a field to a slice of int.
func MapSliceToInt(slice interface{}, fieldName string) (s []int, err error) {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(slice)
		for i := 0; i < val.Len(); i++ {
			fields := structs.Fields(val.Index(i).Interface())
			for _, f := range fields {
				if f.IsExported() && f.Name() == fieldName {
					e := f.Value()
					switch e.(type) {
					case int:
						s = append(s, e.(int))
					default:
						return nil, ErrNotInt
					}
				}
			}
		}
	default:
		return nil, ErrNotSlice
	}
	return
}

// MapSliceToFloat maps a field to a slice of int.
func MapSliceToFloat(slice interface{}, fieldName string) (s []float64, err error) {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(slice)
		for i := 0; i < val.Len(); i++ {
			fields := structs.Fields(val.Index(i).Interface())
			for _, f := range fields {
				if f.IsExported() && f.Name() == fieldName {
					e := f.Value()
					switch e.(type) {
					case float64:
						s = append(s, e.(float64))
					default:
						return nil, ErrNotFloat
					}
				}
			}
		}
	default:
		return nil, ErrNotSlice
	}
	return
}

// MapSliceToBool maps a field to a slice of bool.
func MapSliceToBool(slice interface{}, fieldName string) (s []bool, err error) {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		val := reflect.ValueOf(slice)
		for i := 0; i < val.Len(); i++ {
			fields := structs.Fields(val.Index(i).Interface())
			for _, f := range fields {
				if f.IsExported() && f.Name() == fieldName {
					e := f.Value()
					switch e.(type) {
					case bool:
						s = append(s, e.(bool))
					default:
						return nil, ErrNotBool
					}
				}
			}
		}
	default:
		return nil, ErrNotSlice
	}
	return
}

// MapSliceToStringUnsafe maps a field to a slice of string but not returns an error.
// If an error still happens, s will be nil.
func MapSliceToStringUnsafe(slice interface{}, fieldName string) (s []string) {
	s, _ = MapSliceToString(slice, fieldName)
	return
}

// MapSliceToIntUnsafe maps a field to a slice of int but not returns an error.
// If an error still happens, s will be nil.
func MapSliceToIntUnsafe(slice interface{}, fieldName string) (s []int) {
	s, _ = MapSliceToInt(slice, fieldName)
	return
}

// MapSliceToFloatUnsafe maps a field to a slice of float but not returns an error.
// If an error still happens, s will be nil.
func MapSliceToFloatUnsafe(slice interface{}, fieldName string) (s []float64) {
	s, _ = MapSliceToFloat(slice, fieldName)
	return
}

// MapSliceToBoolUnsafe maps a field to a slice of bool but not returns an error.
// If an error still happens, s will be nil.
func MapSliceToBoolUnsafe(slice interface{}, fieldName string) (s []bool) {
	s, _ = MapSliceToBool(slice, fieldName)
	return
}

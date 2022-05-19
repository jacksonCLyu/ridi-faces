package configer

import (
	"fmt"
	"reflect"
	"time"
)

// FiledType custom field type
type FieldType uint8

const (
	// FiledTypeUnknown is the default value for a field type.
	FiledTypeUnknown FieldType = iota
	// FieldTypeString is a field type for a string.
	FieldTypeString
	// FieldTypeStringSlice is a field type for a slice of strings.
	FieldTypeStringSlice
	// FieldTypeInt is a field type for an integer.
	FieldTypeInt
	// FieldTypeIntSlice is a field type for a slice of integers.
	FieldTypeIntSlice
	// FieldTypeUint is a field type for an unsigned integer.
	FieldTypeBool
	// FieldTypeBoolSlice is a field type for a slice of bools.
	FieldTypeBoolSlice
	// FieldTypeFloat is a field type for a float.
	FieldTypeFloat
	// FieldTypeFloatSlice is a field type for a slice of floats.
	FieldTypeFloatSlice
	// FieldTypeDuration is a field type for a duration.
	FieldTypeDuration
	// FieldTypeTime is a field type for a time.
	FieldTypeTime
	// FieldTypeSection is a field type for a section.
	FieldTypeSection
)

// String returns the string representation of the field.
func (t FieldType) String() string {
	switch t {
	case FieldTypeString:
		return "string"
	case FieldTypeStringSlice:
		return "string slice"
	case FieldTypeIntSlice:
		return "int slice"
	case FieldTypeInt:
		return "int"
	case FieldTypeBool:
		return "bool"
	case FieldTypeBoolSlice:
		return "bool slice"
	case FieldTypeFloat:
		return "float"
	case FieldTypeFloatSlice:
		return "float slice"
	case FieldTypeDuration:
		return "duration"
	case FieldTypeTime:
		return "time"
	case FieldTypeSection:
		return "section"
	default:
		return "unknown"
	}
}

// Field is a configuration field.
type Field struct {
	Type  FieldType
	Value any
}

// Atof convert any to Field
func Atof(value any) Field {
	if reflect.TypeOf(value) == nil {
		return Field{Type: FiledTypeUnknown, Value: value}
	}
	valueOf := reflect.ValueOf(value)
	valueKind := valueOf.Kind()
	if (valueKind == reflect.Interface || valueKind == reflect.Ptr) && !valueOf.IsNil() {
		valueOf = valueOf.Elem()
		valueKind = valueOf.Kind()
	}
	switch valueKind {
	case reflect.String:
		return Field{Type: FieldTypeString, Value: valueOf.String()}
	case reflect.Bool:
		return Field{Type: FieldTypeBool, Value: valueOf.Bool()}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return Field{Type: FieldTypeInt, Value: valueOf.Int()}
	case reflect.Float32, reflect.Float64:
		return Field{Type: FieldTypeFloat, Value: valueOf.Float()}
	case reflect.Slice:
		// slice length
		sliceLen := valueOf.Len()
		// slice type
		vvt := valueOf.Type()
		// slice element type
		etm := vvt.Elem()
		// slice element kind
		etk := etm.Kind()
		if etk == reflect.Interface || etk == reflect.Ptr {
			// if slice element type is interface or pointer, get the type of the interface or pointer
			elv := valueOf.Index(0).Elem()
			etk = elv.Kind()
		}
		if etk == reflect.String {
			slice := make([]string, sliceLen)
			for i := 0; i < sliceLen; i++ {
				slice[i] = valueOf.Index(i).Elem().String()
			}
			return Field{Type: FieldTypeStringSlice, Value: slice}
		}
		if etk == reflect.Bool {
			slice := make([]bool, sliceLen)
			for i := 0; i < sliceLen; i++ {
				slice[i] = valueOf.Index(i).Elem().Bool()
			}
			return Field{Type: FieldTypeBoolSlice, Value: slice}
		}
		if etk == reflect.Int || etk == reflect.Int8 || etk == reflect.Int16 || etk == reflect.Int32 || etk == reflect.Int64 || etk == reflect.Uint || etk == reflect.Uint8 || etk == reflect.Uint16 || etk == reflect.Uint32 || etk == reflect.Uint64 || etk == reflect.Uintptr {
			slice := make([]int64, sliceLen)
			for i := 0; i < sliceLen; i++ {
				slice[i] = valueOf.Index(i).Elem().Int()
			}
			return Field{Type: FieldTypeIntSlice, Value: slice}
		}
		if etk == reflect.Float32 || etk == reflect.Float64 {
			slice := make([]float64, sliceLen)
			for i := 0; i < sliceLen; i++ {
				slice[i] = valueOf.Index(i).Elem().Float()
			}
			return Field{Type: FieldTypeFloatSlice, Value: slice}
		}
		return Field{Type: FiledTypeUnknown, Value: value}
	case reflect.Map:
		vMap := value.(map[any]any)
		subMap := make(map[string]Field, len(vMap))
		for k, v := range vMap {
			subMap[fmt.Sprint(k)] = Atof(v)
		}
		return Field{Type: FieldTypeSection, Value: subMap}
	case reflect.TypeOf(time.Duration(0)).Kind():
		return Field{Type: FieldTypeDuration, Value: value}
	case reflect.TypeOf(time.Now()).Kind():
		return Field{Type: FieldTypeTime, Value: value}
	default:
		return Field{Type: FiledTypeUnknown, Value: value}
	}
}

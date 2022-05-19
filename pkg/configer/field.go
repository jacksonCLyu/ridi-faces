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
	typeOf := reflect.TypeOf(value)
	if typeOf == nil {
		return Field{Type: FiledTypeUnknown, Value: value}
	}
	vtk := typeOf.Kind()
	if vtk == reflect.Ptr {
		typeOf = typeOf.Elem()
		vtk = typeOf.Kind()
	}
	vv := reflect.ValueOf(value)
	if vtk == reflect.Interface && !vv.IsNil() {
		vv = vv.Elem()
		typeOf = vv.Type()
		vtk = typeOf.Kind()
	}
	switch vtk {
	case reflect.String:
		return Field{Type: FieldTypeString, Value: value}
	case reflect.Bool:
		return Field{Type: FieldTypeBool, Value: value}
	case reflect.Slice:
		etm := typeOf.Elem()
		etk := etm.Kind()
		if etk == reflect.Interface {
			etv := reflect.ValueOf(etm)
			etk = etv.Kind()
		}
		if etk == reflect.String {
			return Field{Type: FieldTypeStringSlice, Value: value}
		}
		if etk == reflect.Bool {
			return Field{Type: FieldTypeBoolSlice, Value: value}
		}
		if etk == reflect.Int || etk == reflect.Int8 || etk == reflect.Int16 || etk == reflect.Int32 || etk == reflect.Int64 || etk == reflect.Uint || etk == reflect.Uint8 || etk == reflect.Uint16 || etk == reflect.Uint32 || etk == reflect.Uint64 || etk == reflect.Uintptr{
			return Field{Type: FieldTypeIntSlice, Value: value}
		}
		if etk == reflect.Float32 || etk == reflect.Float64 {
			return Field{Type: FieldTypeFloatSlice, Value: value}
		}
		if etk == reflect.Struct {
			return Field{Type: FieldTypeSection, Value: value}
		}
		return Field{Type: FiledTypeUnknown, Value: value}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return Field{Type: FieldTypeInt, Value: value}
	case reflect.Float32, reflect.Float64:
		return Field{Type: FieldTypeFloat, Value: value}
	case reflect.Struct:
		return Field{Type: FieldTypeSection, Value: value}
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
	//switch value := value.(type) {
	//case Field:
	//	return value
	//case string:
	//	return Field{Type: FieldTypeString, Value: value}
	//case []string:
	//	return Field{Type: FieldTypeStringSlice, Value: value}
	//case bool:
	//	return Field{Type: FieldTypeBool, Value: value}
	//case []bool:
	//	return Field{Type: FieldTypeBoolSlice, Value: value}
	//case uint, uint32, uint64, int, int32, int64:
	//	return Field{Type: FieldTypeInt, Value: value}
	//case []uint, []uint32, []uint64, []int, []int32, []int64:
	//	return Field{Type: FieldTypeIntSlice, Value: value}
	//case float32, float64:
	//	return Field{Type: FieldTypeFloat, Value: value}
	//case []float32, []float64:
	//	return Field{Type: FieldTypeFloatSlice, Value: value}
	//case time.Duration:
	//	return Field{Type: FieldTypeDuration, Value: value}
	//case time.Time:
	//	return Field{Type: FieldTypeTime, Value: value}
	//case map[string]any:
	//	subMap := make(map[string]Field)
	//	for key, value := range value {
	//		subMap[key] = Atof(value)
	//	}
	//	return Field{Type: FieldTypeSection, Value: subMap}
	//case map[any]any:
	//	subMap := make(map[string]Field)
	//	for key, value := range value {
	//		subMap[fmt.Sprint(key)] = Atof(value)
	//	}
	//	return Field{Type: FieldTypeSection, Value: subMap}
	//default:
	//	return Field{Type: FiledTypeUnknown, Value: value}
	//}
}

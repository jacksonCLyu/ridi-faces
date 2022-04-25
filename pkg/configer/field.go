package configer

import (
	"fmt"
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
	FieldTypeUint
	// FieldTypeUintSlice is a field type for a slice of unsigned integers.
	FieldTypeUintSlice
	// FieldTypeBool is a field type for a bool.
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
	case FieldTypeUint:
		return "uint"
	case FieldTypeUintSlice:
		return "uint slice"
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
	switch value := value.(type) {
	case Field:
		return value
	case string:
		return Field{Type: FieldTypeString, Value: value}
	case []string:
		return Field{Type: FieldTypeStringSlice, Value: value}
	case bool:
		return Field{Type: FieldTypeBool, Value: value}
	case []bool:
		return Field{Type: FieldTypeBoolSlice, Value: value}
	case int:
	case int32:
	case int64:
		return Field{Type: FieldTypeInt, Value: value}
	case []int:
	case []int32:
	case []int64:
		return Field{Type: FieldTypeIntSlice, Value: value}
	case uint:
	case uint32:
	case uint64:
		return Field{Type: FieldTypeUint, Value: value}
	case []uint:
	case []uint32:
	case []uint64:
		return Field{Type: FieldTypeUintSlice, Value: value}
	case float32:
	case float64:
		return Field{Type: FieldTypeFloat, Value: value}
	case []float32:
		return Field{Type: FieldTypeFloatSlice, Value: value}
	case []float64:
		return Field{Type: FieldTypeFloatSlice, Value: value}
	case time.Duration:
		return Field{Type: FieldTypeDuration, Value: value}
	case time.Time:
		return Field{Type: FieldTypeTime, Value: value}
	case map[string]any:
		subMap := make(map[string]Field)
		for key, value := range value {
			subMap[key] = Atof(value)
		}
		return Field{Type: FieldTypeSection, Value: subMap}
	case map[any]any:
		subMap := make(map[string]Field)
		for key, value := range value {
			subMap[fmt.Sprint(key)] = Atof(value)
		}
		return Field{Type: FieldTypeSection, Value: subMap}
	default:
		return Field{Type: FiledTypeUnknown, Value: value}
	}
}

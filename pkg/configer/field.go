package configer

import (
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
	// FieldTypeInt32 is a field type for an integer.
	FieldTypeInt32
	// FieldTypeInt32Slice is a field type for a slice of integers.
	FieldTypeInt32Slice
	// FieldTypeInt64 is a field type for an integer.
	FieldTypeInt64
	// FieldTypeInt64Slice is a field type for a slice of integers.
	FieldTypeInt64Slice
	// FieldTypeUint is a field type for an unsigned integer.
	FieldTypeUint
	// FieldTypeUintSlice is a field type for a slice of unsigned integers.
	FieldTypeUintSlice
	// FieldTypeUint32 is a field type for an unsigned integer.
	FieldTypeUint32
	// FieldTypeUint32Slice is a field type for a slice of unsigned integers.
	FieldTypeUint32Slice
	// FieldTypeUint64 is a field type for an unsigned integer.
	FieldTypeUint64
	// FieldTypeUint64Slice is a field type for a slice of unsigned integers.
	FieldTypeUint64Slice
	// FieldTypeBool is a field type for a bool.
	FieldTypeBool
	// FieldTypeBoolSlice is a field type for a slice of bools.
	FieldTypeBoolSlice
	// FieldTypeFloat32 is a field type for a float.
	FieldTypeFloat32
	// FieldTypeFloat32Slice is a field type for a slice of floats.
	FieldTypeFloat32Slice
	// FieldTypeFloat64 is a field type for a float.
	FieldTypeFloat64
	// FieldTypeFloat64Slice is a field type for a slice of floats.
	FieldTypeFloat64Slice
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
	case FieldTypeInt:
		return "int"
	case FieldTypeIntSlice:
		return "int slice"
	case FieldTypeInt32:
		return "int32"
	case FieldTypeInt64:
		return "int64"
	case FieldTypeUint:
		return "uint"
	case FieldTypeUint32:
		return "uint32"
	case FieldTypeUint64:
		return "uint64"
	case FieldTypeBool:
		return "bool"
	case FieldTypeFloat32:
		return "float32"
	case FieldTypeFloat64:
		return "float64"
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
		return Field{Type: FieldTypeInt, Value: value}
	case []int:
		return Field{Type: FieldTypeIntSlice, Value: value}
	case int32:
		return Field{Type: FieldTypeInt32, Value: value}
	case []int32:
		return Field{Type: FieldTypeInt32Slice, Value: value}
	case int64:
		return Field{Type: FieldTypeInt64, Value: value}
	case []int64:
		return Field{Type: FieldTypeInt64Slice, Value: value}
	case uint:
		return Field{Type: FieldTypeUint, Value: value}
	case []uint:
		return Field{Type: FieldTypeUintSlice, Value: value}
	case uint32:
		return Field{Type: FieldTypeUint32, Value: value}
	case []uint32:
		return Field{Type: FieldTypeUint32Slice, Value: value}
	case uint64:
		return Field{Type: FieldTypeUint64, Value: value}
	case []uint64:
		return Field{Type: FieldTypeUint64Slice, Value: value}
	case float32:
		return Field{Type: FieldTypeFloat32, Value: value}
	case []float32:
		return Field{Type: FieldTypeFloat32Slice, Value: value}
	case float64:
		return Field{Type: FieldTypeFloat64, Value: value}
	case []float64:
		return Field{Type: FieldTypeFloat64Slice, Value: value}
	case time.Duration:
		return Field{Type: FieldTypeDuration, Value: value}
	case time.Time:
		return Field{Type: FieldTypeTime, Value: value}
	case map[string]interface{}:
		subMap := make(map[string]Field)
		for key, value := range value {
			subMap[key] = Atof(value)
		}
		return Field{Type: FieldTypeSection, Value: subMap}
	default:
		return Field{Type: FiledTypeUnknown, Value: value}
	}
}

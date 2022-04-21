package configer

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

// ConfigField is a configuration field.
type ConfigField struct {
	Type  FieldType
	Value any
}

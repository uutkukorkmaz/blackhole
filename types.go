package blackhole

// ColumnType represents the type of a column in a database table.
type ColumnType string

// Constants defining various column types.
const (
	// Integer column types
	ColumnTypeInt       ColumnType = "integer"
	ColumnTypeBigInt    ColumnType = "bigint"
	ColumnTypeSmallInt  ColumnType = "smallint"
	ColumnTypeTinyInt   ColumnType = "tinyint"
	ColumnTypeMediumInt ColumnType = "mediumint"

	// Floating point column types
	ColumnTypeFloat   ColumnType = "float"
	ColumnTypeDouble  ColumnType = "double"
	ColumnTypeDecimal ColumnType = "decimal"

	// String column types
	ColumnTypeChar       ColumnType = "char"
	ColumnTypeVarchar    ColumnType = "varchar"
	ColumnTypeText       ColumnType = "text"
	ColumnTypeMediumText ColumnType = "mediumtext"
	ColumnTypeLongText   ColumnType = "longtext"

	// Date and time column types
	ColumnTypeDate      ColumnType = "date"
	ColumnTypeDateTime  ColumnType = "datetime"
	ColumnTypeTime      ColumnType = "time"
	ColumnTypeTimestamp ColumnType = "timestamp"

	// Binary column types
	ColumnTypeBinary    ColumnType = "binary"
	ColumnTypeVarBinary ColumnType = "varbinary"

	// Special column types
	ColumnTypeEnum ColumnType = "enum"
	ColumnTypeSet  ColumnType = "set"
	ColumnTypeJson ColumnType = "json"
)

// IsNumeric checks if the column type is a numeric type.
func (ct ColumnType) IsNumeric() bool {
	switch ct {
	case ColumnTypeInt, ColumnTypeBigInt, ColumnTypeSmallInt, ColumnTypeTinyInt, ColumnTypeMediumInt:
		return true
	}
	return false
}

// IsString checks if the column type is a string type.
func (ct ColumnType) IsString() bool {
	switch ct {
	case ColumnTypeChar, ColumnTypeVarchar, ColumnTypeText, ColumnTypeMediumText, ColumnTypeLongText:
		return true
	}
	return false
}

// IsTime checks if the column type is a time-related type.
func (ct ColumnType) IsTime() bool {
	switch ct {
	case ColumnTypeDate, ColumnTypeDateTime, ColumnTypeTime, ColumnTypeTimestamp:
		return true
	}
	return false
}

// IsFloat checks if the column type is a floating-point type.
func (ct ColumnType) IsFloat() bool {
	switch ct {
	case ColumnTypeFloat, ColumnTypeDouble, ColumnTypeDecimal:
		return true
	}
	return false
}

// IsBinary checks if the column type is a binary type.
func (ct ColumnType) IsBinary() bool {
	switch ct {
	case ColumnTypeBinary, ColumnTypeVarBinary:
		return true
	}
	return false
}

// IsEnum checks if the column type is an enum type.
func (ct ColumnType) IsEnum() bool {
	return ct == ColumnTypeEnum
}

// KindOf checks if the column type belongs to a specific primitive type category.
// The supported primitive categories are: "int", "string", "time", "float", "binary".
func (ct ColumnType) KindOf(primitive string) bool {
	switch primitive {
	case "int", "integer":
		return ct.IsNumeric()
	case "string":
		return ct.IsString() || ct.IsEnum()
	case "time":
		return ct.IsTime()
	case "float":
		return ct.IsFloat()
	case "binary":
		return ct.IsBinary()
	}
	return false
}

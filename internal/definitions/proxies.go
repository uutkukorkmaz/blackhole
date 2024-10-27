package definitions

func stringColumn(name string, length int) *Column {
	return NewColumn(name, ColumnTypeVarchar, length)
}

func integerColumn(name string) *Column {
	return NewColumn(name, ColumnTypeInt, 11)
}

func tinyIntColumn(name string) *Column {
	return NewColumn(name, ColumnTypeTinyInt, 4)
}

func smIntColumn(name string) *Column {
	return NewColumn(name, ColumnTypeSmallInt, 6)
}

func medIntColumn(name string) *Column {
	return NewColumn(name, ColumnTypeMediumInt, 9)
}

func bigIntColumn(name string) *Column {
	return NewColumn(name, ColumnTypeBigInt, 0)
}

func floatColumn(name string) *Column {
	return NewColumn(name, ColumnTypeFloat, 8)
}

func doubleColumn(name string) *Column {
	return NewColumn(name, ColumnTypeDouble, 16)
}

func decimalColumn(name string, precision, scale int) *Column {
	return NewColumn(name, ColumnTypeDecimal, 9).Precision(precision).Scale(scale)
}

func boolColumn(name string) *Column {
	return NewColumn(name, ColumnTypeTinyInt, 1)
}

func enumColumn(name string, values ...string) *Column {
	return NewColumn(name, ColumnTypeEnum, 0).Enum(values...)
}

func setColumn(name string, values ...string) *Column {
	return NewColumn(name, ColumnTypeSet, 0).Enum(values...)
}

func charColumn(name string, length int) *Column {
	return NewColumn(name, ColumnTypeChar, length)
}

func textColumn(name string) *Column {
	return NewColumn(name, ColumnTypeText, 0)
}

func medTextColumn(name string) *Column {
	return NewColumn(name, ColumnTypeMediumText, 0)
}

func longTextColumn(name string) *Column {
	return NewColumn(name, ColumnTypeLongText, 0)
}

func binColumn(name string) *Column {
	return NewColumn(name, ColumnTypeBinary, 0)
}

func varBinColumn(name string) *Column {
	return NewColumn(name, ColumnTypeVarBinary, 0)
}

func timestampColumn(name string) *Column {
	return NewColumn(name, ColumnTypeTimestamp, 0)
}

func dateColumn(name string) *Column {
	return NewColumn(name, ColumnTypeDate, 0)
}

func timeColumn(name string) *Column {
	return NewColumn(name, ColumnTypeTime, 0)
}

func dateTimeColumn(name string) *Column {
	return NewColumn(name, ColumnTypeDateTime, 0)
}

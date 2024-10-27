package blackhole

// Column represents a database column with various attributes such as type, length, precision, etc.
type Column struct {
	Definition
	name           string
	dataType       string
	length         int
	precision      int
	scale          int
	unsigned       bool
	primary        bool
	autoIncrements *AutoIncrements
	nullable       *Nullable
	defaultValue   *DefaultValue
	comment        *Comment
	enumValues     *EnumValues
	blueprint      *Blueprint
}

// NewColumn creates a new column instance with the specified name, data type, and length.
func NewColumn(name string, dataType ColumnType, length int) *Column {
	return &Column{
		name:     name,
		dataType: string(dataType),
		length:   length,
		primary:  false,
		unsigned: false,
	}
}

// AutoIncrement sets the column to auto increment.
func (c *Column) AutoIncrement() *Column {
	c.autoIncrements = &AutoIncrements{}
	return c
}

// Length sets the length of the column.
func (c *Column) Length(length int) *Column {
	c.length = length
	return c
}

// Default sets the default value for the column.
// If the default value is not "NULL", the column is set to NOT NULL.
func (c *Column) Default(defaultValue string) *Column {
	c.defaultValue = NewDefaultValue(defaultValue)
	if defaultValue != "NULL" {
		c.NotNull()
	} else {
		c.Nullable()
	}
	return c
}

// DefaultNull sets the default value of the column to NULL.
func (c *Column) DefaultNull() *Column {
	return c.Default("NULL")
}

// Nullable sets the column to allow NULL values.
func (c *Column) Nullable() *Column {
	return c.setNullable(true)
}

// NotNull sets the column to NOT allow NULL values.
func (c *Column) NotNull() *Column {
	return c.setNullable(false)
}

// setNullable sets the nullable state of the column.
func (c *Column) setNullable(state bool) *Column {
	if c.nullable == nil {
		c.nullable = &Nullable{}
	}
	c.nullable.is = state
	return c
}

// Precision sets the precision for a numeric column.
func (c *Column) Precision(precision int) *Column {
	c.precision = precision
	return c
}

// Scale sets the scale for a numeric column.
func (c *Column) Scale(scale int) *Column {
	c.scale = scale
	return c
}

// Enum sets the possible values for an ENUM column.
func (c *Column) Enum(values ...string) *Column {
	c.enumValues = &EnumValues{Values: values}
	return c
}

// Expression generates the column definition SQL using the provided grammar.
func (c *Column) Expression(grammar Grammar) (string, error) {
	return grammar.CompileColumn(c)
}

// GetName returns the name of the column.
func (c *Column) GetName() string {
	return c.name
}

// GetDataType returns the data type of the column.
func (c *Column) GetDataType() ColumnType {
	return ColumnType(c.dataType)
}

// GetLength returns the length of the column.
func (c *Column) GetLength() int {
	return c.length
}

// GetPrecision returns the precision of the column.
func (c *Column) GetPrecision() int {
	return c.precision
}

// GetScale returns the scale of the column.
func (c *Column) GetScale() int {
	return c.scale
}

// GetAutoIncrements returns the auto increment setting of the column.
func (c *Column) GetAutoIncrements() *AutoIncrements {
	return c.autoIncrements
}

// GetNullable returns the nullable setting of the column.
func (c *Column) GetNullable() *Nullable {
	return c.nullable
}

// GetDefaultValue returns the default value of the column.
func (c *Column) GetDefaultValue() *DefaultValue {
	return c.defaultValue
}

// GetComment returns the comment associated with the column.
func (c *Column) GetComment() *Comment {
	return c.comment
}

// GetEnumValues returns the possible ENUM values for the column.
func (c *Column) GetEnumValues() *EnumValues {
	return c.enumValues
}

// IsUnsigned returns whether the column is unsigned.
func (c *Column) IsUnsigned() bool {
	return c.unsigned
}

// Unsigned sets the column to be unsigned.
func (c *Column) Unsigned() *Column {
	c.unsigned = true
	return c
}

// Unique adds a unique index to the column.
func (c *Column) Unique() *Column {
	index := &Index{
		Type:  IndexTypeUnique,
		Table: c.blueprint.GetTable(),
		Columns: []string{
			c.name,
		},
		Algorithm: IndexAlgorithmDefault,
	}
	c.blueprint.AddChild(&Blueprint{
		table:   c.blueprint.GetTable(),
		grammar: c.blueprint.grammar,
		mode:    "alter",
		definitions: []Definition{
			index,
		},
	})

	return c
}

// Index adds a regular index to the column.
func (c *Column) Index() *Column {
	index := &Index{
		Type:  IndexTypeIndex,
		Table: c.blueprint.GetTable(),
		Columns: []string{
			c.name,
		},
		Algorithm: IndexAlgorithmDefault,
	}
	c.blueprint.AddChild(&Blueprint{
		table:   c.blueprint.GetTable(),
		grammar: c.blueprint.grammar,
		mode:    "alter",
		definitions: []Definition{
			index,
		},
	})

	return c
}

// ForeignKey creates a foreign key constraint for the column.
func (c *Column) ForeignKey() *ForeignKey {
	foreignKey := &ForeignKey{
		column: c.name,
	}
	c.blueprint.AddChild(&Blueprint{
		table:   c.blueprint.GetTable(),
		grammar: c.blueprint.grammar,
		mode:    "alter",
		definitions: []Definition{
			foreignKey,
		},
	})

	return foreignKey
}

// IndexUsing adds an index to the column using the specified algorithm.
func (c *Column) IndexUsing(algorithm IndexAlgorithm) *Column {
	index := &Index{
		Type:  IndexTypeIndex,
		Table: c.blueprint.GetTable(),
		Columns: []string{
			c.name,
		},
		Algorithm: algorithm,
	}
	c.blueprint.AddChild(&Blueprint{
		table:   c.blueprint.GetTable(),
		grammar: c.blueprint.grammar,
		mode:    "alter",
		definitions: []Definition{
			index,
		},
	})

	return c
}

// AddComment adds a comment to the column.
func (c *Column) AddComment(comment string) *Column {
	c.comment = NewComment(comment)
	return c
}

// Primary sets the column as a primary key.
func (c *Column) Primary() *Column {
	c.primary = true
	return c
}

// IsPrimary returns whether the column is a primary key.
func (c *Column) IsPrimary() bool {
	return c.primary
}

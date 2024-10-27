package definitions

// Blueprint represents a blueprint for defining database tables or modifying them.
type Blueprint struct {
	mode        string // todo: enum
	table       string
	charSet     string
	collate     string
	grammar     *Grammar
	definitions []Definition
	children    []*Blueprint
}

// NewBlueprint creates a new Blueprint instance with the specified table name.
func NewBlueprint(table string) *Blueprint {
	return &Blueprint{
		table: table,
	}
}

// Create is for creating a new table. It sets the mode to "create" and invokes the callback function.
func (b *Blueprint) Create(callback func(*Blueprint)) *Blueprint {
	b.setMode("create")
	callback(b)
	return b
}

// Alter is for altering an existing table. It sets the mode to "alter" and invokes the callback function.
func (b *Blueprint) Alter(callback func(*Blueprint)) *Blueprint {
	b.setMode("alter")
	callback(b)
	return b
}

// Drop is for dropping an existing table. It sets the mode to "drop" and invokes the callback function.
func (b *Blueprint) Drop(callback func(*Blueprint)) *Blueprint {
	b.setMode("drop")
	callback(b)
	return b
}

// GetTable returns the name of the table associated with the blueprint.
func (b *Blueprint) GetTable() string {
	return b.table
}

// Mode returns the mode (create, alter, drop) of the blueprint.
func (b *Blueprint) Mode() string {
	return b.mode
}

// setMode sets the mode (create, alter, drop) of the blueprint.
func (b *Blueprint) setMode(mode string) {
	b.mode = mode
}

// Grammar sets the grammar for the blueprint.
func (b *Blueprint) Grammar(grammar *Grammar) {
	b.grammar = grammar
}

// Definitions returns the column or index definitions associated with the blueprint.
func (b *Blueprint) Definitions() []Definition {
	return b.definitions
}

// addColumn adds a column definition to the blueprint.
func (b *Blueprint) addColumn(column *Column) {
	column.blueprint = b
	b.definitions = append(b.definitions, column)
}

// String creates a new string (varchar) column and adds it to the blueprint.
func (b *Blueprint) String(column string, len int) *Column {
	col := stringColumn(column, len)
	b.addColumn(col)
	return col
}

// Text creates a new text column and adds it to the blueprint.
func (b *Blueprint) Text(column string) *Column {
	col := textColumn(column)
	b.addColumn(col)
	return col
}

// Char creates a new char column and adds it to the blueprint.
func (b *Blueprint) Char(column string, len int) *Column {
	col := charColumn(column, len)
	b.addColumn(col)
	return col
}

// Binary creates a new binary column and adds it to the blueprint.
func (b *Blueprint) Binary(column string) *Column {
	col := binColumn(column)
	b.addColumn(col)
	return col
}

// Boolean creates a new boolean column and adds it to the blueprint.
func (b *Blueprint) Boolean(column string) *Column {
	col := boolColumn(column)
	b.addColumn(col)
	return col
}

// Date creates a new date column and adds it to the blueprint.
func (b *Blueprint) Date(column string) *Column {
	col := dateColumn(column)
	b.addColumn(col)
	return col
}

// DateTime creates a new datetime column and adds it to the blueprint.
func (b *Blueprint) DateTime(column string) *Column {
	col := dateTimeColumn(column)
	b.addColumn(col)
	return col
}

// Time creates a new time column and adds it to the blueprint.
func (b *Blueprint) Time(column string) *Column {
	col := timeColumn(column)
	b.addColumn(col)
	return col
}

// Int creates a new integer column and adds it to the blueprint.
func (b *Blueprint) Int(column string) *Column {
	col := integerColumn(column)
	b.addColumn(col)
	return col
}

// Id creates a new auto-incrementing primary key column named "id" and adds it to the blueprint.
func (b *Blueprint) Id() *Column {
	col := bigIntColumn("id").Unsigned().Primary().AutoIncrement().NotNull()
	b.addColumn(col)
	return col
}

// TinyInt creates a new tinyint column and adds it to the blueprint.
func (b *Blueprint) TinyInt(column string) *Column {
	col := tinyIntColumn(column)
	b.addColumn(col)
	return col
}

// SmallInt creates a new smallint column and adds it to the blueprint.
func (b *Blueprint) SmallInt(column string) *Column {
	col := smIntColumn(column)
	b.addColumn(col)
	return col
}

// MediumInt creates a new mediumint column and adds it to the blueprint.
func (b *Blueprint) MediumInt(column string) *Column {
	col := medIntColumn(column)
	b.addColumn(col)
	return col
}

// BigInt creates a new bigint column and adds it to the blueprint.
func (b *Blueprint) BigInt(column string) *Column {
	col := bigIntColumn(column)
	b.addColumn(col)
	return col
}

// Float creates a new float column and adds it to the blueprint.
func (b *Blueprint) Float(column string) *Column {
	col := floatColumn(column)
	b.addColumn(col)
	return col
}

// doubleColumn creates a new double column and adds it to the blueprint.
func (b *Blueprint) Double(column string) *Column {
	col := doubleColumn(column)
	b.addColumn(col)
	return col
}

// Enum creates a new enum column with the specified values and adds it to the blueprint.
func (b *Blueprint) Enum(name string, values []string) *Column {
	col := enumColumn(name, values...)
	b.addColumn(col)
	return col
}

// Timestamps adds created_at and updated_at timestamp columns to the blueprint.
func (b *Blueprint) Timestamps() {
	created := timestampColumn("created_at").
		Default("CURRENT_TIMESTAMP")
	updated := timestampColumn("updated_at").
		Default("CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP")
	b.addColumn(created)
	b.addColumn(updated)
}

// Build builds the SQL statement from the blueprint using the associated grammar.
func (b *Blueprint) Build() (string, error) {
	return (*b.grammar).Build(b)
}

// AddChild adds a child blueprint to the current blueprint.
func (b *Blueprint) AddChild(child *Blueprint) {
	b.children = append(b.children, child)
}

// Children returns the child blueprints of the current blueprint.
func (b *Blueprint) Children() []*Blueprint {
	return b.children
}

// AddIndex adds an index definition to the blueprint.
func (b *Blueprint) AddIndex(index *Index) {
	b.definitions = append(b.definitions, index)
}

// Collate sets the collation for the blueprint.
func (b *Blueprint) Collate(collate string) {
	b.collate = collate
}

// CharSet sets the character set for the blueprint.
func (b *Blueprint) CharSet(charSet string) {
	b.charSet = charSet
}

// GetCharSet returns the character set for the blueprint. If not set, it returns the default character set from the grammar.
func (b *Blueprint) GetCharSet() string {
	if b.charSet == "" {
		cs, _ := (*b.grammar).GetDefaultCharset()
		return cs
	}
	return b.charSet
}

// GetCollation returns the collation for the blueprint. If not set, it returns the default collation from the grammar.
func (b *Blueprint) GetCollation() string {
	if b.collate == "" {
		col, _ := (*b.grammar).GetDefaultCollation()
		return col
	}
	return b.collate
}

// SoftDeletes adds a nullable "deleted_at" timestamp column to the blueprint for soft deletion support.
func (b *Blueprint) SoftDeletes() {
	deleted := timestampColumn("deleted_at").
		Nullable()
	b.addColumn(deleted)
}

// ForeignId creates a new unsigned bigint foreign key column and adds it to the blueprint.
func (b *Blueprint) ForeignId(column string) (*ForeignKey, *Column) {
	col := bigIntColumn(column).Unsigned()
	b.addColumn(col)

	fk := NewForeignKey(column, b.GetTable())

	b.AddChild(&Blueprint{
		table:   b.GetTable(),
		grammar: b.grammar,
		mode:    "alter",
		definitions: []Definition{
			fk,
		},
	})

	return fk, col
}

// RenameColumn adds a column rename definition to the blueprint or creates a child blueprint if necessary.
func (b *Blueprint) RenameColumn(old, new string) {
	rename := NewRenameColumn(old, new)
	if b.mode == "alter" {
		b.definitions = append(b.definitions, rename)
		return
	}

	child := &Blueprint{
		table:   b.GetTable(),
		grammar: b.grammar,
		mode:    "alter",
		definitions: []Definition{
			rename,
		},
	}
	b.AddChild(child)
}

// DropColumn adds a drop column definition as a child blueprint.
func (b *Blueprint) DropColumn(table, column string) {
	drop := NewDropColumn(column)

	child := &Blueprint{
		table:   table,
		grammar: b.grammar,
		mode:    "alter",
		definitions: []Definition{
			drop,
		},
	}

	b.AddChild(child)
}

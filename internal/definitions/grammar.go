package definitions

import (
	"fmt"
)

type Grammar interface {
	GetDefaultCollation() (string, error)
	GetDefaultCharset() (string, error)
	CompileCreateDatabase(database string) (string, error)
	CompileDropDatabase(database string) (string, error)
	CompileCreateTable(table Blueprint) (string, error)
	CompileAlterTable(table Blueprint) (string, error)
	CompileDropTable(table Blueprint) (string, error)
	GetDateFormat() string
	CompileColumn(c *Column) (string, error)
	CompileAutoIncrement(a *AutoIncrements) (string, error)
	CompileDefaultValue(d *DefaultValue) (string, error)
	CompileComment(c *Comment) (string, error)
	CompileNullable(n *Nullable) (string, error)
	CompileEnumValues(e *EnumValues) (string, error)
	CompileIndex(i *Index) (string, error)
	Build(b *Blueprint) (string, error)
	CompileForeignKey(f *ForeignKey) (string, error)
	CompileRenameColumn(r *RenameColumn) (string, error)
	CompileDropColumn(column string) (string, error)
}

type BaseGrammar struct{}

// GetDefaultCollation provides a default collation, which can be overridden.
func (bg *BaseGrammar) GetDefaultCollation() (string, error) {
	return "utf8mb4_unicode_ci", nil
}

// GetDefaultCharset provides a default charset, which can be overridden.
func (bg *BaseGrammar) GetDefaultCharset() (string, error) {
	return "utf8mb4", nil
}

// CompileCreateDatabase returns a basic SQL statement for creating a database.
func (bg *BaseGrammar) CompileCreateDatabase(database string) (string, error) {
	return "", fmt.Errorf("blackhole: CompileCreateDatabase not implemented")
}

// CompileDropDatabase returns a basic SQL statement for dropping a database.
func (bg *BaseGrammar) CompileDropDatabase(database string) (string, error) {
	return "", fmt.Errorf("blackhole: CompileDropDatabase not implemented")
}

// CompileCreateTable is a placeholder, expecting the table creation logic to be implemented by specific grammars.
func (bg *BaseGrammar) CompileCreateTable(table Blueprint) (string, error) {
	return "", fmt.Errorf("blackhole: CompileCreateTable not implemented")
}

// CompileAlterTable is a placeholder, expecting the table alteration logic to be implemented by specific grammars.
func (bg *BaseGrammar) CompileAlterTable(table Blueprint) (string, error) {
	return "", fmt.Errorf("blackhole: CompileAlterTable not implemented")
}

// CompileDropTable returns a basic SQL statement for dropping a table.
func (bg *BaseGrammar) CompileDropTable(table Blueprint) (string, error) {
	return "", fmt.Errorf("blackhole: CompileDropTable not implemented")
}

// DefineColumn is a placeholder, expecting column definitions to be handled by specific grammars.
func (bg *BaseGrammar) CompileColumn(c *Column) (string, error) {
	return "", fmt.Errorf("blackhole: CompileColumn not implemented")
}

// CompileAutoIncrement is a placeholder, expecting auto-increment logic to be implemented by specific grammars.
func (bg *BaseGrammar) CompileAutoIncrement(a *AutoIncrements) (string, error) {
	return "", fmt.Errorf("blackhole: CompileAutoIncrement not implemented")
}

// CompileDefaultValue is a placeholder for default value handling.
func (bg *BaseGrammar) CompileDefaultValue(d *DefaultValue) (string, error) {
	return "", fmt.Errorf("blackhole: CompileDefaultValue not implemented")
}

// CompileComment is a placeholder for comments on columns or tables.
func (bg *BaseGrammar) CompileComment(c *Comment) (string, error) {
	return "", fmt.Errorf("blackhole: CompileComment not implemented")
}

// CompileNullable handles nullable constraints for a column.
func (bg *BaseGrammar) CompileNullable(n *Nullable) (string, error) {
	return "", fmt.Errorf("blackhole: CompileNullable not implemented")
}

// CompileEnumValues compiles the list of enum values.
func (bg *BaseGrammar) CompileEnumValues(e *EnumValues) (string, error) {
	return "", fmt.Errorf("blackhole: CompileEnumValues not implemented")
}

// GetDateFormat provides a default date format for the grammar.
func (bg *BaseGrammar) GetDateFormat() string {
	return "2006-01-02 15:04:05"
}

// Build is a placeholder for building a blueprint.
func (bg *BaseGrammar) Build(b *Blueprint) (string, error) {
	return "", fmt.Errorf("blackhole: Build not implemented")
}

// CompileIndex is a placeholder for index handling.
func (bg *BaseGrammar) CompileIndex(i *Index) (string, error) {
	return "", fmt.Errorf("blackhole: CompileIndex not implemented")
}

// CompileForeignId is a placeholder for foreign key handling.
func (bg *BaseGrammar) CompileForeignKey(f *ForeignKey) (string, error) {
	return "", fmt.Errorf("blackhole: CompileForeignKey not implemented")
}

// CompileRenameColumn is a placeholder for renaming columns.
func (bg *BaseGrammar) CompileRenameColumn(r *RenameColumn) (string, error) {
	return "", fmt.Errorf("blackhole: CompileRenameColumn not implemented")
}

// CompileDropColumn is a placeholder for dropping columns.
func (bg *BaseGrammar) CompileDropColumn(column string) (string, error) {
	return "", fmt.Errorf("blackhole: CompileDropColumn not implemented")
}

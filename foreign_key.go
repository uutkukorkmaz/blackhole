package blackhole

import (
	"github.com/gertd/go-pluralize"
	"strings"
)

type ForeignKeyAction string

const (
	NoAction   ForeignKeyAction = "no action"
	Cascade    ForeignKeyAction = "cascade"
	SetNull    ForeignKeyAction = "set null"
	SetDefault ForeignKeyAction = "set default"
	Restrict   ForeignKeyAction = "restrict"
)

// ForeignKey represents a foreign key definition in a database.
type ForeignKey struct {
	Definition
	table            string
	column           string
	autoDiscover     bool
	onDelete         *ForeignKeyAction
	onUpdate         *ForeignKeyAction
	referencedTable  string
	referencedColumn string
}

// NewForeignKey creates a new ForeignKey instance with the specified column and table.
func NewForeignKey(column, table string) *ForeignKey {
	return &ForeignKey{
		table:        table,
		column:       column,
		autoDiscover: true,
	}
}

// GetTable returns the table name of the foreign key.
func (f *ForeignKey) GetTable() string {
	return f.table
}

// GetColumn returns the column name of the foreign key.
func (f *ForeignKey) GetColumn() string {
	return f.column
}

// Expression generates the SQL expression for the foreign key using the provided grammar.
func (f *ForeignKey) Expression(grammar Grammar) (string, error) {
	f.discoverReferences()
	return grammar.CompileForeignKey(f)
}

// discoverReferences automatically discovers the referenced table and column based on the column name if they are not explicitly set.
func (f *ForeignKey) discoverReferences() {
	if f.referencedTable != "" && f.referencedColumn != "" {
		return
	}

	if !f.autoDiscover {
		return
	}

	// Use a pluralizer to infer the referenced table name based on the column name.
	p := pluralize.NewClient()
	parts := strings.Split(f.column, "_")
	columnPart := strings.Join(parts[1:], "_")
	tablePart := parts[0]
	f.referencedTable = p.Plural(tablePart)
	f.referencedColumn = columnPart

	return
}

// On sets the referenced table and column for the foreign key.
func (f *ForeignKey) On(table, column string) *ForeignKey {
	f.referencedTable = table
	f.referencedColumn = column
	return f
}

// ReferencedTable returns the name of the referenced table.
func (f *ForeignKey) ReferencedTable() string {
	return f.referencedTable
}

// ReferencedColumn returns the name of the referenced column.
func (f *ForeignKey) ReferencedColumn() string {
	return f.referencedColumn
}

// Column returns the column name of the foreign key.
func (f *ForeignKey) Column() string {
	return f.column
}

// OnDelete sets the action to be taken when the referenced row is deleted.
func (f *ForeignKey) OnDelete(a ForeignKeyAction) *ForeignKey {
	f.onDelete = &a
	return f
}

// OnUpdate sets the action to be taken when the referenced row is updated.
func (f *ForeignKey) OnUpdate(a ForeignKeyAction) *ForeignKey {
	f.onUpdate = &a
	return f
}

// GetOnDeleteAction returns the action to be taken when the referenced row is deleted.
func (f *ForeignKey) GetOnDeleteAction() *ForeignKeyAction {
	return f.onDelete
}

// GetOnUpdateAction returns the action to be taken when the referenced row is updated.
func (f *ForeignKey) GetOnUpdateAction() *ForeignKeyAction {
	return f.onUpdate
}

// CascadeOnDelete sets the action to cascade when the referenced row is deleted.
func (f *ForeignKey) CascadeOnDelete() *ForeignKey {
	return f.OnDelete(Cascade)
}

// CascadeOnUpdate sets the action to cascade when the referenced row is updated.
func (f *ForeignKey) CascadeOnUpdate() *ForeignKey {
	return f.OnUpdate(Cascade)
}

// SetNullOnDelete sets the action to set the column to null when the referenced row is deleted.
func (f *ForeignKey) SetNullOnDelete() *ForeignKey {
	return f.OnDelete(SetNull)
}

// SetNullOnUpdate sets the action to set the column to null when the referenced row is updated.
func (f *ForeignKey) SetNullOnUpdate() *ForeignKey {
	return f.OnUpdate(SetNull)
}

// RestrictOnDelete sets the action to restrict deletion when the referenced row is deleted.
func (f *ForeignKey) RestrictOnDelete() *ForeignKey {
	return f.OnDelete(Restrict)
}

// RestrictOnUpdate sets the action to restrict update when the referenced row is updated.
func (f *ForeignKey) RestrictOnUpdate() *ForeignKey {
	return f.OnUpdate(Restrict)
}

// NoActionOnDelete sets the action to take no action when the referenced row is deleted.
func (f *ForeignKey) NoActionOnDelete() *ForeignKey {
	return f.OnDelete(NoAction)
}

// NoActionOnUpdate sets the action to take no action when the referenced row is updated.
func (f *ForeignKey) NoActionOnUpdate() *ForeignKey {
	return f.OnUpdate(NoAction)
}

// SetDefaultOnDelete sets the action to set the column to its default value when the referenced row is deleted.
func (f *ForeignKey) SetDefaultOnDelete() *ForeignKey {
	return f.OnDelete(SetDefault)
}

// SetDefaultOnUpdate sets the action to set the column to its default value when the referenced row is updated.
func (f *ForeignKey) SetDefaultOnUpdate() *ForeignKey {
	return f.OnUpdate(SetDefault)
}

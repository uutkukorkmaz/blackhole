package grammars

import (
	def "blackhole/internal/definitions"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type MySqlGrammar struct {
	def.BaseGrammar
}

func NewMySqlGrammar() *MySqlGrammar {
	return &MySqlGrammar{}
}

// GetDefaultCollation provides a default collation for the grammar.
func (m *MySqlGrammar) GetDefaultCollation() (string, error) {
	return "utf8mb4_unicode_ci", nil
}

// GetDateFormat provides a default date format for the grammar.
func (m *MySqlGrammar) GetDateFormat() string {
	return "2006-01-02 15:04:05"
}

// CompileAutoIncrement returns the auto-increment SQL for MySQL.
func (m *MySqlGrammar) CompileAutoIncrement(a *def.AutoIncrements) (string, error) {
	return "auto_increment", nil
}

// CompileNullable returns the nullable SQL for MySQL.
func (m *MySqlGrammar) CompileNullable(n *def.Nullable) (string, error) {
	if n.Is() {
		return "null", nil
	}
	return "not null", nil
}

// CompileDefaultValue returns the default value SQL for MySQL.
func (m *MySqlGrammar) CompileDefaultValue(d *def.DefaultValue) (string, error) {
	return d.Get(), nil
}

// CompileComment returns the comment SQL for MySQL.
func (m *MySqlGrammar) CompileComment(c *def.Comment) (string, error) {
	return fmt.Sprintf("'%s'", c.Get()), nil
}

// CompileColumn returns the column SQL for MySQL.
// It compiles various attributes of the column, such as name, data type, length, nullability, etc.
func (m *MySqlGrammar) CompileColumn(c *def.Column) (string, error) {
	var result string
	dataTypeString := string(c.GetDataType())

	// Handle enum data type if applicable
	if c.GetEnumValues() != nil {
		dts, err := (*c.GetEnumValues()).Expression(m)
		if err != nil {
			return "", err
		}
		dataTypeString = dts
	}

	// Compile column name and data type
	result = fmt.Sprintf("`%s` %s", c.GetName(), dataTypeString)

	// Add length if specified
	if c.GetLength() > 0 {
		result += "(" + strconv.Itoa(c.GetLength()) + ")"
	}

	// Handle unsigned attribute
	if c.IsUnsigned() {
		result += " unsigned"
	}

	// Handle nullable attribute
	if c.GetNullable() != nil {
		nullable, err := c.GetNullable().Expression(m)
		if err != nil {
			return "", err
		}
		result += " " + nullable
	}

	// Handle auto-increment attribute
	if c.GetAutoIncrements() != nil {
		ai, err := c.GetAutoIncrements().Expression(m)
		if err != nil {
			return "", err
		}
		result += " " + ai
	}

	// Handle primary key attribute
	if c.IsPrimary() {
		result += " primary key"
	}

	// Handle default value attribute
	if c.GetDefaultValue() != nil {
		defaultValue, err := c.GetDefaultValue().Expression(m)
		if err != nil {
			return "", err
		}
		if defaultValue != "" && c.GetDataType().KindOf("string") {
			defaultValue = "'" + defaultValue + "'"
		}
		result += " default " + defaultValue
	}

	// Handle comment attribute
	if c.GetComment() != nil {
		comment, err := c.GetComment().Expression(m)
		if err != nil {
			return "", err
		}
		result += " comment " + comment
	}

	return result, nil
}

// Build returns the final runnable SQL for MySQL.
// It constructs the SQL statement based on the blueprint mode (create, drop, alter).
func (m *MySqlGrammar) Build(b *def.Blueprint) (string, error) {
	var sql string
	modeDirective := m.getDirective(b.Mode())
	sql = fmt.Sprintf("%s `%s`", modeDirective, b.GetTable())
	blueprintCompile, err := m.callCompileFunctionsByMode(b)
	if err != nil {
		return "", err
	}
	sql += blueprintCompile + ";"
	return sql, nil
}

// getDirective returns the appropriate SQL directive based on the blueprint mode.
func (m *MySqlGrammar) getDirective(mode string) string {
	switch mode {
	case "create":
		return "create table if not exists"
	case "drop":
		return "drop table if exists"
	case "alter":
		return "alter table"
	}
	panic("blackhole: MySQL build: invalid blueprint mode given : " + mode)
}

// callCompileFunctionsByMode compiles the SQL statement based on the blueprint mode.
func (m *MySqlGrammar) callCompileFunctionsByMode(b *def.Blueprint) (string, error) {
	var sql string
	switch b.Mode() {
	case "create":
		create, err := m.CompileCreateTable(*b)
		if err != nil {
			return "", err
		}
		sql += create
	case "drop":
		drop, err := m.CompileDropTable(*b)
		if err != nil {
			return "", err
		}
		sql += drop
	case "alter":
		alter, err := m.CompileAlterTable(*b)
		if err != nil {
			return "", err
		}
		sql += alter
	}
	return sql, nil
}

// CompileCreateTable returns the SQL for creating a table in MySQL.
// It iterates over the definitions in the blueprint to build the table schema.
func (m *MySqlGrammar) CompileCreateTable(b def.Blueprint) (string, error) {
	var sql string
	sql = "("
	for _, c := range b.Definitions() {
		expression, err := c.Expression(m)
		if err != nil {
			return "", err
		}
		sql += expression + ","
	}

	sql = strings.TrimRight(sql, ",")
	sql += fmt.Sprintf(") default character set %s collate '%s'", b.GetCharSet(), b.GetCollation())

	sql += ";\n"

	// Compile child blueprints if any (e.g., foreign key constraints)
	for _, cb := range b.Children() {
		child, err := m.Build(cb)
		if err != nil {
			return "", err
		}
		sql += child + "\n"
	}

	return strings.TrimRight(sql, ";\n"), nil
}

// CompileDropTable returns the SQL for dropping a table in MySQL.
func (m *MySqlGrammar) CompileDropTable(table def.Blueprint) (string, error) {
	return "", nil
}

// CompileAlterTable returns the SQL for altering a table in MySQL.
// It handles adding new columns or modifying existing columns.
func (m *MySqlGrammar) CompileAlterTable(table def.Blueprint) (string, error) {
	var sql string
	alterDirective := m.getDirective("alter")
	for i, c := range table.Definitions() {
		if i > 0 {
			sql += alterDirective + " `" + table.GetTable() + "`"
		}
		expression, err := c.Expression(m)
		if err != nil {
			return "", err
		}

		// Handle column addition specifically
		if reflect.TypeOf(c) == reflect.TypeOf(&def.Column{}) {
			sql += " add " + expression + ";\n"
			continue
		}

		sql += expression + ";\n"

		// Compile child blueprints if any (e.g., foreign key constraints)
		for _, cb := range table.Children() {
			child, err := m.Build(cb)
			if err != nil {
				return "", err
			}
			sql += child + "\n"
		}
	}

	sql = strings.ReplaceAll(sql, ";;", ";")

	return strings.TrimRight(sql, ";\n"), nil
}

// CompileEnumValues returns the SQL for enum values in MySQL.
// It constructs the enum values as a comma-separated list.
func (m *MySqlGrammar) CompileEnumValues(e *def.EnumValues) (string, error) {
	var sql string
	sql = "enum("
	for _, v := range e.Values {
		sql += "'" + v + "',"
	}
	sql = strings.TrimRight(sql, ",")
	sql += ")"
	return sql, nil
}

// CompileCreateDatabase returns the SQL for creating a database in MySQL.
// This function is not implemented yet.
func (m *MySqlGrammar) CompileCreateDatabase(database string) (string, error) {
	return "", fmt.Errorf("blackhole: MySQL grammar: CompileCreateDatabase not implemented")
}

// CompileDropDatabase returns the SQL for dropping a database in MySQL.
// This function is not implemented yet.
func (m *MySqlGrammar) CompileDropDatabase(database string) (string, error) {
	return "", fmt.Errorf("blackhole: MySQL grammar: CompileDropDatabase not implemented")
}

// CompileIndex returns the SQL for creating an index in MySQL.
// It constructs the index name and optionally specifies the algorithm to use.
func (m *MySqlGrammar) CompileIndex(i *def.Index) (string, error) {
	var sql string
	columnsForName := strings.ReplaceAll(strings.Replace(i.ColumnsString(), ",", "_", -1), "`", "")
	indexName := strings.TrimRight(fmt.Sprintf("%s_%s_%s", i.Table, columnsForName, i.Type), "_")
	using := ""
	if i.Algorithm != def.IndexAlgorithmDefault {
		using = fmt.Sprintf(" using %s", i.Algorithm)
	}
	sql = fmt.Sprintf(" add %s `%s`(%s)%s;", i.Type, indexName, i.ColumnsString(), using)
	return sql, nil
}

// CompileForeignKey returns the SQL for creating a foreign key in MySQL.
// It ensures that both referenced column and table are specified.
func (m *MySqlGrammar) CompileForeignKey(f *def.ForeignKey) (string, error) {
	var sql string
	name := fmt.Sprintf("%s_%s_foreign", f.GetTable(), f.GetColumn())
	if f.ReferencedColumn() == "" || f.ReferencedTable() == "" {
		return "", fmt.Errorf("blackhole: MySQL grammar: CompileForeignKey: referenced column and table are required")
	}
	sql = fmt.Sprintf(" add constraint `%s` foreign key (`%s`) references `%s` (`%s`)", name, f.GetColumn(), f.ReferencedTable(), f.ReferencedColumn())
	if f.GetOnDeleteAction() != nil {
		sql += fmt.Sprintf(" on delete %s", *f.GetOnDeleteAction())
	}
	if f.GetOnUpdateAction() != nil {
		sql += fmt.Sprintf(" on update %s", *f.GetOnUpdateAction())
	}
	sql += ";"
	return sql, nil
}

// CompileRenameColumn returns the SQL for renaming a column in MySQL.
// It generates the SQL statement to rename a column from one name to another.
func (m *MySqlGrammar) CompileRenameColumn(r *def.RenameColumn) (string, error) {
	return fmt.Sprintf(" rename column `%s` to `%s`;", r.From(), r.To()), nil
}

// CompileDropColumn returns the SQL for dropping a column in MySQL.
// It generates the SQL statement to drop the specified column.
func (m *MySqlGrammar) CompileDropColumn(column string) (string, error) {
	return fmt.Sprintf(" drop column `%s`;", column), nil
}

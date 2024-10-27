package blackhole

type DropColumn struct {
	Definition
	column string
}

func NewDropColumn(column string) *DropColumn {
	return &DropColumn{
		column: column,
	}
}

func (d *DropColumn) Expression(grammar Grammar) (string, error) {
	return grammar.CompileDropColumn(d.column)
}

func (d *DropColumn) Column() string {
	return d.column
}

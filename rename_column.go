package blackhole

type RenameColumn struct {
	Definition
	from string
	to   string
}

func NewRenameColumn(from, to string) *RenameColumn {
	return &RenameColumn{
		from: from,
		to:   to,
	}
}

func (r *RenameColumn) From() string {
	return r.from
}

func (r *RenameColumn) To() string {
	return r.to
}

func (r *RenameColumn) Expression(grammar Grammar) (string, error) {
	return grammar.CompileRenameColumn(r)
}

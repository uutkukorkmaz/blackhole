package blackhole

type Nullable struct {
	Definition
	is bool
}

func (n *Nullable) Expression(grammar Grammar) (string, error) {
	return grammar.CompileNullable(n)
}

func (n *Nullable) Is() bool {
	return n.is
}

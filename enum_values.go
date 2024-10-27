package blackhole

type EnumValues struct {
	Definition
	Values []string
}

func (e *EnumValues) Expression(grammar Grammar) (string, error) {
	return grammar.CompileEnumValues(e)
}

package definitions

type AutoIncrements struct{ Definition }

func (a *AutoIncrements) Expression(grammar Grammar) (string, error) {
	return grammar.CompileAutoIncrement(a)
}

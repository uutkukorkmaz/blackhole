package definitions

type Definition interface {
	Expression(grammar Grammar) (string, error)
}

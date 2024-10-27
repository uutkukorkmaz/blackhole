package blackhole

type Definition interface {
	Expression(grammar Grammar) (string, error)
}

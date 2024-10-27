package definitions

type DefaultValue struct {
	Definition
	value string
}

func NewDefaultValue(value string) *DefaultValue {
	return &DefaultValue{value: value}
}

func (d *DefaultValue) Get() string {
	return d.value
}

func (d *DefaultValue) Set(value string) {
	d.value = value
}

func (d *DefaultValue) Expression(grammar Grammar) (string, error) {
	return grammar.CompileDefaultValue(d)
}

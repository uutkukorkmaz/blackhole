package blackhole

type Comment struct {
	Definition
	comment string
}

func NewComment(comment string) *Comment {
	return &Comment{comment: comment}
}

func (c *Comment) Get() string {
	return c.comment
}

func (c *Comment) Set(comment string) {
	c.comment = comment
}

func (c *Comment) Expression(grammar Grammar) (string, error) {
	return grammar.CompileComment(c)
}

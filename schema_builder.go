package blackhole

import (
	"strings"
)

// Schema is a schema builder instance.
type Schema struct {
	grammar    Grammar
	blueprints []*Blueprint
}

// NewSchema creates a new schema instance.
func NewSchema(grammar Grammar) *Schema {
	return &Schema{
		grammar:    grammar,
		blueprints: []*Blueprint{},
	}
}

// Create a new table on the schema.
func (s *Schema) Create(name string, callback func(*Blueprint)) *Schema {
	bp := s.addNewBlueprint(name)
	bp.Create(callback)
	return s
}

// Alter an existing table on the schema.
func (s *Schema) Alter(name string, callback func(*Blueprint)) *Schema {
	bp := s.addNewBlueprint(name)
	bp.Alter(callback)
	return s
}

func (s *Schema) addNewBlueprint(table string) *Blueprint {
	bp := NewBlueprint(table)
	bp.Grammar(&s.grammar)
	s.addBlueprint(bp)

	return bp
}

func (s *Schema) addBlueprint(bp *Blueprint) {
	s.blueprints = append(s.blueprints, bp)
}

// Build the schema into a SQL string.
func (s *Schema) Build() (string, error) {
	var result string
	for _, bp := range s.blueprints {
		sql, err := bp.Build()
		if err != nil {
			return "", err
		}
		result += sql + "\n"
	}
	s.blueprints = []*Blueprint{}
	return strings.TrimRight(result, "\n"), nil
}

// MySQL is a MySQL grammar instance.
var MySQL = NewMySqlGrammar()

// TODO: Add more grammars here.

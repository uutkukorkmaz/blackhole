package blackhole

import (
	"blackhole/internal/definitions"
	"blackhole/internal/definitions/grammars"
	"strings"
)

// Schema is a schema builder instance.
type Schema struct {
	grammar    definitions.Grammar
	blueprints []*definitions.Blueprint
}

// NewSchema creates a new schema instance.
func NewSchema(grammar definitions.Grammar) *Schema {
	return &Schema{
		grammar:    grammar,
		blueprints: []*definitions.Blueprint{},
	}
}

// Create a new table on the schema.
func (s *Schema) Create(name string, callback func(*definitions.Blueprint)) *Schema {
	bp := s.addNewBlueprint(name)
	bp.Create(callback)
	return s
}

// Alter an existing table on the schema.
func (s *Schema) Alter(name string, callback func(*definitions.Blueprint)) *Schema {
	bp := s.addNewBlueprint(name)
	bp.Alter(callback)
	return s
}

func (s *Schema) addNewBlueprint(table string) *definitions.Blueprint {
	bp := definitions.NewBlueprint(table)
	bp.Grammar(&s.grammar)
	s.addBlueprint(bp)

	return bp
}

func (s *Schema) addBlueprint(bp *definitions.Blueprint) {
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
	s.blueprints = []*definitions.Blueprint{}
	return strings.TrimRight(result, "\n"), nil
}

// MySQL is a MySQL grammar instance.
var MySQL = grammars.NewMySqlGrammar()

// TODO: Add more grammars here.

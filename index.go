package blackhole

import "strings"

type IndexType string
type IndexAlgorithm string

const (
	IndexTypePrimary  IndexType = "primary"
	IndexTypeUnique   IndexType = "unique"
	IndexTypeIndex    IndexType = "index"
	IndexTypeSpatial  IndexType = "spatial"
	IndexTypeFullText IndexType = "fulltext"
)

const (
	IndexAlgorithmDefault IndexAlgorithm = ""
	IndexAlgorithmBTree   IndexAlgorithm = "btree"
	IndexAlgorithmHash    IndexAlgorithm = "hash"
)

type Index struct {
	Definition
	Table     string
	Type      IndexType
	Columns   []string
	Algorithm IndexAlgorithm
}

func (i *Index) ColumnsString() string {
	s := make([]string, len(i.Columns))
	for k, v := range i.Columns {
		s[k] = "`" + v + "`"
	}
	return strings.Join(s, ", ")
}

func (i *Index) Using(algorithm IndexAlgorithm) *Index {
	i.Algorithm = algorithm
	return i
}

func (i *Index) Expression(grammar Grammar) (string, error) {
	return grammar.CompileIndex(i)
}

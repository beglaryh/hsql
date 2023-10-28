package hsql

type Predicate int

const (
	AND Predicate = iota + 1
	OR
)

func (predicate Predicate) string() string {
	if predicate == AND {
		return "AND"
	}
	return "OR"
}

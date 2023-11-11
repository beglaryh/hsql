package hsql

type Direction int

const (
	ASC = iota + 1
	DESC
)

func (d Direction) String() string {
	return [...]string{"ASC", "DESC"}[d-1]
}

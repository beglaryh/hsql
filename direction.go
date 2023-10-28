package hsql

type direction int

const (
	ASC = iota + 1
	DESC
)

func (d direction) string() string {
	return [...]string{"ASC", "DESC"}[d-1]
}

package hsql

type Operator int

const (
	Eq Operator = iota + 1
	Like
	ColumnIn
)

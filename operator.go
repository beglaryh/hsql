package hsql

type Operator int

const (
	eq Operator = iota + 1
	like
	columnIn
	valueIn
)

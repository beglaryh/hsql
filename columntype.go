package hsql

type ColumnType int

const (
	String ColumnType = iota + 1
	Date
	TimeStamp
	TimeStampZ
	JsonB
	Boolean
	Integer
	BigInt
	Flot
	BigFloat
	UUID
)

package hsql

type ColumnType int

const (
	String ColumnType = iota + 1
	Date
	TimeStamp
	TimeStampTZ
	JsonB
	JsonArray
	Boolean
	Integer
	BigInt
	Flot
	BigFloat
	UUID
)

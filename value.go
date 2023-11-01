package hsql

type value struct {
	val    string
	vals   []string
	column TableColumn
}

func newValue(v string) value {
	return value{val: v}
}

func newValues(vs []string) value {
	return value{vals: vs}
}

func newValueFromColumn(c TableColumn) value {
	return value{column: c}
}

func (v value) isColumn() bool {
	return len(v.column.name) != 0
}

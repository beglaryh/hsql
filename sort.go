package hsql

type Sort struct {
	column    TableColumn
	direction direction
}

func Asc(column TableColumn) Sort {
	return NewSort(column, ASC)
}

func Desc(column TableColumn) Sort {
	return NewSort(column, DESC)
}

func NewSort(column TableColumn, direction direction) Sort {
	return Sort{column, direction}
}

func (sort Sort) getColumn() TableColumn {
	return sort.column
}

func (sort Sort) getDirection() direction {
	return sort.direction
}

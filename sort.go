package hsql

type Sort struct {
	column    TableColumn
	direction Direction
}

func Asc(column TableColumn) Sort {
	return NewSort(column, ASC)
}

func Desc(column TableColumn) Sort {
	return NewSort(column, DESC)
}

func NewSort(column TableColumn, direction Direction) Sort {
	return Sort{column, direction}
}

func (sort Sort) GetColumn() TableColumn {
	return sort.column
}

func (sort Sort) GetDirection() Direction {
	return sort.direction
}

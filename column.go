package hsql

type TableColumn struct {
	tableName  string
	columnName string
	columnType ColumnType
	foreignKey *TableColumn
}

func NewTableColumn(table string, column string, columnType ColumnType, foreignKey *TableColumn) TableColumn {
	return TableColumn{
		tableName:  table,
		columnName: column,
		columnType: columnType,
		foreignKey: foreignKey,
	}
}

func (column TableColumn) GetTable() string {
	return column.tableName
}

func (column TableColumn) GetName() string {
	return column.columnName
}

func (column TableColumn) GetType() ColumnType {
	return column.columnType
}

func (column TableColumn) GetForeignKey() TableColumn {
	return *column.foreignKey
}

func (column TableColumn) AsTableColumn() string {
	return column.tableName + "." + column.columnName
}

func (column TableColumn) HasForeignKey() bool {
	return column.foreignKey != nil
}

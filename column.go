package hsql

type TableColumn struct {
	table      string
	name       string
	columnType ColumnType
	foreignKey *TableColumn
	mutable    bool
}

func NewTableColumn(table string, name string, columnType ColumnType) TableColumn {
	return TableColumn{
		table:      table,
		name:       name,
		columnType: columnType,
		mutable:    true}
}

func (column TableColumn) GetTable() string {
	return column.table
}

func (column TableColumn) GetName() string {
	return column.name
}

func (column TableColumn) GetType() ColumnType {
	return column.columnType
}

func (column TableColumn) GetForeignKey() TableColumn {
	return *column.foreignKey
}

func (column TableColumn) AsTableColumn() string {
	return column.table + "." + column.name
}

func (column TableColumn) HasForeignKey() bool {
	return column.foreignKey != nil
}

type TableColumnBuilder struct {
	column *TableColumn
}

func NewTableColumnBuilder(table string, name string, columnType ColumnType) *TableColumnBuilder {
	column := NewTableColumn(table, name, columnType)
	return &TableColumnBuilder{column: &column}
}

func (builder *TableColumnBuilder) WithTable(table string) *TableColumnBuilder {
	builder.column.table = table
	return builder
}

func (builder *TableColumnBuilder) WithName(name string) *TableColumnBuilder {
	builder.column.name = name
	return builder
}

func (builder *TableColumnBuilder) WithType(columnType ColumnType) *TableColumnBuilder {
	builder.column.columnType = columnType
	return builder
}

func (builder *TableColumnBuilder) WithForeignKey(column TableColumn) *TableColumnBuilder {
	builder.column.foreignKey = &column
	return builder
}

func (builder *TableColumnBuilder) IsMutable(mutable bool) *TableColumnBuilder {
	builder.column.mutable = mutable
	return builder
}

func (builder *TableColumnBuilder) Build() TableColumn {
	return *builder.column
}

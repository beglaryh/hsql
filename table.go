package hsql

type Table interface {
	GetName() string
	GetColumns() []TableColumn
	GetPrimaryKey() []TableColumn
}

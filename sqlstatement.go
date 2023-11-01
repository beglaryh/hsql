package hsql

type SqlStatement interface {
	Generate() (Sql, error)
}

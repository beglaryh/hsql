package hsql

type Sql struct {
	Sql        string
	Parameters map[string]string
}

func NewSql(sql string, parameters map[string]string) Sql {
	return Sql{sql, parameters}
}

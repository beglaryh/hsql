package hsql

type Sql struct {
	Sql        string
	Parameters map[string]any
}

func NewSql(sql string, parameters map[string]any) Sql {
	return Sql{sql, parameters}
}

package insert

import (
	"errors"
	"github.com/beglaryh/hsql"
	"github.com/beglaryh/hsql/persistence"
	"strconv"
	"strings"
)

type Insert struct {
	table  hsql.Table
	values []persistence.Value
}

func NewInsert() *Insert {
	return &Insert{}
}

func (insert *Insert) Table(table hsql.Table) *Insert {
	insert.table = table
	return insert
}

func (insert *Insert) With(pv persistence.Value) *Insert {
	insert.values = append(insert.values, pv)
	return insert
}

func (insert *Insert) Generate() (*hsql.Sql, error) {
	if insert.table == nil {
		return nil, errors.New("missing table")
	}
	if len(insert.values) == 0 {
		return nil, errors.New("empty insert")
	}
	if !insert.allColumnsBelongToTheSameTable() {
		return nil, errors.New("mismatched column insertions")
	}
	missingColumns := insert.hasNonNullableColumns()
	if len(missingColumns) > 0 {
		return nil, errors.New("Missing non nullable columns:" + strings.Join(missingColumns, ", "))
	}
	params := map[string]string{}
	param := "v"
	paramCount := 0
	sBuilder := strings.Builder{}
	sBuilder.WriteString("INSERT INTO")
	sBuilder.WriteString(" ")
	sBuilder.WriteString(insert.table.GetName())
	sBuilder.WriteString("(")

	for index, e := range insert.values {
		sBuilder.WriteString(e.GetColumn().GetName())
		if index != len(insert.values)-1 {
			sBuilder.WriteString(", ")
		}
	}
	sBuilder.WriteString(")\n")
	sBuilder.WriteString("VALUES (")
	for index, e := range insert.values {
		p := param + strconv.Itoa(paramCount)
		paramCount = paramCount + 1
		params[p] = e.GetValue()
		sBuilder.WriteString(":" + p)
		if index != len(insert.values)-1 {
			sBuilder.WriteString(", ")
		}
	}
	sBuilder.WriteString(")")
	sql := hsql.NewSql(sBuilder.String(), params)
	return &sql, nil
}

func (insert *Insert) allColumnsBelongToTheSameTable() bool {
	table := insert.table.GetName()
	for _, e := range insert.values {
		if e.GetColumn().GetTable() != table {
			return false
		}
	}
	return true
}

func (insert *Insert) hasNonNullableColumns() []string {
	var missingColumns []string
	var nonNullColumns []string
	for _, e := range insert.table.GetColumns() {
		if !e.IsNullable() {
			nonNullColumns = append(nonNullColumns, e.GetName())
		}
	}

	columns := map[string]int{}
	for _, e := range insert.values {
		columns[e.GetColumn().GetName()] = 1
	}

	for _, e := range nonNullColumns {
		_, ok := columns[e]
		if !ok {
			missingColumns = append(missingColumns, e)
		}
	}
	return missingColumns
}

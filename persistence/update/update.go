package update

import (
	"encoding/json"
	"errors"
	"github.com/beglaryh/hsql"
	"github.com/beglaryh/hsql/persistence"
	"strconv"
)

type Update struct {
	table   hsql.Table
	setters []persistence.PersistenceValue
	filters []hsql.Filter
}

func NewUpdate() *Update {
	return &Update{}
}

func (update *Update) Set(value persistence.PersistenceValue) *Update {
	update.setters = append(update.setters, value)
	return update
}

func (update *Update) Table(table hsql.Table) *Update {
	update.table = table
	return update
}

func (update *Update) Where(filter hsql.Filter) *Update {
	update.filters = append(update.filters, filter)
	return update
}

func (update *Update) Generate() (*hsql.Sql, error) {
	if update.table == nil {
		return nil, errors.New("missing table")
	}
	columnSql, columnParams := update.generateColumns()
	conditionSql, conditionParams := update.generateWhere()
	params := map[string]string{}

	for k, v := range columnParams {
		params[k] = v
	}
	for k, v := range conditionParams {
		params[k] = v
	}

	sql := "UPDATE\n\t" + update.table.GetName() + "\n" + "SET" + columnSql + conditionSql
	responseSql := hsql.NewSql(sql, params)
	return &responseSql, nil
}

func (update *Update) generateColumns() (string, map[string]string) {
	s := ""
	params := map[string]string{}
	for index, e := range update.setters {
		param := "v" + strconv.Itoa(index)
		s += "\n\t" + e.GetColumn().AsTableColumn() + " = :" + param
		if index != len(update.setters)-1 {
			s += ","
		}
		params[param] = e.GetValue()
	}
	return s, params
}

func (update *Update) generateWhere() (string, map[string]string) {
	s := ""
	params := map[string]string{}
	for index, e := range update.filters {
		param := "f" + strconv.Itoa(index)
		s += "\n\t" + e.GetColumn().AsTableColumn() + " = :" + param
		if index != len(update.filters)-1 {
			s += ","
		}

		j, _ := json.Marshal(e.GetValue())
		params[param] = string(j)
	}
	return s, params
}

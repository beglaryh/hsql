package hsql

import (
	"errors"
	"strconv"
)

type UpdateValue struct {
	column TableColumn
	value  string
}
type Update struct {
	table   Table
	setters []UpdateValue
	filters []Filter
}

func NewUpdate() *Update {
	return &Update{}
}

func (update *Update) Table(table Table) *Update {
	update.table = table
	return update
}

func (update *Update) Where(filter Filter) *Update {
	update.filters = append(update.filters, filter)
	return update
}

func (update *Update) Generate() (Sql, error) {
	if update.table == nil {
		return Sql{}, errors.New("missing table")
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
	return Sql{sql, params}, nil
}

func (update *Update) generateColumns() (string, map[string]string) {
	s := ""
	params := map[string]string{}
	for index, e := range update.setters {
		param := "v" + strconv.Itoa(index)
		s += "\n\t" + e.column.AsTableColumn() + " = :" + param
		if index != len(update.setters)-1 {
			s += ","
		}
		params[param] = e.value
	}
	return s, params
}

func (update *Update) generateWhere() (string, map[string]string) {
	s := ""
	params := map[string]string{}
	for index, e := range update.filters {
		param := "f" + strconv.Itoa(index)
		s += "\n\t" + e.column.AsTableColumn() + " = :" + param
		if index != len(update.filters)-1 {
			s += ","
		}
		params[param] = e.value.val
	}
	return s, params
}

func (update *Update) Set(column TableColumn, value string) *Update {
	updateValue := UpdateValue{column: column, value: value}
	update.setters = append(update.setters, updateValue)
	return update
}

package query

import (
	"encoding/json"
	"errors"
	"github.com/beglaryh/hsql"
	"strconv"
	"strings"
)

type Query struct {
	selection []hsql.TableColumn
	tables    []hsql.Table
	filters   []hsql.Filter
	sorts     []hsql.Sort
	hasJoin   bool
}

func NewQuery() *Query {
	return &Query{
		selection: []hsql.TableColumn{},
		filters:   []hsql.Filter{},
	}
}

func (query *Query) Select(column ...hsql.TableColumn) *Query {
	for _, e := range column {
		query.selection = append(query.selection, e)
	}
	return query
}

func (query *Query) From(tables ...hsql.Table) *Query {
	for _, table := range tables {
		query.tables = append(query.tables, table)
	}
	if len(query.tables) > 1 {
		query.hasJoin = true
	}
	return query
}

func (query *Query) Where(filter *hsql.Filter) *Query {
	query.filters = append(query.filters, *filter)
	return query
}

func (query *Query) OrderBy(sort hsql.Sort) *Query {
	query.sorts = append(query.sorts, sort)
	return query
}

func (query *Query) Generate() (*hsql.Sql, error) {
	var sql string
	if len(query.selection) == 0 {
		return nil, errors.New("no selection specified")
	}
	filter, params := query.withFilter()
	tables, err := query.withTables()
	if err != nil {
		return nil, err
	}
	sql = strings.Replace(hsql.QUERY_FORMAT, ":COLUMNS", query.withColumns(), 1)
	sql = strings.Replace(sql, ":TABLES", tables, 1)
	sql = strings.Replace(sql, ":WHERE", filter, 1)
	sql = strings.Replace(sql, ":ORDER", query.withSort(), 1)
	sql = strings.Replace(sql, "\n\n", "\n", -1)
	if sql[len(sql)-1] == '\n' {
		sql = sql[0 : len(sql)-1]
	}
	response := hsql.NewSql(sql, params)
	return &response, nil
}

func (query *Query) withColumns() string {
	columns := ""
	for index, column := range query.selection {
		columns = columns + "\t" + column.AsTableColumn()
		if index != len(query.selection)-1 {
			columns = columns + ",\n"
		}
	}
	return columns
}

func (query *Query) withFilter() (string, map[string]string) {
	if len(query.filters) == 0 {
		return "", nil
	}
	if query.hasJoin {
		joinFilters := query.createJoins()
		for _, filter := range joinFilters {
			query.filters = append(query.filters, filter)
		}
	}

	p := "p"
	paramCount := 0
	params := map[string]string{}
	filterString := "WHERE"
	for index, filter := range query.filters {
		filterString += "\n\t"
		if index != 0 {
			filterString += filter.GetPredicate() + " "
		}
		switch filter.GetOperator() {
		case hsql.Eq:
			filterString += filter.GetColumn().AsTableColumn() + " = :"
			otherColumn, ok := filter.GetValue().(hsql.TableColumn)
			if ok {
				filterString += otherColumn.AsTableColumn()
			} else {
				param := p + strconv.Itoa(paramCount)
				filterString += param
				paramCount += 1
				vs, isString := filter.GetValue().(string)
				if isString {
					params[param] = vs
				} else {
					j, _ := json.Marshal(filter.GetValue())
					params[param] = string(j)
				}
			}
			break
		case hsql.Like:
			param := p + strconv.Itoa(paramCount)
			paramCount += 1
			vs, isString := filter.GetValue().(string)
			pv := ""
			if isString {
				pv = vs
			} else {
				j, _ := json.Marshal(filter.GetValue())
				pv = string(j)
			}
			likeString := " LIKE CONCAT ('%', " + param + ", '%')"
			filterString += filter.GetColumn().AsTableColumn() + likeString
			params[param] = pv
			break
		}
	}

	return filterString, params
}

func (query *Query) createJoins() []hsql.Filter {
	var joinFilters []hsql.Filter
	joinColumns := toTableColumns(query.tables)
	for _, column := range joinColumns {
		if column.HasForeignKey() {
			foreign := column.GetForeignKey()
			if column.GetType() == hsql.JsonB {
				joinFilter := hsql.Column(foreign).InColumn(column)
				joinFilters = append(joinFilters, *joinFilter)
			} else {
				joinFilter := hsql.Column(foreign).Eq(column)
				joinFilters = append(joinFilters, *joinFilter)
			}
		}
	}
	return joinFilters
}

func toTableColumns(tables []hsql.Table) []hsql.TableColumn {
	var columns []hsql.TableColumn
	for _, table := range tables {
		for _, column := range table.GetColumns() {
			columns = append(columns, column)
		}
	}
	return columns
}

func (query *Query) withTables() (string, error) {
	tables := ""
	if len(query.tables) == 0 {
		tables := map[string]bool{}
		for _, column := range query.selection {
			tables[column.GetTable()] = true
		}

		if len(tables) > 1 {
			return "", errors.New("join queries require explicit From table expression")
		}
		return "\t" + query.selection[0].GetTable(), nil
	}
	for index, table := range query.tables {
		tables += "\t" + table.GetName()
		if index != len(query.tables)-1 {
			tables += ",\n"
		}
	}
	return tables, nil
}

func getColumnsFromFilter(filters []hsql.Filter) []hsql.TableColumn {
	var tableColumns []hsql.TableColumn

	for _, filter := range filters {
		if len(filter.GetColumn().GetName()) != 0 {
			tableColumns = append(tableColumns, filter.GetColumn())
		}
		value := filter.GetValue()
		otherColumn, ok := value.(hsql.TableColumn)
		if ok {
			tableColumns = append(tableColumns, otherColumn)
		}
		nestColumns := getColumnsFromFilter(filter.GetNestedFilters())
		for _, nestedColumn := range nestColumns {
			tableColumns = append(tableColumns, nestedColumn)
		}
	}

	return tableColumns
}

func (query *Query) withSort() string {
	s := ""
	for index, sort := range query.sorts {
		if index == 0 {
			s += "ORDER BY"
		}
		s += "\n\t" + sort.GetColumn().AsTableColumn() + " " + sort.GetDirection().String()
		if index != len(query.sorts)-1 {
			s += ","
		}
	}
	return s
}

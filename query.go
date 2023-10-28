package hsql

import (
	"errors"
	"github.com/emirpasic/gods/sets"
	"github.com/emirpasic/gods/sets/hashset"
	"strconv"
	"strings"
)

type Query struct {
	selection []TableColumn
	tables    []Table
	filters   []Filter
	sorts     []Sort
	hasJoin   bool
}

func NewQuery() *Query {
	return &Query{
		selection: []TableColumn{},
		filters:   []Filter{},
	}
}

func (query *Query) Select(column TableColumn) *Query {
	query.selection = append(query.selection, column)
	return query
}

func (query *Query) From(tables ...Table) *Query {
	for _, table := range tables {
		query.tables = append(query.tables, table)
	}
	if len(query.tables) > 1 {
		query.hasJoin = true
	}
	return query
}

func (query *Query) Where(filter *Filter) *Query {
	query.filters = append(query.filters, *filter)
	return query
}

func (query *Query) OrderBy(sort Sort) *Query {
	query.sorts = append(query.sorts, sort)
	return query
}

func (query *Query) Generate() (Sql, error) {
	var sql string
	if len(query.selection) == 0 {
		return Sql{}, errors.New("no selection specified")
	}
	filter, params := query.withFilter()
	tables, err := query.withTables()
	if err != nil {
		return Sql{}, err
	}
	sql = strings.Replace(QUERY_FORMAT, ":COLUMNS", query.withColumns(), 1)
	sql = strings.Replace(sql, ":TABLES", tables, 1)
	sql = strings.Replace(sql, ":WHERE", filter, 1)
	sql = strings.Replace(sql, ":ORDER", query.withSort(), 1)
	sql = strings.Replace(sql, "\n\n", "\n", -1)
	if sql[len(sql)-1] == '\n' {
		sql = sql[0 : len(sql)-1]
	}
	return Sql{sql, params}, nil
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
		case eq:
			filterString += filter.getColumn().AsTableColumn() + " = :"
			if filter.getValue().isColumn() {
				filterString += filter.value.column.AsTableColumn()
			} else {
				param := p + strconv.Itoa(paramCount)
				filterString += param
				paramCount += 1
				params[param] = filter.value.val
			}
			break
		case like:
			likeString := " LIKE CONCAT ('%', " + filter.value.val + ", '%')"
			filterString += filter.getColumn().AsTableColumn() + likeString
			break
		}
	}

	return filterString, params
}

func (query *Query) createJoins() []Filter {
	var joinFilters []Filter
	joinColumns := toTableColumns(query.tables)
	for _, column := range joinColumns {
		if column.HasForeignKey() {
			foreign := column.GetForeignKey()
			if column.GetType() == JsonB {
				joinFilter := Column(foreign).InColumn(column)
				joinFilters = append(joinFilters, *joinFilter)
			} else {
				joinFilter := Column(foreign).EqColumn(column)
				joinFilters = append(joinFilters, *joinFilter)
			}
		}
	}
	return joinFilters
}

func toTableColumns(tables []Table) []TableColumn {
	var columns []TableColumn
	for _, table := range tables {
		for _, column := range table.GetColumns() {
			columns = append(columns, column)
		}
	}
	return columns
}

func (query *Query) getAllTables() sets.Set {
	tables := hashset.New()
	for _, column := range query.selection {
		tables.Add(column.GetTable())
	}
	filterColumns := getColumnsFromFilter(query.filters)
	for _, column := range filterColumns {
		tables.Add(column.GetTable())
	}

	return tables
}

func (query *Query) withTables() (string, error) {
	tables := ""
	if len(query.tables) == 0 {
		tables := hashset.New()
		for _, column := range query.selection {
			tables.Add(column.tableName)
		}

		if tables.Size() > 1 {
			return "", errors.New("join queries require explicit From table expression")
		}
		return "\t" + query.selection[0].tableName, nil
	}
	for index, table := range query.tables {
		tables += "\t" + table.GetName()
		if index != len(query.tables)-1 {
			tables += ",\n"
		}
	}
	return tables, nil
}

func getColumnsFromFilter(filters []Filter) []TableColumn {
	var tableColumns []TableColumn

	for _, filter := range filters {
		if len(filter.getColumn().GetName()) != 0 {
			tableColumns = append(tableColumns, filter.getColumn())
		}
		if filter.getValue().isColumn() {
			tableColumns = append(tableColumns, filter.value.column)
		}
		nestColumns := getColumnsFromFilter(filter.nested)
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
		s += "\n\t" + sort.column.AsTableColumn() + " " + sort.direction.string()
		if index != len(query.sorts)-1 {
			s += ","
		}
	}
	return s
}

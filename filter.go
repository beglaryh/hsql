package hsql

import (
	"encoding/json"
	"strings"
	"time"
)

type Filter struct {
	Predicate Predicate
	column    TableColumn
	operator  Operator
	value     any
	nested    []Filter
}

func NewFilter() *Filter {
	return &Filter{}
}

// TODO create filter builder

func Column(column TableColumn) *Filter {
	filter := NewFilter()
	filter.column = column
	return filter
}

func Value(value any) *Filter {
	filter := NewFilter()
	filter.value = value
	return filter
}

func (filter *Filter) Eq(value any) *Filter {
	_, isColumn := value.(TableColumn)
	if !isColumn {
		if filter.column.columnType == Date {
			date := value.(time.Time)
			value = date.Format(time.DateOnly)
		} else if filter.column.columnType == TimeStampTZ {
			date := value.(time.Time)
			value = date.Format(time.RFC3339)
		} else if filter.column.columnType == TimeStamp {
			date := value.(time.Time)
			value = date.Format(DateTimeFormat)
		} else {
			j, _ := json.Marshal(value)
			js := string(j)
			if strings.HasPrefix(js, `"`) {
				js = js[1 : len(js)-1]
			}
			value = js
		}
	}

	filter.value = value
	filter.operator = Eq
	return filter
}

func (filter *Filter) In(value []any) *Filter {
	filter.value = value
	filter.operator = ColumnIn
	return filter
}

func (filter *Filter) InColumn(column TableColumn) *Filter {
	filter.column = column
	filter.operator = ValueIn
	return filter
}

func (filter *Filter) ValueIn(value any) *Filter {
	filter.value = value
	filter.operator = ValueIn
	return filter
}

func (filter *Filter) Like(value string) *Filter {
	filter.value = value
	filter.operator = Like
	return filter
}

func NestedAnd(filters ...Filter) *Filter {
	filter := NewFilter()
	return filter.setNested(AND, filters)
}

func NestedOr(filters ...Filter) *Filter {
	filter := NewFilter()
	return filter.setNested(OR, filters)
}

func (filter *Filter) setNested(predicate Predicate, filters []Filter) *Filter {
	for index := range filters {
		filters[index].Predicate = predicate
	}
	filter.nested = filters
	return filter
}

func (filter *Filter) GetNestedFilters() []Filter {
	return filter.nested
}

func (filter *Filter) GetColumn() TableColumn {
	return filter.column
}

func (filter *Filter) GetPredicate() string {
	if filter.Predicate == 0 {
		return AND.string()
	}
	return filter.Predicate.string()
}

func (filter *Filter) GetOperator() Operator {
	return filter.operator
}

func (filter *Filter) GetValue() any {
	return filter.value
}

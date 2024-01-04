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

type FilterBuilder struct {
	filter *Filter
}

func NewFilter() *Filter {
	return &Filter{}
}

// TODO create filter builder

func Column(column TableColumn) *FilterBuilder {
	return &FilterBuilder{filter: &Filter{column: column}}
}

func (builder *FilterBuilder) Eq(value any) Filter {
	_, isColumn := value.(TableColumn)
	if !isColumn {
		if builder.filter.column.columnType == Date {
			date := value.(time.Time)
			value = date.Format(time.DateOnly)
		} else if builder.filter.column.columnType == TimeStampTZ {
			date := value.(time.Time)
			value = date.Format(time.RFC3339)
		} else if builder.filter.column.columnType == TimeStamp {
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

	builder.filter.value = value
	builder.filter.operator = Eq
	return *builder.filter
}

func (filter *Filter) In(value []any) *Filter {
	filter.value = value
	filter.operator = ColumnIn
	return filter
}

func (builder *FilterBuilder) Like(value string) Filter {
	builder.filter.value = value
	builder.filter.operator = Like
	return *builder.filter
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

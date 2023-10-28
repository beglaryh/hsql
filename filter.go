package hsql

type Filter struct {
	Predicate Predicate
	column    TableColumn
	operator  Operator
	value     value
	nested    []Filter
}

func NewFilter() *Filter {
	return &Filter{}
}

func Column(column TableColumn) *Filter {
	filter := NewFilter()
	filter.column = column
	return filter
}

func Value(value string) *Filter {
	filter := NewFilter()
	filter.value = newValue(value)
	return filter
}

func (filter *Filter) Eq(value string) *Filter {
	filter.value = newValue(value)
	filter.operator = eq
	return filter
}

func (filter *Filter) EqColumn(column TableColumn) *Filter {
	filter.value = newValueFromColumn(column)
	filter.operator = eq
	return filter
}

func (filter *Filter) In(value []string) *Filter {
	filter.value = newValues(value)
	filter.operator = columnIn
	return filter
}

func (filter *Filter) InColumn(column TableColumn) *Filter {
	filter.column = column
	filter.operator = valueIn
	return filter
}

func (filter *Filter) ValueIn(value string) *Filter {
	filter.value = newValue(value)
	filter.operator = valueIn
	return filter
}

func (filter *Filter) Like(value string) *Filter {
	filter.value = newValue(value)
	filter.operator = like
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

func (filter *Filter) getColumn() TableColumn {
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

func (filter *Filter) getValue() value {
	return filter.value
}

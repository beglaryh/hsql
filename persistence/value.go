package persistence

import (
	"github.com/beglaryh/hsql"
	"time"
)

const DateTimeFormat = "2006-01-02T15:04:05"

type PersistenceValue struct {
	column hsql.TableColumn
	value  string
}

type PersistenceValueBuilder struct {
	value *PersistenceValue
}

func Column(column hsql.TableColumn) *PersistenceValueBuilder {
	value := PersistenceValue{column: column}
	return &PersistenceValueBuilder{value: &value}
}

func (builder *PersistenceValueBuilder) Eq(value any) PersistenceValue {
	s, isString := value.(string)
	if isString {
		builder.value.value = s
		return *builder.value
	}

	t, isTime := value.(time.Time)
	timeValue := ""
	if isTime {
		columType := builder.value.column.GetType()
		switch columType {
		case hsql.Date:
			timeValue = t.Format(time.DateOnly)
		case hsql.TimeStamp:
			timeValue = t.Format(DateTimeFormat)
		case hsql.TimeStampZ:
			timeValue = t.Format(time.RFC3339)
		}
		builder.value.value = timeValue
		return *builder.value
	}

	return *builder.value
}

func (pv PersistenceValue) GetColumn() hsql.TableColumn {
	return pv.column
}

func (pv PersistenceValue) GetValue() string {
	return pv.value
}

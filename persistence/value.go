package persistence

import (
	"encoding/json"
	"github.com/beglaryh/hsql"
	"strings"
	"time"
)

type Value struct {
	column hsql.TableColumn
	value  string
}

type ValueBuilder struct {
	value *Value
}

func Column(column hsql.TableColumn) *ValueBuilder {
	value := Value{column: column}
	return &ValueBuilder{value: &value}
}

func (builder *ValueBuilder) Eq(value any) Value {
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
			timeValue = t.Format(hsql.DateTimeFormat)
		case hsql.TimeStampTZ:
			timeValue = t.Format(time.RFC3339)
		}
		builder.value.value = timeValue
		return *builder.value
	}

	v, _ := json.Marshal(value)
	js := string(v)
	if strings.HasPrefix(js, `"`) {
		js = js[1 : len(js)-1]
	}
	builder.value.value = js

	return *builder.value
}

func (pv Value) GetColumn() hsql.TableColumn {
	return pv.column
}

func (pv Value) GetValue() string {
	return pv.value
}

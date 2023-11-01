package test

import (
	"fmt"
	"github.com/beglaryh/hsql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdate1(t *testing.T) {
	sql, err := hsql.NewUpdate().
		Table(NewPersonTable()).
		Set(firstName, "Bob").
		Set(lastName, "Yarn").
		Generate()

	if err != nil {
		t.Fail()
	}
	fmt.Println(sql.Sql)
	assert.Equal(t, sql1, sql.Sql)
}

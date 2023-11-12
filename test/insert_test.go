package test

import (
	. "github.com/beglaryh/hsql/persistence"
	"github.com/beglaryh/hsql/persistence/insert"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert1(t *testing.T) {
	_, err := insert.NewInsert().
		Table(NewPersonTable()).
		With(Column(personId).Eq("ABC")).
		Generate()

	if err == nil {
		t.Fail()
	}
}

func TestInsert2(t *testing.T) {
	_, err := insert.NewInsert().
		Generate()

	if err == nil {
		t.Fail()
	}
}

func TestInsert3(t *testing.T) {
	_, err := insert.NewInsert().
		Table(NewPersonTable()).
		With(Column(personId).Eq("ABC")).
		With(Column(firstName).Eq("John")).
		With(Column(lastName).Eq("Doe")).
		With(Column(companyId).Eq("CDE")).
		Generate()

	if err == nil {
		t.Fail()
	}
}

func TestInsert4(t *testing.T) {
	sql, err := insert.NewInsert().
		Table(NewPersonTable()).
		With(Column(personId).Eq("ABC")).
		With(Column(firstName).Eq("John")).
		With(Column(lastName).Eq("Doe")).
		With(Column(companyForeignKey).Eq("CDE")).
		Generate()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expected_insert, sql.Sql)
}

package test

import (
	"github.com/beglaryh/hsql/persistence"
	"github.com/beglaryh/hsql/persistence/insert"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert1(t *testing.T) {
	_, err := insert.NewInsert().
		Table(NewPersonTable()).
		Column(persistence.With(personId).Eq("ABC")).
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
		Column(persistence.With(personId).Eq("ABC")).
		Column(persistence.With(firstName).Eq("John")).
		Column(persistence.With(lastName).Eq("Doe")).
		Column(persistence.With(companyId).Eq("CDE")).
		Generate()

	if err == nil {
		t.Fail()
	}
}

func TestInsert4(t *testing.T) {
	sql, err := insert.NewInsert().
		Table(NewPersonTable()).
		Column(persistence.With(personId).Eq("ABC")).
		Column(persistence.With(firstName).Eq("John")).
		Column(persistence.With(lastName).Eq("Doe")).
		Column(persistence.With(companyForeignKey).Eq("CDE")).
		Generate()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expected_insert, sql.Sql)
}

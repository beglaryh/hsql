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
		Column(persistence.Column(personId).Eq("ABC")).
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
		Column(persistence.Column(personId).Eq("ABC")).
		Column(persistence.Column(firstName).Eq("John")).
		Column(persistence.Column(lastName).Eq("Doe")).
		Column(persistence.Column(companyId).Eq("CDE")).
		Generate()

	if err == nil {
		t.Fail()
	}
}

func TestInsert4(t *testing.T) {
	sql, err := insert.NewInsert().
		Table(NewPersonTable()).
		Column(persistence.Column(personId).Eq("ABC")).
		Column(persistence.Column(firstName).Eq("John")).
		Column(persistence.Column(lastName).Eq("Doe")).
		Column(persistence.Column(companyForeignKey).Eq("CDE")).
		Generate()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expected_insert, sql.Sql)
}

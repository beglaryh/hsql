package test

import (
	"github.com/beglaryh/hsql"
	"github.com/beglaryh/hsql/query"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQuery(t *testing.T) {
	sql, err := query.NewQuery().
		Select(firstName).
		Select(lastName).
		Select(dateOfBirth).
		From(NewPersonTable()).
		Where(hsql.Column(firstName).Eq("hrach")).
		Where(hsql.Column(lastName).Eq("beglaryan")).
		OrderBy(hsql.Asc(firstName)).
		OrderBy(hsql.Asc(lastName)).
		Generate()

	if err != nil {
		t.Failed()
	}
	expectedParams := map[string]any{}
	expectedParams["p0"] = "hrach"
	expectedParams["p1"] = "beglaryan"
	assert.Equal(t, sql1, sql.Sql)
	assert.Equal(t, expectedParams, sql.Parameters)
}

func TestNewQuery2(t *testing.T) {
	sql, err := query.NewQuery().
		Select(firstName).
		Select(lastName).
		Select(dateOfBirth).
		Select(companyId).
		From(NewPersonTable()).
		From(NewCompanyTable()).
		Where(hsql.Column(firstName).Eq("hrach")).
		Where(hsql.Column(lastName).Eq("beglaryan")).
		OrderBy(hsql.Asc(firstName)).
		OrderBy(hsql.Asc(lastName)).
		Generate()
	if err != nil {
		t.Failed()
	}
	assert.Equal(t, sql2, sql.Sql)
}

func TestNewQuery3(t *testing.T) {
	sql, err := query.NewQuery().
		Select(firstName).
		Select(lastName).
		Select(dateOfBirth).
		Where(hsql.Column(firstName).Like("hrach")).
		Generate()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, sql3, sql.Sql)
}

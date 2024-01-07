package test

import (
	"github.com/beglaryh/hsql"
	"github.com/beglaryh/hsql/persistence"
	"github.com/beglaryh/hsql/persistence/update"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUpdate1(t *testing.T) {
	sql, err := update.NewUpdate().
		Table(NewPersonTable()).
		Set(persistence.Column(firstName).Eq("Bob")).
		Set(persistence.Column(lastName).Eq("Yarn")).
		Generate()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, update1, sql.Sql)
}

func TestUpdate2(t *testing.T) {
	sql, err := update.NewUpdate().
		Table(NewPersonTable()).
		Set(persistence.Column(firstName).Eq("Bob")).
		Set(persistence.Column(lastName).Eq("Yarn")).
		Set(persistence.Column(dateOfBirth).Eq(time.Now())).
		Generate()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "Bob", sql.Parameters["v0"])
	assert.Equal(t, "Yarn", sql.Parameters["v1"])
	assert.Equal(t, time.Now().Format(time.DateOnly), sql.Parameters["v2"])
	assert.Equal(t, update2, sql.Sql)
}

func TestUpdate3(t *testing.T) {
	sql, err := update.NewUpdate().
		Table(NewPersonTable()).
		Set(persistence.Column(firstName).Eq("Bob")).
		Set(persistence.Column(lastName).Eq("Yarn")).
		Set(persistence.Column(dateOfBirth).Eq(time.Now())).
		Where(hsql.Column(personId).Eq("ID")).
		Generate()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "Bob", sql.Parameters["v0"])
	assert.Equal(t, "Yarn", sql.Parameters["v1"])
	assert.Equal(t, "ID", sql.Parameters["f0"])
	assert.Equal(t, time.Now().Format(time.DateOnly), sql.Parameters["v2"])
	assert.Equal(t, update3, sql.Sql)
}

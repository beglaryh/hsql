package test

import (
	"github.com/beglaryh/hsql/persistence"
	"github.com/beglaryh/hsql/persistence/update"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUpdate1(t *testing.T) {
	sql, err := update.NewUpdate().
		Table(NewPersonTable()).
		Set(persistence.With(firstName).Eq("Bob")).
		Set(persistence.With(lastName).Eq("Yarn")).
		Generate()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, update1, sql.Sql)
}

func TestUpdate2(t *testing.T) {
	sql, err := update.NewUpdate().
		Table(NewPersonTable()).
		Set(persistence.With(firstName).Eq("Bob")).
		Set(persistence.With(lastName).Eq("Yarn")).
		Set(persistence.With(dateOfBirth).Eq(time.Now())).
		Generate()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "Bob", sql.Parameters["v0"])
	assert.Equal(t, "Yarn", sql.Parameters["v1"])
	assert.Equal(t, time.Now().Format(time.DateOnly), sql.Parameters["v2"])
	assert.Equal(t, update2, sql.Sql)
}

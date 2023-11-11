package test

import (
	"fmt"
	"github.com/beglaryh/hsql/update"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUpdate1(t *testing.T) {
	sql, err := update.NewUpdate().
		Table(NewPersonTable()).
		Set(update.Column(firstName).Eq("Bob")).
		Set(update.Column(lastName).Eq("Yarn")).
		Generate()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, update1, sql.Sql)
}

func TestUpdate2(t *testing.T) {
	timestamp := time.Now().Format(time.DateOnly)
	fmt.Println(timestamp)
	sql, err := update.NewUpdate().
		Table(NewPersonTable()).
		Set(update.Column(firstName).Eq("Bob")).
		Set(update.Column(lastName).Eq("Yarn")).
		Set(update.Column(dateOfBirth).EqDate(time.Now())).
		Generate()

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "Bob", sql.Parameters["v0"])
	assert.Equal(t, "Yarn", sql.Parameters["v1"])
	assert.Equal(t, time.Now().Format(time.DateOnly), sql.Parameters["v2"])
	assert.Equal(t, update2, sql.Sql)
}

package test

import . "github.com/beglaryh/hsql"

type PersonTable struct {
}

const personTableName = "person"

/* Defines Columns */
var personId = NewColumnBuilder(personTableName, "id", UUID).
	IsMutable(false).
	IsNullable(false).
	Build()
var firstName = NewColumnBuilder(personTableName, "first_name", String).IsNullable(false).Build()
var lastName = NewColumnBuilder(personTableName, "last_name", String).IsNullable(false).Build()
var middleName = NewColumn(personTableName, "last_name", String)
var dateOfBirth = NewColumnBuilder(personTableName, "dob", Date).IsNullable(false).Build()
var status = NewColumnBuilder(personTableName, "status", Boolean).IsNullable(false).Build()
var companyForeignKey = NewColumnBuilder(personTableName, "company_id", UUID).
	WithForeignKey(companyId).
	IsNullable(false).
	Build()

func NewPersonTable() PersonTable {
	return PersonTable{}
}

/* Implement Table Interface */
func (table PersonTable) GetName() string {
	return personTableName
}

func (table PersonTable) GetColumns() []TableColumn {
	return []TableColumn{personId, firstName, lastName, middleName, companyForeignKey}
}

func (table PersonTable) GetPrimaryKey() []TableColumn {
	return []TableColumn{personId}
}

/* Getters */
func (table PersonTable) GetId() TableColumn {
	return personId
}

func (table PersonTable) GetFirstName() TableColumn {
	return firstName
}

func (table PersonTable) GetLastName() TableColumn {
	return lastName
}

func (table PersonTable) GetMiddleName() TableColumn {
	return middleName
}

func (table PersonTable) GetCompanyId() TableColumn {
	return companyForeignKey
}

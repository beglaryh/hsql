package test

import . "github.com/beglaryh/hsql"

type PersonTable struct {
}

const personTableName = "person"

/* Defines Columns */
var personId = NewTableColumnBuilder(personTableName, "id", UUID).
	IsMutable(false).
	Build()
var firstName = NewTableColumn(personTableName, "first_name", String)
var lastName = NewTableColumn(personTableName, "last_name", String)
var middleName = NewTableColumn(personTableName, "last_name", String)
var dateOfBirth = NewTableColumn(personTableName, "dob", Date)
var companyForeignKey = NewTableColumnBuilder(personTableName, "company_id", UUID).
	WithForeignKey(companyId).
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

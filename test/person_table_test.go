package test

import "github.com/beglaryh/hsql"

/* Defines Columns */
var id = hsql.NewTableColumn(tableName, "id", hsql.UUID, nil)
var firstName = hsql.NewTableColumn(tableName, "first_name", hsql.String, nil)
var lastName = hsql.NewTableColumn(tableName, "last_name", hsql.String, nil)
var middleName = hsql.NewTableColumn(tableName, "middle_name", hsql.String, nil)
var dateOfBirth = hsql.NewTableColumn(tableName, "dob", hsql.Date, nil)
var companyForeignKey = hsql.NewTableColumn(tableName, "company_id", hsql.UUID, &companyId)

func NewPersonTable() PersonTable {
	return PersonTable{}
}

/* Implement Table Interface */
func (table PersonTable) GetName() string {
	return tableName
}

func (table PersonTable) GetColumns() []hsql.TableColumn {
	return []hsql.TableColumn{id, firstName, lastName, middleName, companyForeignKey}
}

func (table PersonTable) GetPrimaryKey() []hsql.TableColumn {
	return []hsql.TableColumn{id}
}

/* Getters */
func (table PersonTable) getId() hsql.TableColumn {
	return id
}

func (table PersonTable) getFirstName() hsql.TableColumn {
	return firstName
}

func (table PersonTable) getLastName() hsql.TableColumn {
	return lastName
}

func (table PersonTable) getMiddleName() hsql.TableColumn {
	return middleName
}

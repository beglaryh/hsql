package test

import . "github.com/beglaryh/hsql"

const companyTableName = "company"

var companyId = NewColumnBuilder(companyTableName, "id", UUID).
	IsMutable(false).
	Build()

type CompanyTable struct {
}

func NewCompanyTable() CompanyTable {
	return CompanyTable{}
}

func (ct CompanyTable) GetColumns() []TableColumn {
	return []TableColumn{companyId}
}

func (ct CompanyTable) GetPrimaryKey() []TableColumn {
	return []TableColumn{companyId}
}

func (ct CompanyTable) GetName() string {
	return companyTableName
}

func (ct CompanyTable) GetId() TableColumn {
	return companyId
}

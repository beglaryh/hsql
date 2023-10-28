package test

import "github.com/beglaryh/hsql"

var companyId = hsql.NewTableColumn("company", "id", hsql.UUID, nil)

type CompanyTable struct {
}

func NewCompanyTable() CompanyTable {
	return CompanyTable{}
}

func (ct CompanyTable) GetColumns() []hsql.TableColumn {
	return []hsql.TableColumn{companyId}
}

func (ct CompanyTable) GetPrimaryKey() []hsql.TableColumn {
	return []hsql.TableColumn{companyId}
}

func (ct CompanyTable) GetName() string {
	return "company"
}

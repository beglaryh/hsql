# h-sql
GO SQL ORM

## TableColumn
- Defines the following properties for a sql column:
    - Table association
    - Name of column
    - Type of Column (i.e. String, Int, Date...)
    - Foreign Key

## Table
A Table is defined by the following interface:
```go
type Table interface {
	getName() string
	getColumns() []TableColumn
	getPrimaryKey() []TableColumn
}
```

#### Implementation 
##### Define Columns
```go
const personTableName = "person"

var personId = NewTableColumnBuilder(personTableName, "id", UUID).
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

```
##### Implement Interface
```go
type PersonTable struct {
}

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
```

##### Additional Getters
```go
func (table PersonTable) getId() TableColumn {
	return id
}

func (table PersonTable) getFirstName() TableColumn {
	return firstName
}

func (table PersonTable) getLastName() TableColumn {
	return lastName
}

func (table PersonTable) getMiddleName() TableColumn {
	return middleName
}
```

##  Query Examples
### Simple
```go
query := NewQuery().
    Select(userId)
    Select(firstName).
    Select(lastName).
    Where(Column(firstName).Eq("John"))
    Where(Column(secondName).Eq("Doe"))

var sql Sql = query.Generate()
```

The following SQL will be generated from the above query:
```sql
SELECT
    user.id
    user.first_name
    user.last_name
FROM
    user
WHERE
    user.first_name = :p0
    user.last_name = :p1
```

Of course the parameters will be:
```json
{
    "p0" : "John",
    "p1" : "Doe"
}
```

### With Ordering
```go
query := NewQuery().
    Select(userId)
    Select(firstName).
    Select(lastName).
    Where(Column(firstName).like("John")).
    OrderBy(Asc(firstName))
    OrderBy(Asc(lastName))
    OrderBy(Desc(middleName))
```  

Generated Sql
```sql
SELECT
    user.id,
    user.first_name,
    user.last_name
FROM
    user
WHERE
    user.first_name LIKE %:p0%
ORDER BY
    user.first_name ASC,
    user.last_name ASC
    user.middle_name DESC
```    

### Nested Conditions
```go
query := NewQuery().
    Select(userId)
    Select(firstName).
    Select(lastName).
    Where(Column(firstName).like("John")).
    Where(Or(
        Column(lastName).Eq("Doe"),
        Column(lastName).Eq("Edwards")
    ))
    OrderBy(Asc(firstName))
    OrderBy(Asc(lastName))
    OrderBy(Desc(middleName))
```  

Generated Sql
```sql
SELECT
    user.id,
    user.first_name,
    user.last_name
FROM
    user
WHERE
    user.first_name LIKE %:p0%
    AND (
        user.last_name = :p1
        OR user.last_name = :p2
    )
ORDER BY
    user.first_name ASC,
    user.last_name ASC
    user.middle_name DESC
```    

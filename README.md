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
var tableName = "user"

var id = NewTableColumn(tableName, "id", UUID, nil)
var firstName = NewTableColumn(tableName, "first_name", String, nil)
var lastName = NewTableColumn(tableName, "last_name", String, nil)
var middleName = NewTableColumn(tableName, "middle_name", String, nil)
var dateOfBirth = NewTableColumn(tableName, "dob", Date, nil)

```
##### Implement Interface
```go
type PersonTable struct {
}

/* Defines Columns */

func NewPersonTable() PersonTable {
	return PersonTable{}
}

/* Implement Table Interface */
func (table PersonTable) getName() string {
	return tableName
}

func (table PersonTable) getColumns() []TableColumn {
	return []TableColumn{id, firstName, lastName, middleName}
}

func (table PersonTable) getPrimaryKey() []TableColumn {
	return []TableColumn{id}
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

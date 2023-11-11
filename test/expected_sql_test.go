package test

var sql1 = `SELECT
	person.first_name,
	person.last_name,
	person.dob
FROM
	person
WHERE
	person.first_name = :p0
	AND person.last_name = :p1
ORDER BY
	person.first_name ASC,
	person.last_name ASC`

var sql2 = `SELECT
	person.first_name,
	person.last_name,
	person.dob,
	company.id
FROM
	person,
	company
WHERE
	person.first_name = :p0
	AND person.last_name = :p1
	AND company.id = :person.company_id
ORDER BY
	person.first_name ASC,
	person.last_name ASC`

var sql3 = `SELECT
	person.first_name,
	person.last_name,
	person.dob
FROM
	person
WHERE
	person.first_name LIKE CONCAT ('%', hrach, '%')`

var update1 = `UPDATE
	person
SET
	person.first_name = :v0,
	person.last_name = :v1`

var update2 = `UPDATE
	person
SET
	person.first_name = :v0,
	person.last_name = :v1,
	person.dob = :v2`

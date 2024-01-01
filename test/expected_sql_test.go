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
	person.first_name LIKE CONCAT ('%', p0, '%')`

var sql4 = `SELECT
	person.first_name,
	person.last_name,
	person.dob,
	COUNT(*) OVER() AS query_total_count
FROM
	person
WHERE
	person.first_name LIKE CONCAT ('%', p0, '%')
LIMIT 10 OFFSET 100`

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

var expected_insert = `INSERT INTO person(id, first_name, last_name, company_id)
VALUES (:v0, :v1, :v2, :v3)`

# Module 19 Notes: Database Connectivity

## MySQL Driver

Go requires an external driver to communicate with a specific database.
Import: `_ "github.com/go-sql-driver/mysql"` (side-effect import).

## SQL Operations

- `db.Query()`: Used to retrieve multiple rows (SELECT).
- `rows.Scan()`: Transfers data from SQL rows into Go variables/structs.
- `defer rows.Close()`: Important to prevent database connection leaks.

## Repository Pattern

Separating SQL code into a `repository` folder makes the code easier to test and maintain.

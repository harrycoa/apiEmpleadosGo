package database

import "database/sql"

func InitDB() *sql.DB {
	connectionString := "root:admin@tcp(localhost:3306)/northwind"
	dbConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error()) // Error Handling
	}
	return dbConnection
}

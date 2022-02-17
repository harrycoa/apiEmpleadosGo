package database

import (
	"database/sql"
	"fmt"
	"log"
)

func InitDB() *sql.DB {
	connectionString := "root:root@tcp(localhost:3306)/northwind"
	dbConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error()) // Error Handling
	}
	defer dbConnection.Close()
	// Ping the database to verify that the connection is valid.
	if err := dbConnection.Ping(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("se ejecuto correctamente")
	return dbConnection

}

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"puppy/database"
)

func main() {
	//fmt.Printf("Hola")
	//r := chi.NewRouter()
	//r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Welcome"))
	//})
	//http.ListenAndServe(":3000", r)
	dbConnect := database.InitDB()
	defer dbConnect.Close()
	fmt.Println(dbConnect)
}

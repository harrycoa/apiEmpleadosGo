package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"puppy/database"
)

var dbConnect *sql.DB

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code"`
	Description  string `json:"description"`
}

func main() {
	// Conexion a la BD
	dbConnect := database.InitDB()
	defer dbConnect.Close()

	// Server Chi
	r := chi.NewRouter()

	// Endpoint de Productos
	r.Get("/products", AllProducts)
	http.ListenAndServe(":3000", r)

}

func AllProducts(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id, product_code, COALESCE(description, '') FROM Products`
	results, err := dbConnect.Query(sql)

	if err != nil {
		log.Fatal(err)
	}
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)

		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	respondwithJSON(w, http.StatusOK, products)
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

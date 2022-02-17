package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"puppy/database"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

var databaseConnection *sql.DB

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code"`
	Description  string `json:"description"`
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("iniciando servidor go")
	// Conexion a la BD
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	// Server Chi
	r := chi.NewRouter()

	// Endpoint de Productos
	r.Get("/products", AllProducts)
	http.ListenAndServe(":4000", r)

}

func AllProducts(w http.ResponseWriter, r *http.Request) {
	const query = "SELECT * FROM products"
	results, err := databaseConnection.Query(query)
	catch(err)
	defer results.Close()
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)

		catch(err)
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

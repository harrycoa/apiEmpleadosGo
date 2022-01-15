package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func main(){
	fmt.Printf("Hola")

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	http.ListenAndServe(":3000", r)
}

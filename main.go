package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	var products Products
	var arr_product []Products
	var response Response

	db := connect()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id, name, image, price FROM products"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&products.Id, &products.Name, &products.Image, &products.Price); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_product = append(arr_product, products)
		}
	}
	response.Status = 1
	response.Message = "Success"
	response.Data = arr_product

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getProducts", returnAllProducts).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to Port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}

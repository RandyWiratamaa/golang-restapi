package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Rest API menggunakan Golang By Randy Wiratama")
}

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	fmt.Println("Connected to Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequest()
}

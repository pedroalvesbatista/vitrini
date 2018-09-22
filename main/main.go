package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
}

type Store struct {
}

type Social struct {
}

func GetSocial() {
}

func GetStore() {

}

func GetProduct() {

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/facebook".GetSocial.Methods("GET"))
	router.HandleFunc("/facebook/{id}".GetSocial.Methods("GET"))
	router.HandleFunc("/product".GetProduct.Methods("GET"))
	router.HandleFunc("/store".GetStore.Methods("GET"))
	router.HandleFunc("/store/{id}".GetStore.Methods("GET"))

	http.ListenAndServe(":8000", router)
}

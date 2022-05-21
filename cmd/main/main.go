package main

import (
	"log"
	"net/http"

	"github.com/codrAlxx/alok/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("local host 9010", r))

}

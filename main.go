package main

import (
	"log"
	"net/http"
	"restMux/handlers"
	"restMux/handlers/catalogs"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	mux := mux.NewRouter()
	prefix := "/shark/golang/api/v0/"
	//Methods of practice
	mux.HandleFunc(prefix+"doOnEjemploGet", handlers.DoOnEjemploGet).Methods("GET")
	mux.HandleFunc(prefix+"doOnEjemploGetWithParameter/{id:[0-9]+}", handlers.DoOnEjemploGetWithParameter).Methods("GET")
	mux.HandleFunc(prefix+"doOnEjemploGetQueryString", handlers.DoOnEjemploGetQueryString).Methods("GET")
	mux.HandleFunc(prefix+"doOnEjemploPost", handlers.DoOnEjemploPost).Methods("POST")
	mux.HandleFunc(prefix+"doOnEjemploPostWithBody", handlers.DoOnEjemploPostWithBody).Methods("POST")
	mux.HandleFunc(prefix+"doOnEjemploPut/{id:[0-9]+}", handlers.DoOnEjemploPut).Methods("PUT")
	mux.HandleFunc(prefix+"doOnEjemploDelete/{id:[0-9]+}", handlers.DoOnEjemploDelete).Methods("DELETE")
	handler := cors.AllowAll().Handler(mux)

	//Methods of catalogs
	mux.HandleFunc(prefix+"catalogs", catalogs.HandleCatalogRequest).Methods("GET")

	log.Fatal(http.ListenAndServe(":0426", handler))
}

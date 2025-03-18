package routes

import (
	"restMux/handlers"

	"github.com/gorilla/mux"
)

// Methods of practice
func RegisterExampleRoutes(router *mux.Router) {
	prefix := "/api/v0/examples/"
	router.HandleFunc(prefix+"doOnEjemploGet", handlers.DoOnEjemploGet).Methods("GET")
	router.HandleFunc(prefix+"doOnEjemploGetWithParameter/{id:[0-9]+}", handlers.DoOnEjemploGetWithParameter).Methods("GET")
	router.HandleFunc(prefix+"doOnEjemploGetQueryString", handlers.DoOnEjemploGetQueryString).Methods("GET")
	router.HandleFunc(prefix+"doOnEjemploPost", handlers.DoOnEjemploPost).Methods("POST")
	router.HandleFunc(prefix+"doOnEjemploPostWithBody", handlers.DoOnEjemploPostWithBody).Methods("POST")
	router.HandleFunc(prefix+"doOnEjemploPut/{id:[0-9]+}", handlers.DoOnEjemploPut).Methods("PUT")
	router.HandleFunc(prefix+"doOnEjemploDelete/{id:[0-9]+}", handlers.DoOnEjemploDelete).Methods("DELETE")
}

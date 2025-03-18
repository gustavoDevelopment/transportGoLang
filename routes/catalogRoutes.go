package routes

import (
	"restMux/handlers/catalogs"

	"github.com/gorilla/mux"
)

// Methods of catalogs
func RegisterCatalogRoutes(router *mux.Router) {
	prefix := "/api/v0/catalogs"
	router.HandleFunc(prefix, catalogs.HandleCatalogRequest).Methods("GET")
}

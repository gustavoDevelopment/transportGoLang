package routes

import "github.com/gorilla/mux"

// SetupRoutes configura todas las rutas
func SetupRoutes(router *mux.Router) {
	RegisterCatalogRoutes(router)
	RegisterExampleRoutes(router)
}

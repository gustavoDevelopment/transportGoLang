package main

import (
	"log"
	"net/http"
	"restMux/logger"
	"restMux/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	mux := mux.NewRouter()
	routes.SetupRoutes(mux)
	handler := cors.AllowAll().Handler(mux)
	logger.Log.Info("Servidor corriendo en el puerto 0426")
	log.Fatal(http.ListenAndServe(":0426", handler))
}

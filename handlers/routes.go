package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jdetle/captable-backend/captable"
	"github.com/jdetle/captable-backend/config"
)

func SetupRoutes(cfg *config.Config, router *mux.Router, capTable *captable.CapTable) {
	fs := http.FileServer(http.Dir("./swaggerui"))
	router.PathPrefix("/").Handler(http.StripPrefix("/swaggerui/", fs))

	router.HandleFunc("/captable", CreateCTHandler(cfg, capTable)).Methods("POST")
	router.HandleFunc("/captable/{id}", GetCTHandler(cfg, capTable)).Methods("GET")
	router.HandleFunc("/captable/{id}", UpdateCTHandler(cfg, capTable)).Methods("PUT")
	router.HandleFunc("/captable/{id}", DeleteCTHandler(cfg, capTable)).Methods("DELETE")

}

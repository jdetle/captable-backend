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

	router.HandleFunc("captable/", CreateCTHandler(cfg, capTable)).Methods("POST")
	router.HandleFunc("captable/{id}", GetCTHandler(cfg, capTable)).Methods("GET")
	router.HandleFunc("captable/{id}", UpdateCTHandler(cfg, capTable)).Methods("PUT")
	router.HandleFunc("captable/{id}", DeleteCTHandler(cfg, capTable)).Methods("DELETE")

	router.HandleFunc("shareholders/", CreateShareholderHandler(cfg, capTable)).Methods("POST")
	router.HandleFunc("shareholders/{id}", GetShareholderHandler(cfg, capTable)).Methods("GET")
	router.HandleFunc("shareholders/{id}", UpdateShareholderHandler(cfg, capTable)).Methods("PUT")
	router.HandleFunc("shareholders/{id}", DeleteShareholderHandler(cfg, capTable)).Methods("DELETE")

	router.HandleFunc("chunks/", CreateChunkHandler(cfg, capTable)).Methods("POST")
	router.HandleFunc("chunks/{id}", GetChunkHandler(cfg, capTable)).Methods("GET")
	router.HandleFunc("chunks/{id}", UpdateChunkHandler(cfg, capTable)).Methods("PUT")
	router.HandleFunc("chunks/{id}", DeleteChunkHandler(cfg, capTable)).Methods("DELETE")
}

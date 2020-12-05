// Package classification CapTable API.
//
// The Captable API provides a REST interface to a fictional capital table
//
// Terms Of Service:
//
//     Schemes: http
//     Host: localhost:8000
//     Version: 1.0.0
//     Contact: jdetle@gmail.com.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

// Swagger docs are generated using https://github.com/go-swagger/go-swagger
// install that, and then from the root of the repository run:
//   go generate
//go:generate swagger generate spec -o swaggerui/swagger.json

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jdetle/captable-backend/captable"
	"github.com/jdetle/captable-backend/config"
	"github.com/jdetle/captable-backend/dal/store"
	"github.com/jdetle/captable-backend/handlers"
	"github.com/namsral/flag"

	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		err error
	)
	postgresConnURL := flag.String("postgres-conn-url", "postgres://postgres:postgres@localhost/captable",
		"Postgres connection URL") // POSTGRES_CONN_URL

	cfg := &config.Config{
		PostgresConnURL: *postgresConnURL,
	}
	ctx := context.Background()
	var dal *store.Store
	for i := 1; i <= 10; i++ {
		dal, err = store.NewPostgres(ctx, cfg.PostgresConnURL)
		if err == nil {
			log.Info("connected to postgres")
			break
		}
		log.Errorf("unable to initiate postgres conn attempt %d: %#v", i, err)
		time.Sleep(time.Second)
	}
	if err != nil {
		log.Fatalf("unable to initiate postgres conn to %s: %#v", cfg.PostgresConnURL, err)
	}
	captable, err := captable.New(cfg, dal)
	router := mux.NewRouter()
	handlers.SetupRoutes(cfg, router, captable)
	log.Info("api listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

/*
Package config contains the structs and helpers for managing the configuration.
*/
package config

// Config represents variational config throughout deployments of the api
type Config struct {
	PostgresConnURL string
}

// Package captable handles the business logic between a request and response
package captable

import (
	"github.com/jdetle/captable-backend/config"
	"github.com/jdetle/captable-backend/dal/store"
)

type CapTable struct {
	DAL store.DAL
}

// New returns a new CapTable struct.
func New(cfg *config.Config, dal store.DAL) (*CapTable, error) {
	return &CapTable{
		DAL: dal,
	}, nil
}

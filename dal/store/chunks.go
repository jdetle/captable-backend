package store

import (
	"context"

	"github.com/jdetle/captable-backend/dal/model"
)

const createChunkSQL = `
INSERT INTO chunks (

) VALUES (
	$1,
) RETURNING id
`

const readChunksSQL = `

`

const updateChunksSQL = `

`

// CreateOwnershipChunk attaches a new ownership chunk to a shareholder and a cap table.
func (s *Store) CreateOwnershipChunk(ctx context.Context, oc *model.CreateOwnershipChunk) (*model.OwnershipChunk, error) {
	return nil, nil
}

// UpdateOwnershipChunk modifies an existing ownership chunk attached to a cap table and shareholder.
func (s *Store) UpdateOwnershipChunk(ctx context.Context, oc *model.UpdateOwnershipChunk) (*model.OwnershipChunk, error) {
	return nil, nil
}

// ReadOwnershipChunk reads a chunk from the db
func (s *Store) ReadOwnershipChunk(ctx context.Context, shid int, ctid int) (*model.OwnershipChunk, error) {
	return nil, nil
}

// DeleteOwnershipChunk deletes an ownership chunk from the db.
func (s *Store) DeleteOwnershipChunk(ctx context.Context, shid int, ctid int) error {
	return nil
}

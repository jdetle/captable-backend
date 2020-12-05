package store

import (
	"context"
	"fmt"

	"github.com/jdetle/captable-backend/dal/model"
	log "github.com/sirupsen/logrus"
)

const createChunkSQL = `
INSERT INTO ownership_chunks (
	shares_owned,
    share_price,
    captable_id,
    shareholder_id,
) VALUES (
	$1,
	$2,
	$3,
	$4
) RETURNING id
`

const readChunkSQL = `
SELECT 
	id,
	shares_owned,
	share_price,
	captable_id,
	shareholder_id
FROM ownership_chunks
`

const updateChunkSQL = `
UPDATE ownership_chunks 
SET 
	shares_owned=$1,
	share_price=$2,
	captable_id=$3,
	shareholder_id=$4
WHERE id=$5
`

// CreateOwnershipChunk attaches a new ownership chunk to a shareholder and a cap table.
func (s *Store) CreateOwnershipChunk(ctx context.Context, oc *model.CreateOwnershipChunk) (*model.OwnershipChunk, error) {
	log.Debugf("CREATE OWNERSHIP CHUNK: %#v", oc)

	var id int
	err := s.Conn.QueryRowContext(ctx, createChunkSQL,
		oc.SharesOwned,
		oc.SharePrice,
		oc.CapTableID,
		oc.ShareholderID,
	).Scan(&id)
	if err != nil {
		return nil, err
	}
	resp := &model.OwnershipChunk{
		ID:                   int(id),
		CreateOwnershipChunk: *oc,
	}
	return resp, err
}

// UpdateOwnershipChunk modifies an existing ownership chunk attached to a cap table and shareholder.
func (s *Store) UpdateOwnershipChunk(ctx context.Context, oc *model.UpdateOwnershipChunk) (*model.OwnershipChunk, error) {
	log.Debugf("UPDATE OWNERSHIP CHUNK: %#v", oc)
	result, err := s.Conn.ExecContext(ctx, updateChunkSQL,
		oc.SharesOwned,
		oc.SharePrice,
		oc.CapTableID,
		oc.ShareholderID,
	)
	if err != nil {
		return nil, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("result.RowsAffected: %w", err)
	}
	return s.ReadOwnershipChunk(ctx, oc.ID)
}

// ReadOwnershipChunk reads a chunk from the db
func (s *Store) ReadOwnershipChunk(ctx context.Context, id int) (*model.OwnershipChunk, error) {

	oc := &model.OwnershipChunk{}
	q := readChunkSQL + `WHERE id=$1`
	err := s.Conn.QueryRowContext(ctx, q, id).Scan(
		&oc.SharesOwned,
		&oc.SharePrice,
		&oc.ShareholderID,
		&oc.CapTableID,
	)
	if err != nil {
		return nil, err
	}

	log.Debugf("READ OWNERSHIP CHUNK: %#v", oc)
	return oc, nil
}

// DeleteOwnershipChunk deletes an ownership chunk from the db.
func (s *Store) DeleteOwnershipChunk(ctx context.Context, id int) error {
	// check if it exists before deleting so we can properly 404
	_, err := s.ReadOwnershipChunk(ctx, id)
	if err != nil {
		return err
	}
	result, err := s.Conn.ExecContext(ctx, "DELETE FROM ownership_chunks where id=$1", id)
	log.Debugf("RESULT: %#v", result)
	return err
}

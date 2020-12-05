// Package store contains database access logic.
package store

import (
	"context"
	"fmt"
	"time"

	"github.com/jdetle/captable-backend/dal/model"

	log "github.com/sirupsen/logrus"
)

const createShareholderSQL = `
INSERT INTO shareholders (
	ownership_chunk_ids,
    first_name,
    last_name,
    email
) VALUES (
	$1,
	$2,
	$3,
	$4
) RETURNING id
`

const readShareholderSQL = `
SELECT 
	id,
	ownership_chunk_ids,
	first_name,
	last_name,
	email
FROM shareholders
`

const updateShareholderSQL = `
UPDATE shareholders 
SET 
	ownership_chunk_ids=$1,
	first_name=$2,
	last_name=$3,
	email=$4
WHERE id=$5
`

// CreateShareholder creates the initial shareholder data without ownership chunks.
func (s *Store) CreateShareholder(ctx context.Context, ct *model.CreateShareholderRequest) (*model.Shareholder, error) {
	log.Debugf("CREATE SHAREHOLDER: %#v", ct)

	var id int
	err := s.Conn.QueryRowContext(ctx, createCTSQL).Scan(&id)
	if err != nil {
		return nil, err
	}
	updatedAt := time.Now()
	resp := &model.Shareholder{
		ID:        int(id),
		UpdatedAt: &updatedAt,
	}
	return resp, err
}

// UpdateShareholder is meant for adjusting shareholder info.
func (s *Store) UpdateShareholder(ctx context.Context, sh *model.UpdateShareholderRequest) (*model.Shareholder, error) {
	result, err := s.Conn.ExecContext(ctx, updateShareholderSQL, sh.ID)
	if err != nil {
		return nil, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("result.RowsAffected: %w", err)
	}
	return s.ReadShareholder(ctx, sh.ID)
}

// ReadShareholder will give you current share price, total number of shares, and sharehoders.
func (s *Store) ReadShareholder(ctx context.Context, id int) (*model.Shareholder, error) {
	sh := &model.Shareholder{}
	q := readChunkSQL + `WHERE id=$1`
	err := s.Conn.QueryRowContext(ctx, q, id).Scan(
		&sh.OwnershipChunks,
		&sh.FirstName,
		&sh.LastName,
		&sh.Email,
	)
	if err != nil {
		return nil, err
	}

	log.Debugf("READ OWNERSHIP CHUNK: %#v", sh)
	return sh, nil
}

// DeleteShareholder deletes a cap table from the database if the id exists in the db.
func (s *Store) DeleteShareholder(ctx context.Context, id int) error {
	// check if it exists before deleting so we can properly 404
	_, err := s.ReadShareholder(ctx, id)
	if err != nil {
		return err
	}
	result, err := s.Conn.ExecContext(ctx, "DELETE FROM shareholders where id=$1", id)
	log.Debugf("DELETE SHAREHOLDER RESULT: %#v", result)
	return err
}

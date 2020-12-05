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

) VALUES (
	$1,
) RETURNING id
`

const readShareholderSQL = `

`

const updateShareholderSQL = `

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
func (s *Store) ReadShareholder(ctx context.Context, shid int) (*model.Shareholder, error) {
	return nil, nil
}

// DeleteShareholder deletes a cap table from the database if the id exists in the db.
func (s *Store) DeleteShareholder(ctx context.Context, shid int) error {
	// check if it exists before deleting so we can properly 404
	_, err := s.ReadShareholder(ctx, shid)
	if err != nil {
		return err
	}
	result, err := s.Conn.ExecContext(ctx, "DELETE FROM shareholders where id=$1", shid)
	log.Debugf("RESULT: %#v", result)
	return err
}

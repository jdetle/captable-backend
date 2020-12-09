// Package store contains database access logic.
package store

import (
	"context"
	"fmt"
	"time"

	"github.com/jdetle/captable-backend/dal/model"

	log "github.com/sirupsen/logrus"
)

const createCTSQL = `
INSERT INTO cap_tables (
	total_shares, 
	company_name,
	shareholder_ids,
    share_price,
) VALUES (
	$1,
	$2,
	$3,
	$4,
) RETURNING id
`

const readCTSQL = `	
SELECT 
	id,
	total_shares,
	company_name,
	shareholder_ids,
	share_price
FROM cap_tables
`

const updateCTSQL = `
UPDATE cap_tables
SET
	total_shares=$1,
	company_name=$2,
	shareholder_ids=$3,
	share_price=$4
WHERE id=$5
`

// CreateCT creates the initial cap table in the database with zero shareholders.
func (s *Store) CreateCT(ctx context.Context, ct *model.CreateCapTableRequestWithShareholders) (*model.CapTable, error) {
	log.Debugf("CREATE CAP TABLE: %#v", ct)

	var id int
	var shareholders []model.Shareholder
	err := s.Conn.QueryRowContext(ctx, createCTSQL,
		ct.CompanyName,
		ct.SharePrice,
		ct.TotalShares,
	).Scan(&id)
	tx, err := s.Conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	for _, item := range *ct.Shareholders {
		resp, err := s.CreateShareholder(ctx, &item)
		shareholders = append(shareholders, *resp)
		if err != nil {
			log.Errorf("%#v", err)
			tx.Rollback()
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	updatedAt := time.Now()
	createdAt := time.Now()
	resp := &model.CapTable{
		ID:           int(id),
		UpdatedAt:    &updatedAt,
		CreatedAt:    &createdAt,
		Shareholders: &shareholders,
		CreateCapTableRequest: model.CreateCapTableRequest{
			TotalShares: ct.TotalShares,
			SharePrice:  ct.SharePrice,
			CompanyName: ct.CompanyName,
		},
	}
	return resp, err
}

// UpdateCT is meant for adjusting share prices or the current number of shares.
func (s *Store) UpdateCT(ctx context.Context, ct *model.UpdateCapTableRequest) (*model.CapTable, error) {
	result, err := s.Conn.ExecContext(ctx, updateCTSQL, ct.ID,
		ct.CompanyName,
		ct.SharePrice,
		ct.TotalShares,
		ct.CreatedAt,
		ct.UpdatedAt,
		ct.Shareholders)

	if err != nil {
		return nil, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("result.RowsAffected: %w", err)
	}
	return s.ReadCT(ctx, ct.ID)

}

// ReadCT will give you current share price, total number of shares, and sharehoders.
func (s *Store) ReadCT(ctx context.Context, id int) (*model.CapTable, error) {
	ct := &model.CapTable{}
	q := readCTSQL + `WHERE id=$1`
	err := s.Conn.QueryRowContext(ctx, q, id).
		Scan(
			&ct.ID,
			&ct.CompanyName,
			&ct.SharePrice,
			&ct.TotalShares,
			&ct.CreatedAt,
			&ct.UpdatedAt,
			&ct.Shareholders,
		)
	if err != nil {
		return nil, err
	}
	return ct, nil
}

// DeleteCT deletes a cap table from the database if the id exists in the db.
func (s *Store) DeleteCT(ctx context.Context, id int) error {
	// check if it exists before deleting so we can properly 404
	_, err := s.ReadCT(ctx, id)
	if err != nil {
		return err
	}
	result, err := s.Conn.ExecContext(ctx, "DELETE FROM cap_tables where id=$1", id)
	log.Debugf("RESULT: %#v", result)
	return err
}

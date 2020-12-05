/*
Package store contains structs and methods used to interact with the database.
*/
package store

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jdetle/captable-backend/dal/model"
)

// ErrNoUpdatedRows gives a standard error for a request leading to no changes to the database.
var ErrNoUpdatedRows = errors.New("no rows were updated")

// DAL allows us to mock out our data access layer for testing.
type DAL interface {
	CreateCT(ctx context.Context, ct *model.CreateCapTableRequest) (*model.CapTable, error)
	UpdateCT(ctx context.Context, ct *model.UpdateCapTableRequest) (*model.CapTable, error)
	ReadCT(ctx context.Context, id int) (*model.CapTable, error)
	DeleteCT(ctx context.Context, id int) error

	CreateShareholder(ctx context.Context, ct *model.CreateShareholderRequest) (*model.Shareholder, error)
	UpdateShareholder(ctx context.Context, sh *model.UpdateShareholderRequest) (*model.Shareholder, error)
	ReadShareholder(ctx context.Context, shID int) (*model.Shareholder, error)

	CreateOwnershipChunk(ctx context.Context, chunk *model.CreateOwnershipChunk) (*model.OwnershipChunk, error)
	UpdateOwnershipChunk(ctx context.Context, chunk *model.UpdateOwnershipChunk) (*model.OwnershipChunk, error)

	ReadOwnershipChunk(ctx context.Context, shID int, ctID int) (*model.OwnershipChunk, error)
	DeleteOwnershipChunk(ctx context.Context, shID int, ctID int) error
}

// Store lets us connect to the database.
type Store struct {
	Conn *sql.DB
}

// NewPostgres creates a store.
func NewPostgres(ctx context.Context, dbURL string) (*Store, error) {
	conn, err := sql.Open("pgx", dbURL)
	if err != nil {
		return nil, err
	}
	return &Store{
		Conn: conn,
	}, conn.Ping()
}

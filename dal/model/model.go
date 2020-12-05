/*
Package model contains structs for use in the store and handlers
*/
package model

import (
	"regexp"
	"time"
)

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// CreateCapTableRequest has all the data to create a new cap table.
type CreateCapTableRequest struct {
	TotalShares int     `json:"totalShares"`
	SharePrice  float64 `json:"sharePrice"`
	CompanyName string  `json:"companyName"`
}

// CapTable after it has been created.
type CapTable struct {
	ID           int            `json:"id"`
	UpdatedAt    *time.Time     `json:"updatedAt,omitEmpty"`
	CreatedAt    *time.Time     `json:"createdAt,omitEmpty"`
	Shareholders *[]Shareholder `json:"shareholders,omitEmpty"`
	CreateCapTableRequest
}

// UpdateCapTableRequest is the payload required to update a captable
type UpdateCapTableRequest CapTable

// ShareholderCreate contains all the info to initialize shareholder data.
type CreateShareholderRequest struct {
	Email                string                  `json:"email"`
	CapTableID           int                     `json:"capTableId"`
	FirstName            string                  `json:"firstName"`
	LastName             string                  `json:"lastName"`
	OwnershipChunkCreate *[]CreateOwnershipChunk `json:"ownershipChunks,omitEmpty"`
}

type AddShareholdersRequest struct {
	CapTableID   int                         `json:"capTableId"`
	Shareholders *[]CreateShareholderRequest `json:"shareholders,omitEmpty"`
}

type CreateOwnershipChunk struct {
	SharesOwned int
	OwnerID     int
	OwnerName   *string

	ShareholderID int `json:"shareholderId"`
	CapTableID    int `json:"capTableId"`
}

// OwnershipChunk is the representation of a discrete award of company shares at a given fundraising round.
type OwnershipChunk struct {
	ID int `json:"id"`
	CreateOwnershipChunk
}

type UpdateOwnershipChunk struct {
	ShareholderID int `json:"shareholderId"`
	CapTableID    int `json:"capTableId"`
}

type DeleteOwnershipChunk struct {
	ShareholderID int `json:"shareholderId"`
	CapTableID    int `json:"capTableId"`
}

// Shareholder is shareholder data after its been created
type Shareholder struct {
	ID        int        `json:"id"`
	UpdatedAt *time.Time `json:"updatedAt,omitEmpty"`
	CreatedAt *time.Time `json:"createdAt,omitEmpty"`
	CreateShareholderRequest
}

// UpdateShareholderRequest is the entire shareholder object.
type UpdateShareholderRequest Shareholder

/*
Package model contains structs for use in the store and handlers
*/
package model

import (
	"fmt"
	"strings"
	"time"
)

type CreateOwnershipChunkRequest struct {
	SharesOwned int     `json:"sharesOwned"`
	SharePrice  float64 `json:"sharePrice"`
}

type CreateOwnershipChunk struct {
	CreateOwnershipChunkRequest
	ShareholderID int `json:"shareholderId"`
	CapTableID    int `json:"capTableId"`
}

// OwnershipChunk is the representation of a discrete award of company shares at a given fundraising round.
type OwnershipChunk struct {
	ID        int        `json:"id"`
	UpdatedAt *time.Time `json:"updatedAt,omitEmpty"`
	CreatedAt *time.Time `json:"createdAt,omitEmpty"`
	CreateOwnershipChunk
}

type UpdateOwnershipChunk OwnershipChunk

func (c *CreateOwnershipChunk) Validate() error {
	errMsgs := c.validate()
	if len(errMsgs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errMsgs, ", "))
	}
	return nil
}

func (c *CreateOwnershipChunk) validate() []string {
	errMsgs := []string{}
	if c.CapTableID < 1 {
		errMsgs = append(errMsgs, "captable id must be present")
	}
	if c.ShareholderID < 1 {
		errMsgs = append(errMsgs, "shareholder id must be present")
	}
	if c.SharePrice < 0 {
		errMsgs = append(errMsgs, "share price cannot be less than 0")
	}
	if c.SharesOwned < 1 {
		errMsgs = append(errMsgs, "shares owned must exist")
	}
	return errMsgs
}

func (u *UpdateOwnershipChunk) Validate() error {
	errMsgs := u.validate()
	if len(errMsgs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errMsgs, ", "))
	}
	return nil
}

func (u *UpdateOwnershipChunk) validate() []string {
	errMsgs := u.CreateOwnershipChunk.validate()
	if u.ID < 1 {
		errMsgs = append(errMsgs, "ID must be greater than 0")
	}
	return errMsgs
}

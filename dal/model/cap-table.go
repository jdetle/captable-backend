/*
Package model contains structs for use in the store and handlers
*/
package model

import (
	"fmt"
	"strings"
	"time"
)

// CreateCapTableRequest has all the data to create a new cap table.
type CreateCapTableRequest struct {
	TotalShares int     `json:"totalShares"`
	SharePrice  float64 `json:"sharePrice"`
	CompanyName string  `json:"companyName"`
}

type CreateCapTableRequestWithShareholders struct {
	Shareholders *[]CreateShareholderRequest `json:"shareholders"`
	CreateCapTableRequest
}

// CapTable after it has been created.
type CapTable struct {
	ID           int            `json:"id"`
	UpdatedAt    *time.Time     `json:"updatedAt,omitEmpty"`
	CreatedAt    *time.Time     `json:"createdAt,omitEmpty"`
	Shareholders *[]Shareholder `json`
	CreateCapTableRequest
}

// UpdateCapTableRequest is the payload required to update a captable
type UpdateCapTableRequest CapTable

func (c *CreateCapTableRequest) Validate() error {
	errMsgs := c.validate()
	if len(errMsgs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errMsgs, ", "))
	}
	return nil
}

func (c *CreateCapTableRequest) validate() []string {
	errMsgs := []string{}

	if c.TotalShares < 1 {
		errMsgs = append(errMsgs, "shares must exist to capitalize")
	}
	if c.CompanyName == "" {
		errMsgs = append(errMsgs, "companyname must be present")
	}
	if c.SharePrice < 0 {
		errMsgs = append(errMsgs, "sharePrice must be a positive number")
	}
	return errMsgs
}

func (u *UpdateCapTableRequest) Validate() error {
	errMsgs := u.validate()
	if len(errMsgs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errMsgs, ", "))
	}
	return nil
}

func (u *UpdateCapTableRequest) validate() []string {
	errMsgs := u.CreateCapTableRequest.validate()
	if u.ID < 1 {
		errMsgs = append(errMsgs, "ID must be greater than 0")
	}
	return errMsgs
}

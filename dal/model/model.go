/*
Package model contains structs for use in the store and handlers
*/
package model

import (
	"fmt"
	"regexp"
	"strings"
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

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

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
	Email           string                  `json:"email"`
	CapTableID      int                     `json:"capTableId"`
	FirstName       string                  `json:"firstName"`
	LastName        string                  `json:"lastName"`
	OwnershipChunks *[]CreateOwnershipChunk `json:"ownershipChunks,omitEmpty"`
}

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

// Shareholder is shareholder data after its been created
type Shareholder struct {
	ID        int        `json:"id"`
	UpdatedAt *time.Time `json:"updatedAt,omitEmpty"`
	CreatedAt *time.Time `json:"createdAt,omitEmpty"`
	CreateShareholderRequest
}

// UpdateShareholderRequest is the entire shareholder object.
type UpdateShareholderRequest Shareholder

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

func (c *CreateShareholderRequest) Validate() error {
	errMsgs := c.validate()
	if len(errMsgs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errMsgs, ", "))
	}
	return nil
}

func (c *CreateShareholderRequest) validate() []string {
	errMsgs := []string{}
	if c.FirstName == "" {
		errMsgs = append(errMsgs, "first name must be present")
	}
	if c.LastName == "" {
		errMsgs = append(errMsgs, "last name must be present")
	}
	if !isEmailValid(c.Email) {
		errMsgs = append(errMsgs, "valid email must be supplied")
	}
	return errMsgs
}

func (u *UpdateShareholderRequest) Validate() error {
	errMsgs := u.validate()
	if len(errMsgs) > 0 {
		return fmt.Errorf("validation failed: %s", strings.Join(errMsgs, ", "))
	}
	return nil
}

func (u *UpdateShareholderRequest) validate() []string {
	errMsgs := u.CreateShareholderRequest.validate()
	if u.ID < 1 {
		errMsgs = append(errMsgs, "ID must be greater than 0")
	}
	return errMsgs
}

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

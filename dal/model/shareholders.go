/*
Package model contains structs for use in the store and handlers
*/
package model

import (
	"fmt"
	"strings"
	"time"
)

// ShareholderCreate contains all the info to initialize shareholder data.
type CreateShareholderRequest struct {
	Email           string                  `json:"email"`
	CapTableID      int                     `json:"capTableId"`
	FirstName       string                  `json:"firstName"`
	LastName        string                  `json:"lastName"`
	OwnershipChunks *[]CreateOwnershipChunk `json:"ownershipChunks,omitEmpty"`
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

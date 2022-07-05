package models

import (
	"log"
	"time"

	"github.com/pkg/errors"
)

type Host struct {
	ID        int       `db:"id" json:"-"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Email     string    `db:"email_address" json:"email_address"`
	Phone     string    `db:"phone_number" json:"phone_number"`
	AddressID int       `db:"address_id" json:"-"`
	Address   *Address  `db:"-" belongs_to:"address" json:"address" fk_id:"address_id"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

func (h Host) TableName() string {
	return "hosts"
}

func (h Host) Save() Host {
	log.Printf("Attempting model creation: %v", h)
	vErr, cErr := DB.ValidateAndCreate(&h)
	if vErr.Count() > 0 {
		log.Fatalf("Validation failed for model %v: %s", h, vErr.String())
	}
	if cErr != nil {
		log.Fatalf("Transaction for model failed: %+v", errors.WithStack(cErr).Error())
	}

	return h
}

func (h Host) FullName() string {
	return h.FirstName + " " + h.LastName
}

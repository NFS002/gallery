package models

import (
	"log"
	"time"

	"github.com/pkg/errors"
)

type Artist struct {
	ID        int       `db:"id" json:"-"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Email     string    `db:"email_address" json:"email"`
	Phone     string    `db:"phone_number" json:"phone"`
	AddressID int       `db:"address_id" json:"-"`
	Address   *Address  `db:"-" belongs_to:"address" json:"address"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

func (a Artist) TableName() string {
	return "artists"
}

func (a Artist) Save() Artist {
	log.Printf("Attempting model creation: %v", a)
	vErr, cErr := DB.ValidateAndCreate(&a)
	if vErr.Count() > 0 {
		log.Fatalf("Validation failed for model %v: %s", a, vErr.String())
	}
	if cErr != nil {
		log.Fatalf("Transaction for model failed: %+v", errors.WithStack(cErr).Error())
	}

	return a
}

func (a Artist) FullName() string {
	return a.FirstName + " " + a.LastName
}

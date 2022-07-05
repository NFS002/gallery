package models

import (
	"log"
	"time"

	"github.com/pkg/errors"
)

type Address struct {
	ID        int       `db:"id" json:"-"`
	Number    int       `db:"number" json:"number"`
	Address1  string    `db:"address_1" json:"address_1"`
	Address2  string    `db:"address_2" json:"address_2"`
	City      string    `db:"city" json:"city"`
	Postcode  string    `db:"postcode" json:"postcode"`
	Host      *Host     `db:"-" has_one:"host" json:"-" fk_id:"address_id"`
	Artist    *Artist   `db:"-" has_one:"artist" json:"-" fk_id:"address_id"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

func (a Address) TableName() string {
	return "addresses"
}

func (a Address) Save() Address {
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

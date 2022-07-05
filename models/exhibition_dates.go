package models

import (
	"log"
	"time"

	"github.com/pkg/errors"
)

type ExhibitionDate struct {
	ID           int         `db:"id" json:"-"`
	ExhibitionID int         `db:"exhibition_id" json:"-"`
	Exhibition   *Exhibition `db:"-" belongs_to:"exhibition" json:"-"`
	StartTime    time.Time   `db:"start_time" json:"start_time"`
	EndTime      time.Time   `db:"end_time" json:"end_time"`
	CreatedAt    time.Time   `db:"created_at" json:"-"`
	UpdatedAt    time.Time   `db:"updated_at" json:"-"`
}

func (e ExhibitionDate) TableName() string {
	return "exhibition_dates"
}

func (e ExhibitionDate) Save() ExhibitionDate {
	log.Printf("Attempting model creation: %v", e)
	vErr, cErr := DB.ValidateAndCreate(&e)
	if vErr.Count() > 0 {
		log.Fatalf("Validation failed for model %v: %s", e, vErr.String())
	}
	if cErr != nil {
		log.Fatalf("Transaction for model failed: %+v", errors.WithStack(cErr).Error())
	}

	return e
}

func (e ExhibitionDate) ReadableStartDate() string {
	return e.StartTime.Format(time.RFC822)
}

func (e ExhibitionDate) ReadableEndDate() string {
	return e.EndTime.Format(time.RFC822)
}

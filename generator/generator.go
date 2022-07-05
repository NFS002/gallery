package generator

import (
	"gallery/models"
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

func GenerateAddress() models.Address {
	model := models.Address{
		Number:   gofakeit.Number(1, 1000),
		Address1: gofakeit.Street(),
		Address2: gofakeit.Street(),
		City:     gofakeit.City(),
		Postcode: gofakeit.Regex("^[A-Z1-9]{3}\\s[A-Z1-9]{3}$"),
	}
	return model
}

func GenerateArtist(addressId int) models.Artist {
	model := models.Artist{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Phone:     gofakeit.Phone(),
		AddressID: addressId,
	}
	return model
}

func GenerateHost(addressId int) models.Host {
	model := models.Host{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Phone:     gofakeit.Phone(),
		AddressID: addressId,
	}
	return model
}

func GenerateExhibition(artistId int, hostId int) models.Exhibition {
	model := models.Exhibition{
		Title:    gofakeit.HipsterSentence(5),
		HostID:   hostId,
		ArtistID: artistId,
		Notes:    gofakeit.Paragraph(1, 3, 30, ","),
	}
	return model
}

func GenerateExhibitionDate(exhibitionId int) models.ExhibitionDate {
	minStartTime := time.Now()
	maxStartTime := minStartTime.AddDate(0, 5, 0)
	minEndTime := maxStartTime.AddDate(0, 5, 0)
	maxEndTime := minEndTime.AddDate(0, 5, 0)
	model := models.ExhibitionDate{
		ExhibitionID: exhibitionId,
		StartTime:    gofakeit.DateRange(minStartTime, maxStartTime),
		EndTime:      gofakeit.DateRange(minEndTime, maxEndTime),
	}
	return model
}

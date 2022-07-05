package grifts

import (
	gen "gallery/generator"
	. "github.com/markbates/grift/grift"
)

var _ = Namespace("db", func() {

	_ = Desc("seed", "Task Description")
	_ = Add("seed", func(c *Context) error {
		// Add DB seeding stuff here

		for i := 0; i < 10; i++ {
			addr1 := gen.GenerateAddress()
			addr2 := gen.GenerateAddress()
			addr1 = addr1.Save()
			addr2 = addr2.Save()

			artist := gen.GenerateArtist(addr1.ID)
			host := gen.GenerateHost(addr2.ID)
			artist = artist.Save()
			host = host.Save()

			exhibition := gen.GenerateExhibition(artist.ID, host.ID)
			exhibition = exhibition.Save()

			exhibitionDate1 := gen.GenerateExhibitionDate(exhibition.ID)
			exhibitionDate2 := gen.GenerateExhibitionDate(exhibition.ID)

			exhibitionDate1 = exhibitionDate1.Save()
			exhibitionDate2 = exhibitionDate2.Save()
		}

		return nil
	})
})

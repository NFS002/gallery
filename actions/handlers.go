package actions

import (
	"gallery/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"net/http"
)

const (
	ArtistHeadingText = "Gallery artists"
	HostHeadingText   = "Gallery hosts"
)

func HomeHandler(c buffalo.Context) error {
	c.Set("search_menu_partial_path", "search_menu_items.partial.html")
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}

func getAllExhibitions() (error, []models.Exhibition) {
	DB, dbErr := pop.Connect(app.Env)
	if dbErr != nil {
		return dbErr, nil
	}
	var exhibitions []models.Exhibition
	err1 := DB.Eager().All(&exhibitions)
	if err1 != nil {
		return err1, nil
	}
	err2 := models.LoadAllAddresses(exhibitions)
	if err2 != nil {
		return err2, nil
	}
	return nil, exhibitions
}

func GetAllExhibitionsHandler(c buffalo.Context) error {
	err, exhibitions := getAllExhibitions()
	if err != nil {
		return c.Error(500, err)
	} else {
		c.Set("custom_js_paths", []string{"search.js"})
		c.Set("custom_css_paths", []string{"search.css"})
		c.Set("exhibitions", exhibitions)
	}
	return c.Render(http.StatusOK, r.HTML("home/search.plush.html"))
}

func ArtistsLoginHandler(c buffalo.Context) error {
	c.Set("search_menu_partial_path", "search_menu_items.partial.html")
	c.Set("heading_text", ArtistHeadingText)

	return c.Render(http.StatusOK, r.HTML("home/artists_login.plush.html"))
}

func HostsLoginHandler(c buffalo.Context) error {
	c.Set("search_menu_partial_path", "search_menu_items.partial.html")
	c.Set("heading_text", HostHeadingText)

	return c.Render(http.StatusOK, r.HTML("home/hosts_login.plush.html"))
}

func ArtistsCreatePostHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.String(""))
}

func HostsCreatePostHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.String(""))
}

func LoginPostHandler(c buffalo.Context) error {

	return c.Render(http.StatusOK, r.String(""))
}

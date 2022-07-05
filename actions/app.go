package actions

import (
	obs "gallery/observability"
	"gallery/public"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	csrf "github.com/gobuffalo/mw-csrf"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")

var (
	app *buffalo.App
)

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.

func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Host:        "localhost",
			Env:         ENV,
			SessionName: "_gallery_session",
			Logger:      obs.JSONLogger(logrus.InfoLevel),
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		app.GET("/", HomeHandler)

		app.GET("/exhibitions", GetAllExhibitionsHandler).Name("allExhibitions")

		artistsSubApp := app.Group("/a")
		configureArtistRoutes(artistsSubApp)

		hostsSubApp := app.Group("/h")
		configureHostRoutes(hostsSubApp)

		app.ServeFiles("/", http.FS(public.FS())) // serve files from the public directory
	}

	return app
}

func configureArtistRoutes(subApp *buffalo.App) {
	subApp.GET("/", ArtistsLoginHandler).Name("artistsLogin")
	subApp.POST("/create", ArtistsCreatePostHandler).Name("artistsCreatePost")
	subApp.POST("/login", LoginPostHandler).Name("artistsLoginPost")
}

func configureHostRoutes(subApp *buffalo.App) {
	subApp.GET("/", HostsLoginHandler).Name("hostsLogin")
	subApp.POST("/create", HostsCreatePostHandler).Name("hostsCreatePost")
	subApp.POST("/login", LoginPostHandler).Name("hostsLoginPost")
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}

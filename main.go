package main

import (
	"path"

	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/blitzprog/blitzprog.org/components/css"
	"github.com/blitzprog/blitzprog.org/components/js"
	"github.com/blitzprog/blitzprog.org/layout"
	"github.com/blitzprog/blitzprog.org/pages/contact"
	"github.com/blitzprog/blitzprog.org/pages/home"
	"github.com/blitzprog/blitzprog.org/pages/skills"
	"github.com/blitzprog/blitzprog.org/pages/websites"
)

func main() {
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {
	l := layout.New(app)
	l.Render = fullpage.Render

	l.Page("/", home.Get)
	l.Page("/skills", skills.Get)
	l.Page("/websites", websites.Get)
	l.Page("/contact", contact.Get)

	// Certificate
	app.Security.Load("security/server.crt", "security/server.key")

	// Allow Cloudflare CDN as CSS source
	app.ContentSecurityPolicy.Set("style-src", "'self' https://cdnjs.cloudflare.com")

	// Script bundle
	scriptBundle := js.Bundle()

	app.Get("/scripts", func(ctx *aero.Context) string {
		return ctx.JavaScript(scriptBundle)
	})

	// CSS bundle
	cssBundle := css.Bundle()

	app.Get("/styles", func(ctx *aero.Context) string {
		return ctx.CSS(cssBundle)
	})

	// Static files
	app.Get("/images/*file", func(ctx *aero.Context) string {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return ctx.File(path.Join("images", ctx.Get("file")))
	})

	// Manifest
	app.Get("/manifest.json", func(ctx *aero.Context) string {
		return ctx.JSON(app.Config.Manifest)
	})

	return app
}

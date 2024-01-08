package web

import (
	"ats3fx/web/demo"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewRouter(app *fiber.App) fiber.Router {
	group := app.Group("/")
	return group
}
func NewWebModule() fx.Option {
	return fx.Module(
		"web-module",
		fx.Provide(
			fx.Annotate(NewRouter, fx.As(new(fiber.Router))),
		),
		fx.Invoke(demo.InitDemoRoute),
	)
}

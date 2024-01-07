package api

import (
  "ats3fx/api/demo"
  "github.com/gofiber/fiber/v2"
  "go.uber.org/fx"
)

func NewRouter(app *fiber.App) fiber.Router {
  group := app.Group("/")
  return group
}
func NewRouterModule() fx.Option {
  return fx.Module(
    "app-router",
    fx.Provide(
      fx.Annotate(NewRouter, fx.As(new(fiber.Router))),
    ),
    fx.Invoke(demo.InitDemoRoute),
  )
}

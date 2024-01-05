package router

import (
  "github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) fiber.Router {
  group := app.Group("/")
  InitDemoRoute(group)
  return group
}

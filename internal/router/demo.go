package router

import (
  "ats3/internal/controller"
  "github.com/gofiber/fiber/v2"
)

func InitDemoRoute(group fiber.Router) {
  group.Get("/", controller.DemoController.Test)
  group.Post("/api/a", controller.DemoController.Test)
}

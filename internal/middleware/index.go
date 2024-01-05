package middleware

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
  "github.com/gofiber/fiber/v2/middleware/recover"
)

func InitMiddleware(a *fiber.App) {
  a.Use(
    recover.New(),
    // Add CORS to each route.
    cors.New(),
    // Add simple logger.
    DefaultLogger(),
  )
  NotFoundRoute(a)
}

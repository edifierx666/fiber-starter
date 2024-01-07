package middleware

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
  "github.com/gofiber/fiber/v2/middleware/recover"
)

func InitMiddleware(a *fiber.App) {
  // cfg := swagger.Config{
  //   BasePath: "/",
  //   FilePath: "./docs/swagger.json",
  //   Path:     "swagger",
  //   Title:    a.Config().AppName,
  // }

  a.Use(
    recover.New(),
    // Add CORS to each route.
    cors.New(),
    // Add simple logger.
    // swagger.New(cfg),
    DefaultLogger(),
  )
}

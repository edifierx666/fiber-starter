package core

import (
  "context"
  "fmt"

  "ats3fx/internal/config"
  "ats3fx/internal/core/middleware"
  "ats3fx/internal/g"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
  recover2 "github.com/gofiber/fiber/v2/middleware/recover"
  "go.uber.org/fx"
)

func NewFiberApp(lf fx.Lifecycle, cfg fiber.Config, c *config.Server) *fiber.App {
  app := fiber.New(cfg)

  app.Use(
    recover2.New(),
    // Add CORS to each route.
    cors.New(),
    // Add simple logger.
    middleware.DefaultLogger(),
  )

  lf.Append(
    fx.Hook{
      OnStart: func(c context.Context) error {
        go func() {
          err := app.Listen(fmt.Sprintf(":%v", g.CONFIG.System.Port))
          if err != nil {

          }
        }()
        return nil
      },
      OnStop: func(c context.Context) error {
        return app.Shutdown()
      },
    },
  )

  return app
}

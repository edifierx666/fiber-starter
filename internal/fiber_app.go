package internal

import (
	"context"
	"fmt"

	"ats3fx/internal/config"
	"ats3fx/internal/g"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewFiberApp(lf fx.Lifecycle, cfg fiber.Config, c *config.Server) *fiber.App {
	app := fiber.New(cfg)

	lf.Append(
		fx.Hook{
			OnStart: func(c context.Context) error {
				go func() {
					err := app.Listen(fmt.Sprintf(":%v", g.CONFIG.System.Port))
					if err != nil {
						panic(err)
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

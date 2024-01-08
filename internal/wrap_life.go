package internal

import (
	"ats3fx/internal/config"
	"dario.cat/mergo"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type WrapLife struct {
}

func (w WrapLife) Before() fx.Option {
	return fx.Module(
		"before-module",
		fx.Supply(
			fiber.Config{
				AppName:           "edx",
				Immutable:         true,
				EnablePrintRoutes: true,
			},
		),
		fx.Invoke(
			func(c *config.Server) *config.Server {
				mergo.Merge(
					&c.Zap, &config.Zap{
						Level:         "info",
						Prefix:        "",
						Format:        "console",
						Director:      "",
						EncodeLevel:   "LowercaseColorLevelEncoder",
						StacktraceKey: "stacktrace",
						ShowLine:      true,
						LogInConsole:  true,
					},
				)
				mergo.Merge(
					&c.System, &config.System{
						AppName: "Edx",
						Port:    "9888",
					},
				)
				return c
			},
		),
	)
}

func (w WrapLife) After() fx.Option {
	return fx.Module(
		"after-module",
	)
}

var Wrapper = NewWrapLife()

func NewWrapLife() *WrapLife {
	return &WrapLife{}
}

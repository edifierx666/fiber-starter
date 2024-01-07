package main

import (
  "ats3fx/api"
  "ats3fx/internal/core"
  "ats3fx/internal/g"
  "ats3fx/util"
  "github.com/gofiber/fiber/v2"
  "go.uber.org/fx"
)

func main() {
  fx.New(
    fx.Supply(
      fiber.Config{
        AppName:           "edx",
        Immutable:         true,
        EnablePrintRoutes: true,
      },
    ),
    fx.Provide(
      fx.Annotate(
        util.MemStorage, fx.As(new(fiber.Storage)),
      ),
    ),
    fx.Populate(&g.MemStorage),
    core.NewCoreModule(),
    api.NewRouterModule(),
  ).Run()
}

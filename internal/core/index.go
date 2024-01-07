package core

import (
  "ats3fx/internal/config"
  "ats3fx/internal/g"
  "dario.cat/mergo"
  "go.uber.org/fx"
  "go.uber.org/zap"
)

func NewCoreModule() fx.Option {
  v := Viper()
  return fx.Module(
    "core-module",
    fx.Supply(v),
    fx.Supply(g.CONFIG),
    fx.Provide(
      func(c *config.Server) (logger *zap.Logger) {
        p := &Params{}
        _ = mergo.Merge(p, c.Zap)
        return Zap(p)
      },
    ),
    fx.Populate(&g.Log),
    fx.Decorate(
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
    fx.Populate(&g.APP),
    fx.Provide(NewFiberApp),
    fx.Invoke(),
  )
}

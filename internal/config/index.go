package config

import (
  "go.uber.org/fx"
)

func NewConfigModule() fx.Option {
  return fx.Module(
    "config-module",

  )
}

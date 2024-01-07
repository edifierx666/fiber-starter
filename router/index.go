package router

import "go.uber.org/fx"

func NewRouterModule() fx.Option {

  return fx.Module(
    "router-module",

  )
}

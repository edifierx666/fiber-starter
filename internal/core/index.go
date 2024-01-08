package core

import (
	"ats3fx/internal"
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
		fx.Supply(v, g.CONFIG),
		fx.Provide(
			func(c *config.Server) (logger *zap.Logger) {
				p := &Params{}
				_ = mergo.Merge(p, c.Zap)
				return Zap(p)
			},
			internal.NewFiberApp,
		),
		fx.Populate(&g.Log, &g.APP),
		fx.Invoke(),
	)
}

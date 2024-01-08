package main

import (
	"ats3fx/internal"
	"ats3fx/internal/core"
	"ats3fx/web"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		internal.Wrapper.Before(),
		core.NewCoreModule(),
		web.NewWebModule(),
		internal.Wrapper.After(),
	).Run()
}

package main

import (
  "embed"
  "fmt"

  "ats3/internal/config"
  _ "ats3/internal/config"
  "ats3/internal/core"
  "ats3/internal/g"
  "ats3/internal/middleware"
  "ats3/internal/router"
  "dario.cat/mergo"
  "go.uber.org/zap"
)

//go:embed resource/*
var resources embed.FS

func main() {
  mergo.Merge(&g.CONFIG.Zap, &config.Zap{
    Level:         "info",
    Prefix:        "",
    Format:        "console",
    Director:      "",
    EncodeLevel:   "LowercaseColorLevelEncoder",
    StacktraceKey: "stacktrace",
    ShowLine:      true,
    LogInConsole:  true,
  })
  mergo.Merge(&g.CONFIG.System, &config.System{
    AppName: "Edx",
    Port:    "9888",
  })
  _, log := core.Init()
  g.Log = log
  zap.ReplaceGlobals(g.Log)
  g.APP = core.InitApp()
  middleware.InitMiddleware(g.APP)
  router.InitRoutes(g.APP)
  err := g.APP.Listen(fmt.Sprintf(":%v", g.CONFIG.System.Port))
  if err != nil {
    panic(err)
  }
}

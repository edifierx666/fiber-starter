package main

import (
  "embed"

  _ "ats3/internal/config"
  "ats3/internal/core"
  "ats3/internal/g"
  "go.uber.org/zap"
)

//go:embed resource/*
var resources embed.FS

func main() {
  _, log := core.Init()
  g.Log = log
  zap.ReplaceGlobals(g.Log)
  g.APP = core.InitApp()
  err := g.APP.Listen(":9888")
  if err != nil {
    panic(err)
  }
}

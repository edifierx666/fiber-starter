package main

import (
  "embed"

  _ "ats3/internal/config"
  "ats3/internal/core"
)

//go:embed resource/*
var resources embed.FS

func main() {
  core.Run()
}

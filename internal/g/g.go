package g

import (
  "ats3/internal/config"
  "github.com/gofiber/fiber/v2"
  "go.uber.org/zap"
)

var (
  CONFIG config.Server
  Log    *zap.Logger
  APP    *fiber.App
)

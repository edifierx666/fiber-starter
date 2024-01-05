package config

import (
  "github.com/gofiber/fiber/v2"
)

type Server struct {
  fiber.Config
}

type AppCfg struct {
  ServerCfg Server
}

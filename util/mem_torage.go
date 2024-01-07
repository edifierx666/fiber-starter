package util

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/storage/memory/v2"
)

func MemStorage() fiber.Storage {
  storage := memory.New()

  return storage
}

package core

import (
  "github.com/gofiber/fiber/v2"
  "github.com/spf13/viper"
  "go.uber.org/zap"
)

func init() {

}

func Init() (*viper.Viper, *zap.Logger) {
  v := Viper()
  z := Zap()
  return v, z
}

func InitApp(c ...fiber.Config) *fiber.App {
  var cfg fiber.Config
  if len(c) > 0 {
    cfg = c[0]
  }
  app := fiber.New(cfg)

  return app
}

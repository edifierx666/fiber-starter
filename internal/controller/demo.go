package controller

import "github.com/gofiber/fiber/v2"

type Demo struct {
}

var DemoController = new(Demo)

func (d Demo) Test(c *fiber.Ctx) error {
  return c.JSON(fiber.Map{
    "ok": "ok111111111",
  })
}

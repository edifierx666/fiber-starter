package demo

import (
	"github.com/gofiber/fiber/v2"
)

func InitDemoRoute(group fiber.Router) {

	group.Post(
		"/demo", func(c *fiber.Ctx) error {
			return c.JSON(
				fiber.Map{
					"asd": "Asd",
				},
			)
		},
	)

}

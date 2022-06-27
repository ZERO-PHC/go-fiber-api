package main

import "github.com/gofiber/fiber/v2"

func main() {
	// create instance of fiber
	app := fiber.New()

	//create handler function
	app.Get("/testEndPoint", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Go fiber api created",
		})
	})

	app.Listen(":3000")

}

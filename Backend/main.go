package main

import (
	avatar "ROK/requests"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/avatargender", avatar.Create)
	app.Get("/addassetitem", avatar.Addassets)
	app.Get("/setskincolour", avatar.Setskincolour)

	log.Fatal(app.Listen(":8080"))
}

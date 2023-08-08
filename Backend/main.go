package main

import (
	avatar "ROK/requests"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/avatargender", avatar.Create)
	app.Get("/addassetitem", avatar.Addassets)
	app.Get("/setskincolour", avatar.Setskincolour)

	// app.Get("/fullapi", func(c *fiber.Ctx) error {

	// 	URL := "https://doppelme-avatars.p.rapidapi.com/avatar/1011/"

	// 	req, _ := http.NewRequest("POST", URL, nil)

	// 	req.Header.Add("X-RapidAPI-Key", "0300c63a5emsh24ffe9004fc03a7p199a96jsnaa04a9ed76ca")
	// 	req.Header.Add("X-RapidAPI-Host", "doppelme-avatars.p.rapidapi.com")

	// 	for i := 0; i < 100; i++ {
	// 		res, err := http.DefaultClient.Do(req)
	// 		if err != nil {
	// 			log.Println(err)
	// 			return err
	// 		}

	// 		fmt.Println(res)
	// 	}
	// 	return nil
	// })

	log.Fatal(app.Listen(":8080"))
}

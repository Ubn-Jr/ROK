package avatar

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

// eso verecek
type set_skincolour struct {
	Skincolour string `json:"skin_colour"`
}

type Response3 struct {
	Body []Apiresp `json:"body"`
}

func Setskincolour(c *fiber.Ctx) error {

	var esorequest set_skincolour

	c.BodyParser(&esorequest)

	var getdoppelme Apiresp
	c.BodyParser(&getdoppelme)

	URL := fmt.Sprintf("https://doppelme-avatars.p.rapidapi.com/avatar/%s/skin/%s", getdoppelme.Doppelme_key, esorequest.Skincolour)

	req, _ := http.NewRequest("PUT", URL, nil)

	filePath := "key.json"
	jsonBytes, _ := os.ReadFile(filePath)
	var apiKeyData APIKEY

	json.Unmarshal(jsonBytes, &apiKeyData)

	ApiKEY := apiKeyData.Key
	req.Header.Add("X-RapidAPI-Key", ApiKEY)
	req.Header.Add("X-RapidAPI-Host", "doppelme-avatars.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	bodystring := string(body)

	var response Apiresp

	err = json.Unmarshal([]byte(bodystring), &response)
	if err != nil {
		log.Println(err)
		return err
	}

	err = c.JSON(response.Status)
	if err != nil {
		log.Println(err)
		return err
	}

	err = c.JSON(response.AvatarSrc)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}

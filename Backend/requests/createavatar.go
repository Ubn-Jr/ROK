package avatar

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type APIKEY struct {
	Key string `json:"APIKEY"`
}

// eso verecek
type bodytype_id struct {
	BodytypeId string `json:"bodytype_id"`
}

type Apiresp struct {
	AvatarSrc    string `json:"avatarSrc"`
	Status       string `json:"status"`
	Doppelme_key string `json:"doppelme_key"`
	ThumbnailSrc string `json:"thumbnailSrc"`
}

type Response struct {
	Body []Apiresp `json:"body"`
}

func Create(c *fiber.Ctx) error {
	var esorequest bodytype_id

	c.BodyParser(&esorequest)

	URL := fmt.Sprintf("https://doppelme-avatars.p.rapidapi.com/avatar/%s/", esorequest.BodytypeId)

	req, _ := http.NewRequest("POST", URL, nil)

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

	c.JSON(response.AvatarSrc)

	return err
}

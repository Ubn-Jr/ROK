package avatar

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

// eso verecek
type add_assetitem struct {
	AssetId string `json:"asset_id"`
}

type Response2 struct {
	Body []Apiresp `json:"body"`
}

func Addassets(c *fiber.Ctx) error {

	var esorequest add_assetitem

	c.BodyParser(&esorequest)

	var getdoppelme Apiresp
	c.BodyParser(&getdoppelme)

	URL := fmt.Sprintf("https://doppelme-avatars.p.rapidapi.com/avatar/%s/%s", getdoppelme.Doppelme_key, esorequest.AssetId)

	filePath := "key.json"
	jsonBytes, _ := os.ReadFile(filePath)

	var apiKeyData APIKEY

	err := json.Unmarshal(jsonBytes, &apiKeyData)
	if err != nil {
		fmt.Println(err)
	}

	var apiResp Apiresp

	for i := 0; i < 7; {
		ApiKEY := fmt.Sprintf("Key%d", i)

		apiKeyVal := reflect.ValueOf(apiKeyData.Api_keys[i]).FieldByName(ApiKEY).String()

		fmt.Println("apiKeyVal:", apiKeyVal)

		req, _ := http.NewRequest("PUT", URL, nil)

		req.Header.Add("X-RapidAPI-Key", apiKeyVal)
		req.Header.Add("X-RapidAPI-Host", "doppelme-avatars.p.rapidapi.com")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			return err
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		bodystring := string(body)

		err = json.Unmarshal([]byte(bodystring), &apiResp)
		if err != nil {
			log.Default()
		}

		if res.StatusCode == 429 {
			fmt.Println("apiKeyVal", apiKeyVal, bodystring)
			fmt.Println("i:", i, "API KEY FULL TRIAL")
			i += 1
			continue
		} else if res.StatusCode != 200 {
			fmt.Println("i: ", i, "else ife girdi")

			err = c.JSON(apiResp.AvatarSrc)
			if err != nil {
				log.Println(err)
				return err
			}

			break
		} else {
			fmt.Println("i: ", i, "else girdi")
			var response Apiresp
			err := json.Unmarshal([]byte(bodystring), &response)
			if err != nil {
				log.Println(err)
				return err
			}
			fmt.Println(err)

			err = c.JSON(response.Status)
			if err != nil {
				log.Println(err)
				return err
			}

			err = c.JSON(response.Doppelme_key)
			if err != nil {
				log.Println(err)
				return err
			}

			break
		}
	}
	return err
}

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

		// Get the API key from the APIKEY struct
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

		// res.StatusCode

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("io.ReadAll", err)
		}

		bodystring := string(body)
		fmt.Println(bodystring)

		err = json.Unmarshal([]byte(bodystring), &apiResp)
		if err != nil {
			log.Default()
		}

		// fmt.Println("apiResp.Message", apiResp.Message)
		if res.StatusCode == 429 {
			fmt.Println("apiKeyVal", apiKeyVal, bodystring)
			fmt.Println("i:", i, "API KEY FULL TRIAL")
			i += 1
			continue
		} else if res.StatusCode != 200 {
			fmt.Println("i: ", i, "else ife girdi")

			err = c.JSON(apiResp.Message)
			if err != nil {
				log.Println(err)
				return err
			}
			err = c.JSON(res.StatusCode)
			if err != nil {
				log.Println(err)
				return err
			}
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

			err = c.JSON(response.AvatarSrc)
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

			err = c.JSON(response.AvatarSrc)
			if err != nil {
				log.Println(err)
				return err
			}

			break
		}

	}

	return nil
}

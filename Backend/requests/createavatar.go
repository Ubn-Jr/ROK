package avatar

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type APIKEY struct {
	Api_keys []struct {
		Key0 string `json:"APIKEY0"`
		Key1 string `json:"APIKEY1"`
		Key2 string `json:"APIKEY2"`
		Key3 string `json:"APIKEY3"`
		Key4 string `json:"APIKEY4"`
		Key5 string `json:"APIKEY5"`
		Key6 string `json:"APIKEY6"`
		Key7 string `json:"APIKEY7"`
	} `json:"api_keys"`
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
	Message      string `json:"message"`
}
type Response struct {
	Body []Apiresp `json:"body"`
}

func Create(c *fiber.Ctx) error {
	var esorequest bodytype_id

	c.BodyParser(&esorequest)

	URL := fmt.Sprintf("https://doppelme-avatars.p.rapidapi.com/avatar/%s/", esorequest.BodytypeId)


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

		req, _ := http.NewRequest("POST", URL, nil)

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

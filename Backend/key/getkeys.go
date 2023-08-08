package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type APIKEY struct {
	RowNumber int    `json:"RowNumber"`
	APIKEY    string `json:"APIKEY"`
}

type APIKEYs struct {
	APIKEYs []APIKEY `json:"api_keys"`
}

func GetApiKey(rowNumber int) (string, error) {

	apiKeys, err := ReadJSONFile()
	if err != nil {
		log.Fatal(err)
		//ERROR HANDLING
	}

	for _, apiKey := range apiKeys.APIKEYs {
		if apiKey.RowNumber == rowNumber {
			return apiKey.APIKEY, nil
		}
	}
	return "", errors.New("APIKEY not found")
}

func ReadJSONFile() (*APIKEYs, error) {

	filePath := "key.json"
	jsonBytes, _ := os.ReadFile(filePath)

	jsonData := &APIKEYs{}

	err := json.Unmarshal(jsonBytes, &jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func main() {
	fmt.Println(GetApiKey(0))
}

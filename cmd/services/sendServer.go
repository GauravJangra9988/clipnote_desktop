package services

import (
	"bytes"
	"clipnote/desktop/cmd/token"
	"encoding/json"
	"fmt"
	"os"

	"log"
	"net/http"
)

func SendServer(data string) {

	BE_URL := os.Getenv("BE_URL")

	type ClipboardData struct {
		Data string `json:"data"`
	}

	payload := ClipboardData{
		Data: data,
	}

	token := token.GetToken()

	if token == "" {
		log.Println("login first")
		return
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("error json data")
		return
	}

	client := http.Client{}

	saveEndpoint := fmt.Sprintf("%s/save", BE_URL)

	req, err := http.NewRequest("POST", saveEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		log.Println("Clipboard data send")
	} else {
		log.Println("server responded with status code", resp.StatusCode)
	}
}

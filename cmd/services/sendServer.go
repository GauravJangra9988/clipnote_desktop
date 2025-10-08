package services

import (
	"bytes"
	"clipnote/desktop/cmd/token"
	"encoding/json"

	"log"
	"net/http"
)

func SendServer(data string) {

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

	req, err := http.NewRequest("POST", "http://56.228.21.202:8080/save", bytes.NewBuffer(jsonData))
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

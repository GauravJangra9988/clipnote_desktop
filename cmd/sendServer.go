package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendServer(data string) {

	type ClipboardData struct {
		Data string `json:"data"`
	}

	payload := ClipboardData{
		Data: data,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("error json data")
		return
	}

	resp, err := http.Post("http://localhost:8080/save", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("error sending to server", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Clipboard data send")
	} else {
		fmt.Println("server responded with status code", resp.StatusCode)
	}
}

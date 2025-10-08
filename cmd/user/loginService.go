package user

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"os"
	"strings"

	"clipnote/desktop/cmd/token"
)

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func Login() error {

	BE_URL := os.Getenv("BE_URL")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	payload := map[string]string{
		"email":    email,
		"password": password,
	}

	jsonData, _ := json.Marshal(payload)

	loginEndpoint := fmt.Sprintf("%s/login", BE_URL)

	res, err := http.Post(loginEndpoint, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		fmt.Println("statuscode", res.StatusCode)
		return nil
	}

	defer res.Body.Close()

	var result LoginResponse
	json.NewDecoder(res.Body).Decode(&result)

	fmt.Println(result.Message)
	fmt.Println(result.Token)

	token.SaveToken(result.Token)

	return nil

}

package token

import (
	"encoding/json"

	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Token string `json:"token"`
}

var configPath = filepath.Join(`C:\ProgramData\clipnote`, "config.json")

func SaveToken(token string) {
	log.Println("token save started")

	err := os.MkdirAll(filepath.Dir(configPath), 0700)
	if err != nil {
		log.Println("Error creating dir:", err)
	}

	cfg := Config{Token: token}
	data, _ := json.Marshal(cfg)

	err = os.WriteFile(configPath, data, 0600)
	if err != nil {
		log.Println(err)
	}

}

func GetToken() string {

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Println("Error reading token file:", err)
		return ""
	}

	var cfg Config
	json.Unmarshal(data, &cfg)
	return cfg.Token

}

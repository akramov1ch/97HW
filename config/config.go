package config

import (
	"encoding/json"
	"os"
)

var (
	WeatherAPIKey string
	DatabaseURL   string
	SMTPServer    string
	SMTPPort      string
	SMTPUser      string
	SMTPPassword  string
)

func LoadConfig(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := struct {
		WeatherAPIKey string `json:"WeatherAPIKey"`
		DatabaseURL   string `json:"DatabaseURL"`
		SMTPServer    string `json:"SMTPServer"`
		SMTPPort      string `json:"SMTPPort"`
		SMTPUser      string `json:"SMTPUser"`
		SMTPPassword  string `json:"SMTPPassword"`
	}{}

	if err := decoder.Decode(&config); err != nil {
		return err
	}

	WeatherAPIKey = config.WeatherAPIKey
	DatabaseURL = config.DatabaseURL
	SMTPServer = config.SMTPServer
	SMTPPort = config.SMTPPort
	SMTPUser = config.SMTPUser
	SMTPPassword = config.SMTPPassword

	return nil
}

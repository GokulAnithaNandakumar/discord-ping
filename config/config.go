package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := ioutil.ReadFile("config.json")
	if err == nil {
		err = json.Unmarshal(file, &config)
		if err == nil {
			Token = config.Token
			BotPrefix = config.BotPrefix
			return nil
		}
	}

	// Fallback to environment variables
	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOT_PREFIX")

	if Token == "" || BotPrefix == "" {
		return fmt.Errorf("config not found in file or environment variables")
	}

	return nil
}

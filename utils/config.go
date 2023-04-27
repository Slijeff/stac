package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	IP   string
	Port string
}

func ReadConfig(config_path string) *Config {
	var config Config

	file, err := os.Open(config_path)
	if err != nil {
		fmt.Println("error: ", err)
		panic("error reading config file")
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		fmt.Println("error: ", err)
		panic("error parsing config file")
	}

	return &config
}

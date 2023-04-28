package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	IP   string
	Port string
	Pwd  string `json:"Stac-pwd"`
}

var (
	Config *Configuration
)

func ReadConfig(config_path string) {
	var config Configuration

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

	Config = &config
}

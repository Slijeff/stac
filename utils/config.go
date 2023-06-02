package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	IP      string `json:"ip"`
	Port    string `json:"port"`
	Pwd     string `json:"stac-pwd"`
	Base    string `json:"baseDir"` // where the repos will be cloned
	StacLog string `json:"stacLog"`
	ExecLog string `json:"execLog"`
}

var (
	Config *Configuration
)

func ReadConfig(config_path string) {
	// populate default values
	config := Configuration{
		IP:   "0.0.0.0",
		Port: "8080",
	}

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

	if config.Pwd == "" {
		panic("stac-pwd field must be given in config.json")
	}
	if config.Base == "" {
		panic("baseDir field must be given in config.json")
	}
	if config.StacLog == "" {
		panic("stacLog field must be given in config.json")
	}
	if config.ExecLog == "" {
		panic("execLog field must be given in config.json")
	}

	Config = &config
}

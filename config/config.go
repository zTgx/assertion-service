package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Host     string `json:"host"`
	Key      string `json:"key"`
	NamePath string `json:"name_path"`
}

func GetServerConfig() Config {
	// Open our jsonFile
	jsonFile, err := os.Open("./config/release-config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("Successfully Opened release-config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var config Config
	json.Unmarshal(byteValue, &config)

	return config
}

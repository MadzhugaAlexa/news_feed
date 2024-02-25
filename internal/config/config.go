package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	RequestPeriod int      `json:"request_period"`
	RSS           []string `json:"rss"`
}

func LoadConfig(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("failed to open config")
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)

	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal("failed to load config")
	}

	return cfg
}

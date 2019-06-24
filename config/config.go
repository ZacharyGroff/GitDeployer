package config

import (
	"encoding/json"
	"os"
	"log"
	"io/ioutil"
)

type Config struct {
	Port string `json:"port"`
	Route string `json:"route"`
	ScriptPath string `json:"scriptPath"`
	Validate bool `json:"validate"`
}

func (config *Config) parseConfig(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &config)
	
	return err
}

func NewConfig() *Config {
	config := Config{}
	err := config.parseConfig("config/conf.json")
	
	if err != nil {
		log.Fatal("Error while parsing config")
	}
	
	return &config
}

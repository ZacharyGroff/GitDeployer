package main

import (
	"encoding/json"
	"os"
	"log"
	"io/ioutil"
)

type Config struct {
	Port string `json:"port"`
	Route string `json:"route"`
	Validate bool `json:"validate"`
}

func parseConfig() (*Config, error) {
	file, err := os.Open("conf.json")
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var config Config
	json.Unmarshal(byteValue, &config)
	
	return &config, err
}

func NewConfig() *Config {
	config, err := parseConfig()
	
	if err != nil {
		log.Fatal("Error while parsing config")
	}
	
	return config
}

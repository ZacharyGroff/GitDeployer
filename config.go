package main

import (
	"encoding/json"
	"os"
	"log"
)

type Config struct {
	port string
	route string
	validate bool
}

func parseConfig() (*Config, error) {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)	
	config := Config{}
	err := decoder.Decode(&config)
	
	return &config, err
}

func NewConfig() *Config {
	configPtr, err := parseConfig()
	
	if err != nil {
		log.Fatal("Error while parsing config")
	}
	
	return configPtr
}

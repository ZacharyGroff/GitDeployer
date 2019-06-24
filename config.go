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

func (config *Config) ParseConfig(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &config)
	
	return err
}

func NewConfig() *Config {
	config := Config{}
	err := config.ParseConfig("conf.json")
	
	if err != nil {
		log.Fatal("Error while parsing config")
	}
	
	return &config
}

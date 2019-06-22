package main

import (
	"log"
)

type PushProcessor struct {
	Deployer *Deployer
	Config *Config
}

func NewPushProcessor(deployer *Deployer, config *Config) *PushProcessor {
	return &PushProcessor{deployer, config}
}

func HandleRequest(request Request) {
	log.Printf("New push request: %s\n", request)
}

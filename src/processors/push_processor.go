package processors

import (
	"log"
	"net/http"
	"config"
)

type PushProcessor struct {
	Deployer *Deployer
	Config *config.Config
}

func NewPushProcessor(deployer *Deployer, config *config.Config) *PushProcessor {
	return &PushProcessor{deployer, config}
}

func (pushProcessor PushProcessor) HandleRequest(request *http.Request) {
	log.Printf("New push request: %+v\n", request)
}

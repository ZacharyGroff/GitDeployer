package processors

import (
	"log"
	"github.com/ZacharyGroff/GitHooks/config"
	"github.com/ZacharyGroff/GitHooks/models"
)

type PushProcessor struct {
	Deployer *Deployer
	Config *config.Config
}

func NewPushProcessor(deployer *Deployer, config *config.Config) *PushProcessor {
	return &PushProcessor{deployer, config}
}

func (pushProcessor PushProcessor) HandleMessage(message *models.Message) {
	output := pushProcessor.Deployer.Deploy()
	log.Printf("Successfully deployed script with output:\n%s\n", output)
}

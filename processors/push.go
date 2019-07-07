package processors

import (
	"log"
	"github.com/ZacharyGroff/GitDeployer/config"
	"github.com/ZacharyGroff/GitDeployer/models"
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
